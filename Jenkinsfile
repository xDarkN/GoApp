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
        stage('Run') {
            steps {
                echo 'Running helloworld:'
                sh './main' // Assuming the compiled binary is named 'main'
            }
        }
        stage('Echo mf.txt (Dev Branch Only)') {
            when {
                expression {
                    return env.BRANCH_NAME == 'dev'
                }
            }
            steps {
                echo 'Echoing content of mf.txt:'
                sh 'cat mf.txt'
            }
        }
    }
}
