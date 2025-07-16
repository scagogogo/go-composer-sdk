package composer

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for validate command
	SetupMockOutput("validate", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Validate()
	if err != nil {
		t.Errorf("Validate执行失败: %v", err)
	}
}

func TestValidateWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("validate", "", errors.New("validation failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Validate()
	if err == nil {
		t.Error("验证失败时应该返回错误")
	}
}

func TestValidateStrict(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for strict validate command
	SetupMockOutput("validate --strict", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.ValidateStrict()
	if err != nil {
		t.Errorf("ValidateStrict执行失败: %v", err)
	}

	if !contains(output, "is valid") {
		t.Errorf("输出应包含验证结果，实际为\"%s\"", output)
	}
}

func TestValidateWithNoCheck(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for validate command with no-check-all
	SetupMockOutput("validate --no-check-all", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.ValidateWithNoCheck()
	if err != nil {
		t.Errorf("ValidateWithNoCheck执行失败: %v", err)
	}

	if !contains(output, "is valid") {
		t.Errorf("输出应包含验证结果，实际为\"%s\"", output)
	}
}
