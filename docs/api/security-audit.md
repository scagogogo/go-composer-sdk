# Security & Audit API

The security and audit API provides comprehensive functionality for analyzing dependencies, detecting vulnerabilities, and ensuring project security.

## Security Audit Methods

### Audit

Performs a security audit of project dependencies.

```go
func (c *Composer) Audit() (string, error)
```

**Returns:**
- `string` - Audit results in text format
- `error` - Error if audit fails

**Example:**
```go
auditOutput, err := comp.Audit()
if err != nil {
    log.Printf("Security audit failed: %v", err)
    return
}

fmt.Printf("Security audit results:\n%s\n", auditOutput)
```

### AuditWithJSON

Performs a security audit and returns structured JSON results.

```go
func (c *Composer) AuditWithJSON() (*AuditResult, error)
```

**Returns:**
- `*AuditResult` - Structured audit results
- `error` - Error if audit fails

**Example:**
```go
result, err := comp.AuditWithJSON()
if err != nil {
    log.Printf("Security audit failed: %v", err)
    return
}

fmt.Printf("Found %d vulnerabilities\n", result.Found)
for _, vuln := range result.Vulnerabilities {
    fmt.Printf("🔴 %s\n", vuln.Title)
    fmt.Printf("   Package: %s\n", vuln.Package)
    fmt.Printf("   Severity: %s\n", vuln.Severity)
    fmt.Printf("   Description: %s\n", vuln.Description)
}
```

### GetHighSeverityVulnerabilities

Gets only high and critical severity vulnerabilities.

```go
func (c *Composer) GetHighSeverityVulnerabilities() ([]Vulnerability, error)
```

**Returns:**
- `[]Vulnerability` - Array of high/critical vulnerabilities
- `error` - Error if retrieval fails

**Example:**
```go
highVulns, err := comp.GetHighSeverityVulnerabilities()
if err != nil {
    log.Printf("Failed to get high severity vulnerabilities: %v", err)
    return
}

if len(highVulns) > 0 {
    fmt.Printf("⚠️  Found %d high/critical vulnerabilities:\n", len(highVulns))
    for _, vuln := range highVulns {
        fmt.Printf("- %s (%s): %s\n", vuln.Package, vuln.Severity, vuln.Title)
    }
} else {
    fmt.Println("✅ No high/critical vulnerabilities found")
}
```

### HasVulnerabilities

Checks if the project has any security vulnerabilities.

```go
func (c *Composer) HasVulnerabilities() (bool, error)
```

**Returns:**
- `bool` - True if vulnerabilities exist
- `error` - Error if check fails

**Example:**
```go
hasVulns, err := comp.HasVulnerabilities()
if err != nil {
    log.Printf("Failed to check vulnerabilities: %v", err)
    return
}

if hasVulns {
    fmt.Println("⚠️  Project has security vulnerabilities")
    // Perform detailed audit
    result, _ := comp.AuditWithJSON()
    // Handle vulnerabilities...
} else {
    fmt.Println("✅ No vulnerabilities found")
}
```

## Dependency Analysis

### CheckDependencies

Analyzes and validates project dependencies.

```go
func (c *Composer) CheckDependencies() (string, error)
```

**Returns:**
- `string` - Dependency analysis results
- `error` - Error if analysis fails

**Example:**
```go
analysis, err := comp.CheckDependencies()
if err != nil {
    log.Printf("Dependency analysis failed: %v", err)
    return
}

fmt.Printf("Dependency analysis:\n%s\n", analysis)
```

### GetAbandonedPackages

Gets a list of abandoned packages in the project.

```go
func (c *Composer) GetAbandonedPackages() ([]string, error)
```

**Returns:**
- `[]string` - Array of abandoned package names
- `error` - Error if retrieval fails

**Example:**
```go
abandoned, err := comp.GetAbandonedPackages()
if err != nil {
    log.Printf("Failed to get abandoned packages: %v", err)
    return
}

if len(abandoned) > 0 {
    fmt.Printf("⚠️  Found %d abandoned packages:\n", len(abandoned))
    for _, pkg := range abandoned {
        fmt.Printf("- %s\n", pkg)
    }
    fmt.Println("Consider finding alternatives for these packages.")
} else {
    fmt.Println("✅ No abandoned packages found")
}
```

