# 基本操作示例

本页面展示 Go Composer SDK 的基本操作示例，包括创建实例、检查安装、获取版本等核心功能。

## 创建 Composer 实例

### 使用默认配置

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
    
    fmt.Println("✅ Composer 实例创建成功！")
}
```

### 自定义配置

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // 自定义配置选项
    options := composer.Options{
        ExecutablePath:  "/usr/local/bin/composer",
        WorkingDir:      "/path/to/your/php/project",
        AutoInstall:     true,
        DefaultTimeout:  5 * time.Minute,
    }
    
    comp, err := composer.New(options)
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    fmt.Println("✅ 自定义 Composer 实例创建成功！")
}
```

## 检查 Composer 安装

### 基本检查

```go
func checkComposerInstallation() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 检查 Composer 是否已安装
    if comp.IsInstalled() {
        fmt.Println("✅ Composer 已安装")
        
        // 获取版本信息
        version, err := comp.GetVersion()
        if err != nil {
            log.Printf("获取版本失败: %v", err)
        } else {
            fmt.Printf("📦 Composer 版本: %s\n", version)
        }
    } else {
        fmt.Println("❌ Composer 未安装")
        fmt.Println("💡 提示: 启用 AutoInstall 选项可以自动安装 Composer")
    }
}
```

### 详细检查

```go
func detailedComposerCheck() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    fmt.Println("🔍 正在检查 Composer 安装状态...")
    
    // 检查安装状态
    if !comp.IsInstalled() {
        fmt.Println("❌ Composer 未安装")
        return
    }
    
    fmt.Println("✅ Composer 已安装")
    
    // 获取版本
    if version, err := comp.GetVersion(); err == nil {
        fmt.Printf("📦 版本: %s\n", version)
    }
    
    // 获取可执行文件路径
    execPath := comp.GetExecutablePath()
    fmt.Printf("📁 可执行文件路径: %s\n", execPath)
    
    // 获取工作目录
    workDir := comp.GetWorkingDir()
    fmt.Printf("📂 工作目录: %s\n", workDir)
    
    // 运行诊断
    fmt.Println("\n🔧 运行诊断...")
    if output, err := comp.Diagnose(); err == nil {
        fmt.Printf("诊断结果:\n%s\n", output)
    } else {
        fmt.Printf("诊断失败: %v\n", err)
    }
}
```

## 配置工作目录

### 设置项目目录

```go
func setupWorkingDirectory() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置工作目录
    projectPath := "/path/to/your/php/project"
    comp.SetWorkingDir(projectPath)
    
    fmt.Printf("📂 工作目录已设置为: %s\n", comp.GetWorkingDir())
    
    // 验证目录中是否有 composer.json
    if _, err := os.Stat(filepath.Join(projectPath, "composer.json")); err == nil {
        fmt.Println("✅ 找到 composer.json 文件")
    } else {
        fmt.Println("⚠️  未找到 composer.json 文件")
        fmt.Println("💡 提示: 您可能需要先初始化项目或切换到正确的目录")
    }
}
```

### 动态切换目录

```go
func switchBetweenProjects() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    projects := []string{
        "/path/to/project1",
        "/path/to/project2",
        "/path/to/project3",
    }
    
    for _, project := range projects {
        fmt.Printf("\n🔄 切换到项目: %s\n", project)
        comp.SetWorkingDir(project)
        
        // 检查项目状态
        if _, err := os.Stat(filepath.Join(project, "composer.json")); err == nil {
            fmt.Println("✅ 有效的 Composer 项目")
            
            // 可以在这里执行项目特定的操作
            if err := comp.Validate(); err == nil {
                fmt.Println("✅ composer.json 验证通过")
            } else {
                fmt.Printf("❌ composer.json 验证失败: %v\n", err)
            }
        } else {
            fmt.Println("❌ 不是有效的 Composer 项目")
        }
    }
}
```

## 环境变量配置

### 基本环境配置

```go
func configureEnvironment() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置环境变量
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",           // 移除内存限制
        "COMPOSER_PROCESS_TIMEOUT=600",       // 10分钟超时
        "COMPOSER_CACHE_DIR=/tmp/composer",   // 自定义缓存目录
        "COMPOSER_HOME=/opt/composer",        // 自定义主目录
    })
    
    fmt.Println("✅ 环境变量配置完成")
    
    // 测试配置
    if version, err := comp.GetVersion(); err == nil {
        fmt.Printf("📦 Composer 版本: %s\n", version)
    }
}
```

### 开发环境配置

```go
func setupDevelopmentEnvironment() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 开发环境特定配置
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",       // 开发环境较短超时
        "COMPOSER_DISCARD_CHANGES=true",      // 自动丢弃更改
        "COMPOSER_PREFER_STABLE=false",       // 允许开发版本
        "COMPOSER_MINIMUM_STABILITY=dev",
    })
    
    fmt.Println("🛠️  开发环境配置完成")
    
    // 设置开发项目目录
    comp.SetWorkingDir("./my-dev-project")
    
    fmt.Println("✅ 开发环境就绪")
}
```

## 超时和上下文管理

### 设置超时

