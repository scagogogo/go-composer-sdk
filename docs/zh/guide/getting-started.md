# 快速开始

欢迎使用 Go Composer SDK！本指南将帮助您快速上手。

## 什么是 Go Composer SDK？

Go Composer SDK 是一个全面的 Go 语言库，提供对 PHP Composer 包管理器的完整封装。它允许您直接从 Go 应用程序中管理 PHP 项目依赖、执行 Composer 命令以及处理各种 Composer 相关功能。

## 前置要求

开始之前，请确保您已安装：

- **Go 1.21 或更高版本**
- **PHP**（Composer 运行所需）
- **Composer**（SDK 可以自动安装）

## 安装

使用 `go get` 安装 Go Composer SDK：

```bash
go get github.com/scagogogo/go-composer-sdk
```

## 第一个程序

让我们创建一个简单的程序来演示基本功能：

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
        fmt.Println("Composer 未安装，但 SDK 可以自动安装！")
        return
    }
    
    // 获取并显示 Composer 版本
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("获取 Composer 版本失败: %v", err)
    }
    
    fmt.Printf("✅ Composer 版本: %s\n", version)
    
    // 设置工作目录到您的 PHP 项目
    comp.SetWorkingDir("/path/to/your/php/project")
    
    // 验证 composer.json 文件
    err = comp.Validate()
    if err != nil {
        fmt.Printf("❌ composer.json 验证失败: %v\n", err)
    } else {
        fmt.Println("✅ composer.json 有效")
    }
    
    // 显示已安装的包
    output, err := comp.ShowAllPackages()
    if err != nil {
        log.Printf("获取包列表失败: %v", err)
    } else {
        fmt.Println("📦 已安装的包:")
        fmt.Println(output)
    }
}
```

## 核心概念

### 1. Composer 实例

`Composer` 结构体是所有操作的主要入口点。您可以使用 `New()` 函数和配置选项来创建它：

```go
// 默认配置
comp, err := composer.New(composer.DefaultOptions())

// 自定义配置
options := composer.Options{
    WorkingDir:     "/path/to/php/project",
    AutoInstall:    true,
    DefaultTimeout: 5 * time.Minute,
}
comp, err := composer.New(options)
```

### 2. 错误处理

所有方法都返回应该正确处理的错误：

```go
err := comp.Install(false, false)
if err != nil {
    // 适当处理错误
    log.Printf("安装失败: %v", err)
    return
}
```

### 3. 工作目录

设置工作目录指向您的 PHP 项目：

```go
comp.SetWorkingDir("/path/to/your/php/project")
```

### 4. 环境变量

使用环境变量配置 Composer 行为：

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
})
```

## 常见操作

### 安装依赖

```go
// 安装所有依赖
err := comp.Install(false, false) // noDev=false, optimize=false

// 不安装开发依赖
err := comp.Install(true, false) // noDev=true, optimize=false

// 安装并优化
err := comp.Install(false, true) // noDev=false, optimize=true
```

### 添加包

```go
// 添加包
err := comp.RequirePackage("monolog/monolog", "^3.0")

// 添加开发依赖
err := comp.RequirePackage("phpunit/phpunit", "^10.0")
```

### 更新依赖

```go
// 更新所有包
err := comp.Update(false, false) // noDev=false, optimize=false

// 更新特定包
err := comp.UpdatePackage("symfony/console")
```

### 获取包信息

```go
// 显示所有包
output, err := comp.ShowAllPackages()

// 显示特定包
output, err := comp.ShowPackage("symfony/console")

// 显示依赖树
output, err := comp.ShowDependencyTree("")
```

## 下一步

现在您已经掌握了基础知识，可以探索更高级的功能：

- [配置指南](/zh/guide/configuration) - 了解高级配置选项
- [API 参考](/zh/api/) - 完整的 API 文档
- [示例](/zh/examples/) - 实际使用示例

## 获取帮助

如果遇到任何问题：

1. 查看 [API 参考](/zh/api/) 获取详细的方法文档
2. 查看 [示例](/zh/examples/) 了解常见用例
3. 在 [GitHub](https://github.com/scagogogo/go-composer-sdk/issues) 上搜索或创建问题
4. 加入 [讨论区](https://github.com/scagogogo/go-composer-sdk/discussions)
