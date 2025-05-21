package handlereq

type Project struct {
	Path string
}

type Infrastructure interface {
	createPath() error
}

type Ansible struct {
	Enabled bool
	Host    string
	SSHKey  string
	SSHUser string
}

type Terraform struct {
	Enabled  bool
	Provider string
}
