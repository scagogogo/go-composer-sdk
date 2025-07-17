# Configuration

This guide covers how to configure the Go Composer SDK for different environments and use cases.

## Basic Configuration

### Default Options

The simplest way to get started is with default options:

```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("Failed to create Composer instance: %v", err)
}
```

The default options include:
- **AutoInstall**: `true` - Automatically install Composer if not found
- **DefaultTimeout**: `5 minutes` - Default timeout for operations
- **WorkingDir**: Current directory
- **ExecutablePath**: Auto-detected

### Custom Configuration

For more control, create custom options:

```go
options := composer.Options{
    ExecutablePath:  "/usr/local/bin/composer",
    WorkingDir:      "/path/to/php/project",
    AutoInstall:     false,
    DefaultTimeout:  10 * time.Minute,
}

comp, err := composer.New(options)
if err != nil {
    log.Fatalf("Failed to create Composer instance: %v", err)
}
```

## Configuration Options

### ExecutablePath

Specify a custom path to the Composer executable:

```go
options := composer.DefaultOptions()
options.ExecutablePath = "/opt/composer/composer"
comp, err := composer.New(options)
```

**Use cases:**
- Composer installed in non-standard location
- Multiple Composer versions on the system
- Containerized environments with specific paths

### WorkingDir

Set the working directory for Composer operations:

```go
options := composer.DefaultOptions()
options.WorkingDir = "/var/www/html/my-project"
comp, err := composer.New(options)

// Or set it after creation
comp.SetWorkingDir("/path/to/another/project")
```

**Important notes:**
- Must contain a valid `composer.json` file
- All Composer operations will be relative to this directory
- Can be changed at runtime using `SetWorkingDir()`

### AutoInstall

Control whether Composer should be automatically installed if not found:

```go
// Enable auto-installation (default)
options := composer.DefaultOptions()
options.AutoInstall = true

// Disable auto-installation
options.AutoInstall = false
```

**When to disable:**
- Production environments where you want explicit control
- Systems where automatic installation might fail
- When using custom Composer installations

### DefaultTimeout

Set the default timeout for Composer operations:

```go
options := composer.DefaultOptions()
options.DefaultTimeout = 15 * time.Minute // For slow networks
comp, err := composer.New(options)

// Or set it after creation
comp.SetTimeout(30 * time.Minute)
```

**Considerations:**
- Large projects may need longer timeouts
- Network conditions affect required timeout
- Can be overridden per operation using context

## Environment Variables

### Setting Environment Variables

Configure Composer behavior using environment variables:

```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",           // Unlimited memory
    "COMPOSER_PROCESS_TIMEOUT=600",       // 10 minute timeout
    "COMPOSER_CACHE_DIR=/tmp/composer",   // Custom cache directory
    "COMPOSER_HOME=/opt/composer",        // Custom home directory
    "COMPOSER_DISCARD_CHANGES=true",      // Auto-discard changes
})
```

### Common Environment Variables

#### Memory and Performance
```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",           // Remove memory limit
    "COMPOSER_PROCESS_TIMEOUT=0",         // No process timeout
    "COMPOSER_HTACCESS_PROTECT=0",        // Disable .htaccess protection
})
```

#### Caching
```go
comp.SetEnv([]string{
    "COMPOSER_CACHE_DIR=/var/cache/composer",  // Custom cache location
    "COMPOSER_CACHE_FILES_TTL=86400",          // Cache TTL in seconds
    "COMPOSER_CACHE_REPO_TTL=3600",            // Repository cache TTL
})
```

#### Network and Proxy
```go
comp.SetEnv([]string{
    "HTTP_PROXY=http://proxy.company.com:8080",
    "HTTPS_PROXY=http://proxy.company.com:8080",
    "NO_PROXY=localhost,127.0.0.1,.local",
    "COMPOSER_DISABLE_NETWORK=false",
})
```

#### Security
```go
comp.SetEnv([]string{
    "COMPOSER_ALLOW_SUPERUSER=1",         // Allow running as root
    "COMPOSER_DISABLE_XDEBUG_WARN=1",     // Disable Xdebug warnings
    "COMPOSER_AUDIT_ABANDONED=report",    // Report abandoned packages
})
```

## Platform-Specific Configuration

### Windows Configuration

```go
func configureForWindows() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // Windows-specific paths
    options.ExecutablePath = `C:\ProgramData\ComposerSetup\bin\composer.bat`
    options.WorkingDir = `C:\inetpub\wwwroot\myproject`
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // Windows-specific environment
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + os.Getenv("APPDATA") + "\\Composer",
        "COMPOSER_CACHE_DIR=" + os.Getenv("LOCALAPPDATA") + "\\Composer",
        "COMPOSER_MEMORY_LIMIT=-1",
    })
    
    return comp, nil
}
```

### macOS Configuration

```go
func configureForMacOS() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    // Try Homebrew installation first
    homebrewPath := "/opt/homebrew/bin/composer"
    if _, err := os.Stat(homebrewPath); err == nil {
        options.ExecutablePath = homebrewPath
    }
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // macOS-specific environment
    homeDir, _ := os.UserHomeDir()
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + homeDir + "/.composer",
        "COMPOSER_CACHE_DIR=" + homeDir + "/Library/Caches/composer",
    })
    
    return comp, nil
}
```

### Linux Configuration

```go
func configureForLinux() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    // Linux-specific environment
    homeDir, _ := os.UserHomeDir()
    comp.SetEnv([]string{
        "COMPOSER_HOME=" + homeDir + "/.config/composer",
        "COMPOSER_CACHE_DIR=" + homeDir + "/.cache/composer",
        "COMPOSER_MEMORY_LIMIT=-1",
    })
    
    return comp, nil
}
```

