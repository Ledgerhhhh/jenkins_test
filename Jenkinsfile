def name = "test_jenkins"
def dockerfile = "."
def dockerImage = "${name}:1.1.1.035"

pipeline {
    agent any

    stages {
       stage('Checkout') {
			steps {
				checkout scm
			}
		}
		stage('Building'){
			steps{
				sh "docker version"
			}
		}
	}
}