FROM ubuntu:22.04

# 避免交互式提示
ENV DEBIAN_FRONTEND=noninteractive

# 设置时区
RUN apt-get update && apt-get install -y tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

# 安装基本工具
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    unzip \
    software-properties-common \
    apt-transport-https \
    ca-certificates

# 安装PHP
RUN add-apt-repository ppa:ondrej/php && \
    apt-get update && \
    apt-get install -y php8.1 php8.1-cli php8.1-common php8.1-curl \
    php8.1-mbstring php8.1-xml php8.1-zip php8.1-json php8.1-opcache php8.1-readline

# 安装Go
RUN curl -OL https://golang.org/dl/go1.20.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz && \
    rm go1.20.3.linux-amd64.tar.gz

# 设置Go环境变量
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# 创建工作目录
WORKDIR /app

# 复制Go模块文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制整个项目
COPY . .

# 编译测试程序
RUN cd tests/docker && go build -o /app/tests/docker/test_linux_install

# 设置入口点
ENTRYPOINT ["/app/tests/docker/test_linux_install"] 