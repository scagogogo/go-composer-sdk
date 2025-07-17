package composer

import (
	"os"
	"path/filepath"
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
