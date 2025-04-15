#!/bin/bash

# 切换到项目根目录
cd "$(dirname "$0")/../.." || exit 1

echo "===== 开始在Docker容器中测试Linux安装 ====="

# 构建Docker镜像
echo "正在构建Docker镜像..."
docker build -t composer-linux-test -f tests/docker/Dockerfile .

# 如果构建失败，退出
if [ $? -ne 0 ]; then
  echo "Docker镜像构建失败"
  exit 1
fi

echo "Docker镜像构建成功!"

# 运行测试容器
echo "正在启动测试容器..."
docker run --rm composer-linux-test

# 显示测试结果
if [ $? -eq 0 ]; then
  echo "测试在Linux容器中成功完成!"
else
  echo "测试在Linux容器中失败"
  exit 1
fi

echo "===== 测试完成 =====" 