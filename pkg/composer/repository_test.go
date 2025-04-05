package composer

import (
	"encoding/json"
	"testing"
)

func TestAddRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config命令
	expectedRepo := Repository{
		Type: VcsRepository,
		URL:  "https://github.com/vendor/repo",
	}
	repoJson, _ := json.Marshal(expectedRepo)

	extendMockScript(t, execPath, "config repositories.github "+string(repoJson), "Repository 'github' added to composer.json")

	err = composer.AddRepository("github", expectedRepo)
	if err != nil {
		t.Errorf("AddRepository执行失败: %v", err)
	}
}

func TestRemoveRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config --unset命令
	extendMockScript(t, execPath, "config --unset repositories.github", "Repository 'github' removed from composer.json")

	err = composer.RemoveRepository("github")
	if err != nil {
		t.Errorf("RemoveRepository执行失败: %v", err)
	}
}

func TestListRepositories(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config repositories命令
	extendMockScript(t, execPath, "config repositories", "github: vcs, https://github.com/vendor/repo\npackagist: composer, https://repo.packagist.org")

	output, err := composer.ListRepositories()
	if err != nil {
		t.Errorf("ListRepositories执行失败: %v", err)
	}

	if output == "" || !contains(output, "github") || !contains(output, "packagist") {
		t.Errorf("输出应包含仓库列表，实际为\"%s\"", output)
	}
}

func TestAddPackagistRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config命令
	expectedRepo := Repository{
		Type: PackagistRepository,
		URL:  "https://custom.packagist.org",
	}
	repoJson, _ := json.Marshal(expectedRepo)

	extendMockScript(t, execPath, "config repositories.packagist.org "+string(repoJson), "Repository 'packagist.org' added to composer.json")

	err = composer.AddPackagistRepository("https://custom.packagist.org")
	if err != nil {
		t.Errorf("AddPackagistRepository执行失败: %v", err)
	}
}

func TestDisablePackagistRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config命令
	extendMockScript(t, execPath, "config repositories.packagist.org.url false", "Packagist.org repository disabled")

	err = composer.DisablePackagistRepository()
	if err != nil {
		t.Errorf("DisablePackagistRepository执行失败: %v", err)
	}
}

func TestEnablePackagistRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config命令
	extendMockScript(t, execPath, "config repositories.packagist.org.url https://repo.packagist.org", "Packagist.org repository enabled")

	err = composer.EnablePackagistRepository()
	if err != nil {
		t.Errorf("EnablePackagistRepository执行失败: %v", err)
	}
}

func TestAddVcsRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 复用AddRepository的测试逻辑
	expectedRepo := Repository{
		Type: VcsRepository,
		URL:  "https://github.com/vendor/repo",
	}
	repoJson, _ := json.Marshal(expectedRepo)

	extendMockScript(t, execPath, "config repositories.vcs-repo "+string(repoJson), "Repository 'vcs-repo' added to composer.json")

	err = composer.AddVcsRepository("vcs-repo", "https://github.com/vendor/repo")
	if err != nil {
		t.Errorf("AddVcsRepository执行失败: %v", err)
	}
}

func TestAddPathRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 准备一个带选项的Path仓库
	options := map[string]interface{}{
		"symlink": true,
	}

	expectedRepo := Repository{
		Type:    PathRepository,
		URL:     "../local/package",
		Options: options,
	}
	repoJson, _ := json.Marshal(expectedRepo)

	extendMockScript(t, execPath, "config repositories.local "+string(repoJson), "Repository 'local' added to composer.json")

	err = composer.AddPathRepository("local", "../local/package", options)
	if err != nil {
		t.Errorf("AddPathRepository执行失败: %v", err)
	}
}

func TestAddComposerRepository(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 复用AddRepository的测试逻辑
	expectedRepo := Repository{
		Type: ComposerRepository,
		URL:  "https://composer.example.org",
	}
	repoJson, _ := json.Marshal(expectedRepo)

	extendMockScript(t, execPath, "config repositories.composer-repo "+string(repoJson), "Repository 'composer-repo' added to composer.json")

	err = composer.AddComposerRepository("composer-repo", "https://composer.example.org")
	if err != nil {
		t.Errorf("AddComposerRepository执行失败: %v", err)
	}
}

func TestGetPreferredInstall(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config preferred-install命令
	extendMockScript(t, execPath, "config preferred-install", "dist")

	output, err := composer.GetPreferredInstall()
	if err != nil {
		t.Errorf("GetPreferredInstall执行失败: %v", err)
	}

	if output == "" || !contains(output, "dist") {
		t.Errorf("输出应为dist，实际为\"%s\"", output)
	}
}

func TestSetPreferredInstall(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 有效的preferred-install值测试
	validValues := []string{"dist", "source", "auto"}

	for _, value := range validValues {
		t.Run(value, func(t *testing.T) {
			// 扩展模拟可执行文件以支持config preferred-install命令
			extendMockScript(t, execPath, "config preferred-install "+value, "preferred-install set to "+value)

			err := composer.SetPreferredInstall(value)
			if err != nil {
				t.Errorf("SetPreferredInstall(%s)执行失败: %v", value, err)
			}
		})
	}

	// 无效值测试
	err = composer.SetPreferredInstall("invalid")
	if err == nil {
		t.Error("SetPreferredInstall应拒绝无效值")
	}
}

func TestSetMinimumStability(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config minimum-stability命令
	extendMockScript(t, execPath, "config minimum-stability stable", "minimum-stability set to stable")

	err = composer.SetMinimumStability("stable")
	if err != nil {
		t.Errorf("SetMinimumStability执行失败: %v", err)
	}
}

func TestGetMinimumStability(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持config minimum-stability命令
	extendMockScript(t, execPath, "config minimum-stability", "stable")

	output, err := composer.GetMinimumStability()
	if err != nil {
		t.Errorf("GetMinimumStability执行失败: %v", err)
	}

	if output == "" || !contains(output, "stable") {
		t.Errorf("输出应为stable，实际为\"%s\"", output)
	}
}

func TestSetPreferStable(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试设置为true
	extendMockScript(t, execPath, "config prefer-stable 1", "prefer-stable set to 1")

	err = composer.SetPreferStable(true)
	if err != nil {
		t.Errorf("SetPreferStable(true)执行失败: %v", err)
	}

	// 测试设置为false
	extendMockScript(t, execPath, "config prefer-stable 0", "prefer-stable set to 0")

	err = composer.SetPreferStable(false)
	if err != nil {
		t.Errorf("SetPreferStable(false)执行失败: %v", err)
	}
}
