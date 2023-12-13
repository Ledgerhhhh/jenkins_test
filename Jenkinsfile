def name = "test_jenkins"
def dockerfile = "."
def dockerImage = "${name}:1.1.1.035"
def port = 8080

pipeline {
    agent any

    stages {
       stage('Checkout') {
			steps {
			    // 从SCM中检出代码
				checkout scm
			}
		}
		stage('Building'){
			steps{
				sh "docker build  -f ${dockerfile} -t ${dockerImage} ."
			}
		}
		stage('Running'){
			steps{
				sh "docker run -d --name ${name} -p ${port}:${port} ${dockerImage}"
			}
		}
	}
}