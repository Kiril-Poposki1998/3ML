package handleform

import (
	"fmt"

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
	// if cicd.RunForm() != nil {
	// 	panic("CICD form failed to run")
	// }
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
				huh.NewInput().Title("Provider version").Value(&iac.ProviderVersion).Placeholder("e.g. '~> 3.0.0'"),
			),
		)
		return provider_form.Run()
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
			),
		)
		err := provider_form.Run()
		if err != nil {
			return fmt.Errorf("failed to run Ansible form: %w", err)
		}
		alerts_form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().Title("Enable alerts?").Value(&casc.AlertsEnabled),
			),
		)
		err = alerts_form.Run()
		if err != nil {
			return fmt.Errorf("failed to run alerts form: %w", err)
		}
	}
	return nil
}

// Docker options
func (docker *Docker) RunForm() error {
	if !docker.Enabled {
		return nil
	}
	// TODO Implement Dockerfile logic
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Is there a need for compose file").Value(&docker.ComposeEnabled),
		),
	)
	provider_form.Run()

	// Form for database options
	if !docker.ComposeEnabled {
		return nil
	}
	provider_form = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Is there a need for a database").Value(&docker.DatabaseEnabled),
		),
	)
	provider_form.Run()

	// Form for which database to use
	if !docker.DatabaseEnabled {
		return nil
	}
	provider_form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select a database type").Value(&docker.Databasetype).Options(
				huh.NewOption("PostgreSQL", "PostgreSQL"),
				huh.NewOption("MySQL", "MySQL"),
			),
		),
	)
	provider_form.Run()
	return nil
}

// CI/CD options
// func (cicd *CICD) RunForm() error {
// 	if cicd.Enabled {
// 		provider_form := huh.NewForm(
// 			huh.NewGroup(
// 				huh.NewConfirm().Title("Is there a need for CI/CD?").Value(&cicd.Enabled),
// 			),
// 		)
// 		return provider_form.Run()
// 	}
// 	return nil
// }

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
