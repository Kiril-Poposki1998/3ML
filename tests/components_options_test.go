package tests

import (
	componets "3ML/handleForm"
	"testing"
)

type Form interface {
	RunForm() error
}

func TestAddOptions(t *testing.T) {
	t.Run("should not panic when all RunForm succeed", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Unexpected panic: %v", r)
			}
		}()
		componets.AddOptions(&componets.Project{}, &componets.Terraform{}, &componets.Ansible{}, &componets.Docker{}, &componets.CICD{})
	})

	// t.Run("should not panic when all RunForm succeed", func(t *testing.T) {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			t.Errorf("Unexpected panic: %v", r)
	// 		}
	// 	}()
	// 	componets.AddOptions(&componets.Project{Name: "Project Name", Path: "."}, &componets.Terraform{Enabled: true, Provider: "AWS", ProviderVersion: "3.0.0"}, &componets.Ansible{}, &componets.Docker{}, &componets.CICD{})
	// })
	// 	t.Run("should panic when casc.RunForm fails", func(t *testing.T) {
	// 		defer func() {
	// 			if r := recover(); r == nil {
	// 				t.Errorf("Expected panic but did not get one")
	// 			}
	// 		}()
	// 		AddOptions(&Project{}, &fakeTerraform{}, &failingForm{}, &fakeDocker{}, &fakeCICD{})
	// 	})

	//	t.Run("should panic when docker.RunForm fails", func(t *testing.T) {
	//		defer func() {
	//			if r := recover(); r == nil {
	//				t.Errorf("Expected panic but did not get one")
	//			}
	//		}()
	//		AddOptions(&Project{}, &fakeTerraform{}, &fakeAnsible{}, &failingForm{}, &fakeCICD{})
	//	})
}
