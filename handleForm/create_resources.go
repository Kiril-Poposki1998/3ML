package handleform

import (
	"3ML/resource_config/ansible"
	"bytes"
	"fmt"
	"os"
	"text/template"
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
func (casc Ansible) Create(proj Project, docker Docker) error {
	if casc.Enabled {
		// Create Ansible directory structure
		var casc_path = proj.Path + "/infrastructure/ansible/"
		err := os.MkdirAll(casc_path, os.ModePerm)
		if err != nil {
			return err
		}

		// Get the main.yaml template
		main, err := template.New("main").Parse(ansible.Main)
		if err != nil {
			return fmt.Errorf("failed to parse main template: %w", err)
		}

		// Write docker installation and cronjobs in main.yaml
		if docker.Enabled {
			// Build the ansible yaml content
			out, err := build_ansible_yaml(main, casc.HostName, ansible.AnsibleDocker, ansible.DockerCronJobs)
			if err != nil {
				return fmt.Errorf("failed to build ansible yaml: %w", err)
			}

			// Write main.yaml to the project directory
			err = os.WriteFile(casc_path+"main.yaml", []byte(out), 0600)
			if err != nil {
				return fmt.Errorf("failed to write main.yaml: %w", err)
			}

		} else {
			// Write main.yaml to the project directory
			out, err := build_ansible_yaml(main, casc.HostName, "", "")
			if err != nil {
				return fmt.Errorf("failed to build ansible yaml: %w", err)
			}
			err = os.WriteFile(casc_path+"main.yaml", []byte(out), 0600)
			if err != nil {
				return fmt.Errorf("failed to write main.yaml: %w", err)
			}
		}

		// Build the ansible configuration file
		ansibleConf, err := template.New("ansible_conf").Parse(ansible.AnsibleConf)
		if err != nil {
			return fmt.Errorf("failed to parse ansible configuration template: %w", err)
		}
		var buf bytes.Buffer
		err = ansibleConf.Execute(&buf, map[string]string{
			"user":             casc.SSHUser,
			"private_key_file": casc.SSHKey,
		})
		if err != nil {
			return fmt.Errorf("failed to execute ansible configuration template: %w", err)
		}
		err = os.WriteFile(casc_path+"ansible.cfg", buf.Bytes(), 0600)
		if err != nil {
			return fmt.Errorf("failed to write ansible.cfg: %w", err)
		}

		// Create the hosts file
		ansibleHosts, err := template.New("ansible_hosts").Parse(ansible.AnsiblHosts)
		if err != nil {
			return fmt.Errorf("failed to parse ansible hosts template: %w", err)
		}
		var hostsBuf bytes.Buffer
		err = ansibleHosts.Execute(&hostsBuf, map[string]string{
			"host": casc.HostName,
			"ip":   casc.IPaddr,
			"user": casc.SSHUser,
		})
		if err != nil {
			return fmt.Errorf("failed to execute ansible hosts template: %w", err)
		}
		err = os.WriteFile(casc_path+"hosts", hostsBuf.Bytes(), 0600)
		if err != nil {
			return fmt.Errorf("failed to write hosts file: %w", err)
		}

		// Create the templates directory
		err = os.Mkdir(casc_path+"templates", os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create templates directory: %w", err)
		}
		// Copy the nginx template file
		err = os.WriteFile(casc_path+"templates/template.conf", []byte(ansible.AnsibleNginxTemplate), 0600)
		if err != nil {
			return fmt.Errorf("failed to write nginx template file: %w", err)
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

func build_ansible_yaml(main *template.Template, host string, docker_tasks string, docker_cronjob string) (string, error) {
	var buf bytes.Buffer
	err := main.Execute(&buf, map[string]string{
		"host":           host,
		"DockerTasks":    docker_tasks,
		"DockerCronJobs": docker_cronjob,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}
	return buf.String(), nil
}
