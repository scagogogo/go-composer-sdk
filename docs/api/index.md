# API Reference

Welcome to the Go Composer SDK API Reference. This documentation provides comprehensive details about all available packages, types, and functions.

## Package Overview

The Go Composer SDK is organized into several packages, each serving specific purposes:

### Core Packages

| Package | Description |
|---------|-------------|
| [`composer`](/api/core) | Main package containing the core Composer functionality |
| [`detector`](/api/detector) | Composer installation detection and validation |
| [`installer`](/api/installer) | Automatic Composer installation utilities |
| [`utils`](/api/utilities) | Common utilities and helper functions |

## Main Types

### Composer

The primary interface for all Composer operations:

```go
type Composer struct {
    // Contains filtered or unexported fields
}
```

**Key Methods:**
- Package Management: `Install()`, `Update()`, `RequirePackage()`, `RemovePackage()`
- Project Operations: `CreateProject()`, `Validate()`, `RunScript()`
- Information: `GetVersion()`, `ShowPackage()`, `GetLicenses()`
- Security: `Audit()`, `CheckPlatformReqs()`

### Options

Configuration options for creating a Composer instance:

```go
type Options struct {
    ExecutablePath  string        // Path to composer executable
    WorkingDir      string        // Working directory for operations
    AutoInstall     bool          // Auto-install Composer if not found
    DefaultTimeout  time.Duration // Default timeout for operations
    Detector        *detector.Detector // Custom detector instance
    Installer       *installer.Installer // Custom installer instance
}
```

### ComposerJSON

Represents the structure of a composer.json file:

```go
type ComposerJSON struct {
    Name         string                 `json:"name,omitempty"`
    Description  string                 `json:"description,omitempty"`
    Type         string                 `json:"type,omitempty"`
    License      interface{}            `json:"license,omitempty"`
    Authors      []Author               `json:"authors,omitempty"`
    Require      map[string]string      `json:"require,omitempty"`
    RequireDev   map[string]string      `json:"require-dev,omitempty"`
    Autoload     map[string]interface{} `json:"autoload,omitempty"`
    AutoloadDev  map[string]interface{} `json:"autoload-dev,omitempty"`
    Scripts      map[string]interface{} `json:"scripts,omitempty"`
    Config       map[string]interface{} `json:"config,omitempty"`
    Repositories []Repository           `json:"repositories,omitempty"`
    Extra        map[string]interface{} `json:"extra,omitempty"`
}
```

## Quick Reference

### Creating a Composer Instance

```go
// Default options
comp, err := composer.New(composer.DefaultOptions())

// Custom options
options := composer.Options{
    WorkingDir:     "/path/to/project",
    AutoInstall:    true,
    DefaultTimeout: 5 * time.Minute,
}
comp, err := composer.New(options)
```

### Common Operations

```go
// Check installation
isInstalled := comp.IsInstalled()

// Get version
version, err := comp.GetVersion()

// Install dependencies
err = comp.Install(false, false) // noDev, optimize

// Add package
err = comp.RequirePackage("monolog/monolog", "^3.0")

// Update packages
err = comp.Update(false, false)

// Show package info
info, err := comp.ShowPackage("symfony/console")

// Validate composer.json
err = comp.Validate()

// Run security audit
result, err := comp.Audit()
```

### Error Handling

All methods that can fail return an error as the last return value:

```go
version, err := comp.GetVersion()
if err != nil {
    log.Printf("Failed to get version: %v", err)
    return
}
fmt.Printf("Composer version: %s\n", version)
```

### Context Support

Many operations support context for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := comp.InstallWithContext(ctx, false, false)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Installation timed out")
    } else {
        log.Printf("Installation failed: %v", err)
    }
}
```

## API Categories

### [Core Operations](/api/core)
Basic Composer functionality including instance creation, version management, and command execution.

### [Package Management](/api/package-management)
Installing, updating, requiring, and removing packages. Managing dependencies and package information.

### [Project Management](/api/project-management)
Creating projects, running scripts, validating configurations, and managing project-level settings.

### [Security & Audit](/api/security-audit)
Security auditing, vulnerability detection, and dependency analysis.

### [Platform & Environment](/api/platform-environment)
Platform requirements checking, environment configuration, and system compatibility.

### [Utilities](/api/utilities)
Helper functions for file operations, HTTP requests, and cross-platform compatibility.

### [Detector](/api/detector)
Composer installation detection and path resolution.

### [Installer](/api/installer)
Automatic Composer installation and setup.

## Best Practices

### 1. Always Handle Errors

```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("Failed to create Composer instance: %v", err)
}
```

### 2. Use Context for Long Operations

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

err := comp.UpdateWithContext(ctx, false, false)
```

### 3. Set Working Directory

```go
comp.SetWorkingDir("/path/to/your/php/project")
```

### 4. Configure Environment Variables

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
})
```

### 5. Check Installation Before Operations

```go
if !comp.IsInstalled() {
    log.Fatal("Composer is not installed")
}
```

## Examples

For practical examples and use cases, see the [Examples section](/examples/).

## Type Definitions

### Author

```go
type Author struct {
    Name     string `json:"name,omitempty"`
    Email    string `json:"email,omitempty"`
    Homepage string `json:"homepage,omitempty"`
    Role     string `json:"role,omitempty"`
}
```

### Repository

```go
type Repository struct {
    Type string `json:"type"`
    URL  string `json:"url"`
}
```

### AuditResult

```go
type AuditResult struct {
    Vulnerabilities []Vulnerability `json:"vulnerabilities"`
    Found          int             `json:"found"`
}

type Vulnerability struct {
    ID          string   `json:"id"`
    Title       string   `json:"title"`
    Package     string   `json:"package"`
    Version     string   `json:"version"`
    Severity    string   `json:"severity"`
    Description string   `json:"description"`
    References  []string `json:"references"`
}
```

## Support

- [GitHub Repository](https://github.com/scagogogo/go-composer-sdk)
- [Issue Tracker](https://github.com/scagogogo/go-composer-sdk/issues)
- [Discussions](https://github.com/scagogogo/go-composer-sdk/discussions)
