package docker

var DockerCompose = `
services:
  app1:	
    build: .
    container_name: app1
    env_file:
      - .env
    ports:
      - "127.0.0.1:8000:8000"
    restart: always
	{{- if .DatabaseEnabled -}}
		{{- if eq .Databasetype "PostgreSQL" -}}
		 	{{- .Postgresql -}}
		{{- else if eq .Databasetype "MySQL" -}}
		 	{{- .Mysql -}}
		{{- end -}}
	{{- end -}}
`

var PostgresqlDockerCompose = `
  database:
    image: postgres:latest
    container_name: postgres
    env_file:
      - .env
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "127.0.0.1:5432:5432"
    restart: on-failure

volumes:
  mysql_data:		
`

var MysqlDockerCompose = `
  database:
    image: mysql:latest
    container_name: mysql
    env_file:
      - .env
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
	  - "127.0.0.1:3306:3306"
	restart: on-failure

volumes:
  mysql_data:
`

var DockerComposeEnv = `# Environment variables for Docker Compose
COMPOSE_PROJECT_NAME="{{.ProjectName}}"
COMPOSE_FILE="docker-compose.yml"
COMPOSE_DOCKER_CLI_BUILD=1
COMPOSE_REMOVE_ORPHANS=1
COMPOSE_PROFILE="dev"
`
