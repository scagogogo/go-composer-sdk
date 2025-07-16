package composer

import (
	"errors"
	"testing"
)

func TestDiagnose(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for diagnose command
	SetupMockOutput("diagnose", "Checking composer.json: OK\nChecking platform settings: OK\nChecking git settings: OK\nChecking http connectivity: OK\nChecking disk space: OK\nChecking composer version: OK", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Diagnose()
	if err != nil {
		t.Errorf("Diagnose执行失败: %v", err)
	}

	if !contains(output, "Checking composer.json") {
		t.Errorf("输出应包含诊断信息，实际为\"%s\"", output)
	}

	if !contains(output, "OK") {
		t.Errorf("输出应包含检查结果，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("diagnose", "", errors.New("diagnose command failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Diagnose()
	if err == nil {
		t.Error("Diagnose命令失败时应该返回错误")
	}
}

func TestDiagnoseWithProblems(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with problems
	SetupMockOutput("diagnose", "Checking composer.json: OK\nChecking platform settings: WARNING - Some issues found\nChecking git settings: ERROR - Git not configured\nChecking http connectivity: OK\nChecking disk space: WARNING - Low disk space", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Diagnose()
	if err != nil {
		t.Errorf("Diagnose执行失败: %v", err)
	}

	if !contains(output, "WARNING") {
		t.Errorf("输出应包含警告信息，实际为\"%s\"", output)
	}

	if !contains(output, "ERROR") {
		t.Errorf("输出应包含错误信息，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithEmptyOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with empty result
	SetupMockOutput("diagnose", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Diagnose()
	if err != nil {
		t.Errorf("Diagnose执行失败: %v", err)
	}

	if output != "" {
		t.Errorf("空输出应该返回空字符串，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithVerboseOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for verbose diagnose using options
	SetupMockOutput("diagnose -v", "Checking composer.json: OK\n  - Valid JSON format\n  - All required fields present\nChecking platform settings: OK\n  - PHP version: 8.1.0\n  - Extensions loaded: json, mbstring, openssl", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"v": "", // verbose flag
	}

	output, err := composer.DiagnoseWithOptions(options)
	if err != nil {
		t.Errorf("DiagnoseWithOptions执行失败: %v", err)
	}

	if !contains(output, "Valid JSON format") {
		t.Errorf("详细输出应包含更多信息，实际为\"%s\"", output)
	}

	if !contains(output, "PHP version") {
		t.Errorf("详细输出应包含PHP版本信息，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for diagnose with options
	SetupMockOutput("diagnose --no-check-all --no-check-lock", "Checking composer.json: OK\nChecking platform settings: OK", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"no-check-all":  "",
		"no-check-lock": "",
	}

	output, err := composer.DiagnoseWithOptions(options)
	if err != nil {
		t.Errorf("DiagnoseWithOptions执行失败: %v", err)
	}

	if !contains(output, "Checking composer.json") {
		t.Errorf("输出应包含诊断信息，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithNilOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for diagnose without options
	SetupMockOutput("diagnose", "Checking composer.json: OK\nChecking platform settings: OK", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.DiagnoseWithOptions(nil)
	if err != nil {
		t.Errorf("DiagnoseWithOptions（nil选项）执行失败: %v", err)
	}

	if !contains(output, "Checking composer.json") {
		t.Errorf("输出应包含诊断信息，实际为\"%s\"", output)
	}
}

func TestDiagnoseWithEmptyOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for diagnose without options
	SetupMockOutput("diagnose", "Checking composer.json: OK\nChecking platform settings: OK", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.DiagnoseWithOptions(map[string]string{})
	if err != nil {
		t.Errorf("DiagnoseWithOptions（空选项）执行失败: %v", err)
	}

	if !contains(output, "Checking composer.json") {
		t.Errorf("输出应包含诊断信息，实际为\"%s\"", output)
	}
}
