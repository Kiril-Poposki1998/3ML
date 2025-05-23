package handleform

import "os"

func (p Project) create() error {
	err := os.Mkdir(p.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// #TODO: Create terraform, ansible, docker resources
func (casc Ansible) create() error {
	return nil
}

func (iac Terraform) create() error {
	return nil
}

func (d Docker) create() error {
	return nil
}
