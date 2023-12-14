def name = "test_jenkins"
def dockerfile = "."
def version = "1.1.1.035"
def dockerImage = "${name}:${version}"
def port = 8080
def remoteMirrorWarehouse="registry.cn-hangzhou.aliyuncs.com/ledger_test_namespace/test_dockerfile"

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
				sh "docker build -f ${dockerfile} -t ${dockerImage} ."
			}
		}
		stage('Tag'){
			steps{
				sh "docker tag ${dockerImage} ${remoteMirrorWarehouse}:${version}"
			}
		}
		stage('Publish'){
			steps{
				sh "docker push ${remoteMirrorWarehouse}:${version}"
			}
		}
		stage('RemoveLocalImage'){
			steps{
				sh "docker rmi ${remoteMirrorWarehouse}:${version}"
			}
		}
	}
}