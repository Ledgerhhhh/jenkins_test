def name = "test_jenkins"
def dockerfile = "."
def dockerImage = "${name}:1.1.1.035"

pipeline {
    agent any 
    tools {
        // a bit ugly because there is no `@Symbol` annotation for the DockerTool
        // see the discussion about this in PR 77 and PR 52:
        // https://github.com/jenkinsci/docker-commons-plugin/pull/77#discussion_r280910822
        // https://github.com/jenkinsci/docker-commons-plugin/pull/52
        'org.jenkinsci.plugins.docker.commons.tools.DockerTool' '18.09'
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
