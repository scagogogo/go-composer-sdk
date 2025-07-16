#!/bin/bash

# 切换到Vagrantfile所在目录
cd "$(dirname "$0")" || exit 1

# 检查Vagrant是否安装
if ! command -v vagrant &> /dev/null; then
    echo "错误: 未找到Vagrant命令。请先安装Vagrant: https://www.vagrantup.com/downloads"
    exit 1
fi

# 检查VirtualBox是否安装
if ! command -v VBoxManage &> /dev/null; then
    echo "错误: 未找到VirtualBox。请先安装VirtualBox: https://www.virtualbox.org/wiki/Downloads"
    exit 1
fi

echo "===== 开始在Vagrant环境中测试Linux安装 ====="

# 启动虚拟机（如果未启动）
vagrant status | grep -q "running"
if [ $? -ne 0 ]; then
    echo "正在启动Vagrant虚拟机..."
    vagrant up
    if [ $? -ne 0 ]; then
        echo "启动Vagrant虚拟机失败"
        exit 1
    fi
else
    echo "Vagrant虚拟机已运行"
fi

echo "===== 在虚拟机中运行测试 ====="

# 在VM中运行测试
vagrant ssh -c "cd /home/vagrant/go/src/github.com/scagogogo/go-composer-sdk/tests/docker && go run main.go"

# 显示测试结果
if [ $? -eq 0 ]; then
    echo "测试在Linux环境中成功完成!"
else
    echo "测试在Linux环境中失败"
    exit 1
fi

# 提示是否关闭虚拟机
read -p "测试完成。是否关闭虚拟机？(y/n): " answer
if [[ "$answer" == "y" || "$answer" == "Y" ]]; then
    echo "正在关闭虚拟机..."
    vagrant halt
    echo "虚拟机已关闭"
fi

echo "===== 测试完成 =====" 