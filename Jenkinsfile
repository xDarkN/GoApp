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
                sh 'go install golang.org/x/lint/golint@latest' // Install golint
            }
        }
        stage('Initialize Module') {
            steps {
                echo 'Initializing Go module'
                sh 'go mod init main'
            }
        }
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                echo 'Running vetting'
                sh 'go vet .'
                echo 'Running linting'
                sh '/home/jenkins/go/bin/golint .' // Use the full path to golint binary
            }
        }
        stage('RUN') {
            steps {
            	echo 'Running Go app'
            	sh 'go run main.go'
            }
	}
    }
}

