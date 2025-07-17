# Installer API

The installer API provides functionality for automatically installing Composer on systems where it's not available.

## Package: installer

```go
import "github.com/scagogogo/go-composer-sdk/pkg/installer"
```

## Types

### Installer

The main struct for installing Composer.

```go
type Installer struct {
    // Contains filtered or unexported fields
}
```

### Config

Configuration options for the installer.

```go
type Config struct {
    DownloadURL     string // URL to download Composer installer
    InstallPath     string // Path where Composer will be installed
    UseProxy        bool   // Whether to use HTTP proxy
    ProxyURL        string // HTTP proxy URL
    TimeoutSeconds  int    // Download timeout in seconds
    UseSudo         bool   // Use sudo for installation (Unix systems)
    PreferBrewOnMac bool   // Prefer Homebrew on macOS
}
```

## Functions

### NewInstaller

Creates a new installer instance with custom configuration.

```go
func NewInstaller(config Config) *Installer
```

**Parameters:**
- `config` - Installation configuration

**Returns:**
- `*Installer` - A new installer instance

**Example:**
```go
config := installer.Config{
    DownloadURL:     "https://getcomposer.org/installer",
    InstallPath:     "/usr/local/bin",
    TimeoutSeconds:  300,
    PreferBrewOnMac: true,
}

inst := installer.NewInstaller(config)
```

### DefaultInstaller

Creates an installer instance with default configuration.

```go
func DefaultInstaller() *Installer
```

**Returns:**
- `*Installer` - An installer with default settings

**Example:**
```go
installer := installer.DefaultInstaller()
```

### DefaultConfig

Returns default configuration for the current platform.

```go
func DefaultConfig() Config
```

**Returns:**
- `Config` - Default configuration

**Example:**
```go
config := installer.DefaultConfig()
config.TimeoutSeconds = 600 // Customize as needed
inst := installer.NewInstaller(config)
```

## Methods

### Install

Installs Composer using platform-specific methods.

```go
func (i *Installer) Install() error
```

**Returns:**
- `error` - Error if installation fails

**Example:**
```go
installer := installer.DefaultInstaller()
err := installer.Install()
if err != nil {
    log.Fatalf("Failed to install Composer: %v", err)
}

fmt.Println("✅ Composer installed successfully!")
```

### GetConfig

Gets the current installer configuration.

```go
func (i *Installer) GetConfig() Config
```

**Returns:**
- `Config` - Current configuration

**Example:**
```go
installer := installer.DefaultInstaller()
config := installer.GetConfig()
fmt.Printf("Install path: %s\n", config.InstallPath)
```

### SetConfig

Updates the installer configuration.

```go
func (i *Installer) SetConfig(config Config)
```

**Parameters:**
- `config` - New configuration

**Example:**
```go
installer := installer.DefaultInstaller()
config := installer.GetConfig()
config.TimeoutSeconds = 600
config.UseProxy = true
config.ProxyURL = "http://proxy.example.com:8080"
installer.SetConfig(config)
```

## Platform-Specific Installation

The installer uses different strategies based on the operating system:

### Windows
- Downloads the Composer installer
- Executes the installer with appropriate options
- Installs to `%ProgramFiles%\Composer` by default

### macOS
- **Preferred**: Uses Homebrew if available (`brew install composer`)
- **Fallback**: Downloads and installs manually to `/usr/local/bin`

### Linux/Unix
- Downloads the Composer installer script
- Executes with PHP to install Composer
- Installs to `/usr/local/bin` by default
- May require sudo privileges

## Error Types

```go
var (
    ErrInstallationFailed  = errors.New("installation failed")
    ErrInsufficientRights = errors.New("insufficient rights, please use administrator/sudo privileges")
    ErrUnsupportedPlatform = errors.New("unsupported operating system platform")
    ErrDownloadFailed      = errors.New("download failed")
)
```

## Configuration Examples

### Basic Installation

```go
func installComposer() error {
    installer := installer.DefaultInstaller()
    return installer.Install()
}
```

### Custom Installation Path

```go
func installComposerCustomPath() error {
    config := installer.DefaultConfig()
    config.InstallPath = "/opt/composer"
    
    installer := installer.NewInstaller(config)
    return installer.Install()
}
```

### Installation with Proxy

