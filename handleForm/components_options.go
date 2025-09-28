package handleform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/charmbracelet/huh"
)

func AddOptions(proj *Project, iac *Terraform, casc *Ansible, docker *Docker, cicd *CICD) {
	if iac.RunForm() != nil {
		panic("IaC form failed to run")
	}
	if casc.RunForm() != nil {
		panic("CasC form failed to run")
	}
	if docker.RunForm() != nil {
		panic("Docker form failed to run")
	}
	if cicd.RunForm() != nil {
		panic("CICD form failed to run")
	}
	// fmt.Println(FormatAdvancedOptionsVars(*casc, *iac, *docker))
}

// Terraform options
func (iac *Terraform) RunForm() error {
	if iac.Enabled {
		provider_form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().Title("Select a provider").Value(&iac.Provider).Options(
					huh.NewOption("AWS", "AWS"),
					huh.NewOption("GCP", "GCP"),
					huh.NewOption("Azure", "Azure"),
					huh.NewOption("Digital Ocean", "Digital Ocean"),
				),
			),
		)
		err := provider_form.Run()
		if err != nil {
			return fmt.Errorf("failed to run Terraform form: %w", err)
		}
		version_placeholder, err := FetchVersionPlaceholder(iac.Provider)
		if err != nil {
			return fmt.Errorf("failed to fetch version placeholder: %w", err)
		}
		version_form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("Add a provider version").Value(&iac.ProviderVersion).Placeholder(version_placeholder).Validate(func(v string) error {
					if v == "" {
						return fmt.Errorf("provider version is required")
					} else if !regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(v) {
						return fmt.Errorf("provider version must be in the format X.Y.Z")
					}
					return nil
				}),
			),
		)
		return version_form.Run()
	}
	return nil
}

// Ansible options
func (casc *Ansible) RunForm() error {
	if casc.Enabled {
		provider_form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("Add a host").Value(&casc.HostName),
				huh.NewInput().Title("Add an IP addr").Value(&casc.IPaddr),
				huh.NewInput().Title("Add a SSH private key").Value(&casc.SSHKey).Placeholder("id_rsa"),
				huh.NewInput().Title("Add a SSH user").Value(&casc.SSHUser),
				huh.NewConfirm().Title("Enable alerts?").Value(&casc.AlertsEnabled),
			),
		)
		err := provider_form.Run()
		if err != nil {
			return fmt.Errorf("failed to run Ansible form: %w", err)
		}
	}
	return nil
}

// Docker options
func (docker *Docker) RunForm() error {
	if !docker.Enabled {
		return nil
	}
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Is there a need for compose file").Value(&docker.ComposeEnabled),
			huh.NewConfirm().Title("Is there a need for Dockerfile").Value(&docker.DockerfileEnabled),
		),
	)
	provider_form.Run()

	// Form for database options
	if docker.ComposeEnabled {
		provider_form = huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().Title("Is there a need for a database").Value(&docker.DatabaseEnabled),
			),
		)
		provider_form.Run()
	}

	// Form for which database to use
	if docker.DatabaseEnabled {
		provider_form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().Title("Select a database type").Value(&docker.Databasetype).Options(
					huh.NewOption("PostgreSQL", "PostgreSQL"),
					huh.NewOption("MySQL", "MySQL"),
				),
			),
		)
		provider_form.Run()
	}
	// Form for Dockerfile options
	if docker.DockerfileEnabled {
		dockerfile_form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().Title("Select a Dockerfile type").Value(&docker.DockerfileType).Options(
					huh.NewOption("Python", "Python"),
					huh.NewOption("Node.js", "Node.js"),
					huh.NewOption("Go", "Go"),
					huh.NewOption("Java", "Java"),
				),
			),
		)
		dockerfile_form.Run()
	}
	return nil
}

// CI/CD options
func (cicd *CICD) RunForm() error {
	if cicd.Enabled {
		provider_form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().Title("Is there a need for Discord notification?").Value(&cicd.DiscordNotificationEnabled),
			),
		)
		return provider_form.Run()
	}
	return nil
}

func FormatAdvancedOptionsVars(casc Ansible, iac Terraform, docker Docker, cicd CICD) string {
	advancedOptions := "Advanced Options:\n"
	if iac.Enabled {
		advancedOptions += fmt.Sprintf("Provider: %s\n", iac.Provider)
	}
	if casc.Enabled {
		advancedOptions += fmt.Sprintf("HostName: %s\n", casc.HostName)
		advancedOptions += fmt.Sprintf("IPaddr: %s\n", casc.IPaddr)
		advancedOptions += fmt.Sprintf("SSHKey: %s\n", casc.SSHKey)
		advancedOptions += fmt.Sprintf("SSHUser: %s\n", casc.SSHUser)
	}
	if docker.Enabled {
		advancedOptions += fmt.Sprintf("ComposeEnabled: %t\n", docker.ComposeEnabled)
	}
	if cicd.Enabled {
		advancedOptions += fmt.Sprintf("CICD enabled: %t\n", cicd.Enabled)
	}
	return advancedOptions
}

func FetchVersionPlaceholder(provider string) (string, error) {
	switch provider {
	case "AWS":
		return fmt.Sprintf("Currrent version is %s", FetchandDecodeVersion("https://registry.terraform.io/v1/providers/hashicorp/aws")), nil
	case "GCP":
		return fmt.Sprintf("Currrent version is %s", FetchandDecodeVersion("https://registry.terraform.io/v1/providers/hashicorp/google")), nil
	case "Azure":
		return fmt.Sprintf("Currrent version is %s", FetchandDecodeVersion("https://registry.terraform.io/v1/providers/hashicorp/azurerm")), nil
	case "Digital Ocean":
		return fmt.Sprintf("Currrent version is %s", FetchandDecodeVersion("https://registry.terraform.io/v1/providers/digitalocean/digitalocean")), nil
	default:
		return "eg. 3.5.0", nil // Default placeholder if provider is not recognized
	}
}

func FetchandDecodeVersion(url string) string {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "Cannot fetch provider version use eg. 3.5.0"
	}
	defer resp.Body.Close()

	var data struct {
		Version string `json:"version"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "Cannot decode provider version use eg. 3.5.0"
	}
	return data.Version
}
