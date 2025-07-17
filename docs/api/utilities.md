# Utilities API

The utilities API provides helper functions for file system operations, HTTP requests, and cross-platform compatibility.

## Package: utils

```go
import "github.com/scagogogo/go-composer-sdk/pkg/utils"
```

## File System Utilities

### CheckWritePermission

Checks if a directory has write permissions.

```go
func CheckWritePermission(dir string) error
```

**Parameters:**
- `dir` - Directory path to check

**Returns:**
- `error` - Error if directory is not writable

**Example:**
```go
err := utils.CheckWritePermission("/usr/local/bin")
if err != nil {
    log.Printf("Directory not writable: %v", err)
} else {
    fmt.Println("✅ Directory is writable")
}
```

### EnsureDirectoryExists

Ensures a directory exists, creating it if necessary.

```go
func EnsureDirectoryExists(dir string) error
```

**Parameters:**
- `dir` - Directory path to ensure exists

**Returns:**
- `error` - Error if directory cannot be created

**Example:**
```go
installDir := "/opt/composer"
err := utils.EnsureDirectoryExists(installDir)
if err != nil {
    log.Printf("Failed to create directory: %v", err)
} else {
    fmt.Printf("✅ Directory %s is ready\n", installDir)
}
```

### CreateFileWithContent

Creates a file with specified content and permissions.

```go
func CreateFileWithContent(filePath string, content []byte, perm os.FileMode) error
```

**Parameters:**
- `filePath` - Full path to the file to create
- `content` - Content to write to the file
- `perm` - File permissions (e.g., 0644, 0755)

**Returns:**
- `error` - Error if file creation fails

**Example:**
```go
scriptContent := []byte("#!/bin/sh\necho 'Hello World'")
err := utils.CreateFileWithContent("/usr/local/bin/hello.sh", scriptContent, 0755)
if err != nil {
    log.Printf("Failed to create script: %v", err)
} else {
    fmt.Println("✅ Script created successfully")
}
```

## HTTP Utilities

### DownloadFile

Downloads a file from a URL with configurable options.

```go
func DownloadFile(url string, filepath string, config DownloadConfig) error
```

**Parameters:**
- `url` - URL to download from
- `filepath` - Local file path to save to
- `config` - Download configuration options

**Returns:**
- `error` - Error if download fails

**Example:**
```go
config := utils.DownloadConfig{
    UseProxy:       false,
    TimeoutSeconds: 300,
    UserAgent:      "Go-Composer-SDK/1.0",
}

err := utils.DownloadFile(
    "https://getcomposer.org/installer",
    "/tmp/composer-setup.php",
    config,
)
if err != nil {
    log.Printf("Download failed: %v", err)
} else {
    fmt.Println("✅ File downloaded successfully")
}
```

### DownloadWithProgress

Downloads a file with progress reporting.

```go
func DownloadWithProgress(url string, filepath string, config DownloadConfig, progressCallback func(downloaded, total int64)) error
```

**Parameters:**
- `url` - URL to download from
- `filepath` - Local file path to save to
- `config` - Download configuration options
- `progressCallback` - Function called with progress updates

**Example:**
```go
config := utils.DownloadConfig{
    TimeoutSeconds: 600,
}

progressCallback := func(downloaded, total int64) {
    if total > 0 {
        percent := float64(downloaded) / float64(total) * 100
        fmt.Printf("\rDownloading... %.1f%% (%d/%d bytes)", percent, downloaded, total)
    }
}

err := utils.DownloadWithProgress(
    "https://example.com/large-file.zip",
    "/tmp/large-file.zip",
    config,
    progressCallback,
)
```

## Type Definitions

### DownloadConfig

Configuration options for HTTP downloads.

```go
type DownloadConfig struct {
    UseProxy       bool   // Whether to use HTTP proxy
    ProxyURL       string // HTTP proxy URL
    TimeoutSeconds int    // Request timeout in seconds
    UserAgent      string // Custom User-Agent header
    Headers        map[string]string // Additional HTTP headers
    MaxRetries     int    // Maximum number of retry attempts
    RetryDelay     time.Duration // Delay between retries
}
```

## Platform Utilities

### GetPlatformSpecificPath

Gets platform-specific paths for common directories.

```go
func GetPlatformSpecificPath(pathType string) (string, error)
```

**Parameters:**
- `pathType` - Type of path ("home", "temp", "cache", "config")

**Returns:**
- `string` - Platform-specific path
- `error` - Error if path type is unknown

**Example:**
```go
// Get user home directory
homeDir, err := utils.GetPlatformSpecificPath("home")
if err != nil {
    log.Printf("Failed to get home directory: %v", err)
} else {
    fmt.Printf("Home directory: %s\n", homeDir)
}

// Get temporary directory
tempDir, err := utils.GetPlatformSpecificPath("temp")
if err != nil {
    log.Printf("Failed to get temp directory: %v", err)
} else {
    fmt.Printf("Temp directory: %s\n", tempDir)
}
```

### IsExecutable

Checks if a file is executable on the current platform.

```go
func IsExecutable(filePath string) bool
```

**Parameters:**
- `filePath` - Path to the file to check

**Returns:**
- `bool` - True if file is executable

**Example:**
```go
if utils.IsExecutable("/usr/local/bin/composer") {
    fmt.Println("✅ Composer is executable")
} else {
    fmt.Println("❌ Composer is not executable")
}
```

## Command Execution Utilities

### ExecuteCommand

Executes a system command with configurable options.

```go
func ExecuteCommand(command string, args []string, options ExecuteOptions) (string, error)
```

**Parameters:**
- `command` - Command to execute
- `args` - Command arguments
- `options` - Execution options

**Returns:**
- `string` - Command output
- `error` - Error if execution fails