```go
func timeoutExample() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置默认超时
    comp.SetTimeout(2 * time.Minute)
    
    fmt.Println("⏰ 超时设置为 2 分钟")
    
    // 执行可能耗时的操作
    start := time.Now()
    if version, err := comp.GetVersion(); err == nil {
        duration := time.Since(start)
        fmt.Printf("✅ 获取版本成功: %s (耗时: %v)\n", version, duration)
    } else {
        fmt.Printf("❌ 操作失败: %v\n", err)
    }
}
```

### 使用上下文控制

```go
func contextExample() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 创建带超时的上下文
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    fmt.Println("🔄 使用上下文执行命令...")
    
    // 使用上下文执行命令
    output, err := comp.RunWithContext(ctx, "--version")
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            fmt.Println("⏰ 操作超时")
        } else {
            fmt.Printf("❌ 操作失败: %v\n", err)
        }
        return
    }
    
    fmt.Printf("✅ 命令执行成功:\n%s\n", output)
}
```

## 错误处理最佳实践

### 基本错误处理

```go
func basicErrorHandling() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置一个可能不存在的工作目录
    comp.SetWorkingDir("/nonexistent/directory")
    
    // 尝试执行操作并处理错误
    if version, err := comp.GetVersion(); err != nil {
        fmt.Printf("❌ 获取版本失败: %v\n", err)
        
        // 根据错误类型采取不同的处理策略
        if strings.Contains(err.Error(), "not found") {
            fmt.Println("💡 建议: 检查 Composer 是否已正确安装")
        } else if strings.Contains(err.Error(), "permission") {
            fmt.Println("💡 建议: 检查文件权限")
        } else {
            fmt.Println("💡 建议: 检查网络连接和配置")
        }
    } else {
        fmt.Printf("✅ Composer 版本: %s\n", version)
    }
}
```

### 健壮的错误处理

```go
func robustErrorHandling() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 多步骤操作的错误处理
    steps := []struct {
        name string
        fn   func() error
    }{
        {"检查安装", func() error {
            if !comp.IsInstalled() {
                return fmt.Errorf("Composer 未安装")
            }
            return nil
        }},
        {"获取版本", func() error {
            _, err := comp.GetVersion()
            return err
        }},
        {"验证配置", func() error {
            return comp.Validate()
        }},
    }
    
    for _, step := range steps {
        fmt.Printf("🔄 执行步骤: %s\n", step.name)
        
        if err := step.fn(); err != nil {
            fmt.Printf("❌ 步骤 '%s' 失败: %v\n", step.name, err)
            
            // 决定是否继续执行后续步骤
            if step.name == "检查安装" {
                fmt.Println("🛑 关键步骤失败，停止执行")
                return
            }
            
            fmt.Println("⚠️  非关键步骤失败，继续执行")
            continue
        }
        
        fmt.Printf("✅ 步骤 '%s' 成功\n", step.name)
    }
    
    fmt.Println("🎉 所有步骤执行完成")
}
```

## 完整示例

### 综合基本操作示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    fmt.Println("🚀 Go Composer SDK 基本操作示例")
    fmt.Println("=====================================")
    
    // 1. 创建实例
    fmt.Println("\n1️⃣  创建 Composer 实例")
    comp, err := createComposerInstance()
    if err != nil {
        log.Fatalf("创建实例失败: %v", err)
    }
    
    // 2. 检查安装
    fmt.Println("\n2️⃣  检查 Composer 安装")
    checkInstallation(comp)
    
    // 3. 配置环境
    fmt.Println("\n3️⃣  配置环境")
    configureComposer(comp)
    
    // 4. 执行基本命令
    fmt.Println("\n4️⃣  执行基本命令")
    executeBasicCommands(comp)
    
    fmt.Println("\n🎉 示例执行完成！")
}

func createComposerInstance() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 2 * time.Minute
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    fmt.Println("✅ Composer 实例创建成功")
    return comp, nil
}

func checkInstallation(comp *composer.Composer) {
    if !comp.IsInstalled() {
        fmt.Println("❌ Composer 未安装")
        return
    }
    
    fmt.Println("✅ Composer 已安装")
    
    if version, err := comp.GetVersion(); err == nil {
        fmt.Printf("📦 版本: %s\n", version)
    }
}

func configureComposer(comp *composer.Composer) {
    // 设置环境变量
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",
    })
    
    // 设置工作目录（如果存在）
    if wd, err := os.Getwd(); err == nil {
        comp.SetWorkingDir(wd)
        fmt.Printf("📂 工作目录: %s\n", comp.GetWorkingDir())
    }
    
    fmt.Println("✅ 环境配置完成")
}

func executeBasicCommands(comp *composer.Composer) {
    // 使用上下文执行命令
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    commands := []string{"--version", "--help"}
    
    for _, cmd := range commands {
        fmt.Printf("🔄 执行命令: composer %s\n", cmd)
        
        if output, err := comp.RunWithContext(ctx, cmd); err == nil {
            // 只显示输出的前几行
            lines := strings.Split(output, "\n")
            if len(lines) > 3 {
                lines = lines[:3]
                lines = append(lines, "...")
            }
            fmt.Printf("📄 输出: %s\n", strings.Join(lines, " | "))
        } else {
            fmt.Printf("❌ 命令失败: %v\n", err)
        }
    }
}
```

这个示例展示了 Go Composer SDK 的所有基本操作，包括实例创建、安装检查、环境配置、超时处理和错误处理等核心功能。
