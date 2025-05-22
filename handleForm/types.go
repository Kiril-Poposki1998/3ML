package handleform

import (
	"log"
	"os"
)

type Project struct {
	Path string
}

type Infrastructure interface {
	createPath() error
}

type Ansible struct {
	Enabled  bool
	HostName string
	IPaddr   string
	SSHKey   string
	SSHUser  string
}

type Terraform struct {
	Enabled  bool
	Provider string
}

type Docker struct {
	Enabled        bool
	DevEnabled     bool
	ComposeEnabled bool
}

func SetupProject() (Project, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return Project{Path: ""}, err
	}
	return Project{
		Path: path,
	}, nil
}
