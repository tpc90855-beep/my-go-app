pipeline {
agent any
environment {
IMAGE = "tpc90/my-go-app"
DOCKER_CREDENTIALS = "dockerhub-tpc90" // Jenkins credentials ID (username/password)
SSH_CREDENTIALS = "ssh-deploy" // Jenkins SSH credential ID for the deploy user
DEPLOY_USER = "deploy" // remote server username (create this on 192.168.0.22)
DEPLOY_HOST = "192.168.0.155" // your Jenkins / deploy host
REMOTE_PORT = "8081" // host port to publish the app
}
stages {
stage('Checkout') {
steps {
checkout scm
}
}


stage('Build & Test') {
agent {
docker { image 'golang:1.20-alpine' }
}
steps {
sh 'go version'
sh 'go mod download'
sh 'go test ./... -v'
sh 'go build -o app'
}
}


stage('Docker Build') {
steps {
script {
dockerImage = docker.build("${IMAGE}:${env.BUILD_NUMBER}")
}
}
}


stage('Push to Registry') {
steps {
script {
docker.withRegistry('https://index.docker.io/v1/', DOCKER_CREDENTIALS) {
dockerImage.push()
dockerImage.push('latest')
}
}
}
}