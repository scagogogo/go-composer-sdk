# Linux环境下的Composer自动安装测试

本目录包含用于在Linux环境中测试Composer安装功能的测试代码。

## 测试内容

测试程序会验证以下功能：

1. 自动安装Composer到系统路径
2. 自定义安装目录安装Composer
3. 创建简单PHP项目并安装依赖

## 使用Docker测试

如果您的系统上已安装Docker，可以使用Docker容器来运行测试：

```bash
# 在项目根目录下执行
cd tests/docker
./run_tests.sh
```

脚本会构建一个包含所有必要环境的Docker镜像，并在容器中运行测试。

## 使用Vagrant测试

如果没有Docker，但安装了Vagrant和VirtualBox，可以使用Vagrant创建虚拟机来运行测试：

```bash
# 在项目根目录下执行
cd tests/vagrant
./run_test.sh
```

脚本会使用Vagrant创建一个Ubuntu 22.04虚拟机，并自动安装所有必要的依赖。

## 手动测试

也可以手动测试：

1. 在任意Linux环境中（如Ubuntu 22.04），安装PHP：
   ```bash
   sudo add-apt-repository ppa:ondrej/php
   sudo apt-get update
   sudo apt-get install -y php8.1 php8.1-cli php8.1-common php8.1-curl php8.1-mbstring php8.1-xml php8.1-zip
   ```

2. 安装Go：
   ```bash
   wget -q https://golang.org/dl/go1.20.3.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   ```

3. 编译并运行测试程序：
   ```bash
   cd /path/to/go-composer-sdk
   cd tests/docker
   go build -o test_linux_install
   ./test_linux_install
   ```

## 测试输出

成功的测试会输出类似下面的结果：

```
===== 开始在Linux环境中测试Composer自动安装 =====
已创建临时目录: /tmp/composer-linux-test-123456
PHP路径: /usr/bin/php
PHP版本: PHP 8.1.x (cli)...

===== 测试1: 使用SDK自动安装Composer =====
自动安装成功! Composer版本: 2.x.x
Composer安装路径: /usr/local/bin/composer

===== 测试2: 使用自定义安装目录 =====
开始自定义安装...
自定义安装成功! 文件路径:
- composer.phar: /tmp/composer-linux-test-123456/custom-composer/composer.phar
- composer脚本: /tmp/composer-linux-test-123456/custom-composer/composer
自定义Composer版本: 2.x.x

===== 测试3: 创建简单PHP项目 =====
安装项目依赖...
项目依赖安装成功!

===== 所有测试通过! ===== 