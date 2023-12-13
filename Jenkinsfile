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
				sh "docker build  -f ${dockerfile} -t ${dockerImage} ."
				sh "docker tag ${dockerImage} ${url}/${dockerImage}"
			}
		}
		}
	}
}