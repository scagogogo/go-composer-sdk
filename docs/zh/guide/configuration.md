# 配置

本指南涵盖如何为不同环境和用例配置 Go Composer SDK。

## 基本配置

### 默认选项

开始使用的最简单方法是使用默认选项：

```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("创建 Composer 实例失败: %v", err)
}
```

默认选项包括：
- **AutoInstall**: `true` - 如果未找到则自动安装 Composer
- **DefaultTimeout**: `5 分钟` - 操作的默认超时时间
- **WorkingDir**: 当前目录
- **ExecutablePath**: 自动检测

### 自定义配置

要获得更多控制，请创建自定义选项：

```go
options := composer.Options{
    ExecutablePath:  "/usr/local/bin/composer",
    WorkingDir:      "/path/to/php/project",
    AutoInstall:     false,
    DefaultTimeout:  10 * time.Minute,
}

comp, err := composer.New(options)
if err != nil {
    log.Fatalf("创建 Composer 实例失败: %v", err)
}
```

## 配置选项

### ExecutablePath

指定 Composer 可执行文件的自定义路径：

```go
options := composer.DefaultOptions()
options.ExecutablePath = "/opt/composer/composer"
comp, err := composer.New(options)
```

**使用场景：**
- Composer 安装在非标准位置
- 系统上有多个 Composer 版本
- 容器化环境中的特定路径

### WorkingDir

设置 Composer 操作的工作目录：

```go
options := composer.DefaultOptions()
options.WorkingDir = "/var/www/html/my-project"
comp, err := composer.New(options)

// 或在创建后设置
comp.SetWorkingDir("/path/to/another/project")
```

**重要说明：**
- 必须包含有效的 `composer.json` 文件
- 所有 Composer 操作都将相对于此目录
- 可以在运行时使用 `SetWorkingDir()` 更改

### AutoInstall

控制是否在未找到 Composer 时自动安装：

```go
// 启用自动安装（默认）
options := composer.DefaultOptions()
options.AutoInstall = true

// 禁用自动安装
options.AutoInstall = false
```

**何时禁用：**
- 生产环境中需要明确控制
- 自动安装可能失败的系统
- 使用自定义 Composer 安装时

### DefaultTimeout

设置 Composer 操作的默认超时时间：

```go
options := composer.DefaultOptions()
options.DefaultTimeout = 15 * time.Minute // 适用于慢速网络
comp, err := composer.New(options)

// 或在创建后设置
comp.SetTimeout(30 * time.Minute)
```

**考虑因素：**
- 大型项目可能需要更长的超时时间
- 网络条件影响所需的超时时间
- 可以使用上下文为每个操作覆盖

## 环境变量

### 设置环境变量

