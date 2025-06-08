package handleform

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

// TerminalFormRunner is used as the original form runner
type TerminalFormRunner struct{}

// Interface for running forms
type FormRunner interface {
	RunForm(proj *Project, iac *Terraform, casc *Ansible, docker *Docker, cicd *CICD) error
}

func (r *TerminalFormRunner) RunForm(proj *Project, iac *Terraform, casc *Ansible, docker *Docker, cicd *CICD) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&proj.Name).Title("Name of the project"),
			huh.NewInput().Value(&proj.Path).Placeholder(proj.Path).Title("Add project path"),
		),
		huh.NewGroup(
			huh.NewConfirm().Value(&iac.Enabled).Title("Do you want to use Terraform?"),
			huh.NewConfirm().Value(&casc.Enabled).Title("Do you want to use Ansible?"),
			huh.NewConfirm().Value(&docker.Enabled).Title("Do you want to use Docker?"),
			huh.NewConfirm().Title("Is there a need for CI/CD?").Value(&cicd.Enabled),
		),
	)
	return form.Run()
}

// CreateForm initializes the project and add basic options
func CreateForm(runner FormRunner, proj *Project, iac *Terraform, casc *Ansible, docker *Docker, cicd *CICD) error {
	err := runner.RunForm(proj, iac, casc, docker, cicd)
	if err != nil {
		return err
	}
	// fmt.Println(FormatVars(proj, casc, iac, docker))
	return nil
}

// Format the initial variables into a string
func FormatVars(proj *Project, casc *Ansible, iac *Terraform, docker *Docker, cicd *CICD) string {
	return "Project path: " + proj.Path + "\n" +
		"Ansible enabled: " + strconv.FormatBool(casc.Enabled) + "\n" +
		"Terraform enabled: " + strconv.FormatBool(iac.Enabled) + "\n" +
		"Docker enabled: " + strconv.FormatBool(docker.Enabled) + "\n" +
		"CICD enabled: " + strconv.FormatBool(cicd.Enabled) + "\n"
}
