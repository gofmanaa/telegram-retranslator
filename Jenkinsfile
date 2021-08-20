pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'go mod tidy'
                sh 'make test'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'make build'
            }
        }
        stage('Up') {
            steps {
                echo 'Upping....'
                sh 'make up'
            }
        }
    }
}