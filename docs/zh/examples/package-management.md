# 包管理示例

本页面展示如何使用 Go Composer SDK 进行 PHP 包的管理，包括安装、更新、删除包以及依赖分析等功能。

## 安装依赖

### 基本安装

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
    
    // 设置项目目录
    comp.SetWorkingDir("/path/to/your/php/project")
    
    fmt.Println("📦 开始安装依赖...")
    
    // 安装所有依赖（包括开发依赖）
    err = comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        log.Fatalf("安装依赖失败: %v", err)
    }
    
    fmt.Println("✅ 依赖安装成功！")
}
```

### 生产环境安装

```go
func installForProduction() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/production/project")
    
    // 设置生产环境变量
    comp.SetEnv([]string{
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",
        "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
        "COMPOSER_PREFER_STABLE=true",
    })
    
    fmt.Println("🏭 生产环境依赖安装...")
    
    // 安装依赖，排除开发依赖，启用优化
    err = comp.Install(true, true) // noDev=true, optimize=true
    if err != nil {
        log.Fatalf("生产环境安装失败: %v", err)
    }
    
    fmt.Println("✅ 生产环境依赖安装完成！")
}
```

## 添加包

### 添加单个包

```go
func addSinglePackage() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 添加包
    packageName := "monolog/monolog"
    version := "^3.0"
    
    fmt.Printf("📦 添加包: %s %s\n", packageName, version)
    
    err = comp.RequirePackage(packageName, version)
    if err != nil {
        log.Fatalf("添加包失败: %v", err)
    }
    
    fmt.Println("✅ 包添加成功！")
}
```

### 批量添加包

```go
func addMultiplePackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 定义要添加的包
    packages := map[string]string{
        "symfony/console":        "^6.0",
        "guzzlehttp/guzzle":     "^7.0",
        "doctrine/orm":          "^2.14",
        "symfony/http-foundation": "^6.0",
    }
    
    fmt.Println("📦 批量添加包...")
    
    err = comp.RequirePackages(packages)
    if err != nil {
        log.Fatalf("批量添加包失败: %v", err)
    }
    
    fmt.Println("✅ 所有包添加成功！")
}
```

### 添加开发依赖

```go
func addDevDependencies() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 添加开发依赖
    devPackages := map[string]string{
        "phpunit/phpunit":        "^10.0",
        "squizlabs/php_codesniffer": "^3.7",
        "phpstan/phpstan":        "^1.10",
        "friendsofphp/php-cs-fixer": "^3.0",
    }
    
    fmt.Println("🛠️  添加开发依赖...")
    
    for packageName, version := range devPackages {
        fmt.Printf("  📦 添加开发包: %s %s\n", packageName, version)
        
        err = comp.RequireDevPackage(packageName, version)
        if err != nil {
            log.Printf("❌ 添加 %s 失败: %v", packageName, err)
            continue
        }
        
        fmt.Printf("  ✅ %s 添加成功\n", packageName)
    }
    
    fmt.Println("✅ 开发依赖添加完成！")
}
```

## 更新包

### 更新所有包

```go
func updateAllPackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔄 更新所有包...")
    
    // 更新所有包
    err = comp.Update(false, false) // noDev=false, optimize=false
    if err != nil {
        log.Fatalf("更新失败: %v", err)
    }
    
    fmt.Println("✅ 所有包更新完成！")
}
```

### 更新特定包

```go
func updateSpecificPackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 更新特定包
    packagesToUpdate := []string{
        "symfony/console",
        "monolog/monolog",
        "guzzlehttp/guzzle",
    }
    
    fmt.Println("🔄 更新特定包...")
    
    for _, packageName := range packagesToUpdate {
        fmt.Printf("  🔄 更新包: %s\n", packageName)
        
        err = comp.UpdatePackage(packageName)
        if err != nil {
            log.Printf("❌ 更新 %s 失败: %v", packageName, err)
            continue
        }
        
        fmt.Printf("  ✅ %s 更新成功\n", packageName)
    }
    
    fmt.Println("✅ 特定包更新完成！")
}
```

### 批量更新包

```go
func updatePackagesBatch() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 批量更新包
    packages := []string{
        "symfony/console",
        "symfony/http-foundation",
        "doctrine/orm",
    }
    
    fmt.Println("🔄 批量更新包...")
    
    err = comp.UpdatePackages(packages)
    if err != nil {
        log.Fatalf("批量更新失败: %v", err)
    }
    
    fmt.Println("✅ 批量更新完成！")
}
```

## 删除包

### 删除单个包

```go
func removeSinglePackage() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    packageName := "old-package/deprecated"
    
    fmt.Printf("🗑️  删除包: %s\n", packageName)
    
    err = comp.RemovePackage(packageName)
    if err != nil {
        log.Fatalf("删除包失败: %v", err)
    }
    
    fmt.Println("✅ 包删除成功！")
}
```

### 批量删除包

```go
func removeMultiplePackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 要删除的包列表
    packagesToRemove := []string{
        "old-package/deprecated",
        "unused-library/package",
        "legacy-tool/helper",
    }
    
    fmt.Println("🗑️  批量删除包...")
    
    err = comp.RemovePackages(packagesToRemove)
    if err != nil {
        log.Fatalf("批量删除失败: %v", err)
    }
    
    fmt.Println("✅ 批量删除完成！")
}
```

## 包信息查询

### 查看已安装的包

```go
func listInstalledPackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("📋 查看已安装的包...")
    
    // 获取所有已安装的包
    packages, err := comp.ShowAllPackages()
    if err != nil {
        log.Fatalf("获取包列表失败: %v", err)
    }
    
    fmt.Println("📦 已安装的包:")
    fmt.Println(packages)
}
```

### 查看包详情

```go
func showPackageDetails() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    packageName := "symfony/console"
    
    fmt.Printf("🔍 查看包详情: %s\n", packageName)
    
    details, err := comp.ShowPackage(packageName)
    if err != nil {
        log.Fatalf("获取包详情失败: %v", err)
    }
    
    fmt.Printf("📄 包详情:\n%s\n", details)
}
```

### 检查过时的包

```go
func checkOutdatedPackages() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 检查过时的包...")
    
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Fatalf("检查过时包失败: %v", err)
    }
    
    if outdated == "" {
        fmt.Println("✅ 所有包都是最新的！")
    } else {
        fmt.Println("📦 发现过时的包:")
        fmt.Println(outdated)
    }
}
```

## 依赖分析

### 查看依赖树

```go
func showDependencyTree() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🌳 查看依赖树...")
    
    tree, err := comp.ShowDependencyTree("")
    if err != nil {
        log.Fatalf("获取依赖树失败: %v", err)
    }
    
    fmt.Println("🌳 依赖树:")
    fmt.Println(tree)
}
```

### 分析包依赖原因

```go
func analyzePackageDependencies() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    packageName := "psr/log"
    
    fmt.Printf("🔍 分析为什么安装了包: %s\n", packageName)
    
    reasons, err := comp.WhyPackage(packageName)
    if err != nil {
        log.Fatalf("分析依赖失败: %v", err)
    }
    
    fmt.Printf("📋 依赖原因:\n%s\n", reasons)
}
```

### 分析包冲突

```go
func analyzePackageConflicts() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    packageName := "symfony/console"
    version := "^7.0"
    
    fmt.Printf("🔍 分析为什么不能安装: %s %s\n", packageName, version)
    
    conflicts, err := comp.WhyNotPackage(packageName, version)
    if err != nil {
        log.Fatalf("分析冲突失败: %v", err)
    }
    
    fmt.Printf("⚠️  冲突原因:\n%s\n", conflicts)
}
```

## 完整的包管理工作流

### 项目初始化和包管理

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    fmt.Println("📦 完整的包管理工作流示例")
    fmt.Println("================================")
    
    comp, err := setupComposer()
    if err != nil {
        log.Fatalf("设置 Composer 失败: %v", err)
    }
    
    // 1. 安装基础依赖
    fmt.Println("\n1️⃣  安装基础依赖")
    installBaseDependencies(comp)
    
    // 2. 添加新包
    fmt.Println("\n2️⃣  添加新包")
    addNewPackages(comp)
    
    // 3. 检查和更新
    fmt.Println("\n3️⃣  检查和更新")
    checkAndUpdate(comp)
    
    // 4. 清理不需要的包
    fmt.Println("\n4️⃣  清理不需要的包")
    cleanupPackages(comp)
    
    // 5. 最终验证
    fmt.Println("\n5️⃣  最终验证")
    finalValidation(comp)
    
    fmt.Println("\n🎉 包管理工作流完成！")
}

func setupComposer() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 5 * time.Minute
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 设置环境变量
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
    })
    
    fmt.Println("✅ Composer 设置完成")
    return comp, nil
}

func installBaseDependencies(comp *composer.Composer) {
    fmt.Println("📦 安装基础依赖...")
    
    if err := comp.Install(false, false); err != nil {
        log.Printf("❌ 安装失败: %v", err)
        return
    }
    
    fmt.Println("✅ 基础依赖安装完成")
}

func addNewPackages(comp *composer.Composer) {
    packages := map[string]string{
        "monolog/monolog":    "^3.0",
        "guzzlehttp/guzzle": "^7.0",
    }
    
    fmt.Println("📦 添加新包...")
    
    if err := comp.RequirePackages(packages); err != nil {
        log.Printf("❌ 添加包失败: %v", err)
        return
    }
    
    fmt.Println("✅ 新包添加完成")
}

func checkAndUpdate(comp *composer.Composer) {
    fmt.Println("🔍 检查过时的包...")
    
    if outdated, err := comp.ShowOutdated(); err == nil && outdated != "" {
        fmt.Println("📦 发现过时的包，正在更新...")
        
        if err := comp.Update(false, false); err != nil {
            log.Printf("❌ 更新失败: %v", err)
            return
        }
        
        fmt.Println("✅ 包更新完成")
    } else {
        fmt.Println("✅ 所有包都是最新的")
    }
}

func cleanupPackages(comp *composer.Composer) {
    // 这里可以根据实际需要删除不需要的包
    fmt.Println("🧹 清理不需要的包...")
    fmt.Println("✅ 清理完成")
}

func finalValidation(comp *composer.Composer) {
    fmt.Println("🔍 最终验证...")
    
    if err := comp.Validate(); err != nil {
        log.Printf("❌ 验证失败: %v", err)
        return
    }
    
    fmt.Println("✅ 项目验证通过")
}
```

这个示例展示了使用 Go Composer SDK 进行完整包管理的各种操作，包括安装、更新、删除包以及依赖分析等功能。
