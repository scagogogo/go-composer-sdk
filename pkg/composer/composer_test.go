// 测试composer包的功能

package composer

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/scagogogo/go-composer-sdk/pkg/detector"
)

// 创建一个模拟的Composer可执行文件用于测试
func createMockExecutable(t *testing.T) string {
	tmpDir := t.TempDir()
	execPath := filepath.Join(tmpDir, "composer")

	// 脚本内容：返回版本信息
	content := `#!/bin/sh
if [ "$1" = "--version" ]; then
    echo "Composer version 2.5.0 2023-01-01 12:00:00"
    exit 0
elif [ "$1" = "self-update" ]; then
    echo "Updating to version 2.5.1"
    exit 0
elif [ "$1" = "require" ]; then
    echo "Using version ^1.0 for vendor/package"
    exit 0
elif [ "$1" = "install" ]; then
    echo "Installing dependencies from lock file"
    exit 0
elif [ "$1" = "update" ]; then
    echo "Updating dependencies"
    exit 0
elif [ "$1" = "remove" ]; then
    echo "Removing package"
    exit 0
elif [ "$1" = "dump-autoload" ]; then
    echo "Generating optimized autoload files"
    exit 0
else
    echo "Unknown command: $1"
    exit 1
fi`

	err := os.WriteFile(execPath, []byte(content), 0755)
	if err != nil {
		t.Fatalf("创建模拟Composer可执行文件失败: %v", err)
	}

	return execPath
}

func TestDefaultOptions(t *testing.T) {
	options := DefaultOptions()

	if options.WorkingDir != "" {
		t.Errorf("默认工作目录应为空，实际为\"%s\"", options.WorkingDir)
	}

	if !options.AutoInstall {
		t.Error("默认应允许自动安装")
	}

	if options.DefaultTimeout != 10*time.Minute {
		t.Errorf("默认超时时间应为10分钟，实际为%v", options.DefaultTimeout)
	}
}

func TestNew(t *testing.T) {
	// 创建模拟的可执行文件
	execPath := createMockExecutable(t)

	// 使用自定义选项创建Composer
	options := Options{
		ExecutablePath: execPath,
		WorkingDir:     "/custom/dir",
		AutoInstall:    false,
		DefaultTimeout: 10 * time.Minute,
	}

	composer, err := New(options)
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	if composer.executablePath != execPath {
		t.Errorf("可执行文件路径应为\"%s\"，实际为\"%s\"", execPath, composer.executablePath)
	}

	if composer.workingDir != "/custom/dir" {
		t.Errorf("工作目录应为\"/custom/dir\"，实际为\"%s\"", composer.workingDir)
	}

	if composer.autoInstall != false {
		t.Error("自动安装选项应为false")
	}

	if composer.defaultTimeout != 10*time.Minute {
		t.Errorf("超时时间应为10分钟，实际为%v", composer.defaultTimeout)
	}
}

func TestNewWithDefaultDetectorAndInstaller(t *testing.T) {
	// 创建模拟的可执行文件
	execPath := createMockExecutable(t)

	// 创建一个自定义的检测器，总是返回模拟的可执行文件
	mockDetector := detector.NewDetector()
	mockDetector.SetPossiblePaths([]string{execPath})

	options := Options{
		Detector: mockDetector,
	}

	composer, err := New(options)
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	if composer.executablePath != execPath {
		t.Errorf("可执行文件路径应为\"%s\"，实际为\"%s\"", execPath, composer.executablePath)
	}

	if composer.detector != mockDetector {
		t.Error("未正确使用提供的检测器")
	}

	if composer.installer == nil {
		t.Error("应创建默认安装器")
	}
}

func TestSetWorkingDir(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	newDir := "/new/working/dir"
	composer.SetWorkingDir(newDir)

	if composer.workingDir != newDir {
		t.Errorf("工作目录应为\"%s\"，实际为\"%s\"", newDir, composer.workingDir)
	}
}

func TestSetEnv(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	newEnv := []string{"VAR1=value1", "VAR2=value2"}
	composer.SetEnv(newEnv)

	if len(composer.env) != len(newEnv) {
		t.Errorf("环境变量数量应为%d，实际为%d", len(newEnv), len(composer.env))
	}

	for i, value := range newEnv {
		if composer.env[i] != value {
			t.Errorf("环境变量[%d]应为\"%s\"，实际为\"%s\"", i, value, composer.env[i])
		}
	}
}

func TestGetExecutablePath(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	if composer.GetExecutablePath() != execPath {
		t.Errorf("GetExecutablePath应返回\"%s\"，实际返回\"%s\"", execPath, composer.GetExecutablePath())
	}
}

