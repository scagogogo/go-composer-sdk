# Go Composer SDK 示例代码

本目录包含了 Go Composer SDK 的各种使用示例，展示了如何在 Go 应用程序中使用该 SDK 来操作 PHP Composer。

## 示例结构

示例代码按功能分类存放在以下目录中：

1. `01_basic_usage` - 基本使用示例
   - 初始化 Composer 实例
   - 运行基本命令

2. `02_package_management` - 包管理示例
   - 安装和更新依赖包
   - 添加和移除包
   - 查看包信息
   - 搜索包

3. `03_project_management` - 项目管理示例
   - 创建和初始化项目
   - 运行脚本

4. `04_config_management` - 配置管理示例
   - 获取和设置配置
   - 查看所有配置

5. `05_composer_json` - composer.json 文件操作示例
   - 读取和修改 composer.json 文件

## 如何运行示例

每个示例目录中的 Go 文件都是独立的，可以单独查看和学习。示例代码中已包含详细的注释，解释了每个操作的作用和预期输出。

要运行这些示例，请先确保：

1. 已安装 Go 语言环境（推荐 Go 1.17 或更高版本）
2. 已安装 PHP 和 Composer（一些示例可能需要实际的 Composer 环境）
3. 已克隆本仓库并执行了 `go mod tidy` 安装依赖

然后可以通过以下方式运行示例：

```bash
# 进入项目根目录
cd go-composer-sdk

# 运行特定示例
go run examples/01_basic_usage/01_new_composer.go
```

## 注意事项

- 部分示例代码中包含了注释掉的实际执行代码，这是为了避免在查看示例时意外执行修改操作
- 示例中的路径和项目名称是虚拟的，使用时请根据实际情况进行修改
- 一些高级功能可能需要特定版本的 Composer 或额外的配置 