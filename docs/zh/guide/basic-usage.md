# 基本用法

本指南涵盖 Go Composer SDK 的基本使用模式和常见工作流程。

## 快速开始

### 创建 Composer 实例

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // 使用默认选项创建
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置您的项目目录
    comp.SetWorkingDir("/path/to/your/php/project")
    
    fmt.Println("Composer SDK 已准备就绪！")
}
```

### 基本操作

```go
// 检查 Composer 是否已安装
if !comp.IsInstalled() {
    log.Fatal("Composer 未安装")
}

// 获取版本
version, err := comp.GetVersion()
if err != nil {
    log.Printf("获取版本失败: %v", err)
} else {
    fmt.Printf("Composer 版本: %s\n", version)
}

// 安装依赖
err = comp.Install(false, false) // noDev=false, optimize=false
if err != nil {
    log.Printf("安装失败: %v", err)
}
```

## 常见工作流程

### 项目设置工作流程

```go
func setupNewProject(projectPath string) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir(projectPath)
    
    // 1. 如果需要，初始化项目
    if _, err := os.Stat(filepath.Join(projectPath, "composer.json")); os.IsNotExist(err) {
        initOptions := composer.InitOptions{
            Name:        "mycompany/my-project",
            Description: "我的 PHP 项目",
            Type:        "project",
            License:     "MIT",
        }
        
        if err := comp.InitProject(initOptions); err != nil {
            return fmt.Errorf("初始化项目失败: %w", err)
        }
    }
    
    // 2. 添加基础包
    packages := map[string]string{
        "symfony/console": "^6.0",
        "monolog/monolog": "^3.0",
    }
    
    if err := comp.RequirePackages(packages); err != nil {
        return fmt.Errorf("添加包失败: %w", err)
    }
    
    // 3. 安装依赖
    if err := comp.Install(false, true); err != nil { // optimize=true
        return fmt.Errorf("安装依赖失败: %w", err)
    }
    
    fmt.Println("✅ 项目设置完成！")
    return nil
}
```

### 维护工作流程

```go
func maintainProject(projectPath string) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir(projectPath)
    
    // 1. 验证 composer.json
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("验证失败: %w", err)
    }
    
    // 2. 检查过时的包
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("警告: 检查过时包失败: %v", err)
    } else if outdated != "" {
        fmt.Println("📦 发现过时的包:")
        fmt.Println(outdated)
        
        // 可选择更新
        fmt.Println("正在更新包...")
        if err := comp.Update(false, true); err != nil {
            log.Printf("更新失败: %v", err)
        }
    }
    
    // 3. 安全审计
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("安全审计失败: %v", err)
    } else {
        fmt.Println("🔒 安全审计完成")
        if auditResult != "" {
            fmt.Println(auditResult)
        }
    }
    
    // 4. 清理
    if err := comp.ClearCache(); err != nil {
        log.Printf("缓存清理失败: %v", err)
    }
    
    return nil
}
```

## 处理依赖

### 添加依赖

```go
// 添加单个包
err := comp.RequirePackage("guzzlehttp/guzzle", "^7.0")

// 添加多个包
packages := map[string]string{
    "symfony/http-foundation": "^6.0",
    "doctrine/orm":           "^2.14",
}
err = comp.RequirePackages(packages)

// 添加开发依赖
err = comp.RequireDevPackage("phpunit/phpunit", "^10.0")
```

### 更新依赖

```go
// 更新所有包
err := comp.Update(false, false)

// 更新特定包
err = comp.UpdatePackage("symfony/console")

// 更新多个特定包
packages := []string{"symfony/console", "monolog/monolog"}
err = comp.UpdatePackages(packages)
```

### 删除依赖

```go
// 删除包
err := comp.RemovePackage("old-package/deprecated")

// 删除多个包
packages := []string{"package1", "package2"}
err = comp.RemovePackages(packages)
```

## 信息和分析

### 包信息

```go
// 列出所有已安装的包
packages, err := comp.ShowAllPackages()
if err == nil {
    fmt.Println("已安装的包:")
    fmt.Println(packages)
}

// 显示特定包详情
details, err := comp.ShowPackage("symfony/console")
if err == nil {
    fmt.Printf("包详情:\n%s\n", details)
}

// 显示依赖树
tree, err := comp.ShowDependencyTree("")
if err == nil {
    fmt.Printf("依赖树:\n%s\n", tree)
}
```

### 依赖分析

```go
// 为什么安装了某个包？
reasons, err := comp.WhyPackage("psr/log")
if err == nil {
    fmt.Printf("为什么安装 psr/log:\n%s\n", reasons)
}

