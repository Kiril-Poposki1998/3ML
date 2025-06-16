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
var CheckPortScript = `
#!/bin/bash

# Configuration
URL=$1
PROJECT=$2
WEBHOOK_URL=$3


send_discord_alert() {
  local message="$1"
  curl -H "Content-Type: application/json" \
       -X POST \
       -d "{\"content\": \"$message\"}" \
       "$WEBHOOK_URL"
}

curl --silent --head --fail "$URL" > /dev/null
if [ $? -ne 0 ]; then
        send_discord_alert ":rotating_light: ALERT: Cannot reach project $PROJECT"
fi
`

var CheckDiskSpaceScript = `
#!/bin/bash

# Configuration
WEBHOOK_URL=$2

send_discord_alert() {
  local message="$1"
  curl -H "Content-Type: application/json" \
       -X POST \
       -d "{\"content\": \"$message\"}" \
       "$WEBHOOK_URL"
}

disk_space=$(df -h / | awk 'NR==2 {print $5}' | sed 's/%//')
if [ "$disk_space" -gt 90 ]; then
  send_discord_alert "Disk space on $HOSTNAME is critically low: ${disk_space}% used. Please check the server."
fi
`
