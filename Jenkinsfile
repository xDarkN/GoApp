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
                sh 'go get -u golang.org/x/lint/golint' // Install golint
            }
        }
        
        stage('Initialize Module') {
            steps {
                echo 'Initializing Go module'
                sh 'go mod init yourmodule'
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
                sh 'golint .'
                echo 'Running test'
                sh 'cd test && go test -v'
            }
        }
    }
}

