package handleform

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func AddOptions(proj *Project, iac *Terraform, casc *Ansible, docker *Docker) {
	if iac.RunForm() != nil {
		panic("IaC form failed to run")
	}
	if casc.RunForm() != nil {
		panic("CasC form failed to run")
	}
	if docker.RunForm() != nil {
		panic("Docker form failed to run")
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
				huh.NewInput().Title("Provider version").Value(&iac.ProviderVersion),
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
				huh.NewInput().Title("Add a SSH pub key").Value(&casc.SSHKey),
				huh.NewInput().Title("Add a SSH user").Value(&casc.SSHUser),
			),
		)
		return provider_form.Run()
	}
	return nil
}

// Docker options
func (docker *Docker) RunForm() error {
	if docker.Enabled {
		provider_form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().Title("Is there a need for a dev Dockerfile").Value(&docker.DevEnabled),
				huh.NewConfirm().Title("Is there a need for compose file").Value(&docker.ComposeEnabled),
			),
		)
		provider_form.Run()
	}
	return nil
}

func FormatAdvancedOptionsVars(casc Ansible, iac Terraform, docker Docker) string {
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
		advancedOptions += fmt.Sprintf("DevEnabled: %t\n", docker.DevEnabled)
		advancedOptions += fmt.Sprintf("ComposeEnabled: %t\n", docker.ComposeEnabled)
	}
	return advancedOptions
}
