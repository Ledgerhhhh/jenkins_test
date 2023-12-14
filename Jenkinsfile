def name = "test_jenkins"
def dockerfile = "."
def version = "1.1.1.035"
def dockerImage = "${name}:${version}"
def port = 8080
def remoteMirrorWarehouse="registry.cn-hangzhou.aliyuncs.com/ledger_test_namespace/test_dockerfile"

pipeline {
    agent any

    stages {
        // Stage 1: Checkout Code
        stage('Checkout') {
            steps {
                // 从源代码管理系统 (SCM) 中检出代码
                checkout scm
            }
        }

        // Stage 2: Build Docker Image
        stage('Building') {
            steps {
                // 使用 Docker 构建镜像，-f 指定 Dockerfile 的路径，-t 指定镜像名称和标签
                sh "docker build -f ${dockerfile} -t ${dockerImage} ."
            }
        }

        // Stage 3: Tag Docker Image
        stage('Tag') {
            steps {
                // 为构建的 Docker 镜像创建新的标签，以便推送到远程仓库
                sh "docker tag ${dockerImage} ${remoteMirrorWarehouse}:${version}"
            }
        }

        // Stage 4: Publish Docker Image to Remote Repository
        stage('Publish') {
            steps {
                // 推送构建的 Docker 镜像到远程镜像仓库
                sh "docker push ${remoteMirrorWarehouse}:${version}"
            }
        }

        // Stage 5: Remove Local Docker Image
        stage('RemoveLocalImage') {
            steps {
                // 删除本地构建的 Docker 镜像，推送到远程仓库后通常不再需要本地镜像
                sh "docker rmi ${remoteMirrorWarehouse}:${version}"
            }
        }
    }
}
