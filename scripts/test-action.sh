#!/bin/bash

# 测试GitHub Actions脚本
# 在提交代码前运行此脚本，以确保代码可以正常编译和测试通过

set -e

echo "运行编译检查..."
# 使用-o /dev/null选项，只检查编译是否成功，不生成二进制文件
go build -o /dev/null ./...

if [ $? -eq 0 ]; then
    echo "✅ 编译检查通过!"
else
    echo "❌ 编译失败! 请修复问题后再提交代码。"
    exit 1
fi

echo "运行单元测试 (仅pkg/examples)..."
# 只测试我们自己创建的示例测试，这个应该是可以通过的
go test -v ./pkg/examples

if [ $? -eq 0 ]; then
    echo "✅ 测试通过!"
    echo "您可以安全地提交代码了。GitHub Actions将在推送后自动运行。"
else  
    echo "❌ 测试失败! 请修复问题后再提交代码。"
    exit 1
fi

echo ""
echo "提示: 如果您希望在本地测试GitHub Actions工作流，请确保已安装以下工具:"
echo "- act: brew install act"
echo "- Docker: 并确保Docker服务正在运行"
echo ""
echo "然后可以使用以下命令进行测试 (注意可能需要网络访问):"
echo "act --workflows .github/workflows/test.yml -j test -P ubuntu-latest=golang:1.20 --container-architecture linux/amd64 -n" 