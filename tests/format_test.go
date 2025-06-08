package tests

import (
	handleform "3ML/handleForm"
	"testing"
)

func TestInitFormat(t *testing.T) {
	proj := handleform.Project{
		Name: "test",
	}
	casc := handleform.Ansible{
		Enabled: true,
	}
	iac := handleform.Terraform{
		Enabled: true,
	}
	docker := handleform.Docker{
		Enabled: true,
	}
	cicd := handleform.CICD{
		Enabled: true,
	}
	expected := "Project path: \n" +
		"Ansible enabled: true\n" +
		"Terraform enabled: true\n" +
		"Docker enabled: true\n" +
		"CICD enabled: true\n"
	result := handleform.FormatVars(&proj, &casc, &iac, &docker, &cicd)
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestAdvancedFormat(t *testing.T) {
	casc := handleform.Ansible{
		Enabled:  true,
		HostName: "test-host",
		IPaddr:   "192.168.0.1",
		SSHKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC...",
		SSHUser:  "test-user",
	}
	iac := handleform.Terraform{
		Enabled:  true,
		Provider: "AWS",
	}
	docker := handleform.Docker{
		Enabled:        true,
		ComposeEnabled: true,
	}
	expected := "Advanced Options:\n" +
		"Provider: AWS\n" +
		"HostName: test-host\n" +
		"IPaddr: 192.168.0.1\n" +
		"SSHKey: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC...\n" +
		"SSHUser: test-user\n" +
		"DevEnabled: true\n" +
		"ComposeEnabled: true\n"
	result := handleform.FormatAdvancedOptionsVars(casc, iac, docker)
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
