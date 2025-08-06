package handleform

import (
	"log"
	"os"
)

var Getwd = os.Getwd

// Every resource that is created in the project should implement this interface.
type Resource interface {
	create() error
}

// This is a form for each tool used in the project like ansible, terraform, docker, etc.
type ToolForm interface {
	RunForm() error
}

// Project represents the project structure and contains the path and name of the project.
type Project struct {
	Name         string
	Path         string
	InfraEnabled bool
}

// Ansible represents the CasC configuration for the project.
type Ansible struct {
	Enabled       bool
	HostName      string
	IPaddr        string
	SSHKey        string
	SSHUser       string
	AlertsEnabled bool
}

// Terraform represents the IaC configuration for the project.
type Terraform struct {
	Enabled         bool
	Provider        string
	ProviderVersion string
}

// Docker represents the containerization configuration for the project.
type Docker struct {
	Enabled           bool
	ComposeEnabled    bool
	DatabaseEnabled   bool
	DockerfileEnabled bool
	DockerfileType    string
	Databasetype      string
}

type CICD struct {
	Enabled                    bool
	DiscordNotificationEnabled bool
}

// Project constructor
func SetupProject() (*Project, error) {
	path, err := Getwd()
	if err != nil {
		log.Printf("%v", err)
		return &Project{Path: ""}, err
	}
	return &Project{
		Path: path,
	}, nil
}
