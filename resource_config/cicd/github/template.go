package github

var Template = `
name: "Deployment of {{ .ProjectName }}"

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Setup ssh
        run: mkdir -p ~/.ssh/ && echo "{{"${{"}} secrets.SSH_CONFIG {{"}}"}}" > ~/.ssh/config && echo "{{"${{"}} secrets.SSH_RUNNER_KEY {{"}}"}}" > ~/.ssh/runner_key && chmod 600 ~/.ssh/runner_key && ssh-keyscan -H "{{"${{"}} env.IPaddress {{"}}"}}" >> ~/.ssh/known_hosts

      - name: Install rsync
        run: sudo apt-get update && sudo apt-get install rsync -y

      - name: Transfer files
        run: rsync -av --exclude-from=.rsync_ignore . {{ .SSHName }}:~/{{ .SSHName }}

      - name: Build image
        run: |
          ssh {{ .SSHName }} << 'EOF'
            cd ~/{{ .SSHName }}
            docker compose build
          EOF

      - name: Run new image
        run: |
          ssh {{ .SSHName }} << 'EOF'
            cd ~/{{ .SSHName }}
            docker compose up -d
          EOF

  notify:
    runs-on: ubuntu-latest
    needs: deploy
    if: always()

    steps:
      - name: Notify Discord
        uses: Ilshidur/action-discord@0.3.2
        env:
          DISCORD_WEBHOOK: {{ "${{" }} secrets.DISCORD_WEBHOOK_URL {{ "}}" }}
        with:
          args: "STATUS:{{"${{"}} needs.deploy.result {{"}}"}} Actions for project {{"${{"}} github.repository {{"}}"}} on branch main"
`