### WhyPackage

Explains why a package is installed (dependency chain).

```go
func (c *Composer) WhyPackage(packageName string) (string, error)
```

**Parameters:**
- `packageName` - Name of the package to analyze

**Returns:**
- `string` - Dependency explanation
- `error` - Error if analysis fails

**Example:**
```go
explanation, err := comp.WhyPackage("psr/log")
if err != nil {
    log.Printf("Failed to analyze package: %v", err)
    return
}

fmt.Printf("Why 'psr/log' is installed:\n%s\n", explanation)
```

### WhyNotPackage

Explains why a package cannot be installed.

```go
func (c *Composer) WhyNotPackage(packageName string, version string) (string, error)
```

**Parameters:**
- `packageName` - Name of the package
- `version` - Version constraint to test

**Returns:**
- `string` - Conflict explanation
- `error` - Error if analysis fails

**Example:**
```go
conflicts, err := comp.WhyNotPackage("symfony/console", "^7.0")
if err != nil {
    log.Printf("Failed to analyze conflicts: %v", err)
    return
}

fmt.Printf("Why 'symfony/console ^7.0' cannot be installed:\n%s\n", conflicts)
```

## Platform Security

### CheckPlatformReqs

Checks if platform requirements are satisfied.

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
    // Handle platform requirement issues
} else {
    fmt.Println("✅ All platform requirements satisfied")
}
```

### IsPlatformAvailable

Checks if a specific platform requirement is available.

```go
func (c *Composer) IsPlatformAvailable(platform string, version string) (bool, error)
```

**Parameters:**
- `platform` - Platform name (e.g., "php", "ext-mbstring")
- `version` - Required version

**Returns:**
- `bool` - True if platform requirement is satisfied
- `error` - Error if check fails

**Example:**
```go
// Check PHP version
available, err := comp.IsPlatformAvailable("php", "8.1")
if err != nil {
    log.Printf("Failed to check PHP version: %v", err)
} else if available {
    fmt.Println("✅ PHP 8.1+ is available")
} else {
    fmt.Println("❌ PHP 8.1+ is not available")
}

