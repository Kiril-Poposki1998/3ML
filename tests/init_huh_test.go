package tests

import (
	handleform "3ML/handleForm"
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
