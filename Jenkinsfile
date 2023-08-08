pipeline {
    agent {
        label 'my-go-agent'
    }
    tools {
        go '1.20.7'
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
            }
        }
        stage('Test') {
            steps {
                echo 'Running Tests!'
                sh 'go mod init helloworld'
                sh 'go test'
            }
        }


        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build -o main'
            }
        }
    }
}

