# Go Composer SDK

Go Composer SDK 是一个全面的 Go 语言库，提供对 PHP Composer 包管理器的完整封装。它允许您从 Go 应用程序中管理 PHP 项目的依赖项、执行 Composer 命令以及处理各种 Composer 相关的功能。

## 特性

- **完整的 Composer 命令支持**：支持所有标准的 Composer CLI 命令
- **类型安全的 API**：所有方法都提供强类型接口，支持 IDE 代码补全
- **全面的功能集**：包含依赖管理、仓库配置、认证、安全审计等功能
- **多平台支持**：支持 Windows、macOS 和 Linux
- **模块化设计**：按功能分组的代码结构，易于使用和维护

## 安装

```bash
go get github.com/scagogogo/go-composer-sdk
```

## 快速开始

### 基本用法

```go
package main

import (
	"fmt"
	"github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
	// 使用默认选项创建 Composer 实例
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		panic(err)
	}

	// 安装依赖
	err = c.Install(false, false)
	if err != nil {
		panic(err)
	}

	// 添加新的包依赖
	err = c.RequirePackage("monolog/monolog", "^2.0", false)
	if err != nil {
		panic(err)
	}

	// 获取版本信息
	version, err := c.GetVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Composer 版本: %s\n", version)
}
```

### 高级配置

```go
package main

import (
	"time"
	"github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
	// 使用自定义选项创建 Composer 实例
	customOptions := composer.Options{
		ExecutablePath: "/usr/local/bin/composer", // 指定 Composer 可执行文件路径
		WorkingDir:     "/path/to/your/project",   // 指定工作目录
		AutoInstall:    true,                      // 未找到 Composer 时自动安装
		DefaultTimeout: 5 * time.Minute,           // 设置命令执行超时时间
		Env: map[string]string{
			"COMPOSER_MEMORY_LIMIT": "-1",       // 设置环境变量
			"COMPOSER_NO_INTERACTION": "1",
		},
	}
	
	c, err := composer.New(customOptions)
	if err != nil {
		panic(err)
	}
	
	// ... 使用 Composer 实例
}
```

## API 文档

Go Composer SDK 按功能模块提供了丰富的 API。以下是完整的 API 文档：

### 1. 基础核心

#### 1.1 初始化与配置

```go
// 创建具有默认选项的 Composer 实例
comp, err := composer.New(composer.DefaultOptions())

// 检查 Composer 是否已安装
isInstalled := comp.IsInstalled()

// 获取 Composer 版本
version, err := comp.GetVersion()

// 设置工作目录
comp.SetWorkingDir("/path/to/project")

// 设置环境变量
comp.SetEnv([]string{"COMPOSER_MEMORY_LIMIT=-1"})

// 更新 Composer 自身
err := comp.SelfUpdate()

// 直接运行 Composer 命令
output, err := comp.Run("--version")

// 使用上下文运行命令（支持超时控制）
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
output, err := comp.RunWithContext(ctx, "diagnose")
```

### 2. 包管理

#### 2.1 安装与更新

```go
// 安装依赖
err := comp.Install(false, false) // 参数: noDev, optimize

// 安装依赖（不包含开发依赖）
err := comp.Install(true, false)

// 安装依赖并优化自动加载
err := comp.Install(false, true)

// 使用更多选项安装依赖
options := map[string]string{
    "no-dev": "",
    "optimize-autoloader": "",
    "prefer-dist": "",
}
err := comp.InstallWithOptions(options)

// 更新所有依赖
err := comp.Update([]string{}, false) // 参数: packages, noDev

// 更新特定依赖
err := comp.Update([]string{"monolog/monolog", "symfony/console"}, false)

// 仅更新自动加载配置
err := comp.DumpAutoload(false) // 参数: optimize

// 更新并优化自动加载配置
err := comp.DumpAutoload(true)
```

#### 2.2 添加与移除包

```go
// 添加普通依赖（不指定版本）
err := comp.RequirePackage("monolog/monolog", "", false) // 参数: package, version, dev

// 添加指定版本的依赖
err := comp.RequirePackage("symfony/console", "^5.4", false)

// 添加开发依赖
err := comp.RequirePackage("phpunit/phpunit", "^9.5", true)

// 使用高级选项添加依赖
options := map[string]string{
    "no-update": "",
    "no-progress": "",
}
err := comp.RequirePackageWithOptions("guzzlehttp/guzzle", "^7.0", options)

// 移除普通依赖
err := comp.Remove("monolog/monolog", false) // 参数: package, dev

// 移除开发依赖
err := comp.Remove("phpunit/phpunit", true)
```

#### 2.3 包信息查询

