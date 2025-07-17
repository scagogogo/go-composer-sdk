# Core API

The core API provides the fundamental functionality for creating and managing Composer instances.

## Package: composer

```go
import "github.com/scagogogo/go-composer-sdk/pkg/composer"
```

## Types

### Composer

The main struct that provides all Composer functionality.

```go
type Composer struct {
    // Contains filtered or unexported fields
}
```

### Options

Configuration options for creating a Composer instance.

```go
type Options struct {
    ExecutablePath  string                // Path to composer executable
    WorkingDir      string                // Working directory for operations  
    AutoInstall     bool                  // Auto-install Composer if not found
    DefaultTimeout  time.Duration         // Default timeout for operations
    Detector        *detector.Detector    // Custom detector instance
    Installer       *installer.Installer  // Custom installer instance
}
```

## Functions

### New

Creates a new Composer instance with the specified options.

```go
func New(options Options) (*Composer, error)
```

**Parameters:**
- `options` - Configuration options for the Composer instance

**Returns:**
- `*Composer` - A new Composer instance
- `error` - Error if creation fails

**Example:**
```go
comp, err := composer.New(composer.DefaultOptions())
if err != nil {
    log.Fatalf("Failed to create Composer instance: %v", err)
}
```

### DefaultOptions

Returns default configuration options.

```go
func DefaultOptions() Options
```

**Returns:**
- `Options` - Default configuration with sensible defaults

**Example:**
```go
options := composer.DefaultOptions()
options.WorkingDir = "/path/to/project"
comp, err := composer.New(options)
```

## Core Methods

### IsInstalled

Checks if Composer is installed and accessible.

```go
func (c *Composer) IsInstalled() bool
```

**Returns:**
- `bool` - True if Composer is installed and accessible

**Example:**
```go
if !comp.IsInstalled() {
    log.Fatal("Composer is not installed")
}
```

### GetVersion

Gets the installed Composer version.

```go
func (c *Composer) GetVersion() (string, error)
```

**Returns:**
- `string` - Composer version string
- `error` - Error if version cannot be retrieved

**Example:**
```go
version, err := comp.GetVersion()
if err != nil {
    log.Printf("Failed to get version: %v", err)
    return
}
fmt.Printf("Composer version: %s\n", version)
```

### Run

Executes a raw Composer command with the given arguments.

```go
func (c *Composer) Run(args ...string) (string, error)
```

**Parameters:**
- `args` - Command arguments to pass to Composer

**Returns:**
- `string` - Command output
- `error` - Error if command fails

**Example:**
```go
output, err := comp.Run("--version")
if err != nil {
    log.Printf("Command failed: %v", err)
    return
}
fmt.Println(output)
```

### RunWithContext

Executes a Composer command with context support for cancellation and timeouts.

```go
func (c *Composer) RunWithContext(ctx context.Context, args ...string) (string, error)
```

**Parameters:**
- `ctx` - Context for cancellation and timeout
- `args` - Command arguments to pass to Composer

**Returns:**
- `string` - Command output
- `error` - Error if command fails or context is cancelled

**Example:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

output, err := comp.RunWithContext(ctx, "install")
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Command timed out")
    } else {
        log.Printf("Command failed: %v", err)
    }
    return
}
```

## Configuration Methods

### SetWorkingDir

Sets the working directory for Composer operations.

```go
func (c *Composer) SetWorkingDir(dir string)
```

**Parameters:**
- `dir` - Path to the working directory

**Example:**
```go
comp.SetWorkingDir("/path/to/php/project")
```

### GetWorkingDir

Gets the current working directory.

```go
func (c *Composer) GetWorkingDir() string
```

**Returns:**
- `string` - Current working directory path

**Example:**
```go
workDir := comp.GetWorkingDir()
fmt.Printf("Working directory: %s\n", workDir)
```

### SetEnv

Sets environment variables for Composer operations.

```go
func (c *Composer) SetEnv(env []string)
```

**Parameters:**
- `env` - Array of environment variables in "KEY=VALUE" format

**Example:**
```go
comp.SetEnv([]string{
    "COMPOSER_MEMORY_LIMIT=-1",
    "COMPOSER_PROCESS_TIMEOUT=600",
    "COMPOSER_CACHE_DIR=/tmp/composer-cache",
})
```

### SetTimeout

Sets the default timeout for operations.

```go
func (c *Composer) SetTimeout(timeout time.Duration)
```

**Parameters:**
- `timeout` - Default timeout duration

**Example:**
```go
comp.SetTimeout(5 * time.Minute)
```

## Utility Methods

### SelfUpdate

Updates Composer to the latest version.

```go
func (c *Composer) SelfUpdate() error
```

**Returns:**
- `error` - Error if update fails

**Example:**
```go
err := comp.SelfUpdate()
if err != nil {
    log.Printf("Failed to update Composer: %v", err)
}
```

### ClearCache

Clears the Composer cache.

```go
func (c *Composer) ClearCache() error
```

**Returns:**
- `error` - Error if cache clearing fails

**Example:**
```go
err := comp.ClearCache()
if err != nil {
    log.Printf("Failed to clear cache: %v", err)
}
```

### Diagnose

Runs Composer's diagnostic checks.

```go
func (c *Composer) Diagnose() (string, error)
```

**Returns:**
- `string` - Diagnostic output
- `error` - Error if diagnostics fail

**Example:**
```go
output, err := comp.Diagnose()
if err != nil {
    log.Printf("Diagnostics failed: %v", err)
    return
}
fmt.Println("Diagnostic results:")
fmt.Println(output)
```

## Error Handling

All methods that can fail return an error as the last return value. Always check and handle errors appropriately:

```go
// Good error handling
version, err := comp.GetVersion()
if err != nil {
    log.Printf("Failed to get version: %v", err)
    return
}

// Use the version
fmt.Printf("Composer version: %s\n", version)
```

## Context Usage

For long-running operations, use context-aware methods to enable cancellation and timeouts:

```go
// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

// Use context-aware method
output, err := comp.RunWithContext(ctx, "install", "--no-dev")
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Operation timed out")
    } else {
        log.Printf("Operation failed: %v", err)
    }
}
```

## Best Practices

1. **Always check if Composer is installed** before performing operations
2. **Set the working directory** to your PHP project root
3. **Use context-aware methods** for long-running operations
4. **Handle errors appropriately** - don't ignore them
5. **Configure environment variables** as needed for your use case
6. **Set reasonable timeouts** to prevent hanging operations
