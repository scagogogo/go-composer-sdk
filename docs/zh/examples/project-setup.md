# 项目设置示例

本页面展示如何使用 Go Composer SDK 进行 PHP 项目的初始化、配置和管理。

## 创建新项目

### 从模板创建项目

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 从 Laravel 模板创建新项目
    template := "laravel/laravel"
    projectName := "my-laravel-app"
    version := "" // 使用最新版本
    
    fmt.Printf("🚀 从模板创建项目: %s\n", template)
    
    err = comp.CreateProject(template, projectName, version)
    if err != nil {
        log.Fatalf("创建项目失败: %v", err)
    }
    
    fmt.Printf("✅ 项目 '%s' 创建成功！\n", projectName)
}
```

### 创建不同类型的项目

```go
func createDifferentProjects() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 定义不同类型的项目模板
    projects := []struct {
        template string
        name     string
        description string
    }{
        {"laravel/laravel", "my-laravel-app", "Laravel Web 应用"},
        {"symfony/skeleton", "my-symfony-app", "Symfony 微框架"},
        {"slim/slim-skeleton", "my-slim-api", "Slim API 应用"},
        {"cakephp/app", "my-cake-app", "CakePHP 应用"},
    }
    
    for _, project := range projects {
        fmt.Printf("🚀 创建 %s: %s\n", project.description, project.name)
        
        err = comp.CreateProject(project.template, project.name, "")
        if err != nil {
            log.Printf("❌ 创建 %s 失败: %v", project.name, err)
            continue
        }
        
        fmt.Printf("✅ %s 创建成功\n", project.name)
    }
}
```

## 初始化现有项目

### 基本项目初始化

```go
func initializeExistingProject() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    // 设置项目目录
    projectPath := "/path/to/existing/project"
    comp.SetWorkingDir(projectPath)
    
    // 初始化项目配置
    initOptions := composer.InitOptions{
        Name:        "mycompany/my-project",
        Description: "我的 PHP 项目",
        Type:        "project",
        License:     "MIT",
        Authors: []composer.Author{
            {
                Name:  "Your Name",
                Email: "your.email@example.com",
            },
        },
    }
    
    fmt.Println("🔧 初始化项目...")
    
    err = comp.InitProject(initOptions)
    if err != nil {
        log.Fatalf("初始化项目失败: %v", err)
    }
    
    fmt.Println("✅ 项目初始化完成！")
}
```

### 高级项目初始化

```go
func advancedProjectInitialization() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    projectPath := "/path/to/advanced/project"
    comp.SetWorkingDir(projectPath)
    
    // 高级初始化选项
    initOptions := composer.InitOptions{
        Name:        "mycompany/advanced-project",
        Description: "高级 PHP 项目",
        Type:        "library",
        License:     "MIT",
        Authors: []composer.Author{
            {
                Name:     "Lead Developer",
                Email:    "lead@example.com",
                Homepage: "https://example.com",
                Role:     "Developer",
            },
            {
                Name:  "Contributor",
                Email: "contributor@example.com",
                Role:  "Contributor",
            },
        },
        Keywords: []string{"php", "library", "utility"},
        Homepage: "https://github.com/mycompany/advanced-project",
        Support: composer.Support{
            Issues: "https://github.com/mycompany/advanced-project/issues",
            Wiki:   "https://github.com/mycompany/advanced-project/wiki",
            Source: "https://github.com/mycompany/advanced-project",
        },
        MinimumStability: "stable",
        PreferStable:     true,
    }
    
    fmt.Println("🔧 高级项目初始化...")
    
    err = comp.InitProject(initOptions)
    if err != nil {
        log.Fatalf("高级初始化失败: %v", err)
    }
    
    fmt.Println("✅ 高级项目初始化完成！")
}
```

## 项目配置管理

### 配置自动加载

```go
func configureAutoloading() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔧 配置自动加载...")
    
    // 重新生成自动加载文件
    err = comp.DumpAutoload(false, false) // optimize=false, classmap=false
    if err != nil {
        log.Fatalf("生成自动加载失败: %v", err)
    }
    
    fmt.Println("✅ 自动加载配置完成")
}
```

### 优化自动加载

```go
func optimizeAutoloading() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("⚡ 优化自动加载...")
    
    // 生成优化的自动加载文件
    err = comp.DumpAutoload(true, true) // optimize=true, classmap=true
    if err != nil {
        log.Fatalf("优化自动加载失败: %v", err)
    }
    
    fmt.Println("✅ 自动加载优化完成")
}
```

## 脚本管理

### 运行项目脚本

```go
func runProjectScripts() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 定义要运行的脚本
    scripts := []string{
        "test",
        "lint",
        "build",
        "deploy",
    }
    
    for _, script := range scripts {
        fmt.Printf("🔄 运行脚本: %s\n", script)
        
        err = comp.RunScript(script)
        if err != nil {
            log.Printf("❌ 脚本 '%s' 运行失败: %v", script, err)
            continue
        }
        
        fmt.Printf("✅ 脚本 '%s' 运行成功\n", script)
    }
}
```

### 条件运行脚本

```go
func conditionalScriptExecution() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 检查脚本是否存在并运行
    scriptsToCheck := []struct {
        name        string
        description string
        required    bool
    }{
        {"test", "运行测试", true},
        {"lint", "代码检查", false},
        {"build", "构建项目", false},
        {"post-install-cmd", "安装后脚本", false},
    }
    
    for _, script := range scriptsToCheck {
        fmt.Printf("🔍 检查脚本: %s (%s)\n", script.name, script.description)
        
        err = comp.RunScript(script.name)
        if err != nil {
            if script.required {
                log.Fatalf("❌ 必需脚本 '%s' 失败: %v", script.name, err)
            } else {
                fmt.Printf("⚠️  可选脚本 '%s' 跳过: %v\n", script.name, err)
            }
            continue
        }
        
        fmt.Printf("✅ 脚本 '%s' 执行成功\n", script.name)
    }
}
```

## 项目验证

### 基本验证

```go
func validateProject() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 验证项目配置...")
    
    err = comp.Validate()
    if err != nil {
        log.Fatalf("项目验证失败: %v", err)
    }
    
    fmt.Println("✅ 项目配置验证通过")
}
```

### 严格验证

```go
func strictValidation() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 严格验证项目...")
    
    // 严格验证模式
    err = comp.ValidateStrict()
    if err != nil {
        log.Fatalf("严格验证失败: %v", err)
    }
    
    fmt.Println("✅ 严格验证通过")
}
```

### 验证时跳过检查

```go
func validateWithoutChecks() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 验证项目（跳过某些检查）...")
    
    // 跳过某些检查的验证
    err = comp.ValidateWithNoCheck()
    if err != nil {
        log.Fatalf("验证失败: %v", err)
    }
    
    fmt.Println("✅ 验证完成（已跳过某些检查）")
}
```

## 完整的项目设置工作流

### 新项目完整设置

```go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    fmt.Println("🚀 完整的项目设置工作流")
    fmt.Println("==========================")
    
    projectName := "my-awesome-project"
    projectPath := filepath.Join("./projects", projectName)
    
    // 1. 设置 Composer
    fmt.Println("\n1️⃣  设置 Composer")
    comp := setupComposer()
    
    // 2. 创建项目目录
    fmt.Println("\n2️⃣  创建项目目录")
    createProjectDirectory(projectPath)
    
    // 3. 初始化项目
    fmt.Println("\n3️⃣  初始化项目")
    initializeProject(comp, projectPath, projectName)
    
    // 4. 安装基础依赖
    fmt.Println("\n4️⃣  安装基础依赖")
    installBaseDependencies(comp)
    
    // 5. 配置项目
    fmt.Println("\n5️⃣  配置项目")
    configureProject(comp)
    
    // 6. 验证项目
    fmt.Println("\n6️⃣  验证项目")
    validateProjectSetup(comp)
    
    fmt.Println("\n🎉 项目设置完成！")
    fmt.Printf("📁 项目位置: %s\n", projectPath)
}

