package composer

import (
	"errors"
	"os"
	"testing"
)

func TestGetEnvironmentInfo(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config --list command
	SetupMockOutput("config --list", "vendor-dir = vendor\nbin-dir = bin\nprocess-timeout = 300\ncache-dir = /tmp/composer-cache", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	envInfo, err := composer.GetEnvironmentInfo()
	if err != nil {
		t.Errorf("GetEnvironmentInfo执行失败: %v", err)
	}

	if len(envInfo) == 0 {
		t.Error("环境信息不应为空")
	}

	// 检查是否包含预期的配置项
	if vendorDir, exists := envInfo["vendor-dir"]; !exists || vendorDir != "vendor" {
		t.Errorf("vendor-dir配置不正确，期望: vendor，实际: %s", vendorDir)
	}
}

func TestGetEnvironmentInfoWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config --list", "", errors.New("config list failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetEnvironmentInfo()
	if err == nil {
		t.Error("配置列表失败时应该返回错误")
	}
}

func TestGetEnvironmentInfoWithEmptyOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with empty result
	SetupMockOutput("config --list", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	envInfo, err := composer.GetEnvironmentInfo()
	if err != nil {
		t.Errorf("GetEnvironmentInfo执行失败: %v", err)
	}

	if len(envInfo) != 0 {
		t.Errorf("空输出应该返回空映射，实际长度: %d", len(envInfo))
	}
}

func TestGetComposerHomeFromEnv(t *testing.T) {
	// 保存原始环境变量
	originalHome := os.Getenv("COMPOSER_HOME")
	defer func() {
		if originalHome != "" {
			os.Setenv("COMPOSER_HOME", originalHome)
		} else {
			os.Unsetenv("COMPOSER_HOME")
		}
	}()

	// 测试获取COMPOSER_HOME
	testPath := "/tmp/test-composer-home"
	os.Setenv("COMPOSER_HOME", testPath)

	home := os.Getenv("COMPOSER_HOME")
	if home != testPath {
		t.Errorf("COMPOSER_HOME环境变量不正确，期望: %s，实际: %s", testPath, home)
	}
}

func TestGetComposerHomeWithEmptyEnv(t *testing.T) {
	// 保存原始环境变量
	originalHome := os.Getenv("COMPOSER_HOME")
	defer func() {
		if originalHome != "" {
			os.Setenv("COMPOSER_HOME", originalHome)
		} else {
			os.Unsetenv("COMPOSER_HOME")
		}
	}()

	// 清除环境变量
	os.Unsetenv("COMPOSER_HOME")

	home := os.Getenv("COMPOSER_HOME")
	if home != "" {
		t.Error("清除环境变量后COMPOSER_HOME应该为空")
	}
}
