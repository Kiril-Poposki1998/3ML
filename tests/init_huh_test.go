package tests

import (
	handleform "3ML/handleForm"
	"os"
	"testing"
)

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

func TestCreateForm(t *testing.T) {
	output, err := handleform.CreateForm()
	if err != nil {
		t.Fatalf("Cannot create form, %v", err)
	}
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != ("Proj path: " + path + "\n") {
		t.Errorf("Not the expected output, got %s", output)
	}
}
