# Basic Usage

This guide covers the fundamental usage patterns and common workflows with the Go Composer SDK.

## Quick Start

### Creating a Composer Instance

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create with default options
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Set your project directory
    comp.SetWorkingDir("/path/to/your/php/project")
    
    fmt.Println("Composer SDK is ready!")
}
```

### Basic Operations

```go
// Check if Composer is installed
if !comp.IsInstalled() {
    log.Fatal("Composer is not installed")
}

// Get version
version, err := comp.GetVersion()
if err != nil {
    log.Printf("Failed to get version: %v", err)
} else {
    fmt.Printf("Composer version: %s\n", version)
}

// Install dependencies
err = comp.Install(false, false) // noDev=false, optimize=false
if err != nil {
    log.Printf("Installation failed: %v", err)
}
```

## Common Workflows

### Project Setup Workflow

```go
func setupNewProject(projectPath string) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir(projectPath)
    
    // 1. Initialize project if needed
    if _, err := os.Stat(filepath.Join(projectPath, "composer.json")); os.IsNotExist(err) {
        initOptions := composer.InitOptions{
            Name:        "mycompany/my-project",
            Description: "My PHP project",
            Type:        "project",
            License:     "MIT",
        }
        
        if err := comp.InitProject(initOptions); err != nil {
            return fmt.Errorf("failed to initialize project: %w", err)
        }
    }
    
    // 2. Add essential packages
    packages := map[string]string{
        "symfony/console": "^6.0",
        "monolog/monolog": "^3.0",
    }
    
    if err := comp.RequirePackages(packages); err != nil {
        return fmt.Errorf("failed to add packages: %w", err)
    }
    
    // 3. Install dependencies
    if err := comp.Install(false, true); err != nil { // optimize=true
        return fmt.Errorf("failed to install dependencies: %w", err)
    }
    
    fmt.Println("✅ Project setup completed!")
    return nil
}
```

### Maintenance Workflow

```go
func maintainProject(projectPath string) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir(projectPath)
    
    // 1. Validate composer.json
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // 2. Check for outdated packages
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("Warning: Failed to check outdated packages: %v", err)
    } else if outdated != "" {
        fmt.Println("📦 Outdated packages found:")
        fmt.Println(outdated)
        
        // Optionally update
        fmt.Println("Updating packages...")
        if err := comp.Update(false, true); err != nil {
            log.Printf("Update failed: %v", err)
        }
    }
    
    // 3. Security audit
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("Security audit failed: %v", err)
    } else {
        fmt.Println("🔒 Security audit completed")
        if auditResult != "" {
            fmt.Println(auditResult)
        }
    }
    
    // 4. Clean up
    if err := comp.ClearCache(); err != nil {
        log.Printf("Cache cleanup failed: %v", err)
    }
    
    return nil
}
```

## Working with Dependencies

### Adding Dependencies

```go
// Add a single package
err := comp.RequirePackage("guzzlehttp/guzzle", "^7.0")

// Add multiple packages
packages := map[string]string{
    "symfony/http-foundation": "^6.0",
    "doctrine/orm":           "^2.14",
}
err = comp.RequirePackages(packages)

// Add development dependency
err = comp.RequireDevPackage("phpunit/phpunit", "^10.0")
```

### Updating Dependencies

```go
// Update all packages
err := comp.Update(false, false)

// Update specific package
err = comp.UpdatePackage("symfony/console")

// Update multiple specific packages
packages := []string{"symfony/console", "monolog/monolog"}
err = comp.UpdatePackages(packages)
```

### Removing Dependencies

```go
// Remove a package
err := comp.RemovePackage("old-package/deprecated")

// Remove multiple packages
packages := []string{"package1", "package2"}
err = comp.RemovePackages(packages)
```

## Information and Analysis

### Package Information

```go
// List all installed packages
packages, err := comp.ShowAllPackages()
if err == nil {
    fmt.Println("Installed packages:")
    fmt.Println(packages)
}

// Show specific package details
details, err := comp.ShowPackage("symfony/console")
if err == nil {
    fmt.Printf("Package details:\n%s\n", details)
}

// Show dependency tree
tree, err := comp.ShowDependencyTree("")
if err == nil {
    fmt.Printf("Dependency tree:\n%s\n", tree)
}
```

### Dependency Analysis

```go
// Why is a package installed?
reasons, err := comp.WhyPackage("psr/log")
if err == nil {
    fmt.Printf("Why psr/log is installed:\n%s\n", reasons)
}

