package github

var AnsibleTemplate = `
name: "Deployment of {{ .ProjectName }}"

on:
  push:
    branches:
      - main
    paths:
      - 'infrastructure/ansible/**'

jobs:
  deploy_ansible:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Setup ssh
        run: mkdir -p ~/.ssh/ && echo "{{"${{"}} secrets.SSH_CONFIG {{"}}"}}" > ~/.ssh/config && echo "{{"${{"}} secrets.SSH_RUNNER_KEY {{"}}"}}" > ~/.ssh/runner_key && chmod 600 ~/.ssh/runner_key && ssh-keyscan -H {{ .IPaddress }} >> ~/.ssh/known_hosts 

      - name: Install Ansible
        run: |
          sudo apt-get update
          sudo apt-get install -y ansible

      - name: Run ansible playbook
        run: |
            cd infrastructure/ansible
            ansible-playbook --private-key ~/.ssh/runner_key main.yaml

  {{ if .DiscordNotifyEnabled }}
  notify:
    runs-on: ubuntu-latest
    needs: deploy_ansible
    if: always()

    steps:
      - name: Notify Discord
        uses: Ilshidur/action-discord@0.3.2
        env:
          DISCORD_WEBHOOK: {{ "${{" }} secrets.DISCORD_WEBHOOK_URL {{ "}}" }}
        with:
          args: "STATUS:{{"${{"}} needs.deploy.result {{"}}"}} Actions for project {{"${{"}} github.repository {{"}}"}} on branch main"
  {{ end }}
`
