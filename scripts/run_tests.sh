#!/bin/bash

# 设置颜色输出
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 打印标题
echo -e "${YELLOW}开始运行Composer安装器测试...${NC}"
echo "=================================="

# 进入项目根目录
cd "$(dirname "$0")/.."

# 创建utils/mock目录（如果不存在）
mkdir -p pkg/installer/utils/mock

# 执行所有测试
echo -e "${YELLOW}执行单元测试...${NC}"
go test -v ./pkg/installer/...

# 检查测试状态
TEST_STATUS=$?
if [ $TEST_STATUS -eq 0 ]; then
    echo -e "${GREEN}全部测试通过!${NC}"
else 
    echo -e "${RED}测试失败!${NC}"
fi

echo "=================================="
echo -e "${YELLOW}测试完成.${NC}"

exit $TEST_STATUS 