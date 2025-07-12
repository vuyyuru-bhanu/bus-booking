pipeline {
    agent any

    environment {
        DOCKER_COMPOSE_FILE = 'docker-compose.yaml'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Images') {
            steps {
                sh 'docker-compose build'
            }
        }

        stage('Run Backend Tests') {
            steps {
                dir('backend') {
                    sh 'go test ./...'
                }
            }
        }

        stage('Run Frontend Tests') {
            steps {
                dir('frontend') {
                    sh 'npm install'
                    sh 'npm run test -- --watch=false --browsers=ChromeHeadless'
                }
            }
        }

        stage('Deploy') {
            steps {
                sh 'docker-compose up -d'
            }
        }
    }

    post {
        always {
            echo 'Cleaning up...'
            sh 'docker-compose down'
        }
    }
}