## Environment-Specific Configurations

### Development Environment

```go
func configureForDevelopment() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 2 * time.Minute // Shorter timeout for dev
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=300",
        "COMPOSER_DISCARD_CHANGES=true",      // Auto-discard in dev
        "COMPOSER_PREFER_STABLE=false",       // Allow dev versions
        "COMPOSER_MINIMUM_STABILITY=dev",
    })
    
    return comp, nil
}
```

### Production Environment

```go
func configureForProduction() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.AutoInstall = false              // Don't auto-install in prod
    options.DefaultTimeout = 10 * time.Minute // Longer timeout for prod
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=600",
        "COMPOSER_DISCARD_CHANGES=false",     // Don't auto-discard in prod
        "COMPOSER_PREFER_STABLE=true",        // Only stable versions
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",  // Optimize for performance
        "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
        "COMPOSER_APCU_AUTOLOADER=true",      // Use APCu if available
    })
    
    return comp, nil
}
```

### CI/CD Environment

```go
func configureForCI() (*composer.Composer, error) {
    options := composer.DefaultOptions()
    options.DefaultTimeout = 15 * time.Minute // Longer timeout for CI
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, err
    }
    
    comp.SetEnv([]string{
        "COMPOSER_NO_INTERACTION=1",          // Non-interactive mode
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=900",       // 15 minutes
        "COMPOSER_CACHE_DIR=/tmp/composer",   // Temporary cache
        "COMPOSER_PREFER_STABLE=true",
        "COMPOSER_OPTIMIZE_AUTOLOADER=true",
        "COMPOSER_AUDIT_ABANDONED=report",    // Report abandoned packages
    })
    
    return comp, nil
}
```

## Advanced Configuration

### Custom Detector and Installer

```go
func configureWithCustomComponents() (*composer.Composer, error) {
    // Create custom detector
    det := detector.NewDetector()
    det.AddPossiblePath("/opt/php/composer")
    det.AddPossiblePath("/usr/local/php/bin/composer")
    
    // Create custom installer
    installerConfig := installer.Config{
        DownloadURL:     "https://getcomposer.org/installer",
        InstallPath:     "/opt/composer",
        TimeoutSeconds:  600,
        UseProxy:        true,
        ProxyURL:        "http://proxy.company.com:8080",
    }
    inst := installer.NewInstaller(installerConfig)
    
    // Use custom components
    options := composer.Options{
        WorkingDir:      "/var/www/project",
        AutoInstall:     true,
        DefaultTimeout:  10 * time.Minute,
        Detector:        det,
        Installer:       inst,
    }
    
    return composer.New(options)
}
```

### Configuration Validation

```go
func validateConfiguration(comp *composer.Composer) error {
    // Check if Composer is accessible
    if !comp.IsInstalled() {
        return fmt.Errorf("composer is not installed or not accessible")
    }
    
    // Verify version
    version, err := comp.GetVersion()
    if err != nil {
        return fmt.Errorf("failed to get composer version: %w", err)
    }
    
    fmt.Printf("✅ Composer version: %s\n", version)
    
    // Check working directory
    workDir := comp.GetWorkingDir()
    composerJSON := filepath.Join(workDir, "composer.json")
    if _, err := os.Stat(composerJSON); os.IsNotExist(err) {
        return fmt.Errorf("composer.json not found in working directory: %s", workDir)
    }
    
    fmt.Printf("✅ Working directory: %s\n", workDir)
    
    // Validate composer.json
    if err := comp.Validate(); err != nil {
        return fmt.Errorf("composer.json validation failed: %w", err)
    }
    
    fmt.Println("✅ Configuration is valid")
    return nil
}
```

## Configuration Best Practices

### 1. Environment Detection

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

### 2. Configuration from File

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
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
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
    
    // Set environment variables
    var envVars []string
    for key, value := range config.Environment {
        envVars = append(envVars, fmt.Sprintf("%s=%s", key, value))
    }
    comp.SetEnv(envVars)
    
    return comp, nil
}
```

### 3. Runtime Configuration Updates

```go
func updateConfigurationAtRuntime(comp *composer.Composer) {
    // Update timeout based on operation
    comp.SetTimeout(30 * time.Minute)
    
    // Update working directory
    comp.SetWorkingDir("/path/to/different/project")
    
    // Update environment for specific operation
    comp.SetEnv([]string{
        "COMPOSER_MEMORY_LIMIT=-1",
        "COMPOSER_PROCESS_TIMEOUT=1800", // 30 minutes
    })
}
```

## Troubleshooting Configuration

### Common Issues

1. **Composer not found**: Check `ExecutablePath` and `AutoInstall` settings
2. **Permission denied**: Ensure proper file permissions and user rights
3. **Timeout errors**: Increase `DefaultTimeout` or use context with longer timeout
4. **Memory errors**: Set `COMPOSER_MEMORY_LIMIT=-1`
5. **Network issues**: Configure proxy settings in environment variables

### Debug Configuration

```go
func debugConfiguration(comp *composer.Composer) {
    fmt.Printf("Working Directory: %s\n", comp.GetWorkingDir())
    
    // Check if Composer is accessible
    if comp.IsInstalled() {
        version, _ := comp.GetVersion()
        fmt.Printf("Composer Version: %s\n", version)
    } else {
        fmt.Println("❌ Composer not accessible")
    }
    
    // Run diagnostics
    output, err := comp.Diagnose()
    if err != nil {
        fmt.Printf("Diagnostics failed: %v\n", err)
    } else {
        fmt.Printf("Diagnostics:\n%s\n", output)
    }
}
```
