#使用golang镜像编译项目(golang:1.20-alpine这个是带有go环境的很小的基础镜像)
FROM golang:1.20-alpine AS builder

#作者
MAINTAINER ledger

#指定docker中的工作目录
WORKDIR /go/src

#设置代理
ENV GOPROXY https://goproxy.cn,direct

#把本地项目拷贝到wordir中
ADD . /go/src

#执行编译
# CGO_ENABLED=0: 这个环境变量用于禁用 CGO（C语言调用Go的功能）
# GOOS=linux: 这个环境变量设置目标操作系统为 Linux
# GOARCH=amd64: 这个环境变量设置目标架构为 AMD64
# go build: 这是用于构建 Go 代码的命令,会自动下载依赖包
# -ldflags="-s -w": 这个参数用于将二进制文件中的符号表和调试信息去除
# -installsuffix cgo: 这个标志用于在构建时添加安装后缀，它通常用于区分使用 CGO 和不使用 CGO 两种情况
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo main.go

#拉取alpine镜像（空间小）
FROM alpine AS prod

#将镜像源设置为阿里云。
#更新 apk 软件包索引。
#安装 tzdata 和 vim 软件包。
#设置时区为 Asia/Shanghai。
#清理 apk 缓存。
RUN echo -e 'https://mirrors.aliyun.com/alpine/v3.6/main/\nhttps://mirrors.aliyun.com/alpine/v3.6/community/' > /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata vim \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/cache/apk/*

EXPOSE 8080

# 拷贝主程序
COPY --from=builder /go/src/main /main
# 拷贝配置文件
COPY --from=builder /go/src/config.yml ./config.yml

ENTRYPOINT ["/main"]
