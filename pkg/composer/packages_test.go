package composer

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestShowPackage(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for show command
	SetupMockOutput("show vendor/package", "Package vendor/package: v1.0.0", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.ShowPackage("vendor/package")
	if err != nil {
		t.Errorf("ShowPackage执行失败: %v", err)
	}

	if output == "" || !contains(output, "Package vendor/package") {
		t.Errorf("输出应包含包信息，实际为\"%s\"", output)
	}
}

func TestShowPackageEdgeCases(t *testing.T) {
	// 测试空包名
	ClearMockOutputs()
	SetupMockOutput("show ", "No package specified", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.ShowPackage("")
	if err != nil {
		t.Errorf("空包名ShowPackage执行失败: %v", err)
	}

	// 测试包含特殊字符的包名
	ClearMockOutputs()
	SetupMockOutput("show vendor/package-with_special.chars", "Package vendor/package-with_special.chars\nVersion: 1.0.0", nil)

	_, err = composer.ShowPackage("vendor/package-with_special.chars")
	if err != nil {
		t.Errorf("特殊字符包名ShowPackage执行失败: %v", err)
	}

	// 测试非常长的包名
	longPackageName := "vendor/" + strings.Repeat("very-long-package-name", 10)
	ClearMockOutputs()
	SetupMockOutput("show "+longPackageName, "Package "+longPackageName+"\nVersion: 1.0.0", nil)

	_, err = composer.ShowPackage(longPackageName)
	if err != nil {
		t.Errorf("长包名ShowPackage执行失败: %v", err)
	}

	// 测试包含Unicode字符的包名
	unicodePackageName := "vendor/包名-with-中文"
	ClearMockOutputs()
	SetupMockOutput("show "+unicodePackageName, "Package "+unicodePackageName+"\nVersion: 1.0.0", nil)

	_, err = composer.ShowPackage(unicodePackageName)
	if err != nil {
		t.Errorf("Unicode包名ShowPackage执行失败: %v", err)
	}

	// 测试包含斜杠的包名
	slashPackageName := "vendor/sub/package"
	ClearMockOutputs()
	SetupMockOutput("show "+slashPackageName, "Package "+slashPackageName+"\nVersion: 1.0.0", nil)

	_, err = composer.ShowPackage(slashPackageName)
	if err != nil {
		t.Errorf("包含斜杠的包名ShowPackage执行失败: %v", err)
	}
}

func TestShowPackageWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试包不存在的错误
	SetupMockOutput("show nonexistent/package", "", errors.New("Package nonexistent/package not found"))
	_, err = composer.ShowPackage("nonexistent/package")
	if err == nil {
		t.Error("不存在的包应该返回错误")
	}

	// 测试网络错误
	SetupMockOutput("show vendor/package", "", errors.New("Could not fetch package information"))
	_, err = composer.ShowPackage("vendor/package")
	if err == nil {
		t.Error("网络错误应该返回错误")
	}

	// 测试无效包名格式
	SetupMockOutput("show invalid-package-name", "", errors.New("Invalid package name"))
	_, err = composer.ShowPackage("invalid-package-name")
	if err == nil {
		t.Error("无效包名格式应该返回错误")
	}

	// 测试权限错误
	SetupMockOutput("show private/package", "", errors.New("Access denied to private package"))
	_, err = composer.ShowPackage("private/package")
	if err == nil {
		t.Error("私有包权限错误应该返回错误")
	}

	// 测试composer.json不存在
	SetupMockOutput("show vendor/package", "", errors.New("composer.json not found"))
	_, err = composer.ShowPackage("vendor/package")
	if err == nil {
		t.Error("composer.json不存在应该返回错误")
	}
}

func TestSearchWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试搜索失败
	SetupMockOutput("search nonexistent-term", "", errors.New("Search failed"))
	_, err = composer.Search("nonexistent-term")
	if err == nil {
		t.Error("搜索失败应该返回错误")
	}

	// 测试网络连接失败
	SetupMockOutput("search logger", "", errors.New("Network connection failed"))
	_, err = composer.Search("logger")
	if err == nil {
		t.Error("网络连接失败应该返回错误")
	}

	// 测试API限制错误
	SetupMockOutput("search popular-package", "", errors.New("API rate limit exceeded"))
	_, err = composer.Search("popular-package")
	if err == nil {
		t.Error("API限制错误应该返回错误")
	}
}