// 为什么不能安装某个包？
conflicts, err := comp.WhyNotPackage("symfony/console", "^7.0")
if err == nil {
    fmt.Printf("symfony/console ^7.0 的冲突:\n%s\n", conflicts)
}
```

## 错误处理模式

### 基本错误处理

```go
func handleComposerOperation() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return fmt.Errorf("创建 composer 实例失败: %w", err)
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // 始终检查 Composer 是否可用
    if !comp.IsInstalled() {
        return fmt.Errorf("composer 未安装，请先安装")
    }
    
    // 执行操作并处理错误
    if err := comp.Install(false, false); err != nil {
        return fmt.Errorf("安装失败: %w", err)
    }
    
    return nil
}
```

### 带重试的健壮错误处理

```go
func robustInstall(comp *composer.Composer, maxRetries int) error {
    for attempt := 1; attempt <= maxRetries; attempt++ {
        fmt.Printf("安装尝试 %d/%d...\n", attempt, maxRetries)
        
        err := comp.Install(false, false)
        if err == nil {
            fmt.Println("✅ 安装成功！")
            return nil
        }
        
        fmt.Printf("❌ 尝试 %d 失败: %v\n", attempt, err)
        
        if attempt < maxRetries {
            // 重试前等待
            time.Sleep(time.Duration(attempt) * time.Second)
            
            // 重试前清除缓存
            if clearErr := comp.ClearCache(); clearErr != nil {
                log.Printf("清除缓存失败: %v", clearErr)
            }
        }
    }
    
    return fmt.Errorf("经过 %d 次尝试后安装失败", maxRetries)
}
```

## 上下文和超时

### 使用上下文进行取消

```go
func installWithCancellation() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // 创建可取消的上下文
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // 在 goroutine 中开始安装
    errChan := make(chan error, 1)
    go func() {
        errChan <- comp.InstallWithContext(ctx, false, false)
    }()
    
    // 等待完成或用户取消
    select {
    case err := <-errChan:
        if err != nil {
            return fmt.Errorf("安装失败: %w", err)
        }
        fmt.Println("✅ 安装完成！")
        return nil
        
    case <-time.After(30 * time.Second):
        // 用户决定在 30 秒后取消
        cancel()
        return fmt.Errorf("用户取消了安装")
    }
}
```

### 超时处理

```go
func installWithTimeout(timeoutMinutes int) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // 创建带超时的上下文
    ctx, cancel := context.WithTimeout(
        context.Background(),
        time.Duration(timeoutMinutes)*time.Minute,
    )
    defer cancel()
    
    err = comp.InstallWithContext(ctx, false, false)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return fmt.Errorf("安装在 %d 分钟后超时", timeoutMinutes)
        }
        return fmt.Errorf("安装失败: %w", err)
    }
    
    return nil
}
```

## 环境配置

### 开发环境

```go
func setupDevelopmentEnvironment(comp *composer.Composer) {
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",
        "COMPOSER_DISCARD_CHANGES=true",
        "COMPOSER_PREFER_STABLE=false",
    })
}
```

### 生产环境

```go
func setupProductionEnvironment(comp *composer.Composer) {
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",
        "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
        "COMPOSER_PREFER_STABLE=true",
    })
}
```

## 最佳实践

### 1. 操作前始终验证

```go
func safeComposerOperation(comp *composer.Composer) error {
    // 检查 Composer 可用性
    if !comp.IsInstalled() {
        return fmt.Errorf("composer 不可用")
    }
    
    // 验证项目
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("项目验证失败: %w", err)
    }
    
    // 继续操作
    return comp.Install(false, false)
}
```

### 2. 使用适当的超时

```go
func operationWithAppropriateTimeout(comp *composer.Composer, operation string) error {
    var timeout time.Duration
    
    switch operation {
    case "install", "update":
        timeout = 10 * time.Minute
    case "require", "remove":
        timeout = 5 * time.Minute
    default:
        timeout = 2 * time.Minute
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    return comp.InstallWithContext(ctx, false, false)
}
```

### 3. 处理不同环境

```go
func createComposerForEnvironment() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // 根据环境调整
    if os.Getenv("CI") == "true" {
        options.DefaultTimeout = 15 * time.Minute
    } else if os.Getenv("APP_ENV") == "production" {
        options.AutoInstall = false
    }
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // 设置环境特定变量
    if os.Getenv("CI") == "true" {
        comp.SetEnv([]string{
            "COMPOSER_NO_INTERACTION=1",
            "COMPOSER_PREFER_STABLE=true",
        })
    }
    
    return comp, nil
}
```

### 4. 日志记录和监控

```go
func monitoredOperation(comp *composer.Composer) error {
    start := time.Now()
    
    log.Printf("在以下位置开始 Composer 操作: %s", comp.GetWorkingDir())
    
    err := comp.Install(false, false)
    
    duration := time.Since(start)
    if err != nil {
        log.Printf("❌ 操作在 %v 后失败: %v", duration, err)
        return err
    }
    
    log.Printf("✅ 操作在 %v 内成功完成", duration)
    return nil
}
```

这涵盖了大多数 Composer 操作所需的基本使用模式。关键是始终适当处理错误，对长时间运行的操作使用超时，并根据您的特定需求配置环境。
