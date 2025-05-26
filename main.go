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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runner := &handleform.TerminalFormRunner{}
	proj, err := handleform.SetupProject()
	check(err)
	err = handleform.CreateForm(runner, proj, &iac, &casc, &docker)
	check(err)
	handleform.AddOptions(proj, &iac, &casc, &docker)
	err = proj.Create()
	check(err)
	err = casc.Create(*proj, docker)
	check(err)
	err = iac.Create(*proj)
	check(err)
}
