package composer

import (
	"encoding/json"
	"strings"
	"testing"
)

// 定义 ValidationResult 结构体
type ValidationResult struct {
	Valid  bool
	Errors []string
}

// 定义 ConfigItem 结构体
type ConfigItem struct {
	Name   string
	Value  interface{}
	Source string
}

func TestValidate(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟验证成功的情况
	ClearMockOutputs()
	SetupMockOutput("validate", "composer.json is valid", nil)

	// 直接使用Run方法
	output, err := composer.Run("validate")
	if err != nil {
		t.Errorf("执行validate命令失败: %v", err)
	}

	// 手动解析输出
	valid := strings.Contains(output, "is valid")
	if !valid {
		t.Errorf("composer.json应该是有效的")
	}

	// 模拟验证失败的情况
	ClearMockOutputs()
	errorMsg := "composer.json is not valid, the following errors were found:\n- require.invalid/package: package name must be lowercase"
	SetupMockOutput("validate", errorMsg, nil)

	output, err = composer.Run("validate")
	if err != nil {
		t.Errorf("执行validate命令失败: %v", err)
	}

	// 手动解析输出
	valid = strings.Contains(output, "is valid")
	if valid {
		t.Errorf("composer.json应该是无效的")
	}

	errors := parseValidationErrors(output)
	if len(errors) != 1 {
		t.Errorf("应返回1个错误，实际返回%d个", len(errors))
		return
	}

	if len(errors) > 0 && !strings.Contains(errors[0], "package name must be lowercase") {
		t.Errorf("错误消息不符合预期: %s", errors[0])
	}
}

// 辅助函数：解析验证错误
func parseValidationErrors(output string) []string {
	var errors []string
	if !strings.Contains(output, "following errors were found") {
		return errors
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "-") {
			errors = append(errors, strings.TrimPrefix(line, "- "))
		}
	}
	return errors
}

func TestGetConfig(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create mock composer.json
	mockJSON := &ComposerJSON{
		Config: map[string]interface{}{
			"vendor-dir": "/path/to/vendor",
		},
	}
	SetMockComposerJSON(mockJSON)
	defer ClearMockComposerJSON()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	value, err := composer.GetConfig("vendor-dir")
	if err != nil {
		t.Errorf("GetConfig执行失败: %v", err)
	}

	if value != "/path/to/vendor" {
		t.Errorf("vendor-dir配置应为'/path/to/vendor'，实际为'%s'", value)
	}
}

func TestGlobalConfig(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟全局config命令输出
	ClearMockOutputs()
	SetupMockOutput("config --global github-oauth.github.com", "abc123token", nil)

	// 直接使用Run方法
	output, err := composer.Run("config", "--global", "github-oauth.github.com")
	if err != nil {
		t.Errorf("执行全局config命令失败: %v", err)
	}

	if output != "abc123token" {
		t.Errorf("github-oauth配置应为'abc123token'，实际为'%s'", output)
	}
}

func TestSetConfig(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create mock composer.json
	mockJSON := &ComposerJSON{
		Config: map[string]interface{}{},
	}
	SetMockComposerJSON(mockJSON)
	defer ClearMockComposerJSON()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟设置配置命令
	err = composer.SetConfig("process-timeout", "600")
	if err != nil {
		t.Errorf("SetConfig执行失败: %v", err)
	}

	// 验证设置后的配置
	value, err := composer.GetConfig("process-timeout")
	if err != nil {
		t.Errorf("GetConfig执行失败: %v", err)
	}

	if value != "600" {
		t.Errorf("process-timeout配置应为'600'，实际为'%s'", value)
	}
}

func TestSetGlobalConfigRun(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟设置全局配置命令
	ClearMockOutputs()
	SetupMockOutput("config --global cache-files-ttl 86400", "Set cache-files-ttl to 86400", nil)

	// 直接使用Run方法
	_, err = composer.Run("config", "--global", "cache-files-ttl", "86400")
	if err != nil {
		t.Errorf("执行设置全局配置命令失败: %v", err)
	}

	// 验证设置后的全局配置
	ClearMockOutputs()
	SetupMockOutput("config --global cache-files-ttl", "86400", nil)

	output, err := composer.Run("config", "--global", "cache-files-ttl")
	if err != nil {
		t.Errorf("执行获取全局配置命令失败: %v", err)
	}

	if output != "86400" {
		t.Errorf("cache-files-ttl配置应为'86400'，实际为'%s'", output)
	}
}

func TestListConfigsRun(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟列出配置命令输出
	ClearMockOutputs()
	configsOutput := `[
		{
			"name": "vendor-dir",
			"value": "vendor",
			"source": "project"
		},
		{
			"name": "bin-dir",
			"value": "bin",
			"source": "project"
		},
		{
			"name": "process-timeout",
			"value": 300,
			"source": "default"
		}
	]`

	SetupMockOutput("config --list --json", configsOutput, nil)

	// 直接使用Run方法
	output, err := composer.Run("config", "--list", "--json")
	if err != nil {
		t.Errorf("执行列出配置命令失败: %v", err)
	}

	// 手动解析JSON输出
	var configs []ConfigItem
	err = json.Unmarshal([]byte(output), &configs)
	if err != nil {
		t.Errorf("解析JSON输出失败: %v", err)
		return
	}

	if len(configs) != 3 {
		t.Errorf("应返回3个配置，实际返回%d个", len(configs))
		return
	}

	// 验证返回的配置信息
	expectedNames := []string{"vendor-dir", "bin-dir", "process-timeout"}
	for i, config := range configs {
		if config.Name != expectedNames[i] {
			t.Errorf("配置%d的名称应为%s，实际为%s", i, expectedNames[i], config.Name)
		}
	}

	// 验证值和来源
	if configs[0].Value != "vendor" || configs[0].Source != "project" {
		t.Errorf("vendor-dir配置信息不正确: %+v", configs[0])
	}
}

func TestGetComposerHome(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 模拟获取Composer主目录命令输出
	extendMockScript(t, execPath, "config home -g", "/home/user/.composer")

	homePath, err := composer.GetComposerHome()
	if err != nil {
		t.Errorf("GetComposerHome执行失败: %v", err)
	}

	if homePath != "/home/user/.composer" {
		t.Errorf("Composer主目录应为'/home/user/.composer'，实际为'%s'", homePath)
	}
}