func setupComposer() *composer.Composer {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 5 * time.Minute
    options.AutoInstall = true
    
    comp, err := composer.New(options)
    if err != nil {
        log.Fatalf("设置 Composer 失败: %v", err)
    }
    
    // 设置环境变量
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_PREFER_STABLE=true",
    })
    
    fmt.Println("✅ Composer 设置完成")
    return comp
}

func createProjectDirectory(projectPath string) {
    err := os.MkdirAll(projectPath, 0755)
    if err != nil {
        log.Fatalf("创建项目目录失败: %v", err)
    }
    
    fmt.Printf("✅ 项目目录创建: %s\n", projectPath)
}

func initializeProject(comp *composer.Composer, projectPath, projectName string) {
    comp.SetWorkingDir(projectPath)
    
    initOptions := composer.InitOptions{
        Name:        fmt.Sprintf("mycompany/%s", projectName),
        Description: fmt.Sprintf("%s - 一个很棒的 PHP 项目", projectName),
        Type:        "project",
        License:     "MIT",
        Authors: []composer.Author{
            {
                Name:  "Developer",
                Email: "developer@example.com",
            },
        },
        Keywords:         []string{"php", "project", "awesome"},
        MinimumStability: "stable",
        PreferStable:     true,
    }
    
    err := comp.InitProject(initOptions)
    if err != nil {
        log.Fatalf("初始化项目失败: %v", err)
    }
    
    fmt.Println("✅ 项目初始化完成")
}

