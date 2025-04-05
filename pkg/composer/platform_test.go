package composer

import (
	"testing"
)

func TestCheckPlatform(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持check-platform --format=json命令
	jsonOutput := `{
		"platform": {
			"php": {"name": "php", "version": "8.1.0", "available": true},
			"ext-json": {"name": "ext-json", "version": "1.0.0", "available": true},
			"ext-missing": {"name": "ext-missing", "version": null, "available": false}
		}
	}`
	extendMockScript(t, execPath, "check-platform --format=json", jsonOutput)

	platforms, err := composer.CheckPlatform()
	if err != nil {
		t.Errorf("CheckPlatform执行失败: %v", err)
	}

	if len(platforms) != 3 {
		t.Errorf("应返回3个平台信息，实际返回%d个", len(platforms))
		return
	}

	// 验证返回的平台信息
	for _, platform := range platforms {
		switch platform.Name {
		case "php":
			if !platform.Available || platform.Version != "8.1.0" {
				t.Errorf("PHP平台信息不正确: %+v", platform)
			}
		case "ext-json":
			if !platform.Available {
				t.Errorf("ext-json平台信息不正确: %+v", platform)
			}
		case "ext-missing":
			if platform.Available {
				t.Errorf("ext-missing平台信息不正确: %+v", platform)
			}
		default:
			t.Errorf("未预期的平台: %s", platform.Name)
		}
	}
}

func TestCheckPlatformWithLock(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持check-platform --lock --format=json命令
	jsonOutput := `{
		"platform": {},
		"lock": {
			"php": {"name": "php", "version": "8.0.0", "available": true, "required": ">=7.4"},
			"ext-pdo": {"name": "ext-pdo", "version": "1.0.0", "available": true, "required": "*"}
		}
	}`
	extendMockScript(t, execPath, "check-platform --lock --format=json", jsonOutput)

	platforms, err := composer.CheckPlatformWithLock()
	if err != nil {
		t.Errorf("CheckPlatformWithLock执行失败: %v", err)
	}

	if len(platforms) != 2 {
		t.Errorf("应返回2个平台信息，实际返回%d个", len(platforms))
		return
	}

	// 验证返回的平台信息
	for _, platform := range platforms {
		switch platform.Name {
		case "php":
			if !platform.Available || platform.Version != "8.0.0" || platform.Required != ">=7.4" {
				t.Errorf("PHP平台信息不正确: %+v", platform)
			}
		case "ext-pdo":
			if !platform.Available || platform.Required != "*" {
				t.Errorf("ext-pdo平台信息不正确: %+v", platform)
			}
		default:
			t.Errorf("未预期的平台: %s", platform.Name)
		}
	}
}

func TestIsPlatformAvailable(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 先模拟CheckPlatform的返回
	jsonOutput := `{
		"platform": {
			"php": {"name": "php", "version": "8.1.0", "available": true},
			"ext-json": {"name": "ext-json", "version": "1.0.0", "available": true},
			"ext-missing": {"name": "ext-missing", "version": null, "available": false}
		}
	}`
	extendMockScript(t, execPath, "check-platform --format=json", jsonOutput)

	// 测试已在需求中列出的平台
	available, err := composer.IsPlatformAvailable("php", "")
	if err != nil {
		t.Errorf("IsPlatformAvailable(php)执行失败: %v", err)
	}
	if !available {
		t.Error("PHP平台应该可用")
	}

	available, err = composer.IsPlatformAvailable("ext-missing", "")
	if err != nil {
		t.Errorf("IsPlatformAvailable(ext-missing)执行失败: %v", err)
	}
	if available {
		t.Error("ext-missing平台应该不可用")
	}

	// 测试未在需求中列出的平台（需要直接调用check-platform）
	extendMockScript(t, execPath, "check-platform ext-curl:7.0", "ext-curl 7.0 is available")

	available, err = composer.IsPlatformAvailable("ext-curl", "7.0")
	if err != nil {
		t.Errorf("IsPlatformAvailable(ext-curl, 7.0)执行失败: %v", err)
	}
	if !available {
		t.Error("ext-curl 7.0平台应该可用")
	}

	extendMockScript(t, execPath, "check-platform ext-imagick", "ext-imagick is not available")

	available, err = composer.IsPlatformAvailable("ext-imagick", "")
	if err != nil {
		t.Errorf("IsPlatformAvailable(ext-imagick)执行失败: %v", err)
	}
	if available {
		t.Error("ext-imagick平台应该不可用")
	}
}

func TestGetPHPVersion(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持run --php-show-version命令
	extendMockScript(t, execPath, "run --php-show-version", "PHP 8.1.2 (cli)")

	version, err := composer.GetPHPVersion()
	if err != nil {
		t.Errorf("GetPHPVersion执行失败: %v", err)
	}

	if version != "8.1.2" {
		t.Errorf("PHP版本应为8.1.2，实际为%s", version)
	}

	// 测试格式不同的输出
	extendMockScript(t, execPath, "run --php-show-version", "PHP 7.4.0")

	version, err = composer.GetPHPVersion()
	if err != nil {
		t.Errorf("GetPHPVersion执行失败: %v", err)
	}

	if version != "7.4.0" {
		t.Errorf("PHP版本应为7.4.0，实际为%s", version)
	}
}

func TestGetExtensions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持run --show-extensions命令
	extendMockScript(t, execPath, "run --show-extensions", "Loaded extensions:\nCore\ndate\njson\nmysqli\npdo\npdo_mysql\nzip")

	extensions, err := composer.GetExtensions()
	if err != nil {
		t.Errorf("GetExtensions执行失败: %v", err)
	}

	expectedExtensions := []string{"Core", "date", "json", "mysqli", "pdo", "pdo_mysql", "zip"}
	if len(extensions) != len(expectedExtensions) {
		t.Errorf("应返回%d个扩展，实际返回%d个", len(expectedExtensions), len(extensions))
		return
	}

	// 验证返回的扩展列表
	for i, ext := range extensions {
		if ext != expectedExtensions[i] {
			t.Errorf("第%d个扩展应为%s，实际为%s", i, expectedExtensions[i], ext)
		}
	}
}

func TestHasExtension(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持run --show-extensions命令
	extendMockScript(t, execPath, "run --show-extensions", "Loaded extensions:\nCore\ndate\njson\npdo")

	// 测试存在的扩展
	has, err := composer.HasExtension("json")
	if err != nil {
		t.Errorf("HasExtension(json)执行失败: %v", err)
	}
	if !has {
		t.Error("应检测到json扩展")
	}

	// 测试不存在的扩展
	has, err = composer.HasExtension("imagick")
	if err != nil {
		t.Errorf("HasExtension(imagick)执行失败: %v", err)
	}
	if has {
		t.Error("不应检测到imagick扩展")
	}
}
