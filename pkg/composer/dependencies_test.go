package composer

import (
	"errors"
	"testing"
)

func TestSuggests(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for suggests command
	SetupMockOutput("suggests", "vendor/suggested-package: For enhanced functionality\nvendor/another-package: For additional features", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err != nil {
		t.Errorf("Suggests执行失败: %v", err)
	}
}

func TestSuggestsWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("suggests", "", errors.New("suggests command failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err == nil {
		t.Error("Suggests命令失败时应该返回错误")
	}
}

func TestSuggestsWithEmptyOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with empty result
	SetupMockOutput("suggests", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err != nil {
		t.Errorf("Suggests执行失败: %v", err)
	}
}
