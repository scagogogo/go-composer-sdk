# Detector API

The detector API provides functionality for detecting and validating Composer installations on the system.

## Package: detector

```go
import "github.com/scagogogo/go-composer-sdk/pkg/detector"
```

## Types

### Detector

The main struct for detecting Composer installations.

```go
type Detector struct {
    // Contains filtered or unexported fields
}
```

## Functions

### NewDetector

Creates a new Composer detector instance.

```go
func NewDetector() *Detector
```

**Returns:**
- `*Detector` - A new detector instance

**Example:**
```go
detector := detector.NewDetector()
```

## Methods

### Detect

Attempts to detect Composer executable on the system.

```go
func (d *Detector) Detect() (string, error)
```

**Returns:**
- `string` - Path to the Composer executable
- `error` - Error if Composer is not found

**Example:**
```go
detector := detector.NewDetector()
composerPath, err := detector.Detect()
if err != nil {
    log.Printf("Composer not found: %v", err)
    return
}

fmt.Printf("Composer found at: %s\n", composerPath)
```

### IsInstalled

Checks if Composer is installed and accessible.

```go
func (d *Detector) IsInstalled() bool
```

**Returns:**
- `bool` - True if Composer is installed

**Example:**
```go
detector := detector.NewDetector()
if detector.IsInstalled() {
    fmt.Println("✅ Composer is installed")
} else {
    fmt.Println("❌ Composer is not installed")
}
```

### SetPossiblePaths

Sets custom paths to search for Composer.

```go
func (d *Detector) SetPossiblePaths(paths []string)
```

**Parameters:**
- `paths` - Array of paths to search for Composer

**Example:**
```go
detector := detector.NewDetector()
customPaths := []string{
    "/usr/local/bin/composer",
    "/opt/composer/composer",
    "./vendor/bin/composer",
}
detector.SetPossiblePaths(customPaths)
```

### AddPossiblePath

Adds a single path to the search list.

```go
func (d *Detector) AddPossiblePath(path string)
```

**Parameters:**
- `path` - Path to add to the search list

**Example:**
```go
detector := detector.NewDetector()
detector.AddPossiblePath("/custom/path/to/composer")
```

## Detection Strategy

The detector uses the following strategy to find Composer:

1. **Environment Variable**: Checks `COMPOSER_PATH` environment variable
2. **Platform-Specific Paths**: Searches common installation locations for the current OS
3. **System PATH**: Uses `which` (Unix) or `where` (Windows) commands
4. **Current Directory**: Looks for `composer.phar` in the current directory

### Platform-Specific Paths

#### Windows
- `%APPDATA%\Composer\composer.phar`
- `%ProgramFiles%\Composer\composer.phar`
- `%ProgramFiles(x86)%\Composer\composer.phar`
- `composer.phar`
- `composer.bat`
- `composer`

#### macOS
- `/usr/local/bin/composer`
- `/opt/homebrew/bin/composer`
- `~/.composer/vendor/bin/composer`
- `~/composer.phar`

#### Linux/Unix
- `/usr/local/bin/composer`
- `/usr/bin/composer`
- `~/.composer/vendor/bin/composer`
- `~/composer.phar`

## Error Handling

The detector defines specific error types:

```go
var (
    ErrExecutableNotFound = errors.New("composer executable not found")
)
```

**Example:**
```go
detector := detector.NewDetector()
path, err := detector.Detect()
if err != nil {
    if errors.Is(err, detector.ErrExecutableNotFound) {
        fmt.Println("Composer executable not found")
        // Handle installation or provide instructions
    } else {
        fmt.Printf("Detection error: %v\n", err)
    }
    return
}
```

## Advanced Usage

### Custom Detection Logic

```go
func findComposer() (string, error) {
    detector := detector.NewDetector()
    
    // Try standard detection first
    path, err := detector.Detect()
    if err == nil {
        return path, nil
    }
    
    // Add custom paths and try again
    customPaths := []string{
        "/opt/php/composer",
        "/usr/local/php/composer",
        "./tools/composer",
    }
    
    for _, customPath := range customPaths {
        detector.AddPossiblePath(customPath)
    }
    
    return detector.Detect()
}
```

### Validation with Version Check

```go
func validateComposerInstallation() error {
    detector := detector.NewDetector()
    
    if !detector.IsInstalled() {
        return fmt.Errorf("composer is not installed")
    }
    
    path, err := detector.Detect()
    if err != nil {
        return fmt.Errorf("failed to detect composer: %w", err)
    }
    
    // Create a composer instance to verify it works
    comp, err := composer.New(composer.Options{
        ExecutablePath: path,
    })
    if err != nil {
        return fmt.Errorf("failed to create composer instance: %w", err)
    }
    
    // Try to get version to verify it's working
    version, err := comp.GetVersion()
    if err != nil {
        return fmt.Errorf("composer is not working properly: %w", err)
    }
    
    fmt.Printf("✅ Composer %s is working correctly\n", version)
    return nil
}
```

## Best Practices

1. **Always check installation** before using Composer
2. **Handle detection errors** gracefully
3. **Use custom paths** when Composer is in non-standard locations
4. **Validate functionality** after detection
5. **Provide clear error messages** to users

## Integration Example

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/detector"
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create detector
    det := detector.NewDetector()
    
    // Check if Composer is installed
    if !det.IsInstalled() {
        log.Fatal("Composer is not installed. Please install Composer first.")
    }
    
    // Detect Composer path
    composerPath, err := det.Detect()
    if err != nil {
        log.Fatalf("Failed to detect Composer: %v", err)
    }
    
    fmt.Printf("Found Composer at: %s\n", composerPath)
    
    // Create Composer instance with detected path
    comp, err := composer.New(composer.Options{
        ExecutablePath: composerPath,
    })
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Verify it's working
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("Composer is not working: %v", err)
    }
    
    fmt.Printf("✅ Composer %s is ready to use!\n", version)
}
```
