package main

import (
	handleform "3ML/handleForm"
)

var (
	proj   handleform.Project
	casc   handleform.Ansible
	iac    handleform.Terraform
	docker handleform.Docker
	cicd   handleform.CICD
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create project form
	runner := &handleform.TerminalFormRunner{}
	proj, err := handleform.SetupProject()
	check(err)

	// Run forms for configuring elements
	err = handleform.CreateForm(runner, proj, &iac, &casc, &docker, &cicd)
	check(err)
	handleform.AddOptions(proj, &iac, &casc, &docker, &cicd)

	// Create resources
	err = proj.Create()
	check(err)
	err = casc.Create(*proj, docker)
	check(err)
	err = iac.Create(*proj)
	check(err)
	err = docker.Create(*proj)
	check(err)
	err = cicd.Create(*proj, casc)
	check(err)
}
