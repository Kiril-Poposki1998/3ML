package tests

import (
	handleform "3ML/handleForm"
	"os"
	"testing"
)

func TestCreateResourceProject(t *testing.T) {
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Failed to setup project: %v", err)
	}
	proj.Name = "test_project"
	proj.Path, err = os.Getwd()
	proj.Path += "/" + proj.Name
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = proj.Create()
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}
	os.RemoveAll(proj.Path + "/test_project")
}

func TestCreateResourceAnsible(t *testing.T) {
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Failed to setup project: %v", err)
	}
	proj.Name = "test_project"
	proj.Path, err = os.Getwd()
	proj.Path += "/" + proj.Name
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = proj.Create()
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	ansibleConfig := handleform.Ansible{
		Enabled:  true,
		HostName: "localhost",
		IPaddr:   "192.168.0.1",
		SSHKey:   "ssh_key.pem",
		SSHUser:  "user",
	}
	err = ansibleConfig.Create(*proj, handleform.Docker{Enabled: true})
	if err != nil {
		t.Fatalf("Failed to create Ansible resources: %v", err)
	}

	err = ansibleConfig.Create(*proj, handleform.Docker{Enabled: false})
	if err != nil {
		t.Fatalf("Failed to create Ansible resources: %v", err)
	}

	proj.Path = "/"
	err = ansibleConfig.Create(*proj, handleform.Docker{Enabled: false})
	if err == nil {
		t.Fatal("Function did not error out")
	}
	os.RemoveAll(proj.Path + "/test_project")
}

func TestCreateResourceInfra(t *testing.T) {
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Failed to setup project: %v", err)
	}
	proj.Name = "test_project"
	proj.Path, err = os.Getwd()
	proj.Path += "/" + proj.Name
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = proj.Create()
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	infrastructureConfig := handleform.Terraform{
		Enabled:  true,
		Provider: "AWS",
	}
	err = infrastructureConfig.Create(*proj)
	if err != nil {
		t.Fatalf("Failed to create Infrastructure resources: %v", err)
	}

	os.RemoveAll(proj.Path + "/test_project")
}

func TestCreateResourceDocker(t *testing.T) {
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Failed to setup project: %v", err)
	}
	proj.Name = "test_project"
	proj.Path, err = os.Getwd()
	proj.Path += "/" + proj.Name
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = proj.Create()
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	dockerConfig := handleform.Docker{
		Enabled: true,
	}
	err = dockerConfig.Create(*proj)
	if err != nil {
		t.Fatalf("Failed to create Docker resources: %v", err)
	}

	os.RemoveAll(proj.Path + "/test_project")
}

func TestCreateResouceCICD(t *testing.T) {
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Failed to setup project: %v", err)
	}
	proj.Name = "test_project"
	proj.Path, err = os.Getwd()
	proj.Path += "/" + proj.Name
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = proj.Create()
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	cicdConfig := handleform.CICD{
		Enabled: true,
	}
	err = cicdConfig.Create(*proj, handleform.Ansible{Enabled: true})
	if err != nil {
		t.Fatalf("Failed to create CICD resources: %v", err)
	}

	os.RemoveAll(proj.Path + "/test_project")
}
