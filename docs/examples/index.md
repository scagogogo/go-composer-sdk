# Examples

This section provides practical examples and use cases for the Go Composer SDK. Each example includes complete, runnable code that demonstrates real-world usage patterns.

## Quick Examples

### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create Composer instance
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
    
    fmt.Println("Composer is ready!")
}
```

### Installing Dependencies

```go
func installDependencies(comp *composer.Composer) error {
    fmt.Println("Installing dependencies...")
    
    // Install all dependencies
    err := comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        return fmt.Errorf("failed to install dependencies: %w", err)
    }
    
    fmt.Println("Dependencies installed successfully!")
    return nil
}
```

### Adding a Package

```go
func addPackage(comp *composer.Composer) error {
    packageName := "monolog/monolog"
    version := "^3.0"
    
    fmt.Printf("Adding package %s %s...\n", packageName, version)
    
    err := comp.RequirePackage(packageName, version)
    if err != nil {
        return fmt.Errorf("failed to add package: %w", err)
    }
    
    fmt.Printf("Package %s added successfully!\n", packageName)
    return nil
}
```

## Complete Examples

### 1. Project Setup Automation

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    projectDir := "/path/to/new/php/project"
    
    // Create project directory
    err := os.MkdirAll(projectDir, 0755)
    if err != nil {
        log.Fatalf("Failed to create project directory: %v", err)
    }
    
    // Create Composer instance
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    comp.SetWorkingDir(projectDir)
    
    // Initialize new project
    initOptions := composer.InitOptions{
        Name:        "mycompany/my-project",
        Description: "My awesome PHP project",
        Type:        "project",
        License:     "MIT",
        Authors: []composer.Author{
            {
                Name:  "John Doe",
                Email: "john@example.com",
            },
        },
        MinimumStability: "stable",
        PreferStable:     true,
    }
    
    err = comp.InitProject(initOptions)
    if err != nil {
        log.Fatalf("Failed to initialize project: %v", err)
    }
    
    // Add essential packages
    packages := map[string]string{
        "symfony/console":    "^6.0",
        "monolog/monolog":    "^3.0",
        "guzzlehttp/guzzle":  "^7.0",
    }
    
    err = comp.RequirePackages(packages)
    if err != nil {
        log.Fatalf("Failed to add packages: %v", err)
    }
    
    // Add development dependencies
    devPackages := map[string]string{
        "phpunit/phpunit":       "^10.0",
        "symfony/var-dumper":    "^6.0",
        "friendsofphp/php-cs-fixer": "^3.0",
    }
    
    for pkg, version := range devPackages {
        err = comp.RequireDevPackage(pkg, version)
        if err != nil {
            log.Printf("Warning: Failed to add dev package %s: %v", pkg, err)
        }
    }
    
    // Generate optimized autoloader
    err = comp.DumpAutoloadOptimized()
    if err != nil {
        log.Printf("Warning: Failed to generate optimized autoloader: %v", err)
    }
    
    fmt.Println("Project setup completed successfully!")
}
```

### 2. Dependency Management

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func manageDependencies() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    // Check for outdated packages
    fmt.Println("Checking for outdated packages...")
    outdated, err := comp.ShowOutdated()
    if err != nil {
        log.Printf("Failed to check outdated packages: %v", err)
    } else if outdated != "" {
        fmt.Printf("Outdated packages found:\n%s\n", outdated)
        
        // Update with timeout
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
        defer cancel()
        
        fmt.Println("Updating packages...")
        err = comp.UpdateWithContext(ctx, false, true) // noDev=false, optimize=true
        if err != nil {
            log.Printf("Failed to update packages: %v", err)
        } else {
            fmt.Println("Packages updated successfully!")
        }
    } else {
        fmt.Println("All packages are up to date!")
    }
    
    // Show dependency tree
    fmt.Println("\nDependency tree:")
    tree, err := comp.ShowDependencyTree("")
    if err != nil {
        log.Printf("Failed to show dependency tree: %v", err)
    } else {
        fmt.Println(tree)
    }
}
```

### 3. Security Audit

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func performSecurityAudit() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    comp.SetWorkingDir("/path/to/php/project")
    
    // Perform security audit
    fmt.Println("Performing security audit...")
    auditResult, err := comp.Audit()
    if err != nil {
        log.Fatalf("Security audit failed: %v", err)
    }
    
    // Parse audit results
    var result composer.AuditResult
    err = json.Unmarshal([]byte(auditResult), &result)
    if err != nil {
        log.Printf("Failed to parse audit results: %v", err)
        fmt.Printf("Raw audit output:\n%s\n", auditResult)
        return
    }
    
    // Display results
    if result.Found == 0 {
        fmt.Println("✅ No security vulnerabilities found!")
    } else {
        fmt.Printf("⚠️  Found %d security vulnerabilities:\n\n", result.Found)
        
        for _, vuln := range result.Vulnerabilities {
            fmt.Printf("🔴 %s\n", vuln.Title)
            fmt.Printf("   Package: %s\n", vuln.Package)
            fmt.Printf("   Version: %s\n", vuln.Version)
            fmt.Printf("   Severity: %s\n", vuln.Severity)
            fmt.Printf("   Description: %s\n", vuln.Description)
            if len(vuln.References) > 0 {
                fmt.Printf("   References:\n")
                for _, ref := range vuln.References {
                    fmt.Printf("     - %s\n", ref)
                }
            }
            fmt.Println()
        }
        
        fmt.Println("Please update the affected packages to resolve these vulnerabilities.")
    }
}
```

