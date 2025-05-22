package handleform

import (
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
}

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

func cascForm() {
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Add a host").Value(&casc.HostName),
			huh.NewInput().Title("Add an IP addr").Value(&casc.IPaddr),
			huh.NewInput().Title("Add a SSH key").Value(&casc.SSHKey),
			huh.NewInput().Title("Add a SSH user").Value(&casc.SSHUser),
		),
	)
	provider_form.Run()
}

func dockerForm() {
	provider_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Is there a need for a dev Dockerfile").Value(&docker.DevEnabled),
			huh.NewConfirm().Title("Is there a need for compose file").Value(&docker.DevEnabled),
		),
	)
	provider_form.Run()
}