func installBaseDependencies(comp *composer.Composer) {
    // 安装基础依赖
    baseDependencies := map[string]string{
        "monolog/monolog":    "^3.0",
        "symfony/console":    "^6.0",
        "guzzlehttp/guzzle": "^7.0",
    }
    
    fmt.Println("📦 安装基础依赖...")
    
    err := comp.RequirePackages(baseDependencies)
    if err != nil {
        log.Fatalf("安装基础依赖失败: %v", err)
    }
    
    // 安装开发依赖
    devDependencies := map[string]string{
        "phpunit/phpunit":        "^10.0",
        "squizlabs/php_codesniffer": "^3.7",
    }
    
    fmt.Println("🛠️  安装开发依赖...")
    
    for pkg, version := range devDependencies {
        err = comp.RequireDevPackage(pkg, version)
        if err != nil {
            log.Printf("❌ 安装开发依赖 %s 失败: %v", pkg, err)
        }
    }
    
    fmt.Println("✅ 依赖安装完成")
}

func configureProject(comp *composer.Composer) {
    // 生成优化的自动加载
    fmt.Println("⚡ 优化自动加载...")
    
    err := comp.DumpAutoload(true, true)
    if err != nil {
        log.Printf("❌ 自动加载优化失败: %v", err)
    } else {
        fmt.Println("✅ 自动加载优化完成")
    }
    
    // 清理缓存
    fmt.Println("🧹 清理缓存...")
    
    err = comp.ClearCache()
    if err != nil {
        log.Printf("❌ 清理缓存失败: %v", err)
    } else {
        fmt.Println("✅ 缓存清理完成")
    }
}

func validateProjectSetup(comp *composer.Composer) {
    // 验证项目配置
    err := comp.Validate()
    if err != nil {
        log.Fatalf("项目验证失败: %v", err)
    }
    
    fmt.Println("✅ 项目验证通过")
    
    // 显示项目信息
    fmt.Println("\n📋 项目信息:")
    
    if packages, err := comp.ShowAllPackages(); err == nil {
        fmt.Println("📦 已安装的包:")
        // 只显示前几行
        lines := strings.Split(packages, "\n")
        for i, line := range lines {
            if i >= 5 {
                fmt.Println("  ...")
                break
            }
            if strings.TrimSpace(line) != "" {
                fmt.Printf("  %s\n", line)
            }
        }
    }
}
```

### 现有项目迁移设置

```go
func migrateExistingProject() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    existingProjectPath := "/path/to/existing/project"
    comp.SetWorkingDir(existingProjectPath)
    
    fmt.Println("🔄 迁移现有项目...")
    
    // 1. 备份现有的 composer.json
    fmt.Println("💾 备份现有配置...")
    // 这里可以添加备份逻辑
    
    // 2. 更新依赖
    fmt.Println("🔄 更新依赖...")
    err = comp.Update(false, false)
    if err != nil {
        log.Printf("❌ 更新依赖失败: %v", err)
    }
    
    // 3. 优化项目
    fmt.Println("⚡ 优化项目...")
    err = comp.DumpAutoload(true, true)
    if err != nil {
        log.Printf("❌ 优化失败: %v", err)
    }
    
    // 4. 验证迁移结果
    fmt.Println("🔍 验证迁移结果...")
    err = comp.Validate()
    if err != nil {
        log.Fatalf("迁移验证失败: %v", err)
    }
    
    fmt.Println("✅ 项目迁移完成")
}
```

这个示例展示了使用 Go Composer SDK 进行完整项目设置的各种操作，包括项目创建、初始化、配置和验证等功能。
