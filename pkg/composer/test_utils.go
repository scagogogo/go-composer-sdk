package composer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// MockOutput 存储模拟命令的输出和错误
type MockOutput struct {
	Output string
	Error  error
}

// 全局模拟输出映射，用于存储命令到输出的映射
var (
	mockOutputs = make(map[string]MockOutput)
	mockMutex   sync.RWMutex
)

// SetupMockOutputAdvanced 设置模拟命令的输出（高级版本）
func SetupMockOutputAdvanced(command, output string, err error) {
	mockMutex.Lock()
	defer mockMutex.Unlock()
	mockOutputs[command] = MockOutput{
		Output: output,
		Error:  err,
	}
}

// ClearMockOutputsAdvanced 清除所有模拟输出（高级版本）
func ClearMockOutputsAdvanced() {
	mockMutex.Lock()
	defer mockMutex.Unlock()
	mockOutputs = make(map[string]MockOutput)
}

// GetMockOutput 获取模拟命令的输出
func GetMockOutput(command string) (MockOutput, bool) {
	mockMutex.RLock()
	defer mockMutex.RUnlock()
	output, exists := mockOutputs[command]
	return output, exists
}

// createAdvancedMockExecutable 创建一个模拟的Composer可执行文件用于测试
func createAdvancedMockExecutable(t *testing.T) string {
	tmpDir := t.TempDir()
	execPath := filepath.Join(tmpDir, "composer")

	// 脚本内容：模拟Composer命令
	content := `#!/bin/sh

# 构建完整的命令字符串
FULL_COMMAND="$*"

# 检查是否有模拟输出
case "$FULL_COMMAND" in
    "--version")
        echo "Composer version 2.4.4 2022-10-27 14:39:29"
        exit 0
        ;;
    "install")
        echo "Loading composer repositories with package information"
        echo "Installing dependencies (including require-dev) from lock file"
        echo "Nothing to install or update"
        echo "Generating autoload files"
        exit 0
        ;;
    "update")
        echo "Loading composer repositories with package information"
        echo "Updating dependencies (including require-dev)"
        echo "Nothing to update"
        exit 0
        ;;
    "dump-autoload")
        echo "Generating autoload files"
        exit 0
        ;;
    "dump-autoload --optimize")
        echo "Generating optimized autoload files"
        exit 0
        ;;
    "clear-cache")
        echo "Cache directory does not exist (cache-dir): /tmp/composer-cache"
        echo "Cache cleared successfully"
        exit 0
        ;;
    "self-update")
        echo "You are already using composer version 2.4.4 (stable channel)."
        exit 0
        ;;
    *)
        echo "Unknown command: $FULL_COMMAND"
        exit 1
        ;;
esac`

	err := os.WriteFile(execPath, []byte(content), 0755)
	if err != nil {
		t.Fatalf("创建模拟Composer可执行文件失败: %v", err)
	}

	return execPath
}

// createFullMockExecutable 创建一个支持更多命令的完整模拟可执行文件
func createFullMockExecutable(t *testing.T) string {
	tmpDir := t.TempDir()
	execPath := filepath.Join(tmpDir, "composer")

	// 更复杂的脚本，支持更多命令
	content := `#!/bin/sh

# 构建完整的命令字符串
FULL_COMMAND="$*"

# 检查常用命令
case "$FULL_COMMAND" in
    "--version")
        echo "Composer version 2.4.4 2022-10-27 14:39:29"
        exit 0
        ;;
    "install")
        echo "Loading composer repositories with package information"
        echo "Installing dependencies (including require-dev) from lock file"
        echo "Nothing to install or update"
        echo "Generating autoload files"
        exit 0
        ;;
    "update")
        echo "Loading composer repositories with package information"
        echo "Updating dependencies (including require-dev)"
        echo "Nothing to update"
        exit 0
        ;;
    "require "*")
        PACKAGE=$(echo "$FULL_COMMAND" | sed 's/require //')
        echo "Using version ^1.0 for $PACKAGE"
        echo "./composer.json has been updated"
        echo "Loading composer repositories with package information"
        echo "Updating dependencies (including require-dev)"
        echo "Package operations: 1 install, 0 updates, 0 removals"
        echo "  - Installing $PACKAGE (v1.0.0): Downloading (100%)"
        echo "Writing lock file"
        echo "Generating autoload files"
        exit 0
        ;;
    "remove "*")
        PACKAGE=$(echo "$FULL_COMMAND" | sed 's/remove //')
        echo "./composer.json has been updated"
        echo "Loading composer repositories with package information"
        echo "Updating dependencies (including require-dev)"
        echo "Package operations: 0 installs, 0 updates, 1 removal"
        echo "  - Removing $PACKAGE (v1.0.0)"
        echo "Writing lock file"
        echo "Generating autoload files"
        exit 0
        ;;
    "show "*")
        PACKAGE=$(echo "$FULL_COMMAND" | sed 's/show //')
        if [ "$PACKAGE" = "show" ]; then
            echo "vendor/package1    v1.0.0  Package description"
            echo "vendor/package2    v2.0.0  Another package"
        else
            echo "name     : $PACKAGE"
            echo "descrip. : Package description"
            echo "keywords : library, package"
            echo "versions : * v1.0.0"
            echo "type     : library"
            echo "license  : MIT"
            echo "source   : [git] https://github.com/vendor/package.git"
        fi
        exit 0
        ;;
    "search "*")
        QUERY=$(echo "$FULL_COMMAND" | sed 's/search //')
        echo "Found 5 packages matching $QUERY:"
        echo ""
        echo "vendor/package1 Package 1 description"
        echo "vendor/package2 Package 2 description"
        exit 0
        ;;
    "validate")
        echo "./composer.json is valid"
        exit 0
        ;;
    "diagnose")
        echo "Checking composer.json: OK"
        echo "Checking platform settings: OK"
        echo "Checking git settings: OK"
        echo "Checking http connectivity: OK"
        echo "Checking disk space: OK"
        echo "Checking composer version: OK"
        exit 0
        ;;
    "audit")
        echo "No security vulnerability advisories found"
        exit 0
        ;;
    "audit --format=json")
        echo '{"vulnerabilities":[],"found":0}'
        exit 0
        ;;
    "licenses")
        echo "Name: vendor/package1"
        echo "Version: 1.0.0"
        echo "Licenses: MIT"
        echo "Authors: John Doe <john@example.com>"
        echo ""
        echo "Name: vendor/package2"
        echo "Version: 2.0.0"
        echo "Licenses: Apache-2.0"
        echo "Authors: Jane Smith <jane@example.com>"
        exit 0
        ;;
    "fund")
        echo "The following packages were found in your dependencies which publish funding information:"
        echo ""
        echo "vendor/package1"
        echo "  https://github.com/sponsors/vendor"
        echo "  https://patreon.com/vendor"
        exit 0
        ;;
    "config "*")
        CONFIG_CMD=$(echo "$FULL_COMMAND" | sed 's/config //')
        echo "Configuration updated: $CONFIG_CMD"
        exit 0
        ;;
    "dump-autoload")
        echo "Generating autoload files"
        exit 0
        ;;
    "dump-autoload --optimize")
        echo "Generating optimized autoload files"
        exit 0
        ;;
    "clear-cache")
        echo "Cache directory does not exist (cache-dir): /tmp/composer-cache"
        echo "Cache cleared successfully"
        exit 0
        ;;
    "self-update")
        echo "You are already using composer version 2.4.4 (stable channel)."
        exit 0
        ;;
    *)
        echo "Unknown command: $FULL_COMMAND"
        exit 1
        ;;
esac`

	err := os.WriteFile(execPath, []byte(content), 0755)
	if err != nil {
		t.Fatalf("创建高级模拟Composer可执行文件失败: %v", err)
	}

	return execPath
}

