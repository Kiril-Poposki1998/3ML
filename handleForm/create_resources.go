package handleform

import (
	"3ML/resource_config/ansible"
	"os"
)

// Create project directory structure and infrastructure if enabled.
func (p Project) Create() error {
	err := os.Mkdir(p.Path, os.ModePerm)
	if err != nil {
		return err
	}
	if p.InfraEnabled {
		err = os.Mkdir(p.Path+"/infrastructure", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create the necessary files and directories for ansible
func (casc Ansible) Create(proj Project) error {
	if casc.Enabled {
		// Create Ansible directory structure
		var casc_path = proj.Path + "/infrastructure/ansible/"
		err := os.MkdirAll(casc_path, os.ModePerm)
		if err != nil {
			return err
		}
		// Read from config/ansible/main.yaml and create main.yaml in the project directory
		main := ansible.Main
		err = os.WriteFile(casc_path+"main.yaml", []byte(main), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: Implement create for terraform
func (iac Terraform) Create() error {
	return nil
}

// TODO: implement create for docker
func (d Docker) Create() error {
	return nil
}
