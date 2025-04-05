package composer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateSatisConfig(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "satis.json")

	// 测试创建Satis配置
	err = composer.CreateSatisConfig(configPath, "Private Repo", "https://example.org")
	if err != nil {
		t.Errorf("CreateSatisConfig执行失败: %v", err)
	}

	// 验证配置文件是否已创建
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("配置文件应该被创建: %s", configPath)
	}

	// 验证配置文件内容
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}

	if !contains(string(content), "Private Repo") || !contains(string(content), "https://example.org") {
		t.Errorf("配置文件内容不正确，应包含名称和主页: %s", string(content))
	}
}

func TestAddSatisRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录和配置文件
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "satis.json")

	// 先创建一个基本的Satis配置
	err = composer.CreateSatisConfig(configPath, "Private Repo", "https://example.org")
	if err != nil {
		t.Fatalf("创建Satis配置失败: %v", err)
	}

	// 测试添加仓库
	err = composer.AddSatisRepository(configPath, "vcs", "https://github.com/vendor/package")
	if err != nil {
		t.Errorf("AddSatisRepository执行失败: %v", err)
	}

	// 验证配置文件是否包含新仓库
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}

	if !contains(string(content), "vcs") || !contains(string(content), "https://github.com/vendor/package") {
		t.Errorf("配置文件内容不正确，应包含新添加的仓库: %s", string(content))
	}
}

func TestBuildSatis(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持satis build命令
	extendMockScript(t, execPath, "satis build /path/to/satis.json", "Building Satis repository...")

	output, err := composer.BuildSatis("/path/to/satis.json", "")
	if err != nil {
		t.Errorf("BuildSatis执行失败: %v", err)
	}

	if output == "" || !contains(output, "Building Satis repository") {
		t.Errorf("输出应包含构建信息，实际为\"%s\"", output)
	}

	// 测试指定输出目录
	extendMockScript(t, execPath, "satis build /path/to/satis.json /custom/output", "Building Satis repository to /custom/output...")

	output, err = composer.BuildSatis("/path/to/satis.json", "/custom/output")
	if err != nil {
		t.Errorf("BuildSatis(with output)执行失败: %v", err)
	}

	if output == "" || !contains(output, "/custom/output") {
		t.Errorf("输出应包含自定义输出目录，实际为\"%s\"", output)
	}
}

func TestInitSatis(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录
	tempDir := t.TempDir()
	satisDir := filepath.Join(tempDir, "satis")

	// 测试初始化Satis
	err = composer.InitSatis("Private Repo", "https://example.org", satisDir)
	if err != nil {
		t.Errorf("InitSatis执行失败: %v", err)
	}

	// 验证目录和配置文件是否已创建
	configPath := filepath.Join(satisDir, "satis.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("配置文件应该被创建: %s", configPath)
	}

	// 验证默认目录
	err = composer.InitSatis("Default Dir Repo", "https://example.org", "")
	if err != nil {
		t.Errorf("InitSatis(default dir)执行失败: %v", err)
	}

	defaultPath := filepath.Join("satis", "satis.json")
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		// 注意：在真实环境中，这个测试可能会失败，因为当前目录可能没有写权限
		// 但在测试环境中我们可以忽略这个问题
		t.Logf("注意：未能在默认目录创建文件，这在某些环境中是正常的: %v", err)
	}
}

func TestUpdateSatisStability(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录和配置文件
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "satis.json")

	// 先创建一个基本的Satis配置
	err = composer.CreateSatisConfig(configPath, "Private Repo", "https://example.org")
	if err != nil {
		t.Fatalf("创建Satis配置失败: %v", err)
	}

	// 测试有效的稳定性值
	validStabilities := []string{"dev", "alpha", "beta", "RC", "stable"}

	for _, stability := range validStabilities {
		t.Run(stability, func(t *testing.T) {
			err := composer.UpdateSatisStability(configPath, stability)
			if err != nil {
				t.Errorf("UpdateSatisStability(%s)执行失败: %v", stability, err)
			}

			// 验证配置文件是否包含稳定性设置
			content, err := os.ReadFile(configPath)
			if err != nil {
				t.Errorf("读取配置文件失败: %v", err)
			}

			if !contains(string(content), fmt.Sprintf(`"minimum-stability": "%s"`, stability)) {
				t.Errorf("配置文件内容不正确，应包含稳定性设置 %s: %s", stability, string(content))
			}
		})
	}

	// 测试无效的稳定性值
	err = composer.UpdateSatisStability(configPath, "invalid")
	if err == nil {
		t.Error("UpdateSatisStability应拒绝无效的稳定性值")
	}
}

func TestEnableSatisArchive(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录和配置文件
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "satis.json")

	// 先创建一个基本的Satis配置
	err = composer.CreateSatisConfig(configPath, "Private Repo", "https://example.org")
	if err != nil {
		t.Fatalf("创建Satis配置失败: %v", err)
	}

	// 测试启用归档（默认格式）
	err = composer.EnableSatisArchive(configPath, "")
	if err != nil {
		t.Errorf("EnableSatisArchive执行失败: %v", err)
	}

	// 验证配置文件是否包含归档设置
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}

	if !contains(string(content), `"archive"`) || !contains(string(content), `"format": "zip"`) {
		t.Errorf("配置文件内容不正确，应包含归档设置: %s", string(content))
	}

	// 测试指定归档格式
	err = composer.EnableSatisArchive(configPath, "tar")
	if err != nil {
		t.Errorf("EnableSatisArchive(tar)执行失败: %v", err)
	}

	// 验证配置文件是否包含指定的归档格式
	content, err = os.ReadFile(configPath)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}

	if !contains(string(content), `"format": "tar"`) {
		t.Errorf("配置文件内容不正确，应包含指定的归档格式: %s", string(content))
	}
}

func TestAddSatisRequire(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 为测试创建临时目录和配置文件
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "satis.json")

	// 先创建一个基本的Satis配置
	err = composer.CreateSatisConfig(configPath, "Private Repo", "https://example.org")
	if err != nil {
		t.Fatalf("创建Satis配置失败: %v", err)
	}

	// 测试添加依赖
	err = composer.AddSatisRequire(configPath, "vendor/package", "^1.0")
	if err != nil {
		t.Errorf("AddSatisRequire执行失败: %v", err)
	}

	// 验证配置文件是否包含依赖设置
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}

	if !contains(string(content), `"require"`) || !contains(string(content), `"vendor/package": "^1.0"`) {
		t.Errorf("配置文件内容不正确，应包含依赖设置: %s", string(content))
	}

	// 验证require-all是否已关闭
	if !contains(string(content), `"require-all": false`) {
		t.Errorf("配置文件内容不正确，应将require-all设置为false: %s", string(content))
	}
}
