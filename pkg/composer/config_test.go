package composer

import (
	"encoding/json"
	"errors"
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

func TestGetConfigEdgeCases(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试空配置键（应该返回nil，不报错）
	// 设置一个空的mock配置
	mockJSON := &ComposerJSON{
		Config: map[string]interface{}{},
	}
	SetMockComposerJSON(mockJSON)
	defer ClearMockComposerJSON()

	value, err := composer.GetConfig("")
	if err != nil {
		t.Errorf("空配置键GetConfig执行失败: %v", err)
	}
	if value != nil {
		t.Errorf("空配置键应该返回nil，实际为'%v'", value)
	}

	// 测试包含特殊字符的配置键
	mockJSONSpecial := &ComposerJSON{
		Config: map[string]interface{}{
			"vendor-dir.special":     "vendor-special",
			"config-with_underscore": "underscore-value",
			"config-with-dash":       "dash-value",
		},
	}
	SetMockComposerJSON(mockJSONSpecial)
	defer ClearMockComposerJSON()

	value, err = composer.GetConfig("vendor-dir.special")
	if err != nil {
		t.Errorf("特殊字符配置键GetConfig执行失败: %v", err)
	}
	if value != "vendor-special" {
		t.Errorf("特殊字符配置键输出不正确，期望: 'vendor-special'，实际: '%s'", value)
	}

	// 测试包含下划线的配置键
	value, err = composer.GetConfig("config-with_underscore")
	if err != nil {
		t.Errorf("下划线配置键GetConfig执行失败: %v", err)
	}
	if value != "underscore-value" {
		t.Errorf("下划线配置键输出不正确，期望: 'underscore-value'，实际: '%s'", value)
	}

	// 测试不存在的配置键（应该返回nil，不报错）
	value, err = composer.GetConfig("nonexistent-key")
	if err != nil {
		t.Errorf("不存在的配置键GetConfig执行失败: %v", err)
	}
	if value != nil {
		t.Errorf("不存在的配置键应该返回nil，实际为'%v'", value)
	}

	// 测试nil配置值
	mockJSONWithNil := &ComposerJSON{
		Config: map[string]interface{}{
			"nil-config": nil,
		},
	}
	SetMockComposerJSON(mockJSONWithNil)
	defer ClearMockComposerJSON()

	value, err = composer.GetConfig("nil-config")
	if err != nil {
		t.Errorf("nil配置值GetConfig执行失败: %v", err)
	}
	if value != nil {
		t.Errorf("nil配置值应返回nil，实际为'%v'", value)
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

// 测试边界情况和错误处理
func TestGetConfigWithEmptyKey(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetConfig("")
	if err == nil {
		t.Error("空配置键应该返回错误")
	}
}

func TestSetConfigWithEmptyKey(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.SetConfig("", "value")
	if err == nil {
		t.Error("空配置键应该返回错误")
	}
}

func TestUnsetConfigWithEmptyKey(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.UnsetConfig("")
	if err == nil {
		t.Error("空配置键应该返回错误")
	}
}

func TestGetConfigWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config repositories.packagist", "", errors.New("config not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetConfig("repositories.packagist")
	if err == nil {
		t.Error("配置不存在时应该返回错误")
	}
}

func TestSetConfigWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config repositories.packagist composer https://packagist.org", "", errors.New("config set failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.SetConfig("repositories.packagist", "composer https://packagist.org")
	if err == nil {
		t.Error("配置设置失败时应该返回错误")
	}
}

func TestUnsetConfigWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config --unset repositories.packagist", "", errors.New("config unset failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.UnsetConfig("repositories.packagist")
	if err == nil {
		t.Error("配置删除失败时应该返回错误")
	}
}

func TestGetConfigWithGlobalAndError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config --global repositories.packagist", "", errors.New("global config not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetConfigWithGlobal("repositories.packagist", true)
	if err == nil {
		t.Error("全局配置不存在时应该返回错误")
	}
}

func TestSetConfigWithGlobalAndError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config --global repositories.packagist composer https://packagist.org", "", errors.New("global config set failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.SetConfigWithGlobal("repositories.packagist", "composer https://packagist.org", true)
	if err == nil {
		t.Error("全局配置设置失败时应该返回错误")
	}
}

func TestValidateComposerJsonWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("validate --strict", "", errors.New("validation failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.ValidateComposerJson(true, false)
	if err == nil {
		t.Error("验证失败时应该返回错误")
	}
}

func TestCheckPlatformReqsWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("check-platform-reqs", "", errors.New("platform check failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.CheckPlatformReqs()
	if err == nil {
		t.Error("平台检查失败时应该返回错误")
	}
}

func TestClearCacheWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("clear-cache", "", errors.New("cache clear failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.ClearCache()
	if err == nil {
		t.Error("缓存清理失败时应该返回错误")
	}
}

func TestGetComposerHomeWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("config --global home", "", errors.New("home config not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetComposerHome()
	if err == nil {
		t.Error("获取Composer主目录失败时应该返回错误")
	}
}
