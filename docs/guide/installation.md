# Installation

This guide covers different ways to install and set up Go Composer SDK in your project.

## System Requirements

### Go Version
- **Go 1.21 or later** is required
- The SDK uses modern Go features and follows current best practices

### PHP and Composer
- **PHP 7.4 or later** (required by Composer)
- **Composer 2.0 or later** (can be auto-installed by the SDK)

### Operating Systems
- **Windows** (Windows 10/11, Windows Server 2016+)
- **macOS** (macOS 10.15+)
- **Linux** (Ubuntu 18.04+, CentOS 7+, Debian 9+, and other distributions)

## Installation Methods

### Method 1: Using go get (Recommended)

The simplest way to install Go Composer SDK:

```bash
go get github.com/scagogogo/go-composer-sdk
```

This will download and install the latest version of the SDK and all its dependencies.

### Method 2: Using go mod

If you're using Go modules (recommended for new projects):

1. Initialize your Go module (if not already done):
```bash
go mod init your-project-name
```

2. Add the SDK to your project:
```bash
go get github.com/scagogogo/go-composer-sdk
```

3. Import in your Go code:
```go
import "github.com/scagogogo/go-composer-sdk/pkg/composer"
```

### Method 3: Specific Version

To install a specific version:

```bash
go get github.com/scagogogo/go-composer-sdk@v1.0.0
```

## Verifying Installation

Create a simple test file to verify the installation:

```go
// test_installation.go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Try to create a Composer instance
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    fmt.Println("✅ Go Composer SDK installed successfully!")
    
    // Check if Composer is available
    if comp.IsInstalled() {
        version, err := comp.GetVersion()
        if err == nil {
            fmt.Printf("✅ Composer is available: %s\n", version)
        }
    } else {
        fmt.Println("ℹ️  Composer not found, but SDK can auto-install it")
    }
}
```

Run the test:

```bash
go run test_installation.go
```

## Setting Up Composer

The Go Composer SDK can automatically detect and install Composer if it's not already available on your system.

### Automatic Installation

The SDK will automatically install Composer when you create a new instance with `AutoInstall: true` (which is the default):

```go
options := composer.DefaultOptions() // AutoInstall is true by default
comp, err := composer.New(options)
if err != nil {
    log.Fatalf("Failed to create Composer instance: %v", err)
}
```

### Manual Composer Installation

If you prefer to install Composer manually:

#### On Windows:
1. Download the Composer installer from [getcomposer.org](https://getcomposer.org/download/)
2. Run the installer and follow the instructions
3. Verify installation: `composer --version`

#### On macOS:
```bash
# Using Homebrew (recommended)
brew install composer

# Or using the installer
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer
```

#### On Linux:
```bash
# Download and install
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer

# Make it executable
sudo chmod +x /usr/local/bin/composer

# Verify installation
composer --version
```

## Configuration

### Environment Variables

You can configure Composer behavior using environment variables:

```bash
# Increase memory limit
export COMPOSER_MEMORY_LIMIT=-1

# Set custom home directory
export COMPOSER_HOME=/path/to/composer/home

# Configure HTTP proxy
export HTTP_PROXY=http://proxy.example.com:8080
export HTTPS_PROXY=http://proxy.example.com:8080
```

### Custom Composer Path

If Composer is installed in a non-standard location:

```go
options := composer.Options{
    ExecutablePath: "/custom/path/to/composer",
    WorkingDir:     "/path/to/php/project",
    AutoInstall:    false, // Don't auto-install since we have a custom path
}
comp, err := composer.New(options)
```

## Troubleshooting

### Common Issues

#### 1. "composer not found" Error

**Solution**: Enable auto-installation or install Composer manually:

```go
options := composer.DefaultOptions()
options.AutoInstall = true // Enable auto-installation
comp, err := composer.New(options)
```

#### 2. Permission Denied Errors

**Solution**: Ensure proper permissions or use sudo/administrator privileges:

```bash
# On Linux/macOS
sudo chown -R $USER:$USER ~/.composer

# On Windows, run as Administrator
```

#### 3. Network/Proxy Issues

**Solution**: Configure proxy settings:

```go
comp.SetEnv([]string{
    "HTTP_PROXY=http://your-proxy:8080",
    "HTTPS_PROXY=http://your-proxy:8080",
    "NO_PROXY=localhost,127.0.0.1",
})
```

#### 4. PHP Not Found

**Solution**: Ensure PHP is installed and in your PATH:

```bash
# Check PHP installation
php --version

# Add PHP to PATH if needed (example for Windows)
set PATH=%PATH%;C:\php
```

### Getting Help

If you encounter issues not covered here:

1. Check the [GitHub Issues](https://github.com/scagogogo/go-composer-sdk/issues)
2. Create a new issue with:
   - Your operating system and version
   - Go version (`go version`)
   - PHP version (`php --version`)
   - Composer version (`composer --version`)
   - Complete error message
   - Minimal code example that reproduces the issue

## Next Steps

Once you have Go Composer SDK installed:

1. Read the [Getting Started Guide](/guide/getting-started)
2. Explore the [Configuration Options](/guide/configuration)
3. Check out [Basic Usage Examples](/guide/basic-usage)
4. Browse the complete [API Reference](/api/)
