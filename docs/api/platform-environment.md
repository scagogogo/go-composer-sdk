# Platform & Environment API

The platform and environment API provides functionality for checking system requirements, platform compatibility, and environment configuration.

## Platform Requirements

### CheckPlatformReqs

Checks if all platform requirements are satisfied.

```go
func (c *Composer) CheckPlatformReqs() error
```

**Returns:**
- `error` - Error if requirements are not met

**Example:**
```go
err := comp.CheckPlatformReqs()
if err != nil {
    log.Printf("Platform requirements not satisfied: %v", err)
    // Handle missing requirements
} else {
    fmt.Println("✅ All platform requirements satisfied")
}
```

### CheckPlatformReqsWithDetails

Checks platform requirements and returns detailed information.

```go
func (c *Composer) CheckPlatformReqsWithDetails() (*PlatformRequirements, error)
```

**Returns:**
- `*PlatformRequirements` - Detailed platform requirement information
- `error` - Error if check fails

**Example:**
```go
reqs, err := comp.CheckPlatformReqsWithDetails()
if err != nil {
    log.Printf("Failed to check platform requirements: %v", err)
    return
}

fmt.Printf("PHP Version: %s (required: %s)\n", reqs.PHP.Current, reqs.PHP.Required)
for _, ext := range reqs.Extensions {
    status := "✅"
    if !ext.Available {
        status = "❌"
    }
    fmt.Printf("%s Extension %s: %s\n", status, ext.Name, ext.Version)
}
```

## PHP Environment

### GetPHPVersion

Gets the current PHP version.

```go
func (c *Composer) GetPHPVersion() (string, error)
```

**Returns:**
- `string` - PHP version string
- `error` - Error if version cannot be retrieved

**Example:**
```go
phpVersion, err := comp.GetPHPVersion()
if err != nil {
    log.Printf("Failed to get PHP version: %v", err)
} else {
    fmt.Printf("PHP Version: %s\n", phpVersion)
}
```

### CheckPHPExtension

Checks if a specific PHP extension is available.

```go
func (c *Composer) CheckPHPExtension(extension string) (bool, error)
```

**Parameters:**
- `extension` - Name of the PHP extension to check

**Returns:**
- `bool` - True if extension is available
- `error` - Error if check fails

**Example:**
```go
extensions := []string{"mbstring", "openssl", "pdo", "json", "curl"}

for _, ext := range extensions {
    available, err := comp.CheckPHPExtension(ext)
    if err != nil {
        log.Printf("Failed to check extension %s: %v", ext, err)
        continue
    }
    
    status := "✅"
    if !available {
        status = "❌"
    }
    fmt.Printf("%s PHP Extension: %s\n", status, ext)
}
```

### GetPHPConfiguration

Gets PHP configuration information.

```go
func (c *Composer) GetPHPConfiguration() (*PHPConfig, error)
```

**Returns:**
- `*PHPConfig` - PHP configuration details
- `error` - Error if configuration cannot be retrieved

**Example:**
```go
config, err := comp.GetPHPConfiguration()
if err != nil {
    log.Printf("Failed to get PHP configuration: %v", err)
    return
}

fmt.Printf("PHP Version: %s\n", config.Version)
fmt.Printf("Memory Limit: %s\n", config.MemoryLimit)
fmt.Printf("Max Execution Time: %s\n", config.MaxExecutionTime)
fmt.Printf("Upload Max Filesize: %s\n", config.UploadMaxFilesize)
```

## System Information

### GetSystemInfo

Gets comprehensive system information.

```go
func (c *Composer) GetSystemInfo() (*SystemInfo, error)
```

**Returns:**
- `*SystemInfo` - System information
- `error` - Error if information cannot be retrieved

**Example:**
```go
sysInfo, err := comp.GetSystemInfo()
if err != nil {
    log.Printf("Failed to get system info: %v", err)
    return
}

fmt.Printf("Operating System: %s\n", sysInfo.OS)
fmt.Printf("Architecture: %s\n", sysInfo.Arch)
fmt.Printf("PHP Version: %s\n", sysInfo.PHPVersion)
fmt.Printf("Composer Version: %s\n", sysInfo.ComposerVersion)
```

### CheckDiskSpace

Checks available disk space in the working directory.

```go
func (c *Composer) CheckDiskSpace() (*DiskSpace, error)
```

**Returns:**
- `*DiskSpace` - Disk space information
- `error` - Error if check fails

**Example:**
```go
diskSpace, err := comp.CheckDiskSpace()
if err != nil {
    log.Printf("Failed to check disk space: %v", err)
    return
}

fmt.Printf("Total Space: %s\n", formatBytes(diskSpace.Total))
fmt.Printf("Available Space: %s\n", formatBytes(diskSpace.Available))
fmt.Printf("Used Space: %s\n", formatBytes(diskSpace.Used))

if diskSpace.Available < 100*1024*1024 { // Less than 100MB
    fmt.Println("⚠️  Warning: Low disk space!")
}
```

