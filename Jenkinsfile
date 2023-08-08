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
                sh 'echo "Build Stage!" && go run ./main.go'
            }
        }

        stage('Test') {
            steps {
                // Run tests from main_test.go
                sh 'echo "Test Stage!" && go test -v ./main_test.go'
            }
        }

        stage('Deploy') {
            steps {
            	
                sh 'echo "Deploy Stage!" && 'go run main.go'
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

