def name = "test_jenkins"
def dockerfile = "."
def dockerImage = "${name}:1.1.1.035"

pipeline {
    agent any 
    tools {
        'docker' // Use 'docker' instead of 'org.jenkinsci.plugins.docker.commons.tools.DockerTool' '18.09'
    }
    environment {
        DOCKER_CERT_PATH = credentials('id-for-a-docker-cred')
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Building') {
            steps {
                script {
                    sh "docker version"
                    // Add your additional build steps here
                }
            }
        }
    }
}
