# 安装

本指南涵盖了在项目中安装和设置 Go Composer SDK 的不同方法。

## 系统要求

### Go 版本
- **Go 1.21 或更高版本** 是必需的
- SDK 使用现代 Go 功能并遵循当前最佳实践

### PHP 和 Composer
- **PHP 7.4 或更高版本**（Composer 所需）
- **Composer 2.0 或更高版本**（SDK 可以自动安装）

### 操作系统
- **Windows**（Windows 10/11，Windows Server 2016+）
- **macOS**（macOS 10.15+）
- **Linux**（Ubuntu 18.04+，CentOS 7+，Debian 9+ 和其他发行版）

## 安装方法

### 方法 1：使用 go get（推荐）

安装 Go Composer SDK 的最简单方法：

```bash
go get github.com/scagogogo/go-composer-sdk
```

这将下载并安装最新版本的 SDK 及其所有依赖项。

### 方法 2：使用 go mod

如果您使用 Go 模块（推荐用于新项目）：

1. 初始化您的 Go 模块（如果尚未完成）：
```bash
go mod init your-project-name
```

2. 将 SDK 添加到您的项目：
```bash
go get github.com/scagogogo/go-composer-sdk
```

3. 在您的 Go 代码中导入：
```go
import "github.com/scagogogo/go-composer-sdk/pkg/composer"
```

### 方法 3：特定版本

要安装特定版本：

```bash
go get github.com/scagogogo/go-composer-sdk@v1.0.0
```

## 验证安装

创建一个简单的测试文件来验证安装：

```go
// test_installation.go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // 尝试创建 Composer 实例
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("创建 Composer 实例失败: %v", err)
    }
    
    fmt.Println("✅ Go Composer SDK 安装成功！")
    
    // 检查 Composer 是否可用
    if comp.IsInstalled() {
        version, err := comp.GetVersion()
        if err == nil {
            fmt.Printf("✅ Composer 可用: %s\n", version)
        }
    } else {
        fmt.Println("ℹ️  未找到 Composer，但 SDK 可以自动安装")
    }
}
```

运行测试：

```bash
go run test_installation.go
```

## 设置 Composer

Go Composer SDK 可以在系统上没有 Composer 时自动检测并安装 Composer。

### 自动安装

SDK 将在您使用 `AutoInstall: true`（默认值）创建新实例时自动安装 Composer：

```go
options := composer.DefaultOptions() // AutoInstall 默认为 true
comp, err := composer.New(options)
if err != nil {
    log.Fatalf("创建 Composer 实例失败: %v", err)
}
```

### 手动 Composer 安装

如果您更喜欢手动安装 Composer：

#### 在 Windows 上：
1. 从 [getcomposer.org](https://getcomposer.org/download/) 下载 Composer 安装程序
2. 运行安装程序并按照说明操作
3. 验证安装：`composer --version`

#### 在 macOS 上：
```bash
# 使用 Homebrew（推荐）
brew install composer

# 或使用安装程序
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer
```

#### 在 Linux 上：
```bash
# 下载并安装
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer

# 使其可执行
sudo chmod +x /usr/local/bin/composer

# 验证安装
composer --version
```

## 配置

### 环境变量

您可以使用环境变量配置 Composer 行为：

```bash
# 增加内存限制
export COMPOSER_MEMORY_LIMIT=-1

# 设置自定义主目录
export COMPOSER_HOME=/path/to/composer/home

# 配置 HTTP 代理
export HTTP_PROXY=http://proxy.example.com:8080
export HTTPS_PROXY=http://proxy.example.com:8080
```

### 自定义 Composer 路径

如果 Composer 安装在非标准位置：

```go
options := composer.Options{
    ExecutablePath: "/custom/path/to/composer",
    WorkingDir:     "/path/to/php/project",
    AutoInstall:    false, // 不自动安装，因为我们有自定义路径
}
comp, err := composer.New(options)
```

## 故障排除

### 常见问题

#### 1. "composer not found" 错误

**解决方案**：启用自动安装或手动安装 Composer：

```go
options := composer.DefaultOptions()
options.AutoInstall = true // 启用自动安装
comp, err := composer.New(options)
```

#### 2. 权限被拒绝错误

**解决方案**：确保适当的权限或使用 sudo/管理员权限：

```bash
# 在 Linux/macOS 上
sudo chown -R $USER:$USER ~/.composer

# 在 Windows 上，以管理员身份运行
```

#### 3. 网络/代理问题

**解决方案**：配置代理设置：

```go
comp.SetEnv([]string{
    "HTTP_PROXY=http://your-proxy:8080",
    "HTTPS_PROXY=http://your-proxy:8080",
    "NO_PROXY=localhost,127.0.0.1",
})
```

#### 4. 未找到 PHP

**解决方案**：确保 PHP 已安装并在您的 PATH 中：

```bash
# 检查 PHP 安装
php --version

# 如果需要，将 PHP 添加到 PATH（Windows 示例）
set PATH=%PATH%;C:\php
```

### 获取帮助

如果您遇到此处未涵盖的问题：

1. 检查 [GitHub Issues](https://github.com/scagogogo/go-composer-sdk/issues)
2. 创建新问题，包含：
   - 您的操作系统和版本
   - Go 版本（`go version`）
   - PHP 版本（`php --version`）
   - Composer 版本（`composer --version`）
   - 完整的错误消息
   - 重现问题的最小代码示例

## 下一步

一旦您安装了 Go Composer SDK：

1. 阅读[快速开始指南](/zh/guide/getting-started)
2. 探索[配置选项](/zh/guide/configuration)
3. 查看[基本用法示例](/zh/guide/basic-usage)
4. 浏览完整的 [API 参考](/zh/api/)
