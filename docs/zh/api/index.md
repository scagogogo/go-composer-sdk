# API 参考

欢迎查看 Go Composer SDK API 参考文档。本文档提供了所有可用包、类型和函数的详细信息。

## 包概览

Go Composer SDK 组织为几个包，每个包都有特定的用途：

### 核心包

| 包 | 描述 |
|---------|-------------|
| [`composer`](/zh/api/core) | 包含核心 Composer 功能的主包 |
| [`detector`](/zh/api/detector) | Composer 安装检测和验证 |
| [`installer`](/zh/api/installer) | 自动 Composer 安装工具 |
| [`utils`](/zh/api/utilities) | 通用工具和辅助函数 |

## 主要类型

### Composer

所有 Composer 操作的主要接口：

```go
type Composer struct {
    // 包含过滤或未导出的字段
}
```

**主要方法：**
- 包管理：`Install()`、`Update()`、`RequirePackage()`、`RemovePackage()`
- 项目操作：`CreateProject()`、`Validate()`、`RunScript()`
- 信息获取：`GetVersion()`、`ShowPackage()`、`GetLicenses()`
- 安全功能：`Audit()`、`CheckPlatformReqs()`

### Options

创建 Composer 实例的配置选项：

```go
type Options struct {
    ExecutablePath  string        // composer 可执行文件路径
    WorkingDir      string        // 操作的工作目录
    AutoInstall     bool          // 如果未找到则自动安装 Composer
    DefaultTimeout  time.Duration // 操作的默认超时时间
    Detector        *detector.Detector // 自定义检测器实例
    Installer       *installer.Installer // 自定义安装器实例
}
```

## 快速参考

### 创建 Composer 实例

```go
// 默认选项
comp, err := composer.New(composer.DefaultOptions())

// 自定义选项
options := composer.Options{
    WorkingDir:     "/path/to/project",
    AutoInstall:    true,
    DefaultTimeout: 5 * time.Minute,
}
comp, err := composer.New(options)
```

### 常见操作

```go
// 检查安装
isInstalled := comp.IsInstalled()

// 获取版本
version, err := comp.GetVersion()

// 安装依赖
err = comp.Install(false, false) // noDev, optimize

// 添加包
err = comp.RequirePackage("monolog/monolog", "^3.0")

// 更新包
err = comp.Update(false, false)

// 显示包信息
info, err := comp.ShowPackage("symfony/console")

// 验证 composer.json
err = comp.Validate()

// 运行安全审计
result, err := comp.Audit()
```

## API 分类

### [核心操作](/zh/api/core)
基本的 Composer 功能，包括实例创建、版本管理和命令执行。

### [包管理](/zh/api/package-management)
安装、更新、添加和删除包。管理依赖和包信息。

### [项目管理](/zh/api/project-management)
创建项目、运行脚本、验证配置和管理项目级设置。

### [安全审计](/zh/api/security-audit)
安全审计、漏洞检测和依赖分析。

### [平台环境](/zh/api/platform-environment)
平台要求检查、环境配置和系统兼容性。

### [工具函数](/zh/api/utilities)
文件操作、HTTP 请求和跨平台兼容性的辅助函数。

### [检测器](/zh/api/detector)
Composer 安装检测和路径解析。

### [安装器](/zh/api/installer)
自动 Composer 安装和设置。

## 最佳实践

### 1. 始终处理错误

```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("创建 Composer 实例失败: %v", err)
}
```

### 2. 对长时间运行的操作使用上下文

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

err := comp.UpdateWithContext(ctx, false, false)
```

### 3. 设置工作目录

```go
comp.SetWorkingDir("/path/to/your/php/project")
```

### 4. 配置环境变量

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
})
```

### 5. 在操作前检查安装

```go
if !comp.IsInstalled() {
    log.Fatal("Composer 未安装")
}
```

## 示例

有关实际示例和用例，请参阅[示例部分](/zh/examples/)。

## 支持

- [GitHub 仓库](https://github.com/scagogogo/go-composer-sdk)
- [问题跟踪](https://github.com/scagogogo/go-composer-sdk/issues)
- [讨论区](https://github.com/scagogogo/go-composer-sdk/discussions)
