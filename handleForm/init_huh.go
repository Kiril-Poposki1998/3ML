package handleform

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	proj          Project
	casc          Ansible
	iac           Terraform
	docker        Docker
	dockerCompose DockerCompose
	err           error
)

func CreateForm() (string, error) {
	// Create a new form
	proj, err = SetupProject()
	if err != nil {
		return "", err
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&proj.Path).Placeholder(proj.Path).Title("Add project path"),
		),
		huh.NewGroup(
			huh.NewConfirm().Value(&casc.Enabled).Title("Do you want to use Terraform?"),
			huh.NewConfirm().Value(&iac.Enabled).Title("Do you want to use Ansible?"),
			huh.NewConfirm().Value(&docker.Enabled).Title("Do you want to use Dockerfile?"),
			huh.NewConfirm().Value(&dockerCompose.Enabled).Title("Do you want to use docker compose?"),
		),
	)
	// Run the form
	form.Run()
	return formatVars(proj, casc, iac), nil
}

func formatVars(proj Project, casc Ansible, iac Terraform) string {
	return "Project path: " + proj.Path + "\n" +
		"Ansible enabled: " + strconv.FormatBool(casc.Enabled) + "\n" +
		"Terraform enabled: " + strconv.FormatBool(iac.Enabled) + "\n" +
		"Docker enabled: " + strconv.FormatBool(docker.Enabled) + "\n" +
		"Docker Compose enabled: " + strconv.FormatBool(dockerCompose.Enabled) + "\n"
}
