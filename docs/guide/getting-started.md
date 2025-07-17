# Getting Started

Welcome to Go Composer SDK! This guide will help you get up and running quickly.

## What is Go Composer SDK?

Go Composer SDK is a comprehensive Go library that provides a complete wrapper around the PHP Composer package manager. It allows you to manage PHP project dependencies, execute Composer commands, and handle various Composer-related functionality directly from your Go applications.

## Prerequisites

Before you begin, ensure you have:

- **Go 1.21 or later** installed on your system
- **PHP** installed (required for Composer to work)
- **Composer** installed (the SDK can auto-install it if needed)

## Installation

Install the Go Composer SDK using `go get`:

```bash
go get github.com/scagogogo/go-composer-sdk
```

## Your First Program

Let's create a simple program that demonstrates the basic functionality:

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
    
    // Check if Composer is installed
    if !comp.IsInstalled() {
        fmt.Println("Composer is not installed, but the SDK can auto-install it!")
        return
    }
    
    // Get and display Composer version
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("Failed to get Composer version: %v", err)
    }
    
    fmt.Printf("✅ Composer version: %s\n", version)
    
    // Set working directory to your PHP project
    comp.SetWorkingDir("/path/to/your/php/project")
    
    // Validate the composer.json file
    err = comp.Validate()
    if err != nil {
        fmt.Printf("❌ composer.json validation failed: %v\n", err)
    } else {
        fmt.Println("✅ composer.json is valid")
    }
    
    // Show installed packages
    output, err := comp.ShowAllPackages()
    if err != nil {
        log.Printf("Failed to get package list: %v", err)
    } else {
        fmt.Println("📦 Installed packages:")
        fmt.Println(output)
    }
}
```

## Key Concepts

### 1. Composer Instance

The `Composer` struct is the main entry point for all operations. You create it using the `New()` function with configuration options:

```go
// Default configuration
comp, err := composer.New(composer.DefaultOptions())

// Custom configuration
options := composer.Options{
    WorkingDir:     "/path/to/php/project",
    AutoInstall:    true,
    DefaultTimeout: 5 * time.Minute,
}
comp, err := composer.New(options)
```

### 2. Error Handling

All methods return errors that should be properly handled:

```go
err := comp.Install(false, false)
if err != nil {
    // Handle the error appropriately
    log.Printf("Installation failed: %v", err)
    return
}
```

### 3. Working Directory

Set the working directory to point to your PHP project:

```go
comp.SetWorkingDir("/path/to/your/php/project")
```

### 4. Environment Variables

Configure Composer behavior using environment variables:

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
})
```

## Common Operations

### Installing Dependencies

```go
// Install all dependencies
err := comp.Install(false, false) // noDev=false, optimize=false

// Install without dev dependencies
err := comp.Install(true, false) // noDev=true, optimize=false

// Install with optimization
err := comp.Install(false, true) // noDev=false, optimize=true
```

### Adding Packages

```go
// Add a package
err := comp.RequirePackage("monolog/monolog", "^3.0")

// Add a dev dependency
err := comp.RequirePackage("phpunit/phpunit", "^10.0")
```

### Updating Dependencies

```go
// Update all packages
err := comp.Update(false, false) // noDev=false, optimize=false

// Update specific package
err := comp.UpdatePackage("symfony/console")
```

### Getting Package Information

```go
// Show all packages
output, err := comp.ShowAllPackages()

// Show specific package
output, err := comp.ShowPackage("symfony/console")

// Show dependency tree
output, err := comp.ShowDependencyTree("")
```

## Next Steps

Now that you have the basics down, explore more advanced features:

- [Configuration Guide](/guide/configuration) - Learn about advanced configuration options
- [API Reference](/api/) - Complete API documentation
- [Examples](/examples/) - Real-world usage examples

## Getting Help

If you encounter any issues:

1. Check the [API Reference](/api/) for detailed method documentation
2. Look at the [Examples](/examples/) for common use cases
3. Search or create an issue on [GitHub](https://github.com/scagogogo/go-composer-sdk/issues)
4. Join the [Discussions](https://github.com/scagogogo/go-composer-sdk/discussions)