// Check extension
available, err = comp.IsPlatformAvailable("ext-mbstring", "*")
if err != nil {
    log.Printf("Failed to check mbstring extension: %v", err)
} else if available {
    fmt.Println("✅ mbstring extension is available")
} else {
    fmt.Println("❌ mbstring extension is not available")
}
```

## Type Definitions

### AuditResult

```go
type AuditResult struct {
    Vulnerabilities []Vulnerability `json:"vulnerabilities"`
    Found          int             `json:"found"`
}
```

### Vulnerability

```go
type Vulnerability struct {
    ID          string   `json:"id"`
    Title       string   `json:"title"`
    Package     string   `json:"package"`
    Version     string   `json:"version"`
    Severity    string   `json:"severity"`
    Description string   `json:"description"`
    References  []string `json:"references"`
    Link        string   `json:"link"`
    CVE         string   `json:"cve,omitempty"`
    CWE         string   `json:"cwe,omitempty"`
}
```

## Security Best Practices

### 1. Regular Security Audits

```go
func performRegularAudit(comp *composer.Composer) error {
    // Check for vulnerabilities
    hasVulns, err := comp.HasVulnerabilities()
    if err != nil {
        return err
    }
    
    if hasVulns {
        // Get detailed audit results
        result, err := comp.AuditWithJSON()
        if err != nil {
            return err
        }
        
        // Focus on high severity issues first
        highVulns, err := comp.GetHighSeverityVulnerabilities()
        if err != nil {
            return err
        }
        
        if len(highVulns) > 0 {
            log.Printf("URGENT: %d high/critical vulnerabilities found", len(highVulns))
            // Send alerts, create tickets, etc.
        }
        
        // Log all vulnerabilities for tracking
        for _, vuln := range result.Vulnerabilities {
            log.Printf("Vulnerability: %s in %s (%s)", vuln.Title, vuln.Package, vuln.Severity)
        }
    }
    
    return nil
}
```

### 2. Abandoned Package Detection

```go
func checkAbandonedPackages(comp *composer.Composer) error {
    abandoned, err := comp.GetAbandonedPackages()
    if err != nil {
        return err
    }
    
    if len(abandoned) > 0 {
        log.Printf("Warning: %d abandoned packages found", len(abandoned))
        for _, pkg := range abandoned {
            log.Printf("Abandoned package: %s", pkg)
            // Research alternatives, plan migration
        }
    }
    
    return nil
}
```

### 3. Platform Requirement Validation

```go
func validatePlatformRequirements(comp *composer.Composer) error {
    // Check overall platform requirements
    err := comp.CheckPlatformReqs()
    if err != nil {
        log.Printf("Platform requirements not satisfied: %v", err)
        return err
    }
    
    // Check specific critical requirements
    criticalRequirements := map[string]string{
        "php":           "8.1",
        "ext-mbstring":  "*",
        "ext-openssl":   "*",
        "ext-pdo":       "*",
        "ext-json":      "*",
    }
    
    for platform, version := range criticalRequirements {
        available, err := comp.IsPlatformAvailable(platform, version)
        if err != nil {
            log.Printf("Failed to check %s: %v", platform, err)
            continue
        }
        
        if !available {
            log.Printf("Critical requirement missing: %s %s", platform, version)
            return fmt.Errorf("missing critical requirement: %s %s", platform, version)
        }
    }
    
    return nil
}
```

### 4. Comprehensive Security Check

```go
func comprehensiveSecurityCheck(comp *composer.Composer) error {
    fmt.Println("🔍 Starting comprehensive security check...")
    
    // 1. Platform requirements
    fmt.Println("Checking platform requirements...")
    if err := validatePlatformRequirements(comp); err != nil {
        return fmt.Errorf("platform requirements check failed: %w", err)
    }
    fmt.Println("✅ Platform requirements satisfied")
    
    // 2. Security audit
    fmt.Println("Performing security audit...")
    if err := performRegularAudit(comp); err != nil {
        return fmt.Errorf("security audit failed: %w", err)
    }
    fmt.Println("✅ Security audit completed")
    
    // 3. Abandoned packages
    fmt.Println("Checking for abandoned packages...")
    if err := checkAbandonedPackages(comp); err != nil {
        return fmt.Errorf("abandoned package check failed: %w", err)
    }
    fmt.Println("✅ Abandoned package check completed")
    
    // 4. Dependency analysis
    fmt.Println("Analyzing dependencies...")
    analysis, err := comp.CheckDependencies()
    if err != nil {
        return fmt.Errorf("dependency analysis failed: %w", err)
    }
    if analysis != "" {
        fmt.Printf("Dependency analysis results:\n%s\n", analysis)
    }
    fmt.Println("✅ Dependency analysis completed")
    
    fmt.Println("🎉 Comprehensive security check completed successfully!")
    return nil
}
```

## Integration with CI/CD

### GitHub Actions Example

```yaml
name: Security Audit

on:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  security-audit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          
      - name: Run Security Audit
        run: |
          go run security-audit.go
          
      - name: Upload Security Report
        if: failure()
        uses: actions/upload-artifact@v3
        with:
          name: security-report
          path: security-report.json
```

## Error Handling

Always handle security-related errors appropriately:

```go
// Good error handling for security operations
result, err := comp.AuditWithJSON()
if err != nil {
    // Log the error but don't expose sensitive details
    log.Printf("Security audit failed: %v", err)
    
    // In production, you might want to:
    // - Send alerts to security team
    // - Create incident tickets
    // - Fail the deployment pipeline
    return fmt.Errorf("security audit failed")
}

// Process results even if no vulnerabilities found
if result.Found == 0 {
    log.Println("No vulnerabilities found in security audit")
} else {
    log.Printf("Security audit found %d vulnerabilities", result.Found)
    // Handle vulnerabilities appropriately
}
```
