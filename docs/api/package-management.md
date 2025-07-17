# Package Management API

The package management API provides comprehensive functionality for managing PHP packages and dependencies.

## Installation Methods

### Install

Installs all dependencies defined in composer.json.

```go
func (c *Composer) Install(noDev bool, optimize bool) error
```

**Parameters:**
- `noDev` - Skip installing dev dependencies if true
- `optimize` - Generate optimized autoloader if true

**Returns:**
- `error` - Error if installation fails

**Example:**
```go
// Install all dependencies including dev
err := comp.Install(false, false)

// Install only production dependencies with optimization
err := comp.Install(true, true)

if err != nil {
    log.Printf("Installation failed: %v", err)
}
```

### InstallWithContext

Installs dependencies with context support for cancellation and timeouts.

```go
func (c *Composer) InstallWithContext(ctx context.Context, noDev bool, optimize bool) error
```

**Parameters:**
- `ctx` - Context for cancellation and timeout
- `noDev` - Skip installing dev dependencies if true
- `optimize` - Generate optimized autoloader if true

**Returns:**
- `error` - Error if installation fails or context is cancelled

**Example:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
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

### InstallWithOptions

Installs dependencies with advanced options.

```go
func (c *Composer) InstallWithOptions(options InstallOptions) error
```

**Parameters:**
- `options` - Advanced installation options

**Example:**
```go
options := InstallOptions{
    NoDev:           false,
    Optimize:        true,
    PreferDist:      true,
    NoScripts:       false,
    NoProgress:      false,
    NoSuggest:       true,
    PreferStable:    true,
    PreferLowest:    false,
}

err := comp.InstallWithOptions(options)
```

## Update Methods

### Update

Updates all packages to their latest versions within constraints.

```go
func (c *Composer) Update(noDev bool, optimize bool) error
```

**Parameters:**
- `noDev` - Skip updating dev dependencies if true
- `optimize` - Generate optimized autoloader if true

**Returns:**
- `error` - Error if update fails

**Example:**
```go
// Update all packages
err := comp.Update(false, false)

// Update only production packages with optimization
err := comp.Update(true, true)
```

### UpdateWithContext

Updates packages with context support.

```go
func (c *Composer) UpdateWithContext(ctx context.Context, noDev bool, optimize bool) error
```

### UpdatePackage

Updates a specific package.

```go
func (c *Composer) UpdatePackage(packageName string) error
```

**Parameters:**
- `packageName` - Name of the package to update

**Example:**
```go
err := comp.UpdatePackage("symfony/console")
if err != nil {
    log.Printf("Failed to update package: %v", err)
}
```

### UpdatePackages

Updates multiple specific packages.

```go
func (c *Composer) UpdatePackages(packageNames []string) error
```

**Parameters:**
- `packageNames` - Array of package names to update

**Example:**
```go
packages := []string{"symfony/console", "monolog/monolog", "doctrine/orm"}
err := comp.UpdatePackages(packages)
```

## Package Addition Methods

### RequirePackage

Adds a new package dependency.

```go
func (c *Composer) RequirePackage(packageName string, version string) error
```

**Parameters:**
- `packageName` - Name of the package to add
- `version` - Version constraint (e.g., "^3.0", "~2.1", "1.0.0")

**Returns:**
- `error` - Error if package addition fails

**Example:**
```go
// Add a package with version constraint
err := comp.RequirePackage("monolog/monolog", "^3.0")

// Add a specific version
err := comp.RequirePackage("symfony/console", "6.3.0")

// Add latest version
err := comp.RequirePackage("guzzlehttp/guzzle", "*")
```

### RequireDevPackage

Adds a development dependency.

```go
func (c *Composer) RequireDevPackage(packageName string, version string) error
```

**Parameters:**
- `packageName` - Name of the dev package to add
- `version` - Version constraint

**Example:**
```go
// Add development dependencies
err := comp.RequireDevPackage("phpunit/phpunit", "^10.0")
err = comp.RequireDevPackage("symfony/var-dumper", "^6.0")
```

### RequirePackages

Adds multiple packages at once.

```go
func (c *Composer) RequirePackages(packages map[string]string) error
```

**Parameters:**
- `packages` - Map of package names to version constraints

**Example:**
```go
packages := map[string]string{
    "symfony/console":    "^6.0",
    "monolog/monolog":    "^3.0",
    "guzzlehttp/guzzle":  "^7.0",
}

err := comp.RequirePackages(packages)
```

## Package Removal Methods

### RemovePackage

Removes a package dependency.

```go
func (c *Composer) RemovePackage(packageName string) error
```

**Parameters:**
- `packageName` - Name of the package to remove

**Returns:**
- `error` - Error if package removal fails

**Example:**
```go
err := comp.RemovePackage("old-package/deprecated")
if err != nil {
    log.Printf("Failed to remove package: %v", err)
}
```

### RemovePackages

Removes multiple packages at once.

```go
func (c *Composer) RemovePackages(packageNames []string) error
```

**Parameters:**
- `packageNames` - Array of package names to remove