func TestSearch(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for search command
	SetupMockOutput("search logger", "Found 5 packages matching logger", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Search("logger")
	if err != nil {
		t.Errorf("Search执行失败: %v", err)
	}

	if output == "" || !contains(output, "Found") {
		t.Errorf("输出应包含搜索结果，实际为\"%s\"", output)
	}
}

func TestShowAllPackages(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持show命令（不带参数）
	extendMockScript(t, execPath, "show", "vendor/package1    v1.0.0  Package description\nvendor/package2    v2.0.0  Another package")

	output, err := composer.ShowAllPackages()
	if err != nil {
		t.Errorf("ShowAllPackages执行失败: %v", err)
	}

	if output == "" || !contains(output, "vendor/package") {
		t.Errorf("输出应包含所有包的列表，实际为\"%s\"", output)
	}
}

func TestShowDependencyTree(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持show --tree命令
	extendMockScript(t, execPath, "show --tree", "vendor/package v1.0.0\n└──vendor/dependency v2.0.0")

	// 测试无包名参数
	output, err := composer.ShowDependencyTree("")
	if err != nil {
		t.Errorf("ShowDependencyTree（无包名）执行失败: %v", err)
	}

	if output == "" || !contains(output, "vendor/package") {
		t.Errorf("输出应包含依赖树，实际为\"%s\"", output)
	}

	// 测试有包名参数
	extendMockScript(t, execPath, "show --tree vendor/package", "vendor/package v1.0.0\n└──vendor/dependency v2.0.0")

	output, err = composer.ShowDependencyTree("vendor/package")
	if err != nil {
		t.Errorf("ShowDependencyTree（有包名）执行失败: %v", err)
	}

	if output == "" || !contains(output, "vendor/package") {
		t.Errorf("输出应包含指定包的依赖树，实际为\"%s\"", output)
	}
}

func TestShowReverseDependencies(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持depends命令
	extendMockScript(t, execPath, "depends vendor/package", "project1 requires vendor/package (^1.0)\nproject2 requires vendor/package (^2.0)")

	output, err := composer.ShowReverseDependencies("vendor/package")
	if err != nil {
		t.Errorf("ShowReverseDependencies执行失败: %v", err)
	}

	if output == "" || !contains(output, "requires vendor/package") {
		t.Errorf("输出应包含反向依赖信息，实际为\"%s\"", output)
	}
}

func TestWhyPackage(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持why命令
	extendMockScript(t, execPath, "why vendor/package", "project/name requires vendor/package (^1.0)")

	output, err := composer.WhyPackage("vendor/package")
	if err != nil {
		t.Errorf("WhyPackage执行失败: %v", err)
	}

	if output == "" || !contains(output, "requires vendor/package") {
		t.Errorf("输出应包含依赖原因，实际为\"%s\"", output)
	}
}

func TestOutdatedPackages(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持outdated命令
	extendMockScript(t, execPath, "outdated", "vendor/package 1.0.0 < 1.1.0")

	output, err := composer.OutdatedPackages()
	if err != nil {
		t.Errorf("OutdatedPackages执行失败: %v", err)
	}

	if output == "" || !contains(output, "vendor/package") {
		t.Errorf("输出应包含过期包信息，实际为\"%s\"", output)
	}
}

func TestOutdatedPackagesDirect(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持outdated --direct命令
	extendMockScript(t, execPath, "outdated --direct", "vendor/package 1.0.0 < 1.1.0")

	output, err := composer.OutdatedPackagesDirect()
	if err != nil {
		t.Errorf("OutdatedPackagesDirect执行失败: %v", err)
	}

	if output == "" || !contains(output, "vendor/package") {
		t.Errorf("输出应包含直接依赖的过期包信息，实际为\"%s\"", output)
	}
}

func TestRequirePackageWithOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持require命令
	extendMockScript(t, execPath, "require --update-with-dependencies --no-scripts vendor/package:^1.0", "Using version ^1.0 for vendor/package")

	options := map[string]string{
		"update-with-dependencies": "",
		"no-scripts":               "",
	}

	err = composer.RequirePackageWithOptions("vendor/package", "^1.0", options)
	if err != nil {
		t.Errorf("RequirePackageWithOptions执行失败: %v", err)
	}
}

func TestBumpPackages(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持bump命令
	extendMockScript(t, execPath, "bump vendor/package1 vendor/package2", "Bumped vendor/package1 to ^2.0\nBumped vendor/package2 to ^3.0")

	packages := []string{"vendor/package1", "vendor/package2"}
	err = composer.BumpPackages(packages)
	if err != nil {
		t.Errorf("BumpPackages执行失败: %v", err)
	}
}

func TestBumpPackagesWithOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持带选项的bump命令
	extendMockScript(t, execPath, "bump --dev-only --no-scripts vendor/package1 vendor/package2", "Bumped vendor/package1 to ^2.0\nBumped vendor/package2 to ^3.0")

	packages := []string{"vendor/package1", "vendor/package2"}
	options := map[string]string{
		"dev-only":   "",
		"no-scripts": "",
	}

	err = composer.BumpPackagesWithOptions(packages, options)
	if err != nil {
		t.Errorf("BumpPackagesWithOptions执行失败: %v", err)
	}
}

