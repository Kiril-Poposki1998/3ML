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
