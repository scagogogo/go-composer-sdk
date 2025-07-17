# CI/CD 集成示例

本页面展示如何将 Go Composer SDK 集成到 CI/CD 流水线中，实现自动化的 PHP 项目构建、测试和部署。

## GitHub Actions 集成

### 基本 GitHub Actions 工作流

```yaml
# .github/workflows/php-ci.yml
name: PHP CI with Go Composer SDK

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        php-version: [8.1, 8.2, 8.3]
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup PHP
      uses: shivammathur/setup-php@v2
      with:
        php-version: ${{ matrix.php-version }}
        extensions: mbstring, xml, ctype, iconv, intl, pdo_sqlite
        coverage: xdebug
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Install Go Composer SDK
      run: go mod init ci-test && go get github.com/scagogogo/go-composer-sdk
    
    - name: Run PHP CI with Go SDK
      run: |
        cat > ci.go << 'EOF'
        package main
        
        import (
            "fmt"
            "log"
            "os"
            
            "github.com/scagogogo/go-composer-sdk/pkg/composer"
        )
        
        func main() {
            comp, err := composer.New(composer.DefaultOptions())
            if err != nil {
                log.Fatalf("创建 Composer 实例失败: %v", err)
            }
            
            comp.SetWorkingDir(".")
            
            // CI 环境配置
            comp.SetEnv([]string{
                "COMPOSER_NO_INTERACTION=1",
                "COMPOSER_PREFER_STABLE=true",
                "COMPOSER_OPTIMIZE_AUTOLOADER=true",
            })
            
            // 验证项目
            fmt.Println("🔍 验证项目配置...")
            if err := comp.Validate(); err != nil {
                log.Fatalf("项目验证失败: %v", err)
            }
            
            // 安装依赖
            fmt.Println("📦 安装依赖...")
            if err := comp.Install(false, true); err != nil {
                log.Fatalf("依赖安装失败: %v", err)
            }
            
            // 安全审计
            fmt.Println("🔒 执行安全审计...")
            if result, err := comp.Audit(); err != nil {
                log.Fatalf("安全审计失败: %v", err)
            } else if result != "" {
                fmt.Printf("⚠️  安全问题:\n%s\n", result)
                if os.Getenv("FAIL_ON_SECURITY") == "true" {
                    os.Exit(1)
                }
            }
            
            // 运行测试
            fmt.Println("🧪 运行测试...")
            if err := comp.RunScript("test"); err != nil {
                log.Fatalf("测试失败: %v", err)
            }
            
            fmt.Println("✅ CI 流程完成")
        }
        EOF
        
        go run ci.go
```

### 高级 GitHub Actions 工作流

