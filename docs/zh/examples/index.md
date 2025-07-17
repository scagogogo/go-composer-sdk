# 示例

本节提供 Go Composer SDK 的实际示例和用例。每个示例都包含完整的、可运行的代码，演示真实世界的使用模式。

## 快速示例

### 基本设置

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // 创建 Composer 实例
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
    
    fmt.Println("Composer 已准备就绪！")
}
```

### 安装依赖

```go
func installDependencies(comp *composer.Composer) error {
    fmt.Println("正在安装依赖...")
    
    // 安装所有依赖
    err := comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        return fmt.Errorf("安装依赖失败: %w", err)
    }
    
    fmt.Println("依赖安装成功！")
    return nil
}
```

### 添加包

```go
func addPackage(comp *composer.Composer) error {
    packageName := "monolog/monolog"
    version := "^3.0"
    
    fmt.Printf("正在添加包 %s %s...\n", packageName, version)
    
    err := comp.RequirePackage(packageName, version)
    if err != nil {
        return fmt.Errorf("添加包失败: %w", err)
    }
    
    fmt.Printf("包 %s 添加成功！\n", packageName)
    return nil
}
```

## 完整示例

### 1. 项目设置自动化

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    projectDir := "/path/to/new/php/project"
    
    // 创建项目目录
    err := os.MkdirAll(projectDir, 0755)
    if err != nil {
        log.Fatalf("创建项目目录失败: %v", err)
    }
    
    // 创建 Composer 实例
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir(projectDir)
    
    // 初始化新项目
    initOptions := composer.InitOptions{
        Name:        "mycompany/my-project",
        Description: "我的 PHP 项目",
        Type:        "project",
        License:     "MIT",
        Authors: []composer.Author{
            {
                Name:  "张三",
                Email: "zhangsan@example.com",
            },
        },
        MinimumStability: "stable",
        PreferStable:     true,
    }
    
    err = comp.InitProject(initOptions)
    if err != nil {
        log.Fatalf("初始化项目失败: %v", err)
    }
    
    // 添加基础包
    packages := map[string]string{
        "symfony/console":    "^6.0",
        "monolog/monolog":    "^3.0",
        "guzzlehttp/guzzle":  "^7.0",
    }
    
    err = comp.RequirePackages(packages)
    if err != nil {
        log.Fatalf("添加包失败: %v", err)
    }
    
    fmt.Println("✅ 项目设置完成！")
}
```

## 示例分类

### [基本操作](/zh/examples/basic-operations)
涵盖最常见用例的简单示例，如安装包、更新依赖和基本项目管理。

### [包管理](/zh/examples/package-management)
高级包管理场景，包括批量操作、依赖分析和冲突解决。

### [项目设置](/zh/examples/project-setup)
不同类型 PHP 项目的完整项目初始化和设置自动化示例。

### [安全审计](/zh/examples/security-audit)
以安全为重点的示例，展示如何将漏洞扫描和依赖审计集成到您的工作流程中。

### [高级用法](/zh/examples/advanced-usage)
复杂场景，包括自定义配置、错误处理模式和与其他工具的集成。

## 运行示例

1. **安装 SDK**：
   ```bash
   go get github.com/scagogogo/go-composer-sdk
   ```

2. **更新路径**：将 `/path/to/php/project` 替换为您实际的 PHP 项目路径

3. **运行示例**：
   ```bash
   go run example.go
   ```

## 使用示例的提示

- **始终处理错误**：示例展示了正确的错误处理模式
- **使用上下文进行超时**：长时间运行的操作应使用上下文
- **根据您的环境配置**：根据需要调整路径和选项
- **先在开发环境中测试**：在生产环境之前在测试环境中尝试示例
- **检查先决条件**：确保满足 PHP 和 Composer 要求

## 贡献示例

有有用的示例吗？我们很乐意包含它！请：

1. Fork 仓库
2. 将您的示例添加到适当的部分
3. 包含清晰的文档和注释
4. 彻底测试示例
5. 提交 pull request

更多信息，请参阅我们的[贡献指南](https://github.com/scagogogo/go-composer-sdk/blob/main/CONTRIBUTING.md)。