```go
// 显示特定包的信息
output, err := comp.ShowPackage("monolog/monolog")

// 显示所有已安装的包
output, err := comp.ShowAllPackages()

// 显示依赖树
output, err := comp.ShowDependencyTree("")

// 显示特定包的依赖树
output, err := comp.ShowDependencyTree("symfony/console")

// 显示反向依赖关系 (哪些包依赖于指定包)
output, err := comp.ShowReverseDependencies("symfony/polyfill-mbstring")

// 解释为什么需要某个包
output, err := comp.WhyPackage("symfony/polyfill-mbstring")

// 查找过期的包
output, err := comp.OutdatedPackages()

// 查找直接依赖中的过期包
output, err := comp.OutdatedPackagesDirect()

// 搜索包
output, err := comp.Search("logger")
```

### 3. 项目管理

#### 3.1 项目创建与初始化

```go
// 创建新项目 (例如从模板创建 Laravel 项目)
err := comp.CreateProject("laravel/laravel", "my-laravel-project", "^10.0")

// 初始化项目 (创建 composer.json)
err := comp.InitProject()
```

#### 3.2 脚本运行

```go
// 列出可用的脚本
scriptList, err := comp.ListScripts()

// 执行脚本
output, err := comp.ExecuteScript("test")

// 执行脚本并传递额外参数
output, err := comp.RunScript("test", "--verbose")
```

#### 3.3 平台需求检查

```go
// 检查平台需求
output, err := comp.CheckPlatformReqs()

// 验证 composer.json 文件
err := comp.Validate()

// 严格验证 composer.json 文件
err := comp.ValidateComposerJson(true, false) // 参数: strict, withDependencies

// 验证 Schema
output, err := comp.ValidateSchema()

// 检查特定平台是否可用
available, err := comp.IsPlatformAvailable("php", "8.1")

// 获取 PHP 版本
phpVersion, err := comp.GetPHPVersion()

// 获取已安装的 PHP 扩展
extensions, err := comp.GetExtensions()

// 检查是否安装了特定 PHP 扩展
hasExtension, err := comp.HasExtension("mbstring")
```

#### 3.4 依赖分析

```go
// 检查依赖关系
output, err := comp.CheckDependencies()

// 安全审计
output, err := comp.Audit()

// JSON 格式的安全审计
result, err := comp.AuditWithJSON()

// 获取高严重性漏洞
vulns, err := comp.GetHighSeverityVulnerabilities()

// 检查项目是否有漏洞
hasVulns, err := comp.HasVulnerabilities()

// 获取已放弃的包
abandoned, err := comp.GetAbandonedPackages()
```

#### 3.5 完整性检查和诊断

```go
// 诊断系统
output, err := comp.Diagnose()

// 显示已安装包的修改状态
output, err := comp.Status()

// 清除缓存
err := comp.ClearCache()

// 获取 Composer 主目录
homeDir, err := comp.GetComposerHome()
```

### 4. 配置管理

```go
// 获取配置项
value, err := comp.GetConfig("vendor-dir")

// 设置配置项
err := comp.SetConfig("vendor-dir", "vendors")

// 获取全局配置
value, err := comp.GetConfigWithGlobal("github-oauth.github.com", true)

// 设置全局配置
err := comp.SetConfigWithGlobal("github-oauth.github.com", "your-token", true)
```

### 5. Composer.json 操作

```go
// 读取 composer.json
composerJSON, err := comp.ReadComposerJSON()

// 写入 composer.json
err := comp.WriteComposerJSON(composerJSON)

// 添加依赖到 composer.json
err := comp.AddRequire("symfony/console", "^5.0", false)

// 添加开发依赖到 composer.json
err := comp.AddRequire("phpunit/phpunit", "^9.0", true)

// 移除依赖
err := comp.RemoveRequire("symfony/console", false)

// 添加作者
err := comp.AddAuthor("Your Name", "your.email@example.com", "https://example.com", "Developer")

// 添加仓库
repo := composer.Repository{
    Type: composer.ComposerRepository,
    URL:  "https://repo.example.com",
}
err := comp.AddRepository("example", repo)

// 设置属性
err := comp.SetProperty("name", "vendor/project")
err := comp.SetProperty("description", "My awesome project")
err := comp.SetProperty("type", "project")
err := comp.SetProperty("license", "MIT")
err := comp.SetProperty("minimum-stability", "stable")
err := comp.SetProperty("prefer-stable", true)
```

## 示例代码

Go Composer SDK 提供了丰富的示例代码，展示了库的各种功能。示例按以下类别组织：

### 示例目录结构

