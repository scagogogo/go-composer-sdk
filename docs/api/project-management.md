# Project Management API

The project management API provides functionality for creating, configuring, and managing PHP projects.

## Project Creation

### CreateProject

Creates a new PHP project from a Composer package.

```go
func (c *Composer) CreateProject(packageName string, directory string, version string) error
```

**Parameters:**
- `packageName` - Name of the package to create project from
- `directory` - Target directory for the new project
- `version` - Version constraint (optional, use "" for latest)

**Returns:**
- `error` - Error if project creation fails

**Example:**
```go
// Create a Laravel project
err := comp.CreateProject("laravel/laravel", "my-laravel-app", "")

// Create a Symfony project with specific version
err := comp.CreateProject("symfony/skeleton", "my-symfony-app", "^6.0")

if err != nil {
    log.Printf("Failed to create project: %v", err)
}
```

### CreateProjectWithOptions

Creates a project with advanced options.

```go
func (c *Composer) CreateProjectWithOptions(options CreateProjectOptions) error
```

**Parameters:**
- `options` - Advanced project creation options

**Example:**
```go
options := CreateProjectOptions{
    PackageName:      "laravel/laravel",
    Directory:        "my-app",
    Version:          "^10.0",
    PreferDist:       true,
    NoScripts:        false,
    NoProgress:       false,
    KeepVcs:          false,
    NoInstall:        false,
    IgnorePlatformReqs: false,
}

err := comp.CreateProjectWithOptions(options)
```

## Configuration Management

### Validate

Validates the composer.json file in the current project.

```go
func (c *Composer) Validate() error
```

**Returns:**
- `error` - Error if validation fails or file is invalid

**Example:**
```go
err := comp.Validate()
if err != nil {
    log.Printf("composer.json validation failed: %v", err)
} else {
    fmt.Println("composer.json is valid")
}
```

### ValidateWithOptions

Validates composer.json with specific options.

```go
func (c *Composer) ValidateWithOptions(options ValidateOptions) error
```

**Example:**
```go
options := ValidateOptions{
    NoCheckAll:      false,
    NoCheckLock:     false,
    NoCheckPublish:  false,
    WithDependencies: true,
    Strict:          true,
}

err := comp.ValidateWithOptions(options)
```

### GetConfig

Gets a configuration value.

```go
func (c *Composer) GetConfig(key string) (string, error)
```

**Parameters:**
- `key` - Configuration key to retrieve

**Returns:**
- `string` - Configuration value
- `error` - Error if key doesn't exist or retrieval fails

**Example:**
```go
// Get cache directory
cacheDir, err := comp.GetConfig("cache-dir")
if err != nil {
    log.Printf("Failed to get config: %v", err)
} else {
    fmt.Printf("Cache directory: %s\n", cacheDir)
}

// Get other common configs
homeDir, _ := comp.GetConfig("home")
vendorDir, _ := comp.GetConfig("vendor-dir")
```

### SetConfig

Sets a configuration value.

```go
func (c *Composer) SetConfig(key string, value string) error
```

**Parameters:**
- `key` - Configuration key to set
- `value` - Configuration value

**Returns:**
- `error` - Error if setting fails

**Example:**
```go
// Set memory limit
err := comp.SetConfig("memory-limit", "-1")

// Set process timeout
err = comp.SetConfig("process-timeout", "600")

// Set preferred install method
err = comp.SetConfig("preferred-install", "dist")

if err != nil {
    log.Printf("Failed to set config: %v", err)
}
```

### UnsetConfig

Removes a configuration value.

```go
func (c *Composer) UnsetConfig(key string) error
```

**Parameters:**
- `key` - Configuration key to remove

**Example:**
```go
err := comp.UnsetConfig("github-oauth.github.com")
if err != nil {
    log.Printf("Failed to unset config: %v", err)
}
```

## Script Management

### RunScript

Executes a script defined in composer.json.

```go
func (c *Composer) RunScript(scriptName string) error
```

**Parameters:**
- `scriptName` - Name of the script to run

**Returns:**
- `error` - Error if script execution fails

**Example:**
```go
// Run common scripts
err := comp.RunScript("test")
err = comp.RunScript("post-install-cmd")
err = comp.RunScript("build")

if err != nil {
    log.Printf("Script execution failed: %v", err)
}
```

### RunScriptWithArgs

Executes a script with additional arguments.

```go
func (c *Composer) RunScriptWithArgs(scriptName string, args []string) error
```

**Parameters:**
- `scriptName` - Name of the script to run
- `args` - Additional arguments to pass to the script

**Example:**
```go
// Run PHPUnit with specific options
args := []string{"--coverage-html", "coverage/"}
err := comp.RunScriptWithArgs("test", args)

// Run custom script with parameters
err = comp.RunScriptWithArgs("deploy", []string{"production", "--force"})
```

### ListScripts

Lists all available scripts in the project.

```go
func (c *Composer) ListScripts() (string, error)
```

**Returns:**
- `string` - List of available scripts
- `error` - Error if listing fails

**Example:**
```go
scripts, err := comp.ListScripts()
if err != nil {
    log.Printf("Failed to list scripts: %v", err)
} else {
    fmt.Printf("Available scripts:\n%s\n", scripts)
}
```

## Autoloader Management

### DumpAutoload

Generates the autoloader files.

```go
func (c *Composer) DumpAutoload() error
```