**Example:**
```go
packagesToRemove := []string{
    "old-package/deprecated",
    "unused/library",
    "legacy/component",
}

err := comp.RemovePackages(packagesToRemove)
```

### RemoveDevPackage

Removes a development dependency.

```go
func (c *Composer) RemoveDevPackage(packageName string) error
```

## Package Information Methods

### ShowAllPackages

Lists all installed packages.

```go
func (c *Composer) ShowAllPackages() (string, error)
```

**Returns:**
- `string` - List of installed packages
- `error` - Error if listing fails

**Example:**
```go
packages, err := comp.ShowAllPackages()
if err != nil {
    log.Printf("Failed to list packages: %v", err)
    return
}
fmt.Println("Installed packages:")
fmt.Println(packages)
```

### ShowPackage

Shows detailed information about a specific package.

```go
func (c *Composer) ShowPackage(packageName string) (string, error)
```

**Parameters:**
- `packageName` - Name of the package to show

**Returns:**
- `string` - Package information
- `error` - Error if package info cannot be retrieved

**Example:**
```go
info, err := comp.ShowPackage("symfony/console")
if err != nil {
    log.Printf("Failed to get package info: %v", err)
    return
}
fmt.Printf("Package info:\n%s\n", info)
```

### ShowDependencyTree

Shows the dependency tree for packages.

```go
func (c *Composer) ShowDependencyTree(packageName string) (string, error)
```

**Parameters:**
- `packageName` - Package name (empty string for all packages)

**Returns:**
- `string` - Dependency tree
- `error` - Error if tree cannot be generated

**Example:**
```go
// Show full dependency tree
tree, err := comp.ShowDependencyTree("")

// Show tree for specific package
tree, err := comp.ShowDependencyTree("symfony/console")
```

### SearchPackages

Searches for packages in repositories.

```go
func (c *Composer) SearchPackages(query string) (string, error)
```

**Parameters:**
- `query` - Search query

**Returns:**
- `string` - Search results
- `error` - Error if search fails

**Example:**
```go
results, err := comp.SearchPackages("logging")
if err != nil {
    log.Printf("Search failed: %v", err)
    return
}
fmt.Printf("Search results:\n%s\n", results)
```

## Dependency Analysis Methods

### WhyPackage

Shows why a package is installed (which packages depend on it).

```go
func (c *Composer) WhyPackage(packageName string) (string, error)
```

**Parameters:**
- `packageName` - Name of the package to analyze

**Example:**
```go
reasons, err := comp.WhyPackage("psr/log")
if err != nil {
    log.Printf("Failed to analyze package: %v", err)
    return
}
fmt.Printf("Package dependencies:\n%s\n", reasons)
```

### WhyNotPackage

Shows why a package cannot be installed.

```go
func (c *Composer) WhyNotPackage(packageName string, version string) (string, error)
```

**Parameters:**
- `packageName` - Name of the package
- `version` - Version constraint

**Example:**
```go
reasons, err := comp.WhyNotPackage("symfony/console", "^7.0")
if err != nil {
    log.Printf("Analysis failed: %v", err)
    return
}
fmt.Printf("Conflicts:\n%s\n", reasons)
```

### DependsPackage

Shows packages that depend on the specified package.

```go
func (c *Composer) DependsPackage(packageName string) (string, error)
```

### ProhibitsPackage

Shows packages that prohibit the specified package.

```go
func (c *Composer) ProhibitsPackage(packageName string) (string, error)
```

## Outdated Package Methods

### ShowOutdated

Shows outdated packages that can be updated.

```go
func (c *Composer) ShowOutdated() (string, error)
```

**Returns:**
- `string` - List of outdated packages
- `error` - Error if check fails

**Example:**
```go
outdated, err := comp.ShowOutdated()
if err != nil {
    log.Printf("Failed to check outdated packages: %v", err)
    return
}

if outdated != "" {
    fmt.Printf("Outdated packages:\n%s\n", outdated)
} else {
    fmt.Println("All packages are up to date!")
}
```

### ShowOutdatedDirect

Shows only direct dependencies that are outdated.

```go
func (c *Composer) ShowOutdatedDirect() (string, error)
```

## Advanced Options

### InstallOptions

Advanced options for package installation.

```go
type InstallOptions struct {
    NoDev           bool   // Skip dev dependencies
    Optimize        bool   // Generate optimized autoloader
    PreferDist      bool   // Prefer dist over source
    PreferSource    bool   // Prefer source over dist
    NoScripts       bool   // Skip running scripts
    NoProgress      bool   // Disable progress display
    NoSuggest       bool   // Skip package suggestions
    PreferStable    bool   // Prefer stable versions
    PreferLowest    bool   // Prefer lowest versions
    DryRun          bool   // Simulate the operation
    Verbose         bool   // Increase verbosity
}
```

## Best Practices

1. **Use version constraints** instead of exact versions for flexibility
2. **Separate production and development dependencies** appropriately
3. **Use context-aware methods** for long-running operations
4. **Check for outdated packages** regularly
5. **Understand dependency relationships** using analysis methods
6. **Use batch operations** when working with multiple packages
7. **Handle errors gracefully** and provide meaningful feedback