// Why can't a package be installed?
conflicts, err := comp.WhyNotPackage("symfony/console", "^7.0")
if err == nil {
    fmt.Printf("Conflicts for symfony/console ^7.0:\n%s\n", conflicts)
}
```

## Error Handling Patterns

### Basic Error Handling

```go
func handleComposerOperation() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return fmt.Errorf("failed to create composer instance: %w", err)
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // Always check if Composer is available
    if !comp.IsInstalled() {
        return fmt.Errorf("composer is not installed")
    }
    
    // Perform operation with error handling
    if err := comp.Install(false, false); err != nil {
        return fmt.Errorf("installation failed: %w", err)
    }
    
    return nil
}
```

### Robust Error Handling with Retries

```go
func robustInstall(comp *composer.Composer, maxRetries int) error {
    for attempt := 1; attempt <= maxRetries; attempt++ {
        fmt.Printf("Installation attempt %d/%d...\n", attempt, maxRetries)
        
        err := comp.Install(false, false)
        if err == nil {
            fmt.Println("✅ Installation successful!")
            return nil
        }
        
        fmt.Printf("❌ Attempt %d failed: %v\n", attempt, err)
        
        if attempt < maxRetries {
            // Wait before retry
            time.Sleep(time.Duration(attempt) * time.Second)
            
            // Clear cache before retry
            if clearErr := comp.ClearCache(); clearErr != nil {
                log.Printf("Failed to clear cache: %v", clearErr)
            }
        }
    }
    
    return fmt.Errorf("installation failed after %d attempts", maxRetries)
}
```

## Context and Timeouts

### Using Context for Cancellation

```go
func installWithCancellation() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // Create cancellable context
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Start installation in goroutine
    errChan := make(chan error, 1)
    go func() {
        errChan <- comp.InstallWithContext(ctx, false, false)
    }()
    
    // Wait for completion or user cancellation
    select {
    case err := <-errChan:
        if err != nil {
            return fmt.Errorf("installation failed: %w", err)
        }
        fmt.Println("✅ Installation completed!")
        return nil
        
    case <-time.After(30 * time.Second):
        // User decides to cancel after 30 seconds
        cancel()
        return fmt.Errorf("installation cancelled by user")
    }
}
```

### Timeout Handling

```go
func installWithTimeout(timeoutMinutes int) error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/project")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(
        context.Background(),
        time.Duration(timeoutMinutes)*time.Minute,
    )
    defer cancel()
    
    err = comp.InstallWithContext(ctx, false, false)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return fmt.Errorf("installation timed out after %d minutes", timeoutMinutes)
        }
        return fmt.Errorf("installation failed: %w", err)
    }
    
    return nil
}
```

## Environment Configuration

### Development Environment

```go
func setupDevelopmentEnvironment(comp *composer.Composer) {
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",
        "COMPOSER_DISCARD_CHANGES=true",
        "COMPOSER_PREFER_STABLE=false",
    })
}
```

### Production Environment

```go
func setupProductionEnvironment(comp *composer.Composer) {
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",
        "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
        "COMPOSER_PREFER_STABLE=true",
    })
}
```

## Best Practices

### 1. Always Validate Before Operations

```go
func safeComposerOperation(comp *composer.Composer) error {
    // Check Composer availability
    if !comp.IsInstalled() {
        return fmt.Errorf("composer not available")
    }
    
    // Validate project
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("project validation failed: %w", err)
    }
    
    // Proceed with operation
    return comp.Install(false, false)
}
```

### 2. Use Appropriate Timeouts

```go
func operationWithAppropriateTimeout(comp *composer.Composer, operation string) error {
    var timeout time.Duration
    
    switch operation {
    case "install", "update":
        timeout = 10 * time.Minute
    case "require", "remove":
        timeout = 5 * time.Minute
    default:
        timeout = 2 * time.Minute
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    return comp.InstallWithContext(ctx, false, false)
}
```

### 3. Handle Different Environments

```go
func createComposerForEnvironment() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // Adjust based on environment
    if os.Getenv("CI") == "true" {
        options.DefaultTimeout = 15 * time.Minute
    } else if os.Getenv("APP_ENV") == "production" {
        options.AutoInstall = false
    }
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // Set environment-specific variables
    if os.Getenv("CI") == "true" {
        comp.SetEnv([]string{
            "COMPOSER_NO_INTERACTION=1",
            "COMPOSER_PREFER_STABLE=true",
        })
    }
    
    return comp, nil
}
```

### 4. Logging and Monitoring

```go
func monitoredOperation(comp *composer.Composer) error {
    start := time.Now()
    
    log.Printf("Starting Composer operation in: %s", comp.GetWorkingDir())
    
    err := comp.Install(false, false)
    
    duration := time.Since(start)
    if err != nil {
        log.Printf("❌ Operation failed after %v: %v", duration, err)
        return err
    }
    
    log.Printf("✅ Operation completed successfully in %v", duration)
    return nil
}
```

This covers the essential usage patterns you'll need for most Composer operations. The key is to always handle errors appropriately, use timeouts for long-running operations, and configure the environment based on your specific needs.
