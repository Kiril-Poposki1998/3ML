package docker

var DockerCompose = `
services:
	service-a:
		build: .
		container_name: service-a 
		env_file:
			- .env
		ports:
			- "127.0.0.1:8000:8000"
		healthcheck:
			test: ["CMD", "curl", "http://localhost:8000"]
			interval: 5s
			timeout: 5s
			retries: 5
			start_period: 5s
		restart: on-failure
`
