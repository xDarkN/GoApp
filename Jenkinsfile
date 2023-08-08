pipeline {
    agent {
        label 'my-go-agent'
    }
    tools {
        go 'GoLang-1.20.7' // Make sure this matches the configured tool name
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        PATH = "${tool(name: 'GoLang-1.20.7', type: 'org.jenkinsci.plugins.golang.GolangInstallation').getHome()}/bin:${env.PATH}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go mod init' // Initialize a Go module
                sh 'go get -u golang.org/x/lint/golint'
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
                withEnv(["PATH+GO=${WORKSPACE}/bin"]){
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
}

