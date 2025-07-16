package composer

import (
	"errors"
	"testing"
)

func TestExec(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command (note the -- separator)
	SetupMockOutput("exec phpunit --", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.\n\nOK (5 tests, 10 assertions)", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("phpunit")
	if err != nil {
		t.Errorf("Exec执行失败: %v", err)
	}

	if !contains(output, "PHPUnit") {
		t.Errorf("输出应包含PHPUnit信息，实际为\"%s\"", output)
	}
}

func TestExecWithArgs(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command with arguments
	SetupMockOutput("exec phpunit -- --version", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("phpunit", "--version")
	if err != nil {
		t.Errorf("Exec执行失败: %v", err)
	}

	if !contains(output, "PHPUnit 9.5.0") {
		t.Errorf("输出应包含版本信息，实际为\"%s\"", output)
	}
}

func TestExecWithEmptyCommand(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Exec("")
	if err == nil {
		t.Error("空命令应该返回错误")
	}
}

func TestExecWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("exec nonexistent-command", "", errors.New("command not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Exec("nonexistent-command")
	if err == nil {
		t.Error("不存在的命令应该返回错误")
	}
}

func TestExecWithArgsAndEmptyCommand(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Exec("", "--version")
	if err == nil {
		t.Error("空命令应该返回错误")
	}
}

func TestExecWithArgsAndNilArgs(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command without arguments
	SetupMockOutput("exec phpunit --", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.\n\nOK (5 tests, 10 assertions)", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("phpunit")
	if err != nil {
		t.Errorf("Exec（无参数）执行失败: %v", err)
	}

	if !contains(output, "PHPUnit") {
		t.Errorf("输出应包含PHPUnit信息，实际为\"%s\"", output)
	}
}

func TestExecWithArgsAndEmptyArgs(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command without arguments
	SetupMockOutput("exec phpunit --", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.\n\nOK (5 tests, 10 assertions)", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("phpunit")
	if err != nil {
		t.Errorf("Exec（空参数）执行失败: %v", err)
	}

	if !contains(output, "PHPUnit") {
		t.Errorf("输出应包含PHPUnit信息，实际为\"%s\"", output)
	}
}

func TestExecWithMultipleArgs(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command with multiple arguments
	SetupMockOutput("exec phpunit -- --configuration phpunit.xml --testsuite unit", "PHPUnit 9.5.0 by Sebastian Bergmann and contributors.\n\nOK (10 tests, 20 assertions)", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("phpunit", "--configuration", "phpunit.xml", "--testsuite", "unit")
	if err != nil {
		t.Errorf("Exec（多个参数）执行失败: %v", err)
	}

	if !contains(output, "10 tests") {
		t.Errorf("输出应包含测试数量信息，实际为\"%s\"", output)
	}
}

func TestExecWithSpecialCharacters(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command with special characters (note the -- separator)
	SetupMockOutput("exec echo -- 'Hello World!'", "Hello World!", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("echo", "'Hello World!'")
	if err != nil {
		t.Errorf("Exec（特殊字符）执行失败: %v", err)
	}

	if !contains(output, "Hello World!") {
		t.Errorf("输出应包含特殊字符，实际为\"%s\"", output)
	}
}

func TestExecWithLongOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command with long output (note the -- separator)
	longOutput := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5\nLine 6\nLine 7\nLine 8\nLine 9\nLine 10"
	SetupMockOutput("exec command-with-long-output --", longOutput, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("command-with-long-output")
	if err != nil {
		t.Errorf("Exec（长输出）执行失败: %v", err)
	}

	if !contains(output, "Line 1") || !contains(output, "Line 10") {
		t.Errorf("输出应包含完整的长输出，实际为\"%s\"", output)
	}
}

func TestExecWithBinaryCommand(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for exec command with binary
	SetupMockOutput("exec php-cs-fixer -- fix --dry-run", "Loaded config default from \".php-cs-fixer.dist.php\".\nUsing cache file \".php-cs-fixer.cache\".\n\n   1) src/Example.php", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Exec("php-cs-fixer", "fix", "--dry-run")
	if err != nil {
		t.Errorf("Exec（二进制命令）执行失败: %v", err)
	}

	if !contains(output, "php-cs-fixer") {
		t.Errorf("输出应包含命令信息，实际为\"%s\"", output)
	}
}
