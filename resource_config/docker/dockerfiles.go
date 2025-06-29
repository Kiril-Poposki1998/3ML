package docker

var NodeDockerfile = `ARG NODE_VERSION=20
FROM node:${NODE_VERSION}
WORKDIR /app

COPY package*.json ./
RUN npm ci
COPY . .

CMD ["node", "index.js"]
`

var PythonDockerfile = `ARG PYTHON_VERSION=3.11
FROM python:${PYTHON_VERSION}

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

CMD ["python3", "app.py"]
`

var GolangDockerfile = `ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD ["go", "run", "main.go"]
`

var JavaDockerfile = `ARG JAVA_VERSION=17
FROM openjdk:${JAVA_VERSION}-jdk

RUN apk add --no-cache bash curl git

WORKDIR /app

COPY .mvn/ .mvn
COPY mvnw pom.xml ./

RUN sed -i 's/\r$//' mvnw

RUN ./mvnw dependency:go-offline

COPY . .

RUN ./mvnw package -DskipTests

CMD ["java", "-jar", "target/run.jar"]
`
