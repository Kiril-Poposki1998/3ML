package handleform

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	proj   Project
	casc   Ansible
	iac    Terraform
	docker Docker
	err    error
)

// CreateForm initializes the project and add basic options
func CreateForm() error {
	proj, err = SetupProject()
	if err != nil {
		return err
	}
	basic_form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&proj.Path).Placeholder(proj.Path).Title("Add project path"),
		),
		huh.NewGroup(
			huh.NewConfirm().Value(&casc.Enabled).Title("Do you want to use Terraform?"),
			huh.NewConfirm().Value(&iac.Enabled).Title("Do you want to use Ansible?"),
			huh.NewConfirm().Value(&docker.Enabled).Title("Do you want to use Docker?"),
		),
	)
	basic_form.Run()
	fmt.Println(formatVars(proj, casc, iac))
	return nil
}

// Format the initial variables into a string
func formatVars(proj Project, casc Ansible, iac Terraform) string {
	return "Project path: " + proj.Path + "\n" +
		"Ansible enabled: " + strconv.FormatBool(casc.Enabled) + "\n" +
		"Terraform enabled: " + strconv.FormatBool(iac.Enabled) + "\n" +
		"Docker enabled: " + strconv.FormatBool(docker.Enabled) + "\n"
}
