package handleform

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func AddOptions() {
	if iac.Enabled {
		iacForm()
	}
	if casc.Enabled {
		cascForm()
	}
	if docker.Enabled {
		dockerForm()
	}
	fmt.Println(FormatAdvancedOptionsVars(casc, iac, docker))
}

// Terraform options
func iacForm() {
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
	provider_form.Run()
}

// Ansible options
func cascForm() {
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Add a host").Value(&casc.HostName),
			huh.NewInput().Title("Add an IP addr").Value(&casc.IPaddr),
			huh.NewInput().Title("Add a SSH pub key").Value(&casc.SSHKey),
			huh.NewInput().Title("Add a SSH user").Value(&casc.SSHUser),
		),
	)
	provider_form.Run()
}

// Docker options
func dockerForm() {
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Is there a need for a dev Dockerfile").Value(&docker.DevEnabled),
			huh.NewConfirm().Title("Is there a need for compose file").Value(&docker.ComposeEnabled),
		),
	)
	provider_form.Run()
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
