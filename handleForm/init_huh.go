package handleform

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

// TerminalFormRunner is used as the original form runner
type TerminalFormRunner struct{}

// Interface for running forms
type FormRunner interface {
	RunForm(proj *Project, iac *Terraform, casc *Ansible, docker *Docker) error
}

func (r *TerminalFormRunner) RunForm(proj *Project, iac *Terraform, casc *Ansible, docker *Docker) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&proj.Name).Title("Name of the project"),
			huh.NewInput().Value(&proj.Path).Placeholder(proj.Path).Title("Add project path"),
		),
		huh.NewGroup(
			huh.NewConfirm().Value(&casc.Enabled).Title("Do you want to use Terraform?"),
			huh.NewConfirm().Value(&iac.Enabled).Title("Do you want to use Ansible?"),
			huh.NewConfirm().Value(&docker.Enabled).Title("Do you want to use Docker?"),
		),
	)
	return form.Run()
}

// CreateForm initializes the project and add basic options
func CreateForm(runner FormRunner, proj *Project, iac *Terraform, casc *Ansible, docker *Docker) error {
	proj, err := SetupProject()
	if err != nil {
		return err
	}
	err = runner.RunForm(proj, iac, casc, docker)
	if err != nil {
		return err
	}
	// fmt.Println(FormatVars(proj, casc, iac, docker))
	return nil
}

// Format the initial variables into a string
func FormatVars(proj *Project, casc *Ansible, iac *Terraform, docker *Docker) string {
	return "Project path: " + proj.Path + "\n" +
		"Ansible enabled: " + strconv.FormatBool(casc.Enabled) + "\n" +
		"Terraform enabled: " + strconv.FormatBool(iac.Enabled) + "\n" +
		"Docker enabled: " + strconv.FormatBool(docker.Enabled) + "\n"
}