### 4. CI/CD Integration

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func cicdPipeline() {
    // Check if we're in CI environment
    isCI := os.Getenv("CI") == "true"
    
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Set working directory
    workDir := os.Getenv("GITHUB_WORKSPACE")
    if workDir == "" {
        workDir = "."
    }
    comp.SetWorkingDir(workDir)
    
    // Configure for CI environment
    if isCI {
        comp.SetEnv([]string{
            "COMPOSER_NO_INTERACTION=1",
            "COMPOSER_PREFER_STABLE=1",
            "COMPOSER_MEMORY_LIMIT=-1",
        })
    }
    
    // Validate composer.json
    fmt.Println("Validating composer.json...")
    err = comp.Validate()
    if err != nil {
        log.Fatalf("composer.json validation failed: %v", err)
    }
    fmt.Println("✅ composer.json is valid")
    
    // Install dependencies with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()
    
    fmt.Println("Installing dependencies...")
    installOptions := composer.InstallOptions{
        NoDev:        false,
        Optimize:     true,
        PreferDist:   true,
        NoProgress:   isCI,
        NoSuggest:    true,
        PreferStable: true,
    }
    
    err = comp.InstallWithOptions(installOptions)
    if err != nil {
        log.Fatalf("Failed to install dependencies: %v", err)
    }
    fmt.Println("✅ Dependencies installed")
    
    // Security audit
    fmt.Println("Running security audit...")
    auditResult, err := comp.Audit()
    if err != nil {
        log.Printf("Security audit failed: %v", err)
    } else {
        fmt.Println("✅ Security audit completed")
        
        // In CI, you might want to fail the build if vulnerabilities are found
        if isCI {
            var result composer.AuditResult
            if json.Unmarshal([]byte(auditResult), &result) == nil && result.Found > 0 {
                log.Fatalf("Security vulnerabilities found: %d", result.Found)
            }
        }
    }
    
    // Check platform requirements
    fmt.Println("Checking platform requirements...")
    err = comp.CheckPlatformReqs()
    if err != nil {
        log.Printf("Platform requirements check failed: %v", err)
    } else {
        fmt.Println("✅ Platform requirements satisfied")
    }
    
    fmt.Println("CI/CD pipeline completed successfully!")
}
```

## Example Categories

### [Basic Operations](/examples/basic-operations)
Simple examples covering the most common use cases like installing packages, updating dependencies, and basic project management.

### [Package Management](/examples/package-management)
Advanced package management scenarios including batch operations, dependency analysis, and conflict resolution.

### [Project Setup](/examples/project-setup)
Complete project initialization and setup automation examples for different types of PHP projects.

### [Security Audit](/examples/security-audit)
Security-focused examples showing how to integrate vulnerability scanning and dependency auditing into your workflows.

### [Advanced Usage](/examples/advanced-usage)
Complex scenarios including custom configurations, error handling patterns, and integration with other tools.

## Running the Examples

1. **Install the SDK**:
   ```bash
   go get github.com/scagogogo/go-composer-sdk
   ```

2. **Update paths**: Replace `/path/to/php/project` with your actual PHP project path

3. **Run the examples**:
   ```bash
   go run example.go
   ```

## Tips for Using Examples

- **Always handle errors**: The examples show proper error handling patterns
- **Use context for timeouts**: Long-running operations should use context
- **Configure for your environment**: Adjust paths and options as needed
- **Test in development first**: Try examples in a test environment before production
- **Check prerequisites**: Ensure PHP and Composer requirements are met

## Contributing Examples

Have a useful example? We'd love to include it! Please:

1. Fork the repository
2. Add your example to the appropriate section
3. Include clear documentation and comments
4. Test the example thoroughly
5. Submit a pull request

For more information, see our [Contributing Guide](https://github.com/scagogogo/go-composer-sdk/blob/main/CONTRIBUTING.md).