## Environment Configuration

### GetEnvironmentVariables

Gets Composer-related environment variables.

```go
func (c *Composer) GetEnvironmentVariables() map[string]string
```

**Returns:**
- `map[string]string` - Environment variables

**Example:**
```go
envVars := comp.GetEnvironmentVariables()

composerVars := []string{
    "COMPOSER_HOME",
    "COMPOSER_CACHE_DIR",
    "COMPOSER_MEMORY_LIMIT",
    "COMPOSER_PROCESS_TIMEOUT",
    "COMPOSER_DISCARD_CHANGES",
}

for _, varName := range composerVars {
    if value, exists := envVars[varName]; exists {
        fmt.Printf("%s: %s\n", varName, value)
    } else {
        fmt.Printf("%s: (not set)\n", varName)
    }
}
```

### SetEnvironmentVariable

Sets a Composer environment variable.

```go
func (c *Composer) SetEnvironmentVariable(name string, value string)
```

**Parameters:**
- `name` - Environment variable name
- `value` - Environment variable value

**Example:**
```go
// Configure Composer environment
comp.SetEnvironmentVariable("COMPOSER_MEMORY_LIMIT", "-1")
comp.SetEnvironmentVariable("COMPOSER_PROCESS_TIMEOUT", "600")
comp.SetEnvironmentVariable("COMPOSER_CACHE_DIR", "/tmp/composer-cache")
comp.SetEnvironmentVariable("COMPOSER_DISCARD_CHANGES", "true")
```

## Type Definitions

### PlatformRequirements

```go
type PlatformRequirements struct {
    PHP        PHPRequirement        `json:"php"`
    Extensions []ExtensionRequirement `json:"extensions"`
    Satisfied  bool                  `json:"satisfied"`
}

type PHPRequirement struct {
    Required string `json:"required"`
    Current  string `json:"current"`
    Satisfied bool  `json:"satisfied"`
}

type ExtensionRequirement struct {
    Name      string `json:"name"`
    Required  string `json:"required"`
    Version   string `json:"version"`
    Available bool   `json:"available"`
}
```

### PHPConfig

```go
type PHPConfig struct {
    Version            string `json:"version"`
    MemoryLimit        string `json:"memory_limit"`
    MaxExecutionTime   string `json:"max_execution_time"`
    UploadMaxFilesize  string `json:"upload_max_filesize"`
    PostMaxSize        string `json:"post_max_size"`
    DisplayErrors      string `json:"display_errors"`
    ErrorReporting     string `json:"error_reporting"`
    TimeZone           string `json:"timezone"`
}
```

### SystemInfo

```go
type SystemInfo struct {
    OS              string `json:"os"`
    Arch            string `json:"arch"`
    PHPVersion      string `json:"php_version"`
    ComposerVersion string `json:"composer_version"`
    WorkingDir      string `json:"working_dir"`
    HomeDir         string `json:"home_dir"`
    TempDir         string `json:"temp_dir"`
}
```

### DiskSpace

```go
type DiskSpace struct {
    Total     uint64 `json:"total"`
    Available uint64 `json:"available"`
    Used      uint64 `json:"used"`
    Path      string `json:"path"`
}
```

## Platform-Specific Checks

### Windows-Specific Checks

```go
func checkWindowsRequirements(comp *composer.Composer) error {
    // Check if running on Windows
    sysInfo, err := comp.GetSystemInfo()
    if err != nil {
        return err
    }
    
    if sysInfo.OS != "windows" {
        return nil // Not Windows, skip
    }
    
    // Check Windows-specific requirements
    phpConfig, err := comp.GetPHPConfiguration()
    if err != nil {
        return fmt.Errorf("failed to get PHP config: %w", err)
    }
    
    // Check for common Windows issues
    if phpConfig.MemoryLimit == "128M" {
        fmt.Println("⚠️  Warning: PHP memory limit is low (128M)")
        fmt.Println("💡 Consider increasing to 512M or higher")
    }
    
    // Check required extensions for Windows
    windowsExtensions := []string{"openssl", "mbstring", "curl"}
    for _, ext := range windowsExtensions {
        available, err := comp.CheckPHPExtension(ext)
        if err != nil {
            return fmt.Errorf("failed to check extension %s: %w", ext, err)
        }
        if !available {
            return fmt.Errorf("required extension missing: %s", ext)
        }
    }
    
    return nil
}
```

### macOS-Specific Checks

