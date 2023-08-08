pipeline {
    agent {
        label 'my-go-agent'
    }
    tools {
        go '1.20.7' // Make sure this matches the configured tool name
    }
    stages {        
        stage('Example') {
         steps {
         	sh 'go version'
    }  
}
}
}