func TestRunWithContext(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	output, err := composer.RunWithContext(ctx, "--version")
	if err != nil {
		t.Errorf("RunWithContext执行失败: %v", err)
	}

	// 检查输出是否包含预期的版本信息
	if output == "" || !contains(output, "Composer version") {
		t.Errorf("输出应包含版本信息，实际为\"%s\"", output)
	}
}

func TestRun(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Run("--version")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	// 检查输出是否包含预期的版本信息
	if output == "" || !contains(output, "Composer version") {
		t.Errorf("输出应包含版本信息，实际为\"%s\"", output)
	}
}

func TestIsInstalled(t *testing.T) {
	execPath := createMockExecutable(t)

	// 已安装的情况
	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	if !composer.IsInstalled() {
		t.Error("IsInstalled应返回true，表示已安装")
	}

	// 未安装的情况 - 这个需要特殊处理，因为通常New会尝试自动安装
	// 这里我们直接创建一个executablePath为空的Composer实例
	notInstalledComposer := &Composer{executablePath: ""}

	if notInstalledComposer.IsInstalled() {
		t.Error("对于未安装的情况，IsInstalled应返回false")
	}
}

func TestGetVersion(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for version command
	SetupMockOutput("--version", "Composer version 2.5.0 2023-01-01 12:00:00", nil)

	// Create composer with any executable path (won't be used in mocked execution)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	version, err := composer.GetVersion()
	if err != nil {
		t.Errorf("GetVersion执行失败: %v", err)
	}

	expectedVersion := "2.5.0"
	if version != expectedVersion {
		t.Errorf("版本应为\"%s\"，实际为\"%s\"", expectedVersion, version)
	}
}

// 简化的辅助函数，检查字符串是否包含子串
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 以下测试仅验证命令构造和执行，不验证实际效果
// 因为实际效果需要真实的Composer环境

func TestSelfUpdate(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for self-update command
	SetupMockOutput("self-update", "Updating to version 2.6.0\nComposer successfully updated to version 2.6.0.", nil)

	// Create composer with any executable path (won't be used in mocked execution)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.SelfUpdate()
	if err != nil {
		t.Errorf("SelfUpdate执行失败: %v", err)
	}
}

func TestRequirePackage(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used in mocked execution)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试各种参数组合
	testCases := []struct {
		name        string
		packageName string
		version     string
		dev         bool
	}{
		{"基本包", "vendor/package", "", false},
		{"带版本的包", "vendor/package", "^1.0", false},
		{"开发包", "vendor/package", "", true},
		{"带版本的开发包", "vendor/package", "^1.0", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up mock output for require command with appropriate parameters
			mockCmd := "require"
			if tc.dev {
				mockCmd += " --dev"
			}
			if tc.version != "" {
				mockCmd += " " + tc.packageName + ":" + tc.version
			} else {
				mockCmd += " " + tc.packageName
			}

			SetupMockOutput(mockCmd, "Using version "+tc.version+" for "+tc.packageName, nil)

			err := composer.RequirePackage(tc.packageName, tc.version, tc.dev)
			if err != nil {
				t.Errorf("RequirePackage执行失败: %v", err)
			}
		})
	}
}

func TestInstall(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试各种参数组合
	testCases := []struct {
		name     string
		noDev    bool
		optimize bool
	}{
		{"基本安装", false, false},
		{"无开发依赖", true, false},
		{"优化自动加载", false, true},
		{"无开发依赖且优化", true, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := composer.Install(tc.noDev, tc.optimize)
			if err != nil {
				t.Errorf("Install执行失败: %v", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试各种参数组合
	testCases := []struct {
		name     string
		packages []string
		noDev    bool
	}{
		{"更新所有包", []string{}, false},
		{"更新特定包", []string{"vendor/package1", "vendor/package2"}, false},
		{"无开发依赖更新所有包", []string{}, true},
		{"无开发依赖更新特定包", []string{"vendor/package1", "vendor/package2"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := composer.Update(tc.packages, tc.noDev)
			if err != nil {
				t.Errorf("Update执行失败: %v", err)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试各种参数组合
	testCases := []struct {
		name        string
		packageName string
		dev         bool
	}{
		{"移除普通包", "vendor/package", false},
		{"移除开发包", "vendor/package", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := composer.Remove(tc.packageName, tc.dev)
			if err != nil {
				t.Errorf("Remove执行失败: %v", err)
			}
		})
	}
}

func TestDumpAutoload(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试各种参数组合
	testCases := []struct {
		name     string
		optimize bool
	}{
		{"普通自动加载", false},
		{"优化自动加载", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := composer.DumpAutoload(tc.optimize)
			if err != nil {
				t.Errorf("DumpAutoload执行失败: %v", err)
			}
		})
	}
}
