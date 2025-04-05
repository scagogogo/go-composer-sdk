#!/bin/bash

# 测试GitHub Actions脚本
# 在提交代码前运行此脚本，以确保GitHub Actions工作流可以正常运行

set -e

# 检查act是否已安装
if ! command -v act &> /dev/null; then
    echo "错误: act工具未安装。请通过以下命令安装:"
    echo "  macOS: brew install act"
    echo "  其他系统: 请访问 https://github.com/nektos/act#installation"
    exit 1
fi

# 检查docker是否运行
if ! docker info &> /dev/null; then
    echo "错误: Docker服务未运行，请启动Docker服务后再试。"
    exit 1
fi

echo "运行单元测试..."
go test -v ./...

echo "运行GitHub Actions工作流..."
act -j test -P ubuntu-latest=golang:1.20 --container-architecture linux/amd64

if [ $? -eq 0 ]; then
    echo "✅ 测试通过! 现在可以安全地提交代码了。"
else
    echo "❌ 测试失败! 请修复问题后再提交代码。"
    exit 1
fi 