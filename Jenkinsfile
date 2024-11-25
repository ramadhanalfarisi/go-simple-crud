pipeline {
    agent any
    stages {
        stage("Setup Library") {
            steps {
                sh "go mod tidy"
            }
        }
        stage("Build") {
            steps {
                sh "go build -o main"
            }
        }
    }
}