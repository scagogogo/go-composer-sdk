---
layout: home

hero:
  name: "Go Composer SDK"
  text: "Go 语言的 PHP Composer"
  tagline: 全面的 PHP Composer 包管理器 Go 语言库
  image:
    src: /logo.svg
    alt: Go Composer SDK
  actions:
    - theme: brand
      text: 快速开始
      link: /zh/guide/getting-started
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/scagogogo/go-composer-sdk

features:
  - icon: 🚀
    title: 完整的 Composer 支持
    details: 全面支持所有标准 Composer CLI 命令，提供类型安全的 Go API
  - icon: 🛡️
    title: 类型安全
    details: 强类型接口，支持 IDE 代码补全和编译时错误检查
  - icon: 🔧
    title: 功能全面
    details: 依赖管理、仓库配置、身份验证、安全审计等完整功能
  - icon: 🌍
    title: 跨平台支持
    details: 原生支持 Windows、macOS 和 Linux，针对不同平台进行优化
  - icon: 📦
    title: 模块化设计
    details: 代码结构清晰，按功能分组，易于使用和维护
  - icon: ✅
    title: 生产就绪
    details: 通过 GitHub Actions CI/CD 全面测试，确保代码质量和可靠性
---

## 快速开始

安装 Go Composer SDK：

```bash
go get github.com/scagogogo/go-composer-sdk
```

创建 Composer 实例并开始管理 PHP 依赖：

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

## 为什么选择 Go Composer SDK？

- **🎯 专门构建**: 专为需要管理 PHP 项目的 Go 应用程序设计
- **📚 文档完善**: 每个功能都有详细的文档和示例
- **🔒 安全可靠**: 内置安全审计和漏洞检测功能
- **⚡ 高性能**: 高效执行，支持超时处理和上下文控制
- **🧪 充分测试**: 拥有 161+ 个测试用例，确保可靠性

## 功能特性

- **核心 Composer 操作**: 安装、更新、添加、删除包
- **项目管理**: 创建项目、运行脚本、验证配置
- **安全功能**: 审计依赖、检测漏洞
- **平台工具**: 检查 PHP 版本、扩展、平台要求
- **实用工具**: 文件系统操作、HTTP 下载、跨平台支持
- **自动检测**: 自动检测并安装 Composer（如需要）

## 社区

- [GitHub 仓库](https://github.com/scagogogo/go-composer-sdk)
- [问题跟踪](https://github.com/scagogogo/go-composer-sdk/issues)
- [讨论区](https://github.com/scagogogo/go-composer-sdk/discussions)

## 许可证

Go Composer SDK 基于 [MIT 许可证](https://github.com/scagogogo/go-composer-sdk/blob/main/LICENSE) 发布。
