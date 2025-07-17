# Go Composer SDK

[![Go Version](https://img.shields.io/github/go-mod/go-version/scagogogo/go-composer-sdk)](https://golang.org/)
[![License](https://img.shields.io/github/license/scagogogo/go-composer-sdk)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/go-composer-sdk)](https://goreportcard.com/report/github.com/scagogogo/go-composer-sdk)
[![Tests](https://github.com/scagogogo/go-composer-sdk/actions/workflows/test.yml/badge.svg)](https://github.com/scagogogo/go-composer-sdk/actions/workflows/test.yml)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://scagogogo.github.io/go-composer-sdk/)

A comprehensive Go library for PHP Composer package manager. This SDK provides a complete wrapper around Composer functionality, allowing you to manage PHP project dependencies, execute Composer commands, and handle various Composer-related operations directly from your Go applications.

## 📖 Documentation

**Complete documentation is available at: [https://scagogogo.github.io/go-composer-sdk/](https://scagogogo.github.io/go-composer-sdk/)**

- 📚 [API Reference](https://scagogogo.github.io/go-composer-sdk/api/)
- 🚀 [Getting Started Guide](https://scagogogo.github.io/go-composer-sdk/guide/getting-started)
- 💡 [Examples](https://scagogogo.github.io/go-composer-sdk/examples/)
- 🌍 [中文文档](https://scagogogo.github.io/go-composer-sdk/zh/)

## ✨ Features

- **🚀 Complete Composer Support**: Full support for all standard Composer CLI commands
- **🛡️ Type Safety**: Strongly typed interfaces with IDE code completion
- **🔧 Comprehensive Features**: Dependency management, repository configuration, authentication, security auditing
- **🌍 Cross-Platform**: Native support for Windows, macOS, and Linux
- **📦 Modular Design**: Well-organized code structure grouped by functionality
- **✅ Production Ready**: Thoroughly tested with 161+ tests and GitHub Actions CI/CD
- **🔒 Security**: Built-in security auditing and vulnerability detection
- **⚡ Performance**: Efficient execution with proper timeout handling and context support

## 🚀 Quick Start

### Installation

```bash
go get github.com/scagogogo/go-composer-sdk
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create a Composer instance with default options
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Set working directory to your PHP project
    comp.SetWorkingDir("/path/to/your/php/project")
    
    // Check if Composer is installed
    if !comp.IsInstalled() {
        log.Fatal("Composer is not installed")
    }
    
    // Get Composer version
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("Failed to get Composer version: %v", err)
    }
    fmt.Printf("Composer version: %s\n", version)
    
    // Install dependencies
    err = comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        log.Fatalf("Failed to install dependencies: %v", err)
    }
    
    fmt.Println("Dependencies installed successfully!")
}
```

## 📋 Requirements

- **Go 1.21 or later**
- **PHP 7.4 or later** (required by Composer)
- **Composer 2.0 or later** (can be auto-installed by the SDK)

## 🔧 Core Functionality

### Package Management
```go
// Install dependencies
err := comp.Install(false, false)

// Add a package
err = comp.RequirePackage("monolog/monolog", "^3.0")

// Update packages
err = comp.Update(false, false)

// Remove a package
err = comp.RemovePackage("old-package/deprecated")
```

### Project Management
```go
// Create a new project
err := comp.CreateProject("laravel/laravel", "my-app", "")

// Validate composer.json
err = comp.Validate()

// Run scripts
err = comp.RunScript("test")
```

### Security & Audit
```go
// Perform security audit
auditResult, err := comp.Audit()

// Check platform requirements
err = comp.CheckPlatformReqs()
```

### Information & Analysis
```go
// Show package information
info, err := comp.ShowPackage("symfony/console")

// Show dependency tree
tree, err := comp.ShowDependencyTree("")

// Check outdated packages
outdated, err := comp.ShowOutdated()
```

## 🏗️ Architecture

The SDK is organized into several packages:

- **`composer`** - Main package with core Composer functionality
- **`detector`** - Composer installation detection and validation
- **`installer`** - Automatic Composer installation utilities
- **`utils`** - Common utilities and helper functions

## 🧪 Testing

The project includes comprehensive tests with 161+ test cases covering:

- Unit tests for all major functionality
- Integration tests for real-world scenarios
- Cross-platform compatibility tests
- Error handling and edge cases

Run tests:
```bash
go test ./...
```

Run tests with race detection:
```bash
go test -race ./...
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/go-composer-sdk.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Support

- 📖 [Documentation](https://scagogogo.github.io/go-composer-sdk/)
- 🐛 [Issue Tracker](https://github.com/scagogogo/go-composer-sdk/issues)
- 💬 [Discussions](https://github.com/scagogogo/go-composer-sdk/discussions)

## 🙏 Acknowledgments

- [Composer](https://getcomposer.org/) - The PHP package manager that this SDK wraps
- [Go Community](https://golang.org/community/) - For the amazing language and ecosystem
- All [contributors](https://github.com/scagogogo/go-composer-sdk/contributors) who help improve this project

---

**Languages**: [English](README.md) | [简体中文](README.zh.md)

Made with ❤️ by the Go Composer SDK team