func TestReinstall(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持reinstall命令
	extendMockScript(t, execPath, "reinstall vendor/package", "Reinstalling vendor/package (1.0.0)")

	err = composer.Reinstall("vendor/package")
	if err != nil {
		t.Errorf("Reinstall执行失败: %v", err)
	}
}

func TestBrowsePackage(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持browse命令
	extendMockScript(t, execPath, "browse vendor/package", "Opening vendor/package homepage")

	err = composer.BrowsePackage("vendor/package")
	if err != nil {
		t.Errorf("BrowsePackage执行失败: %v", err)
	}
}

func TestBrowsePackageWithOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持带选项的browse命令
	extendMockScript(t, execPath, "browse --homepage --no-browser vendor/package", "Opening vendor/package homepage")

	options := map[string]string{
		"homepage":   "",
		"no-browser": "",
	}

	err = composer.BrowsePackageWithOptions("vendor/package", options)
	if err != nil {
		t.Errorf("BrowsePackageWithOptions执行失败: %v", err)
	}
}

func TestWhyNotPackage(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持prohibits命令
	extendMockScript(t, execPath, "prohibits vendor/package", "vendor/package-conflict conflicts with vendor/package")

	// 测试无版本参数
	output, err := composer.WhyNotPackage("vendor/package", "")
	if err != nil {
		t.Errorf("WhyNotPackage（无版本）执行失败: %v", err)
	}

	if output == "" || !contains(output, "conflicts with") {
		t.Errorf("输出应包含冲突信息，实际为\"%s\"", output)
	}

	// 测试有版本参数
	extendMockScript(t, execPath, "prohibits vendor/package 2.0.0", "vendor/package-conflict conflicts with vendor/package 2.0.0")

	output, err = composer.WhyNotPackage("vendor/package", "2.0.0")
	if err != nil {
		t.Errorf("WhyNotPackage（有版本）执行失败: %v", err)
	}

	if output == "" || !contains(output, "2.0.0") {
		t.Errorf("输出应包含指定版本的冲突信息，实际为\"%s\"", output)
	}
}

// 扩展模拟可执行文件的辅助函数
func extendMockScript(t *testing.T, execPath string, command string, output string) {
	// 获取当前脚本内容
	content, err := os.ReadFile(execPath)
	if err != nil {
		t.Fatalf("读取模拟Composer可执行文件失败: %v", err)
	}

	// 在脚本中添加新命令
	// 将脚本按行分割
	lines := strings.Split(string(content), "\n")

	// 寻找最后一个判断行
	lastIfIndex := -1
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.Contains(lines[i], "elif") || strings.Contains(lines[i], "if") {
			lastIfIndex = i
			break
		}
	}

	// 在最后一个判断之后、else之前插入新的判断
	if lastIfIndex != -1 {
		// 找到else行
		elseIndex := -1
		for i := lastIfIndex; i < len(lines); i++ {
			if strings.Contains(lines[i], "else") {
				elseIndex = i
				break
			}
		}

		// 准备新命令的判断语句
		newCommand := fmt.Sprintf("elif [ \"$1 $2 $3 $4 $5\" = \"%s\" ]; then\n    echo \"%s\"\n    exit 0", command, strings.Replace(output, "\n", "\\n", -1))

		// 插入新命令
		if elseIndex != -1 {
			// 在else之前插入
			lines = append(lines[:elseIndex], append([]string{newCommand}, lines[elseIndex:]...)...)
		} else {
			// 没有找到else，在末尾添加新命令和else语句
			lines = append(lines, newCommand, "else", "    echo \"Unknown command: $1\"", "    exit 1", "fi")
		}
	}

	// 写回脚本文件
	err = os.WriteFile(execPath, []byte(strings.Join(lines, "\n")), 0755)
	if err != nil {
		t.Fatalf("更新模拟Composer可执行文件失败: %v", err)
	}
}

// 测试边界情况和错误处理
func TestShowPackageWithEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.ShowPackage("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestSearchWithEmptyQuery(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Search("")
	if err == nil {
		t.Error("空搜索查询应该返回错误")
	}
}

func TestWhyPackageWithEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.WhyPackage("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestReinstallWithEmptyPackage(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Reinstall("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestBrowsePackageWithEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.BrowsePackage("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestBumpPackagesWithEmptyArray(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.BumpPackages([]string{})
	if err == nil {
		t.Error("空包数组应该返回错误")
	}
}

func TestRequirePackageWithOptionsAndEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.RequirePackageWithOptions("", "^1.0", map[string]string{})
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestBrowsePackageWithOptionsAndEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.BrowsePackageWithOptions("", map[string]string{})
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestWhyNotPackageWithEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.WhyNotPackage("", "")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}
