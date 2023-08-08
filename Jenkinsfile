pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }

        stage('Test') {
            steps {
                // Run tests from main_test.go
                sh 'go test -v ./main_test.go'
            }
        }

        stage('Deploy') {
            steps {
                sh './myapp'
            }
        }
    }

    post {
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}