```
examples/
├── 01_basic_usage/            # 基本用法示例
│   ├── 01_new_composer.go     # 创建 Composer 实例
│   └── 02_run_commands.go     # 运行基本命令
├── 02_package_management/     # 包管理示例
│   ├── 01_install_update.go   # 安装和更新依赖
│   ├── 02_require_remove.go   # 添加和移除包
│   └── 03_show_package.go     # 查看包信息
├── 03_project_management/     # 项目管理示例
│   ├── 01_create_project.go   # 创建和初始化项目
│   ├── 02_run_script.go       # 运行脚本
│   ├── 03_platform_check.go   # 平台需求检查
│   ├── 04_dependency_analysis.go # 依赖分析
│   └── 05_integrity_check.go  # 完整性检查和诊断
└── main.go                    # 运行示例的入口
```

### 运行示例

可以通过以下命令运行示例：

```bash
# 进入项目根目录
cd go-composer-sdk

# 查看所有可用示例
go run examples/main.go

# 运行特定示例 (例如 "基础用法 - 创建 Composer 实例")
go run examples/main.go 1.1
```

## 常见场景

以下是一些常见场景的代码示例：

### 安装并使用特定版本的依赖包

```go
package main

import (
	"github.com/scagogogo/go-composer-sdk/pkg/composer"
	"log"
)

func main() {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}
	
	// 添加 Laravel 框架依赖
	err = comp.RequirePackage("laravel/framework", "^10.0", false)
	if err != nil {
		log.Fatalf("添加依赖失败: %v", err)
	}
	
	// 添加开发依赖
	err = comp.RequirePackage("phpunit/phpunit", "^10.0", true)
	if err != nil {
		log.Fatalf("添加开发依赖失败: %v", err)
	}
	
	// 安装依赖并优化自动加载
	err = comp.Install(false, true)
	if err != nil {
		log.Fatalf("安装依赖失败: %v", err)
	}
	
	log.Println("依赖安装完成")
}
```

### 安全审计与漏洞检查

```go
package main

import (
	"fmt"
	"github.com/scagogogo/go-composer-sdk/pkg/composer"
	"log"
)

func main() {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}
	
	// 执行安全审计
	result, err := comp.AuditWithJSON()
	if err != nil {
		log.Fatalf("安全审计失败: %v", err)
	}
	
	fmt.Printf("找到 %d 个漏洞\n", result.Found)
	
	// 获取并处理高危漏洞
	highSeverityVulns, err := comp.GetHighSeverityVulnerabilities()
	if err != nil {
		log.Fatalf("获取高危漏洞失败: %v", err)
	}
	
	if len(highSeverityVulns) > 0 {
		fmt.Println("项目中存在以下高危漏洞:")
		for _, vuln := range highSeverityVulns {
			fmt.Printf("- %s (%s): %s\n", vuln.Package, vuln.Severity, vuln.Title)
			fmt.Printf("  更多信息: %s\n", vuln.Link)
		}
	} else {
		fmt.Println("未发现高危漏洞")
	}
}
```

### 项目创建与配置

```go
package main

import (
	"github.com/scagogogo/go-composer-sdk/pkg/composer"
	"log"
)

func main() {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}
	
	// 创建一个新的 Laravel 项目
	err = comp.CreateProject("laravel/laravel", "my-laravel-project", "")
	if err != nil {
		log.Fatalf("创建项目失败: %v", err)
	}
	
	// 切换到新项目目录
	comp.SetWorkingDir("my-laravel-project")
	
	// 设置自定义配置
	err = comp.SetConfig("process-timeout", "600")
	if err != nil {
		log.Fatalf("设置配置失败: %v", err)
	}
	
	// 添加自定义仓库
	repo := composer.Repository{
		Type: composer.ComposerRepository,
		URL:  "https://my-private-repo.example.com",
	}
	err = comp.AddRepository("private", repo)
	if err != nil {
		log.Fatalf("添加仓库失败: %v", err)
	}
	
	log.Println("项目创建和配置完成")
}
```

## 故障排除

常见问题及解决方案：

1. **无法找到 Composer**
   - 确保 Composer 已安装并在系统 PATH 中
   - 使用 `ExecutablePath` 选项明确指定 Composer 路径
   - 设置 `AutoInstall: true` 允许 SDK 自动安装

2. **命令超时**
   - 增加 `DefaultTimeout` 选项的值
   - 使用 `COMPOSER_PROCESS_TIMEOUT` 环境变量

3. **内存不足**
   - 设置 `COMPOSER_MEMORY_LIMIT=-1` 环境变量

4. **禁用交互提示**
   - 设置 `COMPOSER_NO_INTERACTION=1` 环境变量

## 贡献

欢迎贡献代码、报告问题或提出功能请求！请确保在提交 Pull Request 之前，您的代码通过了测试。

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。