// containsIgnoreCase 检查字符串是否包含子字符串（不区分大小写）
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// extendAdvancedMockScript 扩展模拟脚本以支持新命令
func extendAdvancedMockScript(t *testing.T, execPath string, command string, output string) {
	// 获取当前脚本内容
	content, err := os.ReadFile(execPath)
	if err != nil {
		t.Fatalf("读取模拟Composer可执行文件失败: %v", err)
	}

	// 将脚本按行分割
	lines := strings.Split(string(content), "\n")

	// 寻找最后一个case语句
	lastCaseIndex := -1
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.Contains(lines[i], "*)") {
			lastCaseIndex = i
			break
		}
	}

	// 在最后一个case之前插入新的case
	if lastCaseIndex != -1 {
		// 准备新命令的case语句
		newCase := fmt.Sprintf("    \"%s\")\n        echo \"%s\"\n        exit 0\n        ;;",
			command, strings.Replace(output, "\n", "\\n", -1))

		// 插入新case
		lines = append(lines[:lastCaseIndex], append([]string{newCase}, lines[lastCaseIndex:]...)...)
	}

	// 写回脚本文件
	err = os.WriteFile(execPath, []byte(strings.Join(lines, "\n")), 0755)
	if err != nil {
		t.Fatalf("更新模拟Composer可执行文件失败: %v", err)
	}
}

// createTestComposerJson 创建一个测试用的composer.json文件
func createTestComposerJson(t *testing.T, dir string) string {
	composerJsonPath := filepath.Join(dir, "composer.json")

	content := `{
    "name": "test/project",
    "description": "Test project for composer SDK",
    "type": "project",
    "license": "MIT",
    "authors": [
        {
            "name": "Test Author",
            "email": "test@example.com"
        }
    ],
    "require": {
        "php": ">=7.4",
        "vendor/package": "^1.0"
    },
    "require-dev": {
        "phpunit/phpunit": "^9.0"
    },
    "autoload": {
        "psr-4": {
            "Test\\": "src/"
        }
    },
    "autoload-dev": {
        "psr-4": {
            "Test\\Tests\\": "tests/"
        }
    }
}`

	err := os.WriteFile(composerJsonPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("创建测试composer.json文件失败: %v", err)
	}

	return composerJsonPath
}

// createTestComposerLock 创建一个测试用的composer.lock文件
func createTestComposerLock(t *testing.T, dir string) string {
	composerLockPath := filepath.Join(dir, "composer.lock")

	content := `{
    "_readme": [
        "This file locks the dependencies of your project to a known state",
        "Read more about it at https://getcomposer.org/doc/01-basic-usage.md#installing-dependencies"
    ],
    "content-hash": "test-hash",
    "packages": [
        {
            "name": "vendor/package",
            "version": "v1.0.0",
            "source": {
                "type": "git",
                "url": "https://github.com/vendor/package.git",
                "reference": "abc123"
            },
            "require": {
                "php": ">=7.4"
            },
            "type": "library",
            "autoload": {
                "psr-4": {
                    "Vendor\\Package\\": "src/"
                }
            },
            "license": [
                "MIT"
            ],
            "authors": [
                {
                    "name": "Package Author",
                    "email": "author@example.com"
                }
            ],
            "description": "Test package"
        }
    ],
    "packages-dev": [],
    "aliases": [],
    "minimum-stability": "stable",
    "stability-flags": [],
    "prefer-stable": false,
    "prefer-lowest": false,
    "platform": {
        "php": ">=7.4"
    },
    "platform-dev": []
}`

	err := os.WriteFile(composerLockPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("创建测试composer.lock文件失败: %v", err)
	}

	return composerLockPath
}
