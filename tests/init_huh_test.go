package tests

import (
	handleform "3ML/handleForm"
	"errors"
	"testing"
)

type MockFormRunner struct {
	RunFunc func(proj *handleform.Project, iac *handleform.Terraform, casc *handleform.Ansible, docker *handleform.Docker) error
}

func (m *MockFormRunner) RunForm(proj *handleform.Project, iac *handleform.Terraform, casc *handleform.Ansible, docker *handleform.Docker) error {
	return m.RunFunc(proj, iac, casc, docker)
}

func TestProjectSetup(t *testing.T) {
	// Test the SetupProject function
	proj, err := handleform.SetupProject()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if proj.Path == "" {
		t.Errorf("Expected project path to be set, got empty string")
	}
}

func TestSetupProject_Error(t *testing.T) {
	// Backup the original function and restore it after the test
	originalGetwd := handleform.Getwd
	defer func() { handleform.Getwd = originalGetwd }()

	// Inject a failing function
	handleform.Getwd = func() (string, error) {
		return "", errors.New("mock error")
	}

	_, err := handleform.SetupProject()
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestCreateForm_Success(t *testing.T) {
	originalGetwd := handleform.Getwd
	defer func() { handleform.Getwd = originalGetwd }()

	handleform.Getwd = func() (string, error) {
		return "/mock/path", nil
	}

	proj := &handleform.Project{}
	iac := &handleform.Terraform{}
	casc := &handleform.Ansible{}
	docker := &handleform.Docker{}

	runner := &MockFormRunner{
		RunFunc: func(p *handleform.Project, i *handleform.Terraform, c *handleform.Ansible, d *handleform.Docker) error {
			// simulate form interaction
			p.Name = "TestProject"
			p.Path = "/mock/path"
			c.Enabled = true
			i.Enabled = false
			d.Enabled = true
			return nil
		},
	}

	err := handleform.CreateForm(runner, proj, iac, casc, docker)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if proj.Path != "/mock/path" {
		t.Errorf("unexpected project values: %+v", proj)
	}
	if !casc.Enabled || docker.Enabled != true || iac.Enabled != false {
		t.Errorf("unexpected service values: casc=%v, iac=%v, docker=%v", casc.Enabled, iac.Enabled, docker.Enabled)
	}
}