**Returns:**
- `error` - Error if autoloader generation fails

**Example:**
```go
err := comp.DumpAutoload()
if err != nil {
    log.Printf("Failed to dump autoload: %v", err)
}
```

### DumpAutoloadOptimized

Generates optimized autoloader files.

```go
func (c *Composer) DumpAutoloadOptimized() error
```

**Returns:**
- `error` - Error if optimized autoloader generation fails

**Example:**
```go
err := comp.DumpAutoloadOptimized()
if err != nil {
    log.Printf("Failed to dump optimized autoload: %v", err)
}
```

### DumpAutoloadWithOptions

Generates autoloader with specific options.

```go
func (c *Composer) DumpAutoloadWithOptions(options AutoloadOptions) error
```

**Example:**
```go
options := AutoloadOptions{
    Optimize:     true,
    Classmap:     true,
    Apcu:         false,
    ApcuPrefix:   "",
    NoDev:        false,
}

err := comp.DumpAutoloadWithOptions(options)
```

## Composer.json Management

### ReadComposerJSON

Reads and parses the composer.json file.

```go
func (c *Composer) ReadComposerJSON() (*ComposerJSON, error)
```

**Returns:**
- `*ComposerJSON` - Parsed composer.json structure
- `error` - Error if reading or parsing fails

**Example:**
```go
composerData, err := comp.ReadComposerJSON()
if err != nil {
    log.Printf("Failed to read composer.json: %v", err)
    return
}

fmt.Printf("Project name: %s\n", composerData.Name)
fmt.Printf("Description: %s\n", composerData.Description)
fmt.Printf("Type: %s\n", composerData.Type)

// Access dependencies
for pkg, version := range composerData.Require {
    fmt.Printf("Requires: %s %s\n", pkg, version)
}
```

### WriteComposerJSON

Writes a ComposerJSON structure to the composer.json file.

```go
func (c *Composer) WriteComposerJSON(data *ComposerJSON) error
```

**Parameters:**
- `data` - ComposerJSON structure to write

**Returns:**
- `error` - Error if writing fails

**Example:**
```go
// Read existing composer.json
composerData, err := comp.ReadComposerJSON()
if err != nil {
    log.Fatal(err)
}

// Modify the data
composerData.Description = "Updated description"
composerData.Require["new/package"] = "^1.0"

// Write back to file
err = comp.WriteComposerJSON(composerData)
if err != nil {
    log.Printf("Failed to write composer.json: %v", err)
}
```

### InitProject

Initializes a new composer.json file in the current directory.

```go
func (c *Composer) InitProject(options InitOptions) error
```

**Parameters:**
- `options` - Project initialization options

**Example:**
```go
options := InitOptions{
    Name:        "vendor/my-package",
    Description: "My awesome PHP package",
    Type:        "library",
    License:     "MIT",
    Authors: []Author{
        {
            Name:  "John Doe",
            Email: "john@example.com",
        },
    },
    MinimumStability: "stable",
    PreferStable:     true,
}

err := comp.InitProject(options)
if err != nil {
    log.Printf("Failed to initialize project: %v", err)
}
```

## Archive Management

### Archive

Creates an archive of the project.

```go
func (c *Composer) Archive(format string, file string) error
```

**Parameters:**
- `format` - Archive format ("tar", "zip")
- `file` - Output file path

**Returns:**
- `error` - Error if archive creation fails

**Example:**
```go
// Create a ZIP archive
err := comp.Archive("zip", "my-project.zip")

// Create a TAR archive
err = comp.Archive("tar", "my-project.tar")

if err != nil {
    log.Printf("Failed to create archive: %v", err)
}
```

### ArchivePackage

Creates an archive of a specific package.

```go
func (c *Composer) ArchivePackage(packageName string, version string, format string, file string) error
```

**Parameters:**
- `packageName` - Name of the package to archive
- `version` - Version to archive
- `format` - Archive format
- `file` - Output file path

**Example:**
```go
err := comp.ArchivePackage("symfony/console", "^6.0", "zip", "symfony-console.zip")
```

## Type Definitions

### CreateProjectOptions

```go
type CreateProjectOptions struct {
    PackageName        string
    Directory          string
    Version            string
    PreferDist         bool
    PreferSource       bool
    NoScripts          bool
    NoProgress         bool
    KeepVcs            bool
    NoInstall          bool
    IgnorePlatformReqs bool
}
```

### ValidateOptions

```go
type ValidateOptions struct {
    NoCheckAll       bool
    NoCheckLock      bool
    NoCheckPublish   bool
    WithDependencies bool
    Strict           bool
}
```

### AutoloadOptions

```go
type AutoloadOptions struct {
    Optimize   bool
    Classmap   bool
    Apcu       bool
    ApcuPrefix string
    NoDev      bool
}
```

### InitOptions

```go
type InitOptions struct {
    Name             string
    Description      string
    Type             string
    License          string
    Authors          []Author
    MinimumStability string
    PreferStable     bool
    Require          map[string]string
    RequireDev       map[string]string
}
```

## Best Practices

1. **Always validate composer.json** after making changes
2. **Use version constraints** appropriately in project templates
3. **Regenerate autoloader** after adding new classes
4. **Keep scripts organized** and well-documented
5. **Use appropriate archive formats** for distribution
6. **Set meaningful project metadata** during initialization
7. **Handle configuration changes** carefully in production environments