使用环境变量配置 Composer 行为：

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",           // 无限内存
    "COMPOSER_PROCESS_TIMEOUT=600",       // 10 分钟超时
    "COMPOSER_CACHE_DIR=/tmp/composer",   // 自定义缓存目录
    "COMPOSER_HOME=/opt/composer",        // 自定义主目录
    "COMPOSER_DISCARD_CHANGES=true",      // 自动丢弃更改
})
```

### 常见环境变量

#### 内存和性能
```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",           // 移除内存限制
    "COMPOSER_PROCESS_TIMEOUT=0",         // 无进程超时
    "COMPOSER_HTACCESS_PROTECT=0",        // 禁用 .htaccess 保护
})
```

#### 缓存
```go
comp.SetEnv([]string{
    "COMPOSER_CACHE_DIR=/var/cache/composer",  // 自定义缓存位置
    "COMPOSER_CACHE_FILES_TTL=86400",          // 缓存 TTL（秒）
    "COMPOSER_CACHE_REPO_TTL=3600",            // 仓库缓存 TTL
})
```

#### 网络和代理
```go
comp.SetEnv([]string{
    "HTTP_PROXY=http://proxy.company.com:8080",
    "HTTPS_PROXY=http://proxy.company.com:8080",
    "NO_PROXY=localhost,127.0.0.1,.local",
    "COMPOSER_DISABLE_NETWORK=false",
})
```

#### 安全
```go
comp.SetEnv([]string{
    "COMPOSER_ALLOW_SUPERUSER=1",         // 允许以 root 身份运行
    "COMPOSER_DISABLE_XDEBUG_WARN=1",     // 禁用 Xdebug 警告
    "COMPOSER_AUDIT_ABANDONED=report",    // 报告废弃的包
})
```

## 平台特定配置

### Windows 配置

```go
func configureForWindows() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // Windows 特定路径
    options.ExecutablePath = `C:\ProgramData\ComposerSetup\bin\composer.bat`
    options.WorkingDir = `C:\inetpub\wwwroot\myproject`
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // Windows 特定环境
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + os.Getenv("APPDATA") + "\\Composer",
        "COMPOSER_CACHE_DIR=" + os.Getenv("LOCALAPPDATA") + "\\Composer",
        "COMPOSER_MEMORY_LIMIT=-1",
    })
    
    return comp, nil
}
```

### macOS 配置

```go
func configureForMacOS() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // 首先尝试 Homebrew 安装
    homebrewPath := "/opt/homebrew/bin/composer"
    if _, err := os.Stat(homebrewPath); err == nil {
        options.ExecutablePath = homebrewPath
    }
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // macOS 特定环境
    homeDir, _ := os.UserHomeDir()
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + homeDir + "/.composer",
        "COMPOSER_CACHE_DIR=" + homeDir + "/Library/Caches/composer",
    })
    
    return comp, nil
}
```

### Linux 配置

```go
func configureForLinux() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // Linux 特定环境
    homeDir, _ := os.UserHomeDir()
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + homeDir + "/.config/composer",
        "COMPOSER_CACHE_DIR=" + homeDir + "/.cache/composer",
        "COMPOSER_MEMORY_LIMIT=-1",
    })
    
    return comp, nil
}
```

## 环境特定配置

### 开发环境

```go
func configureForDevelopment() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 2 * time.Minute // 开发环境较短超时
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",
        "COMPOSER_DISCARD_CHANGES=true",      // 开发环境自动丢弃
        "COMPOSER_PREFER_STABLE=false",       // 允许开发版本
        "COMPOSER_MINIMUM_STABILITY=dev",
    })
    
    return comp, nil
}
```

### 生产环境

```go
func configureForProduction() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.AutoInstall = false              // 生产环境不自动安装
    options.DefaultTimeout = 10 * time.Minute // 生产环境较长超时
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_DISCARD_CHANGES=false",     // 生产环境不自动丢弃
        "COMPOSER_PREFER_STABLE=true",        // 仅稳定版本
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",  // 性能优化
        "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
        "COMPOSER_APCU_AUTOLOADER=true",      // 如果可用则使用 APCu
    })
    
    return comp, nil
}
```

### CI/CD 环境

```go
func configureForCI() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 15 * time.Minute // CI 较长超时
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_NO_INTERACTION=1",          // 非交互模式
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=900",       // 15 分钟
        "COMPOSER_CACHE_DIR=/tmp/composer",   // 临时缓存
        "COMPOSER_PREFER_STABLE=true",
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",
        "COMPOSER_AUDIT_ABANDONED=report",    // 报告废弃的包
    })
    
    return comp, nil
}
```

## 最佳实践

### 1. 环境检测

```go
func createComposerForEnvironment() (*composer.Composer, error) {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }
    
    switch env {
    case "production":
        return configureForProduction()
    case "testing", "ci":
        return configureForCI()
    default:
        return configureForDevelopment()
    }
}
```

### 2. 从文件配置

```go
type Config struct {
    ComposerPath    string            `json:"composer_path"`
    WorkingDir      string            `json:"working_dir"`
    Timeout         int               `json:"timeout_minutes"`
    Environment     map[string]string `json:"environment"`
    AutoInstall     bool              `json:"auto_install"`
}

func loadConfigFromFile(configPath string) (*composer.Composer, error) {
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, fmt.Errorf("读取配置文件失败: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("解析配置失败: %w", err)
    }
    
    options := composer.Options{
        ExecutablePath:  config.ComposerPath,
        WorkingDir:      config.WorkingDir,
        AutoInstall:     config.AutoInstall,
        DefaultTimeout:  time.Duration(config.Timeout) * time.Minute,
    }
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // 设置环境变量
    var envVars []string
    for key, value := range config.Environment {
        envVars = append(envVars, fmt.Sprintf("%s=%s", key, value))
    }
    comp.SetEnv(envVars)
    
    return comp, nil
}
```

## 故障排除配置

### 常见问题

1. **找不到 Composer**：检查 `ExecutablePath` 和 `AutoInstall` 设置
2. **权限被拒绝**：确保适当的文件权限和用户权限
3. **超时错误**：增加 `DefaultTimeout` 或使用更长超时的上下文
4. **内存错误**：设置 `COMPOSER_MEMORY_LIMIT=-1`
5. **网络问题**：在环境变量中配置代理设置

### 调试配置

```go
func debugConfiguration(comp *composer.Composer) {
    fmt.Printf("工作目录: %s\n", comp.GetWorkingDir())
    
    // 检查 Composer 是否可访问
    if comp.IsInstalled() {
        version, _ := comp.GetVersion()
        fmt.Printf("Composer 版本: %s\n", version)
    } else {
        fmt.Println("❌ Composer 不可访问")
    }
    
    // 运行诊断
    output, err := comp.Diagnose()
    if err != nil {
        fmt.Printf("诊断失败: %v\n", err)
    } else {
        fmt.Printf("诊断:\n%s\n", output)
    }
}
```