```yaml
# .github/workflows/advanced-php-ci.yml
name: Advanced PHP CI/CD

on:
  push:
    branches: [ main ]
    tags: [ 'v*' ]
  pull_request:
    branches: [ main ]

env:
  COMPOSER_CACHE_DIR: /tmp/composer-cache

jobs:
  validate:
    name: Validate Project
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup PHP
      uses: shivammathur/setup-php@v2
      with:
        php-version: 8.2
        extensions: mbstring, xml, ctype, iconv, intl
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Cache Composer packages
      uses: actions/cache@v3
      with:
        path: ${{ env.COMPOSER_CACHE_DIR }}
        key: ${{ runner.os }}-composer-${{ hashFiles('**/composer.lock') }}
        restore-keys: ${{ runner.os }}-composer-
    
    - name: Validate with Go SDK
      run: |
        go mod init validation
        go get github.com/scagogogo/go-composer-sdk
        
        cat > validate.go << 'EOF'
        package main
        
        import (
            "fmt"
            "log"
            
            "github.com/scagogogo/go-composer-sdk/pkg/composer"
        )
        
        func main() {
            comp, err := composer.New(composer.DefaultOptions())
            if err != nil {
                log.Fatalf("创建实例失败: %v", err)
            }
            
            comp.SetWorkingDir(".")
            comp.SetEnv([]string{
                "COMPOSER_CACHE_DIR=/tmp/composer-cache",
                "COMPOSER_NO_INTERACTION=1",
            })
            
            fmt.Println("🔍 验证 composer.json...")
            if err := comp.Validate(); err != nil {
                log.Fatalf("验证失败: %v", err)
            }
            
            fmt.Println("🔍 检查平台要求...")
            if err := comp.CheckPlatformReqs(); err != nil {
                log.Fatalf("平台要求检查失败: %v", err)
            }
            
            fmt.Println("✅ 验证完成")
        }
        EOF
        
        go run validate.go

  test:
    name: Test Suite
    runs-on: ubuntu-latest
    needs: validate
    
    strategy:
      matrix:
        php-version: [8.1, 8.2, 8.3]
        dependency-version: [prefer-lowest, prefer-stable]
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup PHP ${{ matrix.php-version }}
      uses: shivammathur/setup-php@v2
      with:
        php-version: ${{ matrix.php-version }}
        extensions: mbstring, xml, ctype, iconv, intl, pdo_sqlite
        coverage: xdebug
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Cache Composer packages
      uses: actions/cache@v3
      with:
        path: ${{ env.COMPOSER_CACHE_DIR }}
        key: ${{ runner.os }}-php${{ matrix.php-version }}-${{ matrix.dependency-version }}-${{ hashFiles('**/composer.lock') }}
    
    - name: Install and Test with Go SDK
      run: |
        go mod init testing
        go get github.com/scagogogo/go-composer-sdk
        
        cat > test.go << 'EOF'
        package main
        
        import (
            "fmt"
            "log"
            "os"
            
            "github.com/scagogogo/go-composer-sdk/pkg/composer"
        )
        
        func main() {
            comp, err := composer.New(composer.DefaultOptions())
            if err != nil {
                log.Fatalf("创建实例失败: %v", err)
            }
            
            comp.SetWorkingDir(".")
            comp.SetEnv([]string{
                "COMPOSER_CACHE_DIR=/tmp/composer-cache",
                "COMPOSER_NO_INTERACTION=1",
                "COMPOSER_OPTIMIZE_AUTOLOADER=true",
            })
            
            // 根据矩阵参数安装依赖
            dependencyVersion := os.Getenv("DEPENDENCY_VERSION")
            fmt.Printf("📦 安装依赖 (%s)...\n", dependencyVersion)
            
            if dependencyVersion == "prefer-lowest" {
                comp.SetEnv(append(comp.GetEnv(), "COMPOSER_PREFER_LOWEST=1"))
            }
            
            if err := comp.Install(false, true); err != nil {
                log.Fatalf("依赖安装失败: %v", err)
            }
            
            // 运行测试
            fmt.Println("🧪 运行测试套件...")
            if err := comp.RunScript("test"); err != nil {
                log.Fatalf("测试失败: %v", err)
            }
            
            fmt.Println("✅ 测试完成")
        }
        EOF
        
        DEPENDENCY_VERSION=${{ matrix.dependency-version }} go run test.go

  security:
    name: Security Audit
    runs-on: ubuntu-latest
    needs: validate
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup PHP
      uses: shivammathur/setup-php@v2
      with:
        php-version: 8.2
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Security Audit with Go SDK
      run: |
        go mod init security
        go get github.com/scagogogo/go-composer-sdk
        
        cat > security.go << 'EOF'
        package main
        
        import (
            "fmt"
            "log"
            "os"
            
            "github.com/scagogogo/go-composer-sdk/pkg/composer"
        )
        
        func main() {
            comp, err := composer.New(composer.DefaultOptions())
            if err != nil {
                log.Fatalf("创建实例失败: %v", err)
            }
            
            comp.SetWorkingDir(".")
            comp.SetEnv([]string{
                "COMPOSER_NO_INTERACTION=1",
                "COMPOSER_AUDIT_ABANDONED=report",
            })
            
            // 安装依赖
            fmt.Println("📦 安装依赖...")
            if err := comp.Install(false, false); err != nil {
                log.Fatalf("依赖安装失败: %v", err)
            }
            
            // 安全审计
            fmt.Println("🔒 执行安全审计...")
            result, err := comp.Audit()
            if err != nil {
                log.Fatalf("安全审计失败: %v", err)
            }
            
            if result != "" {
                fmt.Printf("⚠️  发现安全问题:\n%s\n", result)
                
                // 创建安全报告
                if err := os.WriteFile("security-report.txt", []byte(result), 0644); err != nil {
                    log.Printf("写入安全报告失败: %v", err)
                }
                
                // 在 PR 中，安全问题不应该阻止构建
                if os.Getenv("GITHUB_EVENT_NAME") == "pull_request" {
                    fmt.Println("⚠️  PR 中发现安全问题，但不阻止构建")
                } else {
                    os.Exit(1)
                }
            } else {
                fmt.Println("✅ 未发现安全问题")
            }
        }
        EOF
        
        go run security.go
    
    - name: Upload Security Report
      if: failure()
      uses: actions/upload-artifact@v3
      with:
        name: security-report
        path: security-report.txt

  deploy:
    name: Deploy
    runs-on: ubuntu-latest
    needs: [test, security]
    if: github.ref == 'refs/heads/main' && github.event_name == 'push'
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup PHP
      uses: shivammathur/setup-php@v2
      with:
        php-version: 8.2
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Production Build with Go SDK
      run: |
        go mod init deploy
        go get github.com/scagogogo/go-composer-sdk
        
        cat > deploy.go << 'EOF'
        package main
        
        import (
            "fmt"
            "log"
            
            "github.com/scagogogo/go-composer-sdk/pkg/composer"
        )
        
        func main() {
            comp, err := composer.New(composer.DefaultOptions())
            if err != nil {
                log.Fatalf("创建实例失败: %v", err)
            }
            
            comp.SetWorkingDir(".")
            comp.SetEnv([]string{
                "COMPOSER_NO_INTERACTION=1",
                "COMPOSER_OPTIMIZE_AUTOLOADER=true",
                "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
                "COMPOSER_PREFER_STABLE=true",
            })
            
            // 生产环境安装
            fmt.Println("🏭 生产环境依赖安装...")
            if err := comp.Install(true, true); err != nil { // noDev=true, optimize=true
                log.Fatalf("生产环境安装失败: %v", err)
            }
            
            // 优化自动加载
            fmt.Println("⚡ 优化自动加载...")
            if err := comp.DumpAutoload(true, true); err != nil {
                log.Fatalf("自动加载优化失败: %v", err)
            }
            
            fmt.Println("✅ 生产环境构建完成")
        }
        EOF
        
        go run deploy.go
    
    - name: Deploy to Production
      run: |
        echo "🚀 部署到生产环境..."
        # 这里添加实际的部署逻辑
```