**Example:**
```go
options := utils.ExecuteOptions{
    WorkingDir: "/path/to/project",
    Timeout:    30 * time.Second,
    Env: []string{
        "PATH=/usr/local/bin:/usr/bin:/bin",
        "COMPOSER_HOME=/tmp/composer",
    },
}

output, err := utils.ExecuteCommand("php", []string{"--version"}, options)
if err != nil {
    log.Printf("Command failed: %v", err)
} else {
    fmt.Printf("PHP version: %s\n", output)
}
```

### ExecuteOptions

Options for command execution.

```go
type ExecuteOptions struct {
    WorkingDir string        // Working directory for command
    Timeout    time.Duration // Command timeout
    Env        []string      // Environment variables
    Input      string        // Standard input for command
}
```

## Validation Utilities

### ValidateURL

Validates if a string is a valid URL.

```go
func ValidateURL(urlStr string) error
```

**Parameters:**
- `urlStr` - URL string to validate

**Returns:**
- `error` - Error if URL is invalid

**Example:**
```go
err := utils.ValidateURL("https://getcomposer.org/installer")
if err != nil {
    log.Printf("Invalid URL: %v", err)
} else {
    fmt.Println("✅ URL is valid")
}
```

### ValidateFilePath

Validates if a file path is valid and accessible.

```go
func ValidateFilePath(filePath string) error
```

**Parameters:**
- `filePath` - File path to validate

**Returns:**
- `error` - Error if path is invalid

**Example:**
```go
err := utils.ValidateFilePath("/usr/local/bin/composer")
if err != nil {
    log.Printf("Invalid file path: %v", err)
} else {
    fmt.Println("✅ File path is valid")
}
```

## Archive Utilities

### ExtractArchive

Extracts an archive file (ZIP, TAR, etc.) to a destination directory.

```go
func ExtractArchive(archivePath string, destDir string) error
```

**Parameters:**
- `archivePath` - Path to the archive file
- `destDir` - Destination directory for extraction

**Returns:**
- `error` - Error if extraction fails

**Example:**
```go
err := utils.ExtractArchive("/tmp/composer.zip", "/opt/composer")
if err != nil {
    log.Printf("Extraction failed: %v", err)
} else {
    fmt.Println("✅ Archive extracted successfully")
}
```

### CreateArchive

Creates an archive from a directory.

```go
func CreateArchive(sourceDir string, archivePath string, format string) error
```

**Parameters:**
- `sourceDir` - Source directory to archive
- `archivePath` - Path for the created archive
- `format` - Archive format ("zip", "tar", "tar.gz")

**Returns:**
- `error` - Error if archive creation fails

**Example:**
```go
err := utils.CreateArchive("/path/to/project", "/tmp/project.zip", "zip")
if err != nil {
    log.Printf("Archive creation failed: %v", err)
} else {
    fmt.Println("✅ Archive created successfully")
}
```

## Usage Examples

### Complete File Download with Validation

```go
func downloadComposerInstaller() error {
    // Validate URL
    url := "https://getcomposer.org/installer"
    if err := utils.ValidateURL(url); err != nil {
        return fmt.Errorf("invalid URL: %w", err)
    }
    
    // Ensure download directory exists
    downloadDir := "/tmp/composer-install"
    if err := utils.EnsureDirectoryExists(downloadDir); err != nil {
        return fmt.Errorf("failed to create download directory: %w", err)
    }
    
    // Check write permissions
    if err := utils.CheckWritePermission(downloadDir); err != nil {
        return fmt.Errorf("no write permission: %w", err)
    }
    
    // Download with progress
    filePath := filepath.Join(downloadDir, "installer.php")
    config := utils.DownloadConfig{
        TimeoutSeconds: 300,
        UserAgent:      "Go-Composer-SDK/1.0",
        MaxRetries:     3,
        RetryDelay:     time.Second * 2,
    }
    
    fmt.Println("Downloading Composer installer...")
    err := utils.DownloadWithProgress(url, filePath, config, func(downloaded, total int64) {
        if total > 0 {
            percent := float64(downloaded) / float64(total) * 100
            fmt.Printf("\rProgress: %.1f%%", percent)
        }
    })
    
    if err != nil {
        return fmt.Errorf("download failed: %w", err)
    }
    
    fmt.Println("\n✅ Download completed successfully")
    return nil
}
```

### Cross-Platform Directory Setup

```go
func setupComposerDirectories() error {
    // Get platform-specific paths
    homeDir, err := utils.GetPlatformSpecificPath("home")
    if err != nil {
        return fmt.Errorf("failed to get home directory: %w", err)
    }
    
    cacheDir, err := utils.GetPlatformSpecificPath("cache")
    if err != nil {
        return fmt.Errorf("failed to get cache directory: %w", err)
    }
    
    // Create Composer directories
    composerHome := filepath.Join(homeDir, ".composer")
    composerCache := filepath.Join(cacheDir, "composer")
    
    directories := []string{composerHome, composerCache}
    
    for _, dir := range directories {
        if err := utils.EnsureDirectoryExists(dir); err != nil {
            return fmt.Errorf("failed to create directory %s: %w", dir, err)
        }
        
        if err := utils.CheckWritePermission(dir); err != nil {
            return fmt.Errorf("no write permission for %s: %w", dir, err)
        }
        
        fmt.Printf("✅ Directory ready: %s\n", dir)
    }
    
    return nil
}
```

## Best Practices

1. **Always validate inputs** before performing operations
2. **Check permissions** before attempting file operations
3. **Use appropriate timeouts** for network operations
4. **Handle platform differences** gracefully
5. **Provide progress feedback** for long-running operations
6. **Clean up temporary files** after use
7. **Use proper error handling** and meaningful error messages
