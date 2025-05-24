package main

import (
	handleform "3ML/handleForm"
)

var (
	proj   handleform.Project
	casc   handleform.Ansible
	iac    handleform.Terraform
	docker handleform.Docker
)

func main() {
	runner := &handleform.TerminalFormRunner{}
	proj, err := handleform.SetupProject()
	if err != nil {
		panic(err)
	}
	err = handleform.CreateForm(runner, proj, &iac, &casc, &docker)
	if err != nil {
		panic(err)
	}
	handleform.AddOptions(proj, &iac, &casc, &docker)
}
