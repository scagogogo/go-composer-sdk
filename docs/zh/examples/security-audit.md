# 安全审计示例

本页面展示如何使用 Go Composer SDK 进行 PHP 项目的安全审计，包括漏洞检测、平台要求检查等安全相关功能。

## 基本安全审计

### 执行安全审计

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
    
    fmt.Println("🔒 开始安全审计...")
    
    // 执行安全审计
    auditResult, err := comp.Audit()
    if err != nil {
        log.Fatalf("安全审计失败: %v", err)
    }
    
    if auditResult == "" {
        fmt.Println("✅ 未发现安全漏洞！")
    } else {
        fmt.Println("⚠️  发现安全问题:")
        fmt.Println(auditResult)
    }
}
```

### 详细安全审计

```go
func detailedSecurityAudit() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 执行详细安全审计...")
    
    // 设置审计环境变量
    comp.SetEnv([]string{
        "COMPOSER_AUDIT_ABANDONED=report", // 报告废弃的包
    })
    
    auditResult, err := comp.Audit()
    if err != nil {
        log.Fatalf("详细安全审计失败: %v", err)
    }
    
    fmt.Println("📋 安全审计报告:")
    if auditResult == "" {
        fmt.Println("✅ 项目安全状况良好")
        fmt.Println("  - 未发现已知漏洞")
        fmt.Println("  - 未发现废弃的包")
    } else {
        fmt.Println("⚠️  发现以下安全问题:")
        fmt.Println(auditResult)
        
        // 提供修复建议
        fmt.Println("\n💡 修复建议:")
        fmt.Println("  1. 更新存在漏洞的包到安全版本")
        fmt.Println("  2. 替换已废弃的包")
        fmt.Println("  3. 定期执行安全审计")
    }
}
```

## 平台要求检查

### 检查平台要求

```go
func checkPlatformRequirements() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 检查平台要求...")
    
    err = comp.CheckPlatformReqs()
    if err != nil {
        fmt.Printf("❌ 平台要求检查失败: %v\n", err)
        
        // 提供解决建议
        fmt.Println("\n💡 可能的解决方案:")
        fmt.Println("  1. 升级 PHP 版本")
        fmt.Println("  2. 安装缺失的 PHP 扩展")
        fmt.Println("  3. 更新系统依赖")
        
        return
    }
    
    fmt.Println("✅ 平台要求检查通过")
}
```

### 详细平台信息

```go
func detailedPlatformInfo() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("📋 获取详细平台信息...")
    
    // 获取平台信息
    platformInfo, err := comp.GetPlatformInfo()
    if err != nil {
        log.Printf("❌ 获取平台信息失败: %v", err)
        return
    }
    
    fmt.Println("🖥️  平台信息:")
    fmt.Println(platformInfo)
    
    // 检查平台要求
    fmt.Println("\n🔍 验证平台要求...")
    err = comp.CheckPlatformReqs()
    if err != nil {
        fmt.Printf("❌ 平台要求不满足: %v\n", err)
    } else {
        fmt.Println("✅ 平台要求满足")
    }
}
```

## 依赖安全分析

### 分析依赖安全性

```go
func analyzeDependencySecurity() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("🔍 分析依赖安全性...")
    
    // 1. 获取所有已安装的包
    packages, err := comp.ShowAllPackages()
    if err != nil {
        log.Printf("❌ 获取包列表失败: %v", err)
        return
    }
    
    fmt.Println("📦 已安装的包:")
    fmt.Println(packages)
    
    // 2. 检查过时的包
    fmt.Println("\n🔍 检查过时的包...")
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("❌ 检查过时包失败: %v", err)
    } else if outdated != "" {
        fmt.Println("⚠️  发现过时的包:")
        fmt.Println(outdated)
        fmt.Println("\n💡 建议: 更新过时的包以获得安全修复")
    } else {
        fmt.Println("✅ 所有包都是最新的")
    }
    
    // 3. 执行安全审计
    fmt.Println("\n🔒 执行安全审计...")
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("❌ 安全审计失败: %v", err)
    } else if auditResult != "" {
        fmt.Println("⚠️  发现安全问题:")
        fmt.Println(auditResult)
    } else {
        fmt.Println("✅ 未发现安全问题")
    }
}
```

### 检查特定包的安全性

```go
func checkSpecificPackageSecurity() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 要检查的关键包
    criticalPackages := []string{
        "symfony/symfony",
        "laravel/framework",
        "guzzlehttp/guzzle",
        "monolog/monolog",
        "doctrine/orm",
    }
    
    fmt.Println("🔍 检查关键包的安全性...")
    
    for _, packageName := range criticalPackages {
        fmt.Printf("\n📦 检查包: %s\n", packageName)
        
        // 获取包详情
        details, err := comp.ShowPackage(packageName)
        if err != nil {
            fmt.Printf("  ❌ 包不存在或获取失败: %v\n", err)
            continue
        }
        
        fmt.Printf("  ✅ 包信息获取成功\n")
        
        // 检查为什么安装了这个包
        reasons, err := comp.WhyPackage(packageName)
        if err == nil {
            fmt.Printf("  📋 依赖原因: %s\n", reasons)
        }
    }
    
    // 执行整体安全审计
    fmt.Println("\n🔒 执行整体安全审计...")
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("❌ 安全审计失败: %v", err)
    } else if auditResult != "" {
        fmt.Println("⚠️  发现安全问题:")
        fmt.Println(auditResult)
    } else {
        fmt.Println("✅ 关键包安全检查通过")
    }
}
```

## 自动化安全检查

### 定期安全检查

```go
func scheduledSecurityCheck() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    fmt.Println("⏰ 执行定期安全检查...")
    
    // 安全检查步骤
    securityChecks := []struct {
        name string
        fn   func() error
    }{
        {"平台要求检查", func() error {
            return comp.CheckPlatformReqs()
        }},
        {"项目验证", func() error {
            return comp.Validate()
        }},
        {"安全审计", func() error {
            result, err := comp.Audit()
            if err != nil {
                return err
            }
            if result != "" {
                return fmt.Errorf("发现安全问题: %s", result)
            }
            return nil
        }},
        {"过时包检查", func() error {
            outdated, err := comp.ShowOutdated()
            if err != nil {
                return err
            }
            if outdated != "" {
                fmt.Printf("⚠️  发现过时的包:\n%s\n", outdated)
            }
            return nil
        }},
    }
    
    var failedChecks []string
    
    for _, check := range securityChecks {
        fmt.Printf("🔄 执行: %s\n", check.name)
        
        if err := check.fn(); err != nil {
            fmt.Printf("  ❌ %s 失败: %v\n", check.name, err)
            failedChecks = append(failedChecks, check.name)
        } else {
            fmt.Printf("  ✅ %s 通过\n", check.name)
        }
    }
    
    // 生成安全报告
    fmt.Println("\n📋 安全检查报告:")
    if len(failedChecks) == 0 {
        fmt.Println("✅ 所有安全检查通过")
    } else {
        fmt.Printf("❌ %d 项检查失败:\n", len(failedChecks))
        for _, check := range failedChecks {
            fmt.Printf("  - %s\n", check)
        }
    }
}
```

### CI/CD 安全检查

```go
func cicdSecurityCheck() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 设置 CI/CD 环境变量
    comp.SetEnv([]string{
        "COMPOSER_NO_INTERACTION=1",
        "COMPOSER_AUDIT_ABANDONED=report",
    })
    
    fmt.Println("🚀 CI/CD 安全检查...")
    
    // 1. 验证项目配置
    fmt.Println("1️⃣  验证项目配置")
    if err := comp.Validate(); err != nil {
        log.Fatalf("❌ 项目配置验证失败: %v", err)
    }
    fmt.Println("✅ 项目配置验证通过")
    
    // 2. 检查平台要求
    fmt.Println("2️⃣  检查平台要求")
    if err := comp.CheckPlatformReqs(); err != nil {
        log.Fatalf("❌ 平台要求检查失败: %v", err)
    }
    fmt.Println("✅ 平台要求检查通过")
    
    // 3. 执行安全审计
    fmt.Println("3️⃣  执行安全审计")
    auditResult, err := comp.Audit()
    if err != nil {
        log.Fatalf("❌ 安全审计失败: %v", err)
    }
    
    if auditResult != "" {
        fmt.Println("❌ 发现安全问题:")
        fmt.Println(auditResult)
        
        // 在 CI/CD 中，安全问题应该导致构建失败
        if os.Getenv("CI") == "true" {
            log.Fatal("❌ CI/CD: 由于安全问题，构建失败")
        }
    } else {
        fmt.Println("✅ 安全审计通过")
    }
    
    // 4. 检查过时的包
    fmt.Println("4️⃣  检查过时的包")
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("⚠️  检查过时包失败: %v", err)
    } else if outdated != "" {
        fmt.Println("⚠️  发现过时的包:")
        fmt.Println(outdated)
        
        // 在 CI/CD 中，可以选择是否因过时包而失败
        if os.Getenv("FAIL_ON_OUTDATED") == "true" {
            log.Fatal("❌ CI/CD: 由于过时的包，构建失败")
        }
    } else {
        fmt.Println("✅ 所有包都是最新的")
    }
    
    fmt.Println("🎉 CI/CD 安全检查完成")
}
```

## 完整的安全审计工作流

### 综合安全审计

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    fmt.Println("🔒 完整的安全审计工作流")
    fmt.Println("==========================")
    
    // 1. 设置 Composer
    fmt.Println("\n1️⃣  设置 Composer")
    comp := setupSecureComposer()
    
    // 2. 环境检查
    fmt.Println("\n2️⃣  环境安全检查")
    checkEnvironmentSecurity(comp)
    
    // 3. 依赖安全审计
    fmt.Println("\n3️⃣  依赖安全审计")
    auditDependencies(comp)
    
    // 4. 配置安全检查
    fmt.Println("\n4️⃣  配置安全检查")
    checkConfigurationSecurity(comp)
    
    // 5. 生成安全报告
    fmt.Println("\n5️⃣  生成安全报告")
    generateSecurityReport(comp)
    
    fmt.Println("\n🎉 安全审计工作流完成！")
}

func setupSecureComposer() *composer.Composer {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 3 * time.Minute
    
    comp, err := composer.New(options)
    if err != nil {
        log.Fatalf("设置 Composer 失败: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/your/project")
    
    // 设置安全相关的环境变量
    comp.SetEnv([]string{
        "COMPOSER_AUDIT_ABANDONED=report",
        "COMPOSER_PREFER_STABLE=true",
        "COMPOSER_MINIMUM_STABILITY=stable",
    })
    
    fmt.Println("✅ 安全 Composer 设置完成")
    return comp
}

func checkEnvironmentSecurity(comp *composer.Composer) {
    fmt.Println("🔍 检查环境安全性...")
    
    // 检查 Composer 版本
    if version, err := comp.GetVersion(); err == nil {
        fmt.Printf("📦 Composer 版本: %s\n", version)
    }
    
    // 检查平台要求
    if err := comp.CheckPlatformReqs(); err != nil {
        fmt.Printf("❌ 平台要求检查失败: %v\n", err)
    } else {
        fmt.Println("✅ 平台要求检查通过")
    }
    
    // 验证项目配置
    if err := comp.Validate(); err != nil {
        fmt.Printf("❌ 项目配置验证失败: %v\n", err)
    } else {
        fmt.Println("✅ 项目配置验证通过")
    }
}

func auditDependencies(comp *composer.Composer) {
    fmt.Println("🔒 执行依赖安全审计...")
    
    // 执行安全审计
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("❌ 安全审计失败: %v", err)
        return
    }
    
    if auditResult == "" {
        fmt.Println("✅ 未发现安全漏洞")
    } else {
        fmt.Println("⚠️  发现安全问题:")
        fmt.Println(auditResult)
    }
    
    // 检查过时的包
    fmt.Println("\n🔍 检查过时的包...")
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("❌ 检查过时包失败: %v", err)
    } else if outdated != "" {
        fmt.Println("⚠️  发现过时的包:")
        fmt.Println(outdated)
    } else {
        fmt.Println("✅ 所有包都是最新的")
    }
}

func checkConfigurationSecurity(comp *composer.Composer) {
    fmt.Println("🔧 检查配置安全性...")
    
    // 运行诊断
    if output, err := comp.Diagnose(); err == nil {
        fmt.Println("🔍 诊断结果:")
        fmt.Println(output)
    } else {
        fmt.Printf("❌ 诊断失败: %v\n", err)
    }
}

func generateSecurityReport(comp *composer.Composer) {
    fmt.Println("📋 生成安全报告...")
    
    reportTime := time.Now().Format("2006-01-02 15:04:05")
    
    fmt.Printf(`
🔒 安全审计报告
================
审计时间: %s
项目路径: %s

审计项目:
✅ 环境安全检查
✅ 依赖安全审计  
✅ 配置安全检查
✅ 平台要求验证

建议:
1. 定期执行安全审计
2. 及时更新依赖包
3. 监控安全公告
4. 使用稳定版本

`, reportTime, comp.GetWorkingDir())
    
    fmt.Println("✅ 安全报告生成完成")
}
```

这个示例展示了使用 Go Composer SDK 进行全面安全审计的各种功能，包括漏洞检测、平台要求检查、依赖分析和自动化安全检查等。