## GitLab CI 集成

### GitLab CI 配置

```yaml
# .gitlab-ci.yml
stages:
  - validate
  - test
  - security
  - deploy

variables:
  COMPOSER_CACHE_DIR: "$CI_PROJECT_DIR/.composer-cache"
  GO_VERSION: "1.21"

cache:
  paths:
    - .composer-cache/
    - vendor/

before_script:
  - apt-get update -qq && apt-get install -y -qq git curl
  - curl -sSfL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzC /usr/local
  - export PATH=$PATH:/usr/local/go/bin
  - go version

validate:
  stage: validate
  image: php:8.2-cli
  script:
    - go mod init gitlab-ci
    - go get github.com/scagogogo/go-composer-sdk
    - |
      cat > validate.go << 'EOF'
      package main
      
      import (
          "fmt"
          "log"
          
          "github.com/scagogogo/go-composer-sdk/pkg/composer"
      )
      
      func main() {
          comp, err := composer.New(composer.DefaultOptions())
          if err != nil {
              log.Fatalf("创建实例失败: %v", err)
          }
          
          comp.SetWorkingDir(".")
          comp.SetEnv([]string{
              "COMPOSER_CACHE_DIR=" + "$CI_PROJECT_DIR/.composer-cache",
              "COMPOSER_NO_INTERACTION=1",
          })
          
          fmt.Println("🔍 验证项目...")
          if err := comp.Validate(); err != nil {
              log.Fatalf("验证失败: %v", err)
          }
          
          fmt.Println("✅ 验证完成")
      }
      EOF
    - go run validate.go

test:
  stage: test
  image: php:8.2-cli
  parallel:
    matrix:
      - PHP_VERSION: ["8.1", "8.2", "8.3"]
  before_script:
    - apt-get update -qq && apt-get install -y -qq git curl
    - curl -sSfL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzC /usr/local
    - export PATH=$PATH:/usr/local/go/bin
    - |
      if [ "$PHP_VERSION" != "8.2" ]; then
        curl -sSfL https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions -o /usr/local/bin/install-php-extensions
        chmod +x /usr/local/bin/install-php-extensions
        install-php-extensions @composer
      fi
  script:
    - go mod init test-ci
    - go get github.com/scagogogo/go-composer-sdk
    - |
      cat > test.go << 'EOF'
      package main
      
      import (
          "fmt"
          "log"
          
          "github.com/scagogogo/go-composer-sdk/pkg/composer"
      )
      
      func main() {
          comp, err := composer.New(composer.DefaultOptions())
          if err != nil {
              log.Fatalf("创建实例失败: %v", err)
          }
          
          comp.SetWorkingDir(".")
          comp.SetEnv([]string{
              "COMPOSER_CACHE_DIR=" + "$CI_PROJECT_DIR/.composer-cache",
              "COMPOSER_NO_INTERACTION=1",
          })
          
          fmt.Println("📦 安装依赖...")
          if err := comp.Install(false, true); err != nil {
              log.Fatalf("安装失败: %v", err)
          }
          
          fmt.Println("🧪 运行测试...")
          if err := comp.RunScript("test"); err != nil {
              log.Fatalf("测试失败: %v", err)
          }
          
          fmt.Println("✅ 测试完成")
      }
      EOF
    - go run test.go

security:
  stage: security
  image: php:8.2-cli
  script:
    - go mod init security-ci
    - go get github.com/scagogogo/go-composer-sdk
    - |
      cat > security.go << 'EOF'
      package main
      
      import (
          "fmt"
          "log"
          "os"
          
          "github.com/scagogogo/go-composer-sdk/pkg/composer"
      )
      
      func main() {
          comp, err := composer.New(composer.DefaultOptions())
          if err != nil {
              log.Fatalf("创建实例失败: %v", err)
          }
          
          comp.SetWorkingDir(".")
          comp.SetEnv([]string{
              "COMPOSER_CACHE_DIR=" + "$CI_PROJECT_DIR/.composer-cache",
              "COMPOSER_NO_INTERACTION=1",
              "COMPOSER_AUDIT_ABANDONED=report",
          })
          
          fmt.Println("📦 安装依赖...")
          if err := comp.Install(false, false); err != nil {
              log.Fatalf("安装失败: %v", err)
          }
          
          fmt.Println("🔒 安全审计...")
          result, err := comp.Audit()
          if err != nil {
              log.Fatalf("安全审计失败: %v", err)
          }
          
          if result != "" {
              fmt.Printf("⚠️  安全问题:\n%s\n", result)
              os.WriteFile("security-report.txt", []byte(result), 0644)
              os.Exit(1)
          }
          
          fmt.Println("✅ 安全检查通过")
      }
      EOF
    - go run security.go
  artifacts:
    when: on_failure
    paths:
      - security-report.txt
    expire_in: 1 week

deploy:
  stage: deploy
  image: php:8.2-cli
  script:
    - go mod init deploy-ci
    - go get github.com/scagogogo/go-composer-sdk
    - |
      cat > deploy.go << 'EOF'
      package main
      
      import (
          "fmt"
          "log"
          
          "github.com/scagogogo/go-composer-sdk/pkg/composer"
      )
      
      func main() {
          comp, err := composer.New(composer.DefaultOptions())
          if err != nil {
              log.Fatalf("创建实例失败: %v", err)
          }
          
          comp.SetWorkingDir(".")
          comp.SetEnv([]string{
              "COMPOSER_NO_INTERACTION=1",
              "COMPOSER_OPTIMIZE_AUTOLOADER=true",
              "COMPOSER_CLASSMAP_AUTHORITATIVE=true",
          })
          
          fmt.Println("🏭 生产环境构建...")
          if err := comp.Install(true, true); err != nil {
              log.Fatalf("生产环境安装失败: %v", err)
          }
          
          fmt.Println("✅ 部署准备完成")
      }
      EOF
    - go run deploy.go
    - echo "🚀 部署到生产环境..."
  only:
    - main
  when: manual
```

这个示例展示了如何将 Go Composer SDK 集成到各种 CI/CD 平台中，实现自动化的 PHP 项目构建、测试和部署流程。
