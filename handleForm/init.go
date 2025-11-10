package handleform

import (
	"errors"
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
	var tools_used []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&proj.Name).Title("Name of the project").Validate(func(s string) error {
				// Project needs to be lowercase, without spaces and only letters, numbers, hyphens and underscores
				if s == "" {
					return errors.New("project name cannot be empty")
				}
				for _, char := range s {
					if !(char >= 'a' && char <= 'z') && !(char >= '0' && char <= '9') && char != '-' && char != '_' {
						return errors.New("project name can only contain lowercase letters, numbers, hyphens and underscores")
					}
				}
				return nil
			}),
			huh.NewInput().Value(&proj.Path).Placeholder(proj.Path).Title("Add project path"),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("Select tools to use").Options(
				huh.NewOption("Ansible", "ansible"),
				huh.NewOption("Terraform", "terraform"),
				huh.NewOption("Docker", "docker"),
				huh.NewOption("CICD", "cicd"),
			).Value(&tools_used),
		),
	)
	err := form.Run()
	if err != nil {
		return err
	}
	// Set the project name and path
	for _, tool := range tools_used {
		switch tool {
		case "ansible":
			casc.Enabled = true
		case "terraform":
			iac.Enabled = true
		case "docker":
			docker.Enabled = true
		case "cicd":
			cicd.Enabled = true
		default:
			return errors.New("unknown tool selected: " + tool)
		}
	}
	return nil
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
