package composer

import (
	"errors"
	"testing"
)

func TestGlobalRequire(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global require command
	SetupMockOutput("global require vendor/package", "Using version ^1.0 for vendor/package\n./composer.json has been updated\nLoading composer repositories with package information\nUpdating dependencies (including require-dev)\nPackage operations: 1 install, 0 updates, 0 removals\n  - Installing vendor/package (v1.0.0): Downloading (100%)\nWriting lock file\nGenerating autoload files", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRequire("vendor/package", "")
	if err != nil {
		t.Errorf("GlobalRequire执行失败: %v", err)
	}
}

func TestGlobalRequireWithVersion(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global require command with version
	SetupMockOutput("global require vendor/package:^2.0", "Using version ^2.0 for vendor/package\n./composer.json has been updated", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRequire("vendor/package", "^2.0")
	if err != nil {
		t.Errorf("GlobalRequire执行失败: %v", err)
	}
}

func TestGlobalRequireWithEmptyPackage(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRequire("", "")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestGlobalRequireWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("global require nonexistent/package", "", errors.New("package not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRequire("nonexistent/package", "")
	if err == nil {
		t.Error("不存在的包应该返回错误")
	}
}

func TestGlobalRemove(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global remove command
	SetupMockOutput("global remove vendor/package", "./composer.json has been updated\nLoading composer repositories with package information\nUpdating dependencies (including require-dev)\nPackage operations: 0 installs, 0 updates, 1 removal\n  - Removing vendor/package (v1.0.0)\nWriting lock file\nGenerating autoload files", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRemove("vendor/package")
	if err != nil {
		t.Errorf("GlobalRemove执行失败: %v", err)
	}
}

func TestGlobalRemoveWithEmptyPackage(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalRemove("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestGlobalUpdate(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global update command
	SetupMockOutput("global update", "Loading composer repositories with package information\nUpdating dependencies (including require-dev)\nNothing to update", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalUpdate([]string{})
	if err != nil {
		t.Errorf("GlobalUpdate执行失败: %v", err)
	}
}

func TestGlobalUpdateWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("global update", "", errors.New("update failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.GlobalUpdate([]string{})
	if err == nil {
		t.Error("更新失败时应该返回错误")
	}
}

func TestGlobalList(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global show command
	SetupMockOutput("global show", "vendor/package1    v1.0.0  Package description\nvendor/package2    v2.0.0  Another package", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.GlobalList()
	if err != nil {
		t.Errorf("GlobalList执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含全局包列表，实际为\"%s\"", output)
	}
}

func TestGlobalListWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("global show", "", errors.New("show failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GlobalList()
	if err == nil {
		t.Error("显示失败时应该返回错误")
	}
}

func TestGlobalExecute(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global exec command
	SetupMockOutput("global exec phpunit", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.\n\nOK (5 tests, 10 assertions)", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.GlobalExecute("phpunit")
	if err != nil {
		t.Errorf("GlobalExecute执行失败: %v", err)
	}

	if !contains(output, "PHPUnit") {
		t.Errorf("输出应包含命令执行结果，实际为\"%s\"", output)
	}
}

func TestGlobalExecuteWithEmptyCommand(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GlobalExecute("")
	if err == nil {
		t.Error("空命令应该返回错误")
	}
}

func TestGlobalExecuteWithArgs(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global exec command with arguments
	SetupMockOutput("global exec phpunit --version", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.GlobalExecute("phpunit", "--version")
	if err != nil {
		t.Errorf("GlobalExecute执行失败: %v", err)
	}

	if !contains(output, "PHPUnit 9.5.0") {
		t.Errorf("输出应包含版本信息，实际为\"%s\"", output)
	}
}

func TestGlobalExecuteWithArgsAndEmptyCommand(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GlobalExecute("", "--version")
	if err == nil {
		t.Error("空命令应该返回错误")
	}
}

func TestGlobalConfigUsingRun(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for global config command
	SetupMockOutput("global config repositories.packagist composer https://packagist.org", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Run("global", "config", "repositories.packagist", "composer https://packagist.org")
	if err != nil {
		t.Errorf("Global config执行失败: %v", err)
	}
}

func TestGlobalConfigWithEmptyKey(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Run("global", "config", "", "value")
	if err == nil {
		t.Error("空配置键应该返回错误")
	}
}