```go
func checkMacOSRequirements(comp *composer.Composer) error {
    sysInfo, err := comp.GetSystemInfo()
    if err != nil {
        return err
    }
    
    if sysInfo.OS != "darwin" {
        return nil // Not macOS, skip
    }
    
    // Check if Homebrew PHP is being used
    phpVersion, err := comp.GetPHPVersion()
    if err != nil {
        return err
    }
    
    fmt.Printf("PHP Version on macOS: %s\n", phpVersion)
    
    // Check for macOS-specific paths
    envVars := comp.GetEnvironmentVariables()
    if composerHome, exists := envVars["COMPOSER_HOME"]; exists {
        fmt.Printf("Composer Home: %s\n", composerHome)
    }
    
    return nil
}
```

### Linux-Specific Checks

```go
func checkLinuxRequirements(comp *composer.Composer) error {
    sysInfo, err := comp.GetSystemInfo()
    if err != nil {
        return err
    }
    
    if sysInfo.OS != "linux" {
        return nil // Not Linux, skip
    }
    
    // Check disk space (important for package installations)
    diskSpace, err := comp.CheckDiskSpace()
    if err != nil {
        return fmt.Errorf("failed to check disk space: %w", err)
    }
    
    minSpace := uint64(500 * 1024 * 1024) // 500MB
    if diskSpace.Available < minSpace {
        return fmt.Errorf("insufficient disk space: %d MB available, %d MB required",
            diskSpace.Available/(1024*1024), minSpace/(1024*1024))
    }
    
    return nil
}
```

## Comprehensive Environment Check

```go
func performComprehensiveEnvironmentCheck(comp *composer.Composer) error {
    fmt.Println("🔍 Performing comprehensive environment check...")
    
    // 1. System Information
    fmt.Println("\n📋 System Information:")
    sysInfo, err := comp.GetSystemInfo()
    if err != nil {
        return fmt.Errorf("failed to get system info: %w", err)
    }
    
    fmt.Printf("  OS: %s\n", sysInfo.OS)
    fmt.Printf("  Architecture: %s\n", sysInfo.Arch)
    fmt.Printf("  PHP Version: %s\n", sysInfo.PHPVersion)
    fmt.Printf("  Composer Version: %s\n", sysInfo.ComposerVersion)
    
    // 2. Platform Requirements
    fmt.Println("\n🔧 Platform Requirements:")
    reqs, err := comp.CheckPlatformReqsWithDetails()
    if err != nil {
        return fmt.Errorf("failed to check platform requirements: %w", err)
    }
    
    phpStatus := "✅"
    if !reqs.PHP.Satisfied {
        phpStatus = "❌"
    }
    fmt.Printf("  %s PHP: %s (required: %s)\n", phpStatus, reqs.PHP.Current, reqs.PHP.Required)
    
    for _, ext := range reqs.Extensions {
        extStatus := "✅"
        if !ext.Available {
            extStatus = "❌"
        }
        fmt.Printf("  %s Extension %s: %s\n", extStatus, ext.Name, ext.Version)
    }
    
    // 3. Disk Space
    fmt.Println("\n💾 Disk Space:")
    diskSpace, err := comp.CheckDiskSpace()
    if err != nil {
        return fmt.Errorf("failed to check disk space: %w", err)
    }
    
    fmt.Printf("  Total: %s\n", formatBytes(diskSpace.Total))
    fmt.Printf("  Available: %s\n", formatBytes(diskSpace.Available))
    fmt.Printf("  Used: %s\n", formatBytes(diskSpace.Used))
    
    // 4. Environment Variables
    fmt.Println("\n🌍 Environment Variables:")
    envVars := comp.GetEnvironmentVariables()
    importantVars := []string{
        "COMPOSER_HOME", "COMPOSER_CACHE_DIR", "COMPOSER_MEMORY_LIMIT",
    }
    
    for _, varName := range importantVars {
        if value, exists := envVars[varName]; exists {
            fmt.Printf("  %s: %s\n", varName, value)
        } else {
            fmt.Printf("  %s: (not set)\n", varName)
        }
    }
    
    // 5. Platform-specific checks
    fmt.Println("\n🖥️  Platform-specific checks:")
    switch sysInfo.OS {
    case "windows":
        err = checkWindowsRequirements(comp)
    case "darwin":
        err = checkMacOSRequirements(comp)
    case "linux":
        err = checkLinuxRequirements(comp)
    }
    
    if err != nil {
        return fmt.Errorf("platform-specific check failed: %w", err)
    }
    
    fmt.Println("\n✅ Environment check completed successfully!")
    return nil
}

func formatBytes(bytes uint64) string {
    const unit = 1024
    if bytes < unit {
        return fmt.Sprintf("%d B", bytes)
    }
    div, exp := int64(unit), 0
    for n := bytes / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
```

## Best Practices

1. **Check platform requirements** before performing operations
2. **Monitor disk space** for large installations
3. **Configure environment variables** appropriately
4. **Handle platform differences** gracefully
5. **Provide clear error messages** for missing requirements
6. **Cache system information** when possible
7. **Validate PHP extensions** before using features that depend on them
