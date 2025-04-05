package composer

import (
	"strings"
	"testing"
)

// 定义 RunScriptOptions 结构体
type RunScriptOptions struct {
	Timeout int
}

// 定义 Script 结构体
type Script struct {
	Name        string
	Description string
}

func TestRunScript(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟执行脚本
	extendMockScript(t, execPath, "run test", "Running test script...\nTest completed successfully")

	output, err := composer.RunScript("test")
	if err != nil {
		t.Errorf("RunScript执行失败: %v", err)
	}

	if !strings.Contains(output, "Test completed successfully") {
		t.Errorf("RunScript输出不符合预期: %s", output)
	}
}

func TestRunScriptWithArguments(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟执行带参数的脚本
	extendMockScript(t, execPath, "run test -- --filter=MyTest", "Running test script with filter MyTest...\nTest MyTest completed successfully")

	output, err := composer.RunScript("test", "--filter=MyTest")
	if err != nil {
		t.Errorf("RunScript执行失败: %v", err)
	}

	if !strings.Contains(output, "Test MyTest completed successfully") {
		t.Errorf("RunScript输出不符合预期: %s", output)
	}
}

// 修改为直接使用命令行选项
func TestRunScriptWithOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟执行带选项的脚本
	extendMockScript(t, execPath, "run --timeout=300 test", "Running test script with timeout 300s...\nTest completed successfully")

	// 使用带有选项的命令行调用
	output, err := composer.Run("run", "--timeout=300", "test")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	if !strings.Contains(output, "Test completed successfully") {
		t.Errorf("Run输出不符合预期: %s", output)
	}
}

// 修改为直接使用命令行选项和参数
func TestRunScriptWithOptionsAndArguments(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟执行带选项和参数的脚本
	extendMockScript(t, execPath, "run --timeout=300 test -- --filter=MyTest",
		"Running test script with timeout 300s and filter MyTest...\nTest MyTest completed successfully")

	// 使用带有选项和参数的命令行调用
	output, err := composer.Run("run", "--timeout=300", "test", "--", "--filter=MyTest")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	if !strings.Contains(output, "Test MyTest completed successfully") {
		t.Errorf("Run输出不符合预期: %s", output)
	}
}

func TestListScripts(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟run -l输出
	scriptsOutput := `Scripts:
  test         Run unit tests
  lint         Check coding standards
  build        Build the project
  start        Start development server`

	extendMockScript(t, execPath, "run -l", scriptsOutput)

	// 直接解析列出脚本命令的输出
	output, err := composer.Run("run", "-l")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	// 手动解析输出来检查脚本列表
	scripts := parseScriptsOutput(output)

	if len(scripts) != 4 {
		t.Errorf("应返回4个脚本，实际返回%d个", len(scripts))
		return
	}

	// 验证返回的脚本信息
	expectedScripts := map[string]string{
		"test":  "Run unit tests",
		"lint":  "Check coding standards",
		"build": "Build the project",
		"start": "Start development server",
	}

	for _, script := range scripts {
		expectedDesc, ok := expectedScripts[script.Name]
		if !ok {
			t.Errorf("未预期的脚本: %s", script.Name)
			continue
		}

		if script.Description != expectedDesc {
			t.Errorf("脚本%s的描述应为%s，实际为%s", script.Name, expectedDesc, script.Description)
		}
	}
}

// 辅助函数：解析脚本列表输出
func parseScriptsOutput(output string) []Script {
	var scripts []Script
	lines := strings.Split(output, "\n")

	for i, line := range lines {
		if i == 0 && strings.TrimSpace(line) == "Scripts:" {
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		description := strings.TrimSpace(parts[1])

		scripts = append(scripts, Script{
			Name:        name,
			Description: description,
		})
	}

	return scripts
}

func TestHasScript(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟run -l输出
	scriptsOutput := `Scripts:
  test         Run unit tests
  lint         Check coding standards
  build        Build the project`

	extendMockScript(t, execPath, "run -l", scriptsOutput)

	// 获取脚本列表并检查特定脚本是否存在
	output, err := composer.Run("run", "-l")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	scripts := parseScriptsOutput(output)

	// 测试存在的脚本
	if !hasScript(scripts, "test") {
		t.Error("应检测到test脚本")
	}

	// 测试不存在的脚本
	if hasScript(scripts, "deploy") {
		t.Error("不应检测到deploy脚本")
	}
}

// 辅助函数：检查脚本是否存在
func hasScript(scripts []Script, name string) bool {
	for _, script := range scripts {
		if script.Name == name {
			return true
		}
	}
	return false
}

func TestGetScriptDescription(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟run -l输出
	scriptsOutput := `Scripts:
  test         Run unit tests
  lint         Check coding standards
  build        Build the project`

	extendMockScript(t, execPath, "run -l", scriptsOutput)

	// 获取脚本列表并获取特定脚本的描述
	output, err := composer.Run("run", "-l")
	if err != nil {
		t.Errorf("Run执行失败: %v", err)
	}

	scripts := parseScriptsOutput(output)

	// 测试存在的脚本
	desc := getScriptDescription(scripts, "lint")
	if desc == "" {
		t.Error("未能获取到lint脚本的描述")
	}
	if desc != "Check coding standards" {
		t.Errorf("lint脚本的描述应为'Check coding standards'，实际为'%s'", desc)
	}

	// 测试不存在的脚本
	desc = getScriptDescription(scripts, "deploy")
	if desc != "" {
		t.Errorf("不应获取到deploy脚本的描述，但获取到了'%s'", desc)
	}
}

// 辅助函数：获取脚本描述
func getScriptDescription(scripts []Script, name string) string {
	for _, script := range scripts {
		if script.Name == name {
			return script.Description
		}
	}
	return ""
}