```go
func installComposerWithProxy() error {
    config := installer.DefaultConfig()
    config.UseProxy = true
    config.ProxyURL = "http://proxy.company.com:8080"
    config.TimeoutSeconds = 600 // 10 minutes
    
    installer := installer.NewInstaller(config)
    return installer.Install()
}
```

### macOS with Homebrew Preference

```go
func installComposerMacOS() error {
    config := installer.DefaultConfig()
    config.PreferBrewOnMac = true // Try Homebrew first
    
    installer := installer.NewInstaller(config)
    return installer.Install()
}
```

## Advanced Usage

### Installation with Validation

```go
func installAndValidateComposer() error {
    // Install Composer
    installer := installer.DefaultInstaller()
    err := installer.Install()
    if err != nil {
        return fmt.Errorf("installation failed: %w", err)
    }
    
    // Validate installation
    detector := detector.NewDetector()
    if !detector.IsInstalled() {
        return fmt.Errorf("composer installation validation failed")
    }
    
    // Test functionality
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        return fmt.Errorf("failed to create composer instance: %w", err)
    }
    
    version, err := comp.GetVersion()
    if err != nil {
        return fmt.Errorf("composer is not working properly: %w", err)
    }
    
    fmt.Printf("✅ Composer %s installed and validated successfully!\n", version)
    return nil
}
```

### Retry Installation with Different Configurations

```go
func installComposerWithRetry() error {
    configs := []installer.Config{
        // Try default first
        installer.DefaultConfig(),
        
        // Try with longer timeout
        func() installer.Config {
            config := installer.DefaultConfig()
            config.TimeoutSeconds = 900
            return config
        }(),
        
        // Try different install path
        func() installer.Config {
            config := installer.DefaultConfig()
            config.InstallPath = "/tmp"
            return config
        }(),
    }
    
    for i, config := range configs {
        fmt.Printf("Attempting installation (try %d/%d)...\n", i+1, len(configs))
        
        installer := installer.NewInstaller(config)
        err := installer.Install()
        if err == nil {
            fmt.Println("✅ Installation successful!")
            return nil
        }
        
        fmt.Printf("❌ Installation attempt %d failed: %v\n", i+1, err)
    }
    
    return fmt.Errorf("all installation attempts failed")
}
```

### Platform-Specific Installation Logic

```go
func installComposerPlatformSpecific() error {
    config := installer.DefaultConfig()
    
    switch runtime.GOOS {
    case "windows":
        // Windows-specific configuration
        config.InstallPath = `C:\Tools\Composer`
        
    case "darwin":
        // macOS-specific configuration
        config.PreferBrewOnMac = true
        config.InstallPath = "/usr/local/bin"
        
    case "linux":
        // Linux-specific configuration
        config.UseSudo = true
        config.InstallPath = "/usr/local/bin"
        
    default:
        return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
    }
    
    installer := installer.NewInstaller(config)
    return installer.Install()
}
```

## Best Practices

1. **Use default configuration** when possible for platform compatibility
2. **Handle installation errors** gracefully with fallback options
3. **Validate installation** after completion
4. **Consider proxy settings** in corporate environments
5. **Use appropriate permissions** (sudo on Unix systems when needed)
6. **Set reasonable timeouts** for network operations
7. **Prefer package managers** (like Homebrew) when available

## Integration with Main SDK

```go
func ensureComposerAvailable() (*composer.Composer, error) {
    // Try to create Composer instance with auto-install
    options := composer.DefaultOptions()
    options.AutoInstall = true
    
    comp, err := composer.New(options)
    if err != nil {
        return nil, fmt.Errorf("failed to ensure Composer availability: %w", err)
    }
    
    return comp, nil
}
```

## Error Handling

```go
func handleInstallationError(err error) {
    switch {
    case errors.Is(err, installer.ErrInsufficientRights):
        fmt.Println("❌ Installation failed: Insufficient privileges")
        fmt.Println("💡 Try running with administrator/sudo privileges")
        
    case errors.Is(err, installer.ErrDownloadFailed):
        fmt.Println("❌ Installation failed: Download error")
        fmt.Println("💡 Check your internet connection and proxy settings")
        
    case errors.Is(err, installer.ErrUnsupportedPlatform):
        fmt.Println("❌ Installation failed: Unsupported platform")
        fmt.Println("💡 Please install Composer manually")
        
    default:
        fmt.Printf("❌ Installation failed: %v\n", err)
        fmt.Println("💡 Please check the error message and try again")
    }
}
