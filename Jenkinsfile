pipeline {
    agent {
        label 'my-go-agent'
    }
    tools {
        go '1.20.7' // Make sure this matches the configured tool name
    }
    stages {        
        
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
}

