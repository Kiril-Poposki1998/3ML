package ansible

var AnsibleConf = `
[defaults]
inventory = hosts
remote_user = {{ .user }}
private_key_file = ~/.ssh/{{ .private_key_file }}
`

var AnsiblHosts = `
[{{ .host }}]
{{ .host }} ansible_host={{ .ip }} ansible_connection=ssh
`

var AnsibleNginxTemplate = `
server {
	listen 80;
	server_name _;

	location / {
		proxy_pass http://localhost:8000;
	}
}
`
