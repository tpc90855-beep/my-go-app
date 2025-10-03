pipeline {
  agent any

  environment {
    IMAGE = "tpc90/my-go-app"
  }

  options {
    skipDefaultCheckout(false)
    timestamps()
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build (Go)') {
      steps {
        sh 'go version || true'
        sh 'go env'
        sh 'go mod download'
        sh 'go build -v -o my-go-app ./cmd || go build -v -o my-go-app .'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./... || true'
      }
    }

    stage('Docker Build') {
      steps {
        script {
          IMAGE_TAG = sh(returnStdout: true, script: "git rev-parse --short HEAD").trim()
          sh "docker build -t ${IMAGE}:${IMAGE_TAG} -t ${IMAGE}:latest ."
        }
      }
    }

    stage('Push to Docker Hub') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
          sh 'echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin'
          sh "docker push ${IMAGE}:${IMAGE_TAG}"
          sh "docker push ${IMAGE}:latest"
        }
      }
    }

    stage('Deploy (optional)') {
      when {
        expression { return env.DEPLOY_HOST != null && env.DEPLOY_USER != null }
      }
      steps {
        withCredentials([sshUserPrivateKey(credentialsId: 'deploy-ssh-key', keyFileVariable: 'SSH_KEY', usernameVariable: 'DEPLOY_USER')]) {
          sh '''
            chmod 600 $SSH_KEY
            ssh -o StrictHostKeyChecking=no -i $SSH_KEY ${DEPLOY_USER}@${DEPLOY_HOST} "
              docker pull ${IMAGE}:latest &&
              docker stop my-go-app || true &&
              docker rm my-go-app || true &&
              docker run -d --restart unless-stopped --name my-go-app -p 8080:8080 ${IMAGE}:latest
            "
          '''
        }
      }
    }
  }

  post {
    always {
      archiveArtifacts artifacts: 'my-go-app', allowEmptyArchive: true
      echo "BUILD FINISHED"
    }
    success {
      echo "Pipeline succeeded"
    }
    failure {
      echo "Pipeline failed"
    }
  }
}
