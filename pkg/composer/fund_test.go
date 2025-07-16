package composer

import (
	"errors"
	"testing"
)

func TestFundPackages(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command
	SetupMockOutput("fund", "The following packages were found in your dependencies which publish funding information:\n\nvendor/package1\n  https://github.com/sponsors/vendor\n  https://patreon.com/vendor\n\nvendor/package2\n  https://opencollective.com/package2", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundPackages()
	if err != nil {
		t.Errorf("FundPackages执行失败: %v", err)
	}

	if !contains(output, "funding information") {
		t.Errorf("输出应包含资助信息，实际为\"%s\"", output)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestFundPackagesWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("fund", "", errors.New("fund command failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.FundPackages()
	if err == nil {
		t.Error("Fund命令失败时应该返回错误")
	}
}

func TestFundPackagesWithNoFunding(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with no funding information
	SetupMockOutput("fund", "No funding links were found in your package dependencies.", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundPackages()
	if err != nil {
		t.Errorf("FundPackages执行失败: %v", err)
	}

	if !contains(output, "No funding links") {
		t.Errorf("输出应包含无资助信息的提示，实际为\"%s\"", output)
	}
}

func TestFundWithJSONFormat(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command with JSON format
	SetupMockOutput("fund --format=json", `[{"name": "vendor/package1", "funding": [{"type": "github", "url": "https://github.com/sponsors/vendor"}]}, {"name": "vendor/package2", "funding": [{"type": "opencollective", "url": "https://opencollective.com/package2"}]}]`, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	fundingInfo, err := composer.FundWithJSON()
	if err != nil {
		t.Errorf("FundWithJSON执行失败: %v", err)
	}

	if len(fundingInfo) == 0 {
		t.Error("应该返回资助信息")
	}

	if len(fundingInfo) > 0 && fundingInfo[0].Name != "vendor/package1" {
		t.Errorf("第一个包名应为vendor/package1，实际为%s", fundingInfo[0].Name)
	}
}

func TestFundWithOptionsAndInvalidFormat(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error for invalid format
	SetupMockOutput("fund --format=invalid", "", errors.New("invalid format"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"format": "invalid",
	}

	_, err = composer.FundWithOptions(options)
	if err == nil {
		t.Error("无效格式应该返回错误")
	}
}

func TestFundWithPackage(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command with specific package
	SetupMockOutput("fund vendor/package1", "vendor/package1\n  https://github.com/sponsors/vendor\n  https://patreon.com/vendor", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundWithPackage("vendor/package1")
	if err != nil {
		t.Errorf("FundWithPackage执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含指定包名，实际为\"%s\"", output)
	}

	if !contains(output, "github.com/sponsors") {
		t.Errorf("输出应包含资助链接，实际为\"%s\"", output)
	}
}

func TestFundWithPackageAndEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.FundWithPackage("")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestFundWithPackageAndNoFunding(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for package with no funding
	SetupMockOutput("fund vendor/no-funding", "vendor/no-funding does not publish funding information.", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundWithPackage("vendor/no-funding")
	if err != nil {
		t.Errorf("FundWithPackage执行失败: %v", err)
	}

	if !contains(output, "does not publish funding") {
		t.Errorf("输出应包含无资助信息的提示，实际为\"%s\"", output)
	}
}

func TestFundWithOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command with options
	SetupMockOutput("fund --format=json --no-dev", `{"vendor/package1": [{"type": "github", "url": "https://github.com/sponsors/vendor"}]}`, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"format": "json",
		"no-dev": "",
	}

	output, err := composer.FundWithOptions(options)
	if err != nil {
		t.Errorf("FundWithOptions执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestFundWithNilOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command without options
	SetupMockOutput("fund", "The following packages were found in your dependencies which publish funding information:\n\nvendor/package1\n  https://github.com/sponsors/vendor", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundWithOptions(nil)
	if err != nil {
		t.Errorf("FundWithOptions（nil选项）执行失败: %v", err)
	}

	if !contains(output, "funding information") {
		t.Errorf("输出应包含资助信息，实际为\"%s\"", output)
	}
}

func TestFundWithEmptyOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for fund command without options
	SetupMockOutput("fund", "The following packages were found in your dependencies which publish funding information:\n\nvendor/package1\n  https://github.com/sponsors/vendor", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.FundWithOptions(map[string]string{})
	if err != nil {
		t.Errorf("FundWithOptions（空选项）执行失败: %v", err)
	}

	if !contains(output, "funding information") {
		t.Errorf("输出应包含资助信息，实际为\"%s\"", output)
	}
}
