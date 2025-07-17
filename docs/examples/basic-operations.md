# Basic Operations

This section covers the most common and fundamental operations you'll perform with the Go Composer SDK.

## Installation and Setup

### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create Composer instance with default options
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Set working directory to your PHP project
    comp.SetWorkingDir("/path/to/your/php/project")
    
    // Verify Composer is available
    if !comp.IsInstalled() {
        log.Fatal("Composer is not installed")
    }
    
    fmt.Println("✅ Composer SDK is ready!")
}
```

### Custom Configuration

```go
func setupWithCustomConfig() {
    options := composer.Options{
        WorkingDir:     "/path/to/php/project",
        AutoInstall:    true,
        DefaultTimeout: 5 * time.Minute,
    }
    
    comp, err := composer.New(options)
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Configure environment variables
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_CACHE_DIR=/tmp/composer-cache",
    })
    
    fmt.Println("✅ Custom configuration applied")
}
```

## Package Installation

### Installing All Dependencies

```go
func installDependencies() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Installing dependencies...")
    
    // Install all dependencies (including dev)
    err = comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        return fmt.Errorf("installation failed: %w", err)
    }
    
    fmt.Println("✅ Dependencies installed successfully!")
    return nil
}
```

### Production Installation

```go
func installForProduction() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Installing for production...")
    
    // Install without dev dependencies and with optimization
    err = comp.Install(true, true) // noDev=true, optimize=true
    if err != nil {
        return fmt.Errorf("production installation failed: %w", err)
    }
    
    fmt.Println("✅ Production dependencies installed!")
    return nil
}
```

### Installation with Timeout

```go
func installWithTimeout() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()
    
    fmt.Println("Installing with timeout...")
    
    err = comp.InstallWithContext(ctx, false, false)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return fmt.Errorf("installation timed out after 10 minutes")
        }
        return fmt.Errorf("installation failed: %w", err)
    }
    
    fmt.Println("✅ Installation completed within timeout!")
    return nil
}
```

## Adding Packages

### Adding a Single Package

```go
func addPackage() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packageName := "monolog/monolog"
    version := "^3.0"
    
    fmt.Printf("Adding package %s %s...\n", packageName, version)
    
    err = comp.RequirePackage(packageName, version)
    if err != nil {
        return fmt.Errorf("failed to add package: %w", err)
    }
    
    fmt.Printf("✅ Package %s added successfully!\n", packageName)
    return nil
}
```

### Adding Multiple Packages

```go
func addMultiplePackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packages := map[string]string{
        "symfony/console":    "^6.0",
        "monolog/monolog":    "^3.0",
        "guzzlehttp/guzzle":  "^7.0",
        "doctrine/orm":       "^2.14",
    }
    
    fmt.Println("Adding multiple packages...")
    
    err = comp.RequirePackages(packages)
    if err != nil {
        return fmt.Errorf("failed to add packages: %w", err)
    }
    
    fmt.Println("✅ All packages added successfully!")
    return nil
}
```

### Adding Development Dependencies

```go
func addDevDependencies() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    devPackages := []struct {
        name    string
        version string
    }{
        {"phpunit/phpunit", "^10.0"},
        {"symfony/var-dumper", "^6.0"},
        {"friendsofphp/php-cs-fixer", "^3.0"},
        {"phpstan/phpstan", "^1.0"},
    }
    
    fmt.Println("Adding development dependencies...")
    
    for _, pkg := range devPackages {
        fmt.Printf("Adding dev dependency: %s %s\n", pkg.name, pkg.version)
        err = comp.RequireDevPackage(pkg.name, pkg.version)
        if err != nil {
            log.Printf("Warning: Failed to add %s: %v", pkg.name, err)
            continue
        }
        fmt.Printf("✅ Added: %s\n", pkg.name)
    }
    
    fmt.Println("✅ Development dependencies added!")
    return nil
}
```

## Updating Packages

### Update All Packages

```go
func updateAllPackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Updating all packages...")
    
    err = comp.Update(false, false) // noDev=false, optimize=false
    if err != nil {
        return fmt.Errorf("update failed: %w", err)
    }
    
    fmt.Println("✅ All packages updated successfully!")
    return nil
}
```

### Update Specific Package

```go
func updateSpecificPackage() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packageName := "symfony/console"
    
    fmt.Printf("Updating package: %s\n", packageName)
    
    err = comp.UpdatePackage(packageName)
    if err != nil {
        return fmt.Errorf("failed to update %s: %w", packageName, err)
    }
    
    fmt.Printf("✅ Package %s updated successfully!\n", packageName)
    return nil
}
```

### Update Multiple Specific Packages

```go
func updateSpecificPackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packages := []string{
        "symfony/console",
        "monolog/monolog",
        "guzzlehttp/guzzle",
    }
    
    fmt.Printf("Updating packages: %v\n", packages)
    
    err = comp.UpdatePackages(packages)
    if err != nil {
        return fmt.Errorf("failed to update packages: %w", err)
    }
    
    fmt.Println("✅ Specified packages updated successfully!")
    return nil
}
```

## Removing Packages

### Remove Single Package

```go
func removePackage() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packageName := "old-package/deprecated"
    
    fmt.Printf("Removing package: %s\n", packageName)
    
    err = comp.RemovePackage(packageName)
    if err != nil {
        return fmt.Errorf("failed to remove package: %w", err)
    }
    
    fmt.Printf("✅ Package %s removed successfully!\n", packageName)
    return nil
}
```

### Remove Multiple Packages

```go
func removeMultiplePackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packagesToRemove := []string{
        "old-package/deprecated",
        "unused/library",
        "legacy/component",
    }
    
    fmt.Printf("Removing packages: %v\n", packagesToRemove)
    
    err = comp.RemovePackages(packagesToRemove)
    if err != nil {
        return fmt.Errorf("failed to remove packages: %w", err)
    }
    
    fmt.Println("✅ Packages removed successfully!")
    return nil
}
```

## Getting Information

### Show All Installed Packages

```go
func showInstalledPackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Getting list of installed packages...")
    
    packages, err := comp.ShowAllPackages()
    if err != nil {
        return fmt.Errorf("failed to get package list: %w", err)
    }
    
    fmt.Println("📦 Installed packages:")
    fmt.Println(packages)
    return nil
}
```

### Show Package Details

```go
func showPackageDetails() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    packageName := "symfony/console"
    
    fmt.Printf("Getting details for package: %s\n", packageName)
    
    details, err := comp.ShowPackage(packageName)
    if err != nil {
        return fmt.Errorf("failed to get package details: %w", err)
    }
    
    fmt.Printf("📋 Package details for %s:\n", packageName)
    fmt.Println(details)
    return nil
}
```

### Check for Outdated Packages

```go
func checkOutdatedPackages() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Checking for outdated packages...")
    
    outdated, err := comp.ShowOutdated()
    if err != nil {
        return fmt.Errorf("failed to check outdated packages: %w", err)
    }
    
    if outdated == "" {
        fmt.Println("✅ All packages are up to date!")
    } else {
        fmt.Println("📦 Outdated packages found:")
        fmt.Println(outdated)
    }
    
    return nil
}
```

## Validation and Diagnostics

### Validate composer.json

```go
func validateComposerJSON() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Validating composer.json...")
    
    err = comp.Validate()
    if err != nil {
        return fmt.Errorf("composer.json validation failed: %w", err)
    }
    
    fmt.Println("✅ composer.json is valid!")
    return nil
}
```

### Run Diagnostics

```go
func runDiagnostics() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return err
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    fmt.Println("Running Composer diagnostics...")
    
    output, err := comp.Diagnose()
    if err != nil {
        return fmt.Errorf("diagnostics failed: %w", err)
    }
    
    fmt.Println("🔍 Diagnostic results:")
    fmt.Println(output)
    return nil
}
```

## Complete Basic Workflow

```go
func completeBasicWorkflow() error {
    // 1. Setup
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return fmt.Errorf("setup failed: %w", err)
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    // 2. Validate environment
    if !comp.IsInstalled() {
        return fmt.Errorf("composer is not installed")
    }
    
    version, err := comp.GetVersion()
    if err != nil {
        return fmt.Errorf("failed to get composer version: %w", err)
    }
    fmt.Printf("Using Composer version: %s\n", version)
    
    // 3. Validate project
    fmt.Println("Validating composer.json...")
    err = comp.Validate()
    if err != nil {
        return fmt.Errorf("composer.json validation failed: %w", err)
    }
    
    // 4. Install dependencies
    fmt.Println("Installing dependencies...")
    err = comp.Install(false, false)
    if err != nil {
        return fmt.Errorf("installation failed: %w", err)
    }
    
    // 5. Check for outdated packages
    fmt.Println("Checking for outdated packages...")
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("Warning: Failed to check outdated packages: %v", err)
    } else if outdated != "" {
        fmt.Println("⚠️  Some packages are outdated:")
        fmt.Println(outdated)
    }
    
    // 6. Show installed packages
    fmt.Println("Listing installed packages...")
    packages, err := comp.ShowAllPackages()
    if err != nil {
        log.Printf("Warning: Failed to list packages: %v", err)
    } else {
        fmt.Println("📦 Installed packages:")
        fmt.Println(packages)
    }
    
    fmt.Println("✅ Basic workflow completed successfully!")
    return nil
}
```

## Error Handling Best Practices

```go
func robustPackageOperation() error {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return fmt.Errorf("failed to create composer instance: %w", err)
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    // Always check if Composer is available
    if !comp.IsInstalled() {
        return fmt.Errorf("composer is not installed, please install it first")
    }
    
    // Validate project before operations
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("project validation failed: %w", err)
    }
    
    // Use context for timeout control
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()
    
    // Perform operation with proper error handling
    err = comp.InstallWithContext(ctx, false, false)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return fmt.Errorf("operation timed out after 5 minutes")
        }
        return fmt.Errorf("installation failed: %w", err)
    }
    
    return nil
}
```

These basic operations form the foundation of most Composer workflows. They demonstrate proper error handling, context usage, and common patterns you'll use in your applications.
