package handleform

import (
	"3ML/resource_config/ansible"
	"3ML/resource_config/cicd/github"
	"3ML/resource_config/docker"
	"3ML/resource_config/terraform"
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Create project directory structure and infrastructure if enabled.
func (p Project) Create() error {
	err := os.MkdirAll(p.Path, os.ModePerm)
	if err != nil {
		return err
	}
	if p.InfraEnabled {
		err = os.MkdirAll(p.Path+"/infrastructure", os.ModePerm)
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
		err := os.MkdirAll(casc_path, 0755)
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
			err = os.WriteFile(casc_path+"main.yaml", []byte(out), 0755)
			if err != nil {
				return fmt.Errorf("failed to write main.yaml: %w", err)
			}

		} else {
			// Write main.yaml to the project directory
			out, err := build_ansible_yaml(main, casc.HostName, "", "")
			if err != nil {
				return fmt.Errorf("failed to build ansible yaml: %w", err)
			}
			err = os.WriteFile(casc_path+"main.yaml", []byte(out), 0755)
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
		err = os.WriteFile(casc_path+"ansible.cfg", buf.Bytes(), 0755)
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
		err = os.WriteFile(casc_path+"hosts", hostsBuf.Bytes(), 0755)
		if err != nil {
			return fmt.Errorf("failed to write hosts file: %w", err)
		}

		// Create the templates directory
		err = os.MkdirAll(casc_path+"templates", os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create templates directory: %w", err)
		}
		// Copy the nginx template file
		err = os.WriteFile(casc_path+"templates/template.conf", []byte(ansible.AnsibleNginxTemplate), 0755)
		if err != nil {
			return fmt.Errorf("failed to write nginx template file: %w", err)
		}
	}
	return nil
}

// Create Terraform main.tf file with provider additional config
func (iac Terraform) Create(proj Project) error {
	if iac.Enabled {
		// Create Terraform directory structure
		var iac_path = proj.Path + "/infrastructure/terraform/"
		err := os.MkdirAll(iac_path, os.ModePerm)
		if err != nil {
			return err
		}

		// Insert provider source and options
		var out string
		main, err := template.New("main").Parse(terraform.Main)
		if err != nil {
			return err
		}
		if iac.Provider == "Digital Ocean" {
			out, err = build_terraform_tf(*main, "digitalocean/digitalocean", "digitalocean", iac.ProviderVersion, terraform.DO_Additional)
			if err != nil {
				return err
			}
		} else if iac.Provider == "AWS" {
			out, err = build_terraform_tf(*main, "hashicorp/aws", "aws", iac.ProviderVersion, terraform.AWS_Additional)
			if err != nil {
				return err
			}
		} else if iac.Provider == "Azure" {
			out, err = build_terraform_tf(*main, "hashicorp/azurerm", "azurerm", iac.ProviderVersion, terraform.Azure_Additional)
			if err != nil {
				return err
			}
		} else if iac.Provider == "GCP" {
			out, err = build_terraform_tf(*main, "hashicorp/google", "google", iac.ProviderVersion, terraform.GCP_additional)
			if err != nil {
				return err
			}
		}
		err = os.WriteFile(iac_path+"main.tf", []byte(out), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create dockerfile, dockerfile.dev and docker compose if needed
func (d Docker) Create(proj Project) error {
	dockercompose_template, err := template.New("docker_compose").Parse(docker.DockerCompose)
	if err != nil {
		return fmt.Errorf("failed to parse docker compose template: %w", err)
	}

	// Run docker compose template
	dockerComposeContent, err := build_dockerfile(dockercompose_template, d.DatabaseEnabled, d.Databasetype)
	if err != nil {
		return fmt.Errorf("failed to build docker compose file: %w", err)
	}

	// Create docker compose file
	err = os.WriteFile(proj.Path+"/docker-compose.yaml", []byte(dockerComposeContent), 0600)
	if err != nil {
		return err
	}
	return nil
}

// Create CI/CD files
func (cicd CICD) Create(proj Project, casc Ansible) error {
	if cicd.Enabled {
		// Create the .github/workflows directory
		err := os.MkdirAll(proj.Path+"/.github/workflows", 0700)
		if err != nil {
			return fmt.Errorf("failed to create .github/workflows directory: %w", err)
		}

		// Get the main template
		main, err := template.New("main").Parse(github.Template)
		if err != nil {
			return fmt.Errorf("failed to parse main template: %w", err)
		}

		// Build the content of the GitHub Actions workflow file
		out, err := build_github_workflow(main, proj.Name, casc.HostName, casc.IPaddr)
		if err != nil {
			return fmt.Errorf("failed to build GitHub Actions workflow: %w", err)
		}

		// Write the workflow file
		err = os.WriteFile(proj.Path+"/.github/workflows/deploy.yaml", []byte(out), 0600)
		if err != nil {
			return fmt.Errorf("failed to write deploy.yaml: %w", err)
		}
	}
	return nil
}

func build_github_workflow(main *template.Template, projectName string, sshName string, ipAddress string) (string, error) {
	var buf bytes.Buffer
	err := main.Execute(&buf, map[string]string{
		"ProjectName": projectName,
		"SSHName":     sshName,
		"IPaddress":   ipAddress,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}
	return buf.String(), nil
}

func build_dockerfile(main *template.Template, databaseEnabled bool, databaseType string) (string, error) {
	var buf bytes.Buffer
	err := main.Execute(&buf, map[string]interface{}{
		"DatabaseEnabled": databaseEnabled,
		"Databasetype":    databaseType,
		"Postgresql":      docker.PostgresqlDockerCompose,
		"Mysql":           docker.MysqlDockerCompose,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}
	return buf.String(), nil
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

func build_terraform_tf(main template.Template, provider_source string, provider_name string, version string, additional_info string) (string, error) {
	var buf bytes.Buffer
	err := main.Execute(&buf, map[string]string{
		"provider_name":    provider_name,
		"remote_repo":      provider_source,
		"provider_version": version,
		"additional_info":  additional_info,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}
	return buf.String(), nil
}
