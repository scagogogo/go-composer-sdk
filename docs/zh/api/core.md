# 核心 API

核心 API 提供创建和管理 Composer 实例的基本功能。

## 包: composer

```go
import "github.com/scagogogo/go-composer-sdk/pkg/composer"
```

## 类型

### Composer

提供所有 Composer 功能的主要结构体。

```go
type Composer struct {
    // 包含过滤或未导出的字段
}
```

### Options

创建 Composer 实例的配置选项。

```go
type Options struct {
    ExecutablePath  string                // composer 可执行文件路径
    WorkingDir      string                // 操作的工作目录  
    AutoInstall     bool                  // 如果未找到则自动安装 Composer
    DefaultTimeout  time.Duration         // 操作的默认超时时间
    Detector        *detector.Detector    // 自定义检测器实例
    Installer       *installer.Installer  // 自定义安装器实例
}
```

## 函数

### New

使用指定选项创建新的 Composer 实例。

```go
func New(options Options) (*Composer, error)
```

**参数:**
- `options` - Composer 实例的配置选项

**返回值:**
- `*Composer` - 新的 Composer 实例
- `error` - 创建失败时的错误

**示例:**
```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("创建 Composer 实例失败: %v", err)
}
```

### DefaultOptions

返回默认配置选项。

```go
func DefaultOptions() Options
```

**返回值:**
- `Options` - 具有合理默认值的默认配置

**示例:**
```go
options := composer.DefaultOptions()
options.WorkingDir = "/path/to/project"
comp, err := composer.New(options)
```

## 核心方法

### IsInstalled

检查 Composer 是否已安装且可访问。

```go
func (c *Composer) IsInstalled() bool
```

**返回值:**
- `bool` - 如果 Composer 已安装且可访问则为 true

**示例:**
```go
if !comp.IsInstalled() {
    log.Fatal("Composer 未安装")
}
```

### GetVersion

获取已安装的 Composer 版本。

```go
func (c *Composer) GetVersion() (string, error)
```

**返回值:**
- `string` - Composer 版本字符串
- `error` - 无法检索版本时的错误

**示例:**
```go
version, err := comp.GetVersion()
if err != nil {
    log.Printf("获取版本失败: %v", err)
    return
}
fmt.Printf("Composer 版本: %s\n", version)
```

### Run

使用给定参数执行原始 Composer 命令。

```go
func (c *Composer) Run(args ...string) (string, error)
```

**参数:**
- `args` - 传递给 Composer 的命令参数

**返回值:**
- `string` - 命令输出
- `error` - 命令失败时的错误

**示例:**
```go
output, err := comp.Run("--version")
if err != nil {
    log.Printf("命令失败: %v", err)
    return
}
fmt.Println(output)
```

### RunWithContext

使用上下文支持执行 Composer 命令，支持取消和超时。

```go
func (c *Composer) RunWithContext(ctx context.Context, args ...string) (string, error)
```

**参数:**
- `ctx` - 用于取消和超时的上下文
- `args` - 传递给 Composer 的命令参数

**返回值:**
- `string` - 命令输出
- `error` - 命令失败或上下文被取消时的错误

**示例:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

output, err := comp.RunWithContext(ctx, "install")
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("命令超时")
    } else {
        log.Printf("命令失败: %v", err)
    }
    return
}
```

## 配置方法

### SetWorkingDir

设置 Composer 操作的工作目录。

```go
func (c *Composer) SetWorkingDir(dir string)
```

**参数:**
- `dir` - 工作目录的路径

**示例:**
```go
comp.SetWorkingDir("/path/to/php/project")
```

### GetWorkingDir

获取当前工作目录。

```go
func (c *Composer) GetWorkingDir() string
```

**返回值:**
- `string` - 当前工作目录路径

**示例:**
```go
workDir := comp.GetWorkingDir()
fmt.Printf("工作目录: %s\n", workDir)
```

### SetEnv

设置 Composer 操作的环境变量。

```go
func (c *Composer) SetEnv(env []string)
```

**参数:**
- `env` - "KEY=VALUE" 格式的环境变量数组

**示例:**
```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
    "COMPOSER_CACHE_DIR=/tmp/composer-cache",
})
```

### SetTimeout

设置操作的默认超时时间。

```go
func (c *Composer) SetTimeout(timeout time.Duration)
```

**参数:**
- `timeout` - 默认超时持续时间

**示例:**
```go
comp.SetTimeout(5 * time.Minute)
```

## 实用方法

### SelfUpdate

将 Composer 更新到最新版本。

```go
func (c *Composer) SelfUpdate() error
```

**返回值:**
- `error` - 更新失败时的错误

**示例:**
```go
err := comp.SelfUpdate()
if err != nil {
    log.Printf("更新 Composer 失败: %v", err)
}
```

### ClearCache

清除 Composer 缓存。

```go
func (c *Composer) ClearCache() error
```

**返回值:**
- `error` - 清除缓存失败时的错误

**示例:**
```go
err := comp.ClearCache()
if err != nil {
    log.Printf("清除缓存失败: %v", err)
}
```

### Diagnose

运行 Composer 的诊断检查。

```go
func (c *Composer) Diagnose() (string, error)
```

**返回值:**
- `string` - 诊断输出
- `error` - 诊断失败时的错误

**示例:**
```go
output, err := comp.Diagnose()
if err != nil {
    log.Printf("诊断失败: %v", err)
    return
}
fmt.Println("诊断结果:")
fmt.Println(output)
```

## 错误处理

所有可能失败的方法都将错误作为最后一个返回值。始终检查并适当处理错误：

```go
// 良好的错误处理
version, err := comp.GetVersion()
if err != nil {
    log.Printf("获取版本失败: %v", err)
    return
}

// 使用版本
fmt.Printf("Composer 版本: %s\n", version)
```

## 上下文使用

对于长时间运行的操作，使用支持上下文的方法来启用取消和超时：

```go
// 创建带超时的上下文
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

// 使用支持上下文的方法
output, err := comp.RunWithContext(ctx, "install", "--no-dev")
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("操作超时")
    } else {
        log.Printf("操作失败: %v", err)
    }
}
```

## 最佳实践

1. **在执行操作前始终检查 Composer 是否已安装**
2. **将工作目录设置为您的 PHP 项目根目录**
3. **对长时间运行的操作使用支持上下文的方法**
4. **适当处理错误** - 不要忽略它们
5. **根据您的用例配置环境变量**
6. **设置合理的超时时间**以防止操作挂起
