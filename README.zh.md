# Go Composer SDK

[![Go Version](https://img.shields.io/github/go-mod/go-version/scagogogo/go-composer-sdk)](https://golang.org/)
[![License](https://img.shields.io/github/license/scagogogo/go-composer-sdk)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/go-composer-sdk)](https://goreportcard.com/report/github.com/scagogogo/go-composer-sdk)
[![Tests](https://github.com/scagogogo/go-composer-sdk/actions/workflows/test.yml/badge.svg)](https://github.com/scagogogo/go-composer-sdk/actions/workflows/test.yml)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://scagogogo.github.io/go-composer-sdk/)

全面的 PHP Composer 包管理器 Go 语言库。此 SDK 提供对 Composer 功能的完整封装，允许您直接从 Go 应用程序管理 PHP 项目依赖、执行 Composer 命令以及处理各种 Composer 相关操作。

## 📖 文档

**完整文档请访问：[https://scagogogo.github.io/go-composer-sdk/](https://scagogogo.github.io/go-composer-sdk/)**

- 📚 [API 参考](https://scagogogo.github.io/go-composer-sdk/zh/api/)
- 🚀 [快速开始指南](https://scagogogo.github.io/go-composer-sdk/zh/guide/getting-started)
- 💡 [示例](https://scagogogo.github.io/go-composer-sdk/zh/examples/)
- 🌍 [English Documentation](https://scagogogo.github.io/go-composer-sdk/)

## ✨ 特性

- **🚀 完整的 Composer 支持**：全面支持所有标准 Composer CLI 命令
- **🛡️ 类型安全**：强类型接口，支持 IDE 代码补全
- **🔧 功能全面**：依赖管理、仓库配置、身份验证、安全审计
- **🌍 跨平台支持**：原生支持 Windows、macOS 和 Linux
- **📦 模块化设计**：按功能分组的代码结构，易于使用和维护
- **✅ 生产就绪**：通过 161+ 个测试和 GitHub Actions CI/CD 全面测试
- **🔒 安全功能**：内置安全审计和漏洞检测
- **⚡ 高性能**：高效执行，支持超时处理和上下文控制

## 🚀 快速开始

### 安装

```bash
go get github.com/scagogogo/go-composer-sdk
```

### 基本用法

```go
package main

import (
    "fmt"
    "log"

    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // 使用默认选项创建 Composer 实例
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }

    // 设置工作目录到您的 PHP 项目
    comp.SetWorkingDir("/path/to/your/php/project")

    // 检查 Composer 是否已安装
    if !comp.IsInstalled() {
        log.Fatal("Composer 未安装")
    }

    // 获取 Composer 版本
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("获取 Composer 版本失败: %v", err)
    }
    fmt.Printf("Composer 版本: %s\n", version)

    // 安装依赖
    err = comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        log.Fatalf("安装依赖失败: %v", err)
    }

    fmt.Println("依赖安装成功！")
}
```

## 📋 系统要求

- **Go 1.21 或更高版本**
- **PHP 7.4 或更高版本**（Composer 运行所需）
- **Composer 2.0 或更高版本**（SDK 可以自动安装）

## 🔧 核心功能

### 包管理
```go
// 安装依赖
err := comp.Install(false, false)

// 添加包
err = comp.RequirePackage("monolog/monolog", "^3.0")

// 更新包
err = comp.Update(false, false)

// 删除包
err = comp.RemovePackage("old-package/deprecated")
```

### 项目管理
```go
// 创建新项目
err := comp.CreateProject("laravel/laravel", "my-app", "")

// 验证 composer.json
err = comp.Validate()

// 运行脚本
err = comp.RunScript("test")
```

### 安全审计
```go
// 执行安全审计
auditResult, err := comp.Audit()

// 检查平台要求
err = comp.CheckPlatformReqs()
```

### 信息分析
```go
// 显示包信息
info, err := comp.ShowPackage("symfony/console")

// 显示依赖树
tree, err := comp.ShowDependencyTree("")

// 检查过时的包
outdated, err := comp.ShowOutdated()
```

## 🏗️ 架构

SDK 组织为几个包：

- **`composer`** - 包含核心 Composer 功能的主包
- **`detector`** - Composer 安装检测和验证
- **`installer`** - 自动 Composer 安装工具
- **`utils`** - 通用工具和辅助函数

## 🧪 测试

项目包含全面的测试，拥有 161+ 个测试用例，涵盖：

- 所有主要功能的单元测试
- 真实场景的集成测试
- 跨平台兼容性测试
- 错误处理和边界情况

运行测试：
```bash
go test ./...
```

运行竞态条件检测：
```bash
go test -race ./...
```

## 🤝 贡献

我们欢迎贡献！请查看我们的[贡献指南](CONTRIBUTING.md)了解详情。

### 开发设置

1. Fork 仓库
2. 克隆您的 fork：`git clone https://github.com/yourusername/go-composer-sdk.git`
3. 创建功能分支：`git checkout -b feature/amazing-feature`
4. 进行更改并添加测试
5. 运行测试：`go test ./...`
6. 提交更改：`git commit -m 'Add amazing feature'`
7. 推送到分支：`git push origin feature/amazing-feature`
8. 打开 Pull Request

## 📄 许可证

本项目基于 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🌟 支持

- 📖 [文档](https://scagogogo.github.io/go-composer-sdk/)
- 🐛 [问题跟踪](https://github.com/scagogogo/go-composer-sdk/issues)
- 💬 [讨论区](https://github.com/scagogogo/go-composer-sdk/discussions)

## 🙏 致谢

- [Composer](https://getcomposer.org/) - 此 SDK 封装的 PHP 包管理器
- [Go 社区](https://golang.org/community/) - 提供了出色的语言和生态系统
- 所有帮助改进此项目的[贡献者](https://github.com/scagogogo/go-composer-sdk/contributors)

---

**语言版本**：[English](README.md) | [简体中文](README.zh.md)

由 Go Composer SDK 团队用 ❤️ 制作