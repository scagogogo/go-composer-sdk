package composer

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestAudit(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for audit command
	SetupMockOutput("audit", "No security vulnerabilities found", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Audit()
	if err != nil {
		t.Errorf("Audit执行失败: %v", err)
	}

	if output == "" || !contains(output, "No security vulnerabilities found") {
		t.Errorf("输出应包含审计结果，实际为\"%s\"", output)
	}
}

func TestAuditWithJSON(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for audit command with JSON format
	jsonOutput := `{"vulnerabilities":[{"package":"vendor/package","version":"1.0.0","title":"Security Issue","link":"https://example.com/advisory","advisory":"CVE-2023-1234","severity":"high"}],"found":1}`
	SetupMockOutput("audit --format=json", jsonOutput, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	result, err := composer.AuditWithJSON()
	if err != nil {
		t.Errorf("AuditWithJSON执行失败: %v", err)
	}

	if result == nil {
		t.Error("审计结果不应为nil")
		return
	}

	if result.Found != 1 {
		t.Errorf("应找到1个漏洞，实际为%d", result.Found)
	}

	if len(result.Vulnerabilities) != 1 {
		t.Errorf("漏洞数组应包含1个元素，实际为%d", len(result.Vulnerabilities))
		return
	}

	vuln := result.Vulnerabilities[0]
	if vuln.Package != "vendor/package" {
		t.Errorf("漏洞包名应为vendor/package，实际为%s", vuln.Package)
	}

	if vuln.Severity != "high" {
		t.Errorf("漏洞严重性应为high，实际为%s", vuln.Severity)
	}
}

func TestAuditWithoutDev(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for audit --no-dev command
	SetupMockOutput("audit --no-dev", "No security vulnerabilities found in production dependencies", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.AuditWithoutDev()
	if err != nil {
		t.Errorf("AuditWithoutDev执行失败: %v", err)
	}

	if output == "" || !contains(output, "production dependencies") {
		t.Errorf("输出应包含生产依赖的审计结果，实际为\"%s\"", output)
	}
}

func TestAuditWithFormat(t *testing.T) {
	// 测试不同的格式
	formats := []string{"text", "json", "summary"}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			// Reset mock outputs before test
			ClearMockOutputs()

			// 设置模拟输出
			output := fmt.Sprintf("Audit output in %s format", format)
			if format == "json" {
				output = `{"vulnerabilities":[],"found":0}`
			}

			SetupMockOutput("audit --format="+format, output, nil)

			composer, err := New(Options{ExecutablePath: "/path/to/composer"})
			if err != nil {
				t.Fatalf("创建Composer实例失败: %v", err)
			}

			result, err := composer.AuditWithFormat(format)
			if err != nil {
				t.Errorf("AuditWithFormat(%s)执行失败: %v", format, err)
			}

			if result == "" {
				t.Errorf("审计结果不应为空")
			}
		})
	}
}

func TestHasVulnerabilities(t *testing.T) {
	// 测试有漏洞的情况
	ClearMockOutputs()
	jsonWithVuln := `{"vulnerabilities":[{"package":"vendor/package","version":"1.0.0","title":"Security Issue","advisory":"CVE-2023-1234"}],"found":1}`
	SetupMockOutput("audit --format=json", jsonWithVuln, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	hasVuln, err := composer.HasVulnerabilities()
	if err != nil {
		t.Errorf("HasVulnerabilities执行失败: %v", err)
	}

	if !hasVuln {
		t.Error("应检测到漏洞")
	}

	// 测试无漏洞的情况
	ClearMockOutputs()
	jsonNoVuln := `{"vulnerabilities":[],"found":0}`
	SetupMockOutput("audit --format=json", jsonNoVuln, nil)

	hasVuln, err = composer.HasVulnerabilities()
	if err != nil {
		t.Errorf("HasVulnerabilities执行失败: %v", err)
	}

	if hasVuln {
		t.Error("不应检测到漏洞")
	}
}

func TestGetHighSeverityVulnerabilities(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 准备一个包含不同严重性漏洞的模拟输出
	jsonWithVulns := `{
		"vulnerabilities": [
			{"package": "vendor/package1", "version": "1.0.0", "title": "Critical Issue", "advisory": "CVE-2023-1234", "severity": "critical"},
			{"package": "vendor/package2", "version": "2.0.0", "title": "High Issue", "advisory": "CVE-2023-5678", "severity": "high"},
			{"package": "vendor/package3", "version": "3.0.0", "title": "Medium Issue", "advisory": "CVE-2023-9012", "severity": "medium"}
		],
		"found": 3
	}`

	extendMockScript(t, execPath, "audit --format=json", jsonWithVulns)

	highVulns, err := composer.GetHighSeverityVulnerabilities()
	if err != nil {
		t.Errorf("GetHighSeverityVulnerabilities执行失败: %v", err)
	}

	if len(highVulns) != 2 {
		t.Errorf("应检测到2个高危漏洞，实际为%d", len(highVulns))
	}

	// 验证检测到的确实是高危漏洞
	for _, vuln := range highVulns {
		if vuln.Severity != "high" && vuln.Severity != "critical" {
			t.Errorf("漏洞严重性应为high或critical，实际为%s", vuln.Severity)
		}
	}
}

func TestAuditWithOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试带选项的审计
	options := map[string]string{
		"no-dev": "",
		"format": "json",
	}

	extendMockScript(t, execPath, "audit --no-dev --format=json", `{"vulnerabilities":[],"found":0,"without-dev":true}`)

	output, err := composer.AuditWithOptions(options)
	if err != nil {
		t.Errorf("AuditWithOptions执行失败: %v", err)
	}

	if output == "" || !contains(output, "without-dev") {
		t.Errorf("输出应包含审计结果，实际为\"%s\"", output)
	}
}

func TestAuditLock(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试默认情况
	extendMockScript(t, execPath, "audit", "No security vulnerabilities found")

	output, err := composer.AuditLock("")
	if err != nil {
		t.Errorf("AuditLock(\"\")执行失败: %v", err)
	}

	if output == "" {
		t.Error("审计结果不应为空")
	}

	// 测试指定lock文件路径
	lockPath := "/path/to/composer.lock"
	extendMockScript(t, execPath, "audit "+lockPath, "No security vulnerabilities found in "+lockPath)

	output, err = composer.AuditLock(lockPath)
	if err != nil {
		t.Errorf("AuditLock(%s)执行失败: %v", lockPath, err)
	}

	if output == "" || !contains(output, lockPath) {
		t.Errorf("输出应包含指定lock文件的审计结果，实际为\"%s\"", output)
	}
}

func TestGetAbandonedPackages(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 准备一个包含已放弃包的模拟输出
	jsonWithAbandoned := `{
		"vulnerabilities": [
			{"package": "vendor/abandoned1", "version": "1.0.0", "title": "Abandoned Package", "advisory": "abandoned", "abandoned": true},
			{"package": "vendor/active", "version": "2.0.0", "title": "Active Package", "advisory": "active"}
		],
		"found": 2
	}`

	extendMockScript(t, execPath, "audit --format=json", jsonWithAbandoned)

	abandoned, err := composer.GetAbandonedPackages()
	if err != nil {
		t.Errorf("GetAbandonedPackages执行失败: %v", err)
	}

	if len(abandoned) != 1 {
		t.Errorf("应检测到1个已放弃的包，实际为%d", len(abandoned))
	}

	if len(abandoned) > 0 && abandoned[0].Package != "vendor/abandoned1" {
		t.Errorf("已放弃的包应为vendor/abandoned1，实际为%s", abandoned[0].Package)
	}
}

// 测试边界情况和错误处理
func TestAuditWithEmptyOutput(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持空输出
	extendMockScript(t, execPath, "audit", "")

	output, err := composer.Audit()
	if err != nil {
		t.Errorf("Audit执行失败: %v", err)
	}

	if output != "" {
		t.Errorf("空输出应该返回空字符串，实际为\"%s\"", output)
	}
}

func TestAuditWithInvalidJSON(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持无效JSON输出
	extendMockScript(t, execPath, "audit --format=json", "invalid json")

	_, err = composer.AuditWithJSON()
	if err == nil {
		t.Error("无效JSON应该返回错误")
	}
}

func TestAuditWithFormatAndEmptyFormat(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.AuditWithFormat("")
	if err == nil {
		t.Error("空格式应该返回错误")
	}
}

func TestAuditWithInvalidFormat(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持无效格式错误
	extendMockScript(t, execPath, "audit --format=invalid", "Error: Invalid format 'invalid'")

	_, err = composer.AuditWithFormat("invalid")
	if err != nil {
		// 这里我们期望有错误，因为格式无效
		return
	}

	// 如果没有错误，检查输出是否包含错误信息
	output, _ := composer.AuditWithFormat("invalid")
	if !contains(output, "Error") {
		t.Error("无效格式应该返回错误或错误信息")
	}
}

func TestHasVulnerabilitiesWithInvalidJSON(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with invalid JSON
	SetupMockOutput("audit --format=json", "invalid json", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.HasVulnerabilities()
	if err == nil {
		t.Error("无效JSON应该返回错误")
	}
}

func TestAuditEdgeCases(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试空输出
	SetupMockOutput("audit", "", nil)
	output, err := composer.Audit()
	if err != nil {
		t.Errorf("空输出Audit执行失败: %v", err)
	}
	if output != "" {
		t.Errorf("空输出应返回空字符串，实际为'%s'", output)
	}

	// 测试包含特殊字符的输出
	specialOutput := "Audit output with special chars: !@#$%^&*()_+-=[]{}|;':\",./<>?"
	SetupMockOutput("audit", specialOutput, nil)
	output, err = composer.Audit()
	if err != nil {
		t.Errorf("特殊字符输出Audit执行失败: %v", err)
	}
	if output != specialOutput {
		t.Errorf("特殊字符输出不正确，期望: '%s'，实际: '%s'", specialOutput, output)
	}

	// 测试非常长的输出
	longOutput := strings.Repeat("Very long audit output line. ", 100)
	SetupMockOutput("audit", longOutput, nil)
	output, err = composer.Audit()
	if err != nil {
		t.Errorf("长输出Audit执行失败: %v", err)
	}
	if len(output) != len(longOutput) {
		t.Errorf("长输出长度不正确，期望: %d，实际: %d", len(longOutput), len(output))
	}

	// 测试包含Unicode的输出
	unicodeOutput := "审计输出包含中文字符和emoji 🔒🛡️"
	SetupMockOutput("audit", unicodeOutput, nil)
	output, err = composer.Audit()
	if err != nil {
		t.Errorf("Unicode输出Audit执行失败: %v", err)
	}
	if output != unicodeOutput {
		t.Errorf("Unicode输出不正确，期望: '%s'，实际: '%s'", unicodeOutput, output)
	}

	// 测试多行输出
	multilineOutput := "Line 1\nLine 2\nLine 3\n\nLine 5"
	SetupMockOutput("audit", multilineOutput, nil)
	output, err = composer.Audit()
	if err != nil {
		t.Errorf("多行输出Audit执行失败: %v", err)
	}
	if output != multilineOutput {
		t.Errorf("多行输出不正确")
	}
}

func TestAuditWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试composer.lock不存在
	SetupMockOutput("audit", "", errors.New("composer.lock not found"))
	_, err = composer.Audit()
	if err == nil {
		t.Error("composer.lock不存在应该返回错误")
	}

	// 测试网络连接失败
	SetupMockOutput("audit", "", errors.New("Could not connect to security advisory database"))
	_, err = composer.Audit()
	if err == nil {
		t.Error("网络连接失败应该返回错误")
	}

	// 测试权限错误
	SetupMockOutput("audit", "", errors.New("Permission denied"))
	_, err = composer.Audit()
	if err == nil {
		t.Error("权限错误应该返回错误")
	}

	// 测试无效的composer.lock格式
	SetupMockOutput("audit", "", errors.New("Invalid composer.lock format"))
	_, err = composer.Audit()
	if err == nil {
		t.Error("无效composer.lock格式应该返回错误")
	}

	// 测试API服务不可用
	SetupMockOutput("audit", "", errors.New("Security advisory service unavailable"))
	_, err = composer.Audit()
	if err == nil {
		t.Error("API服务不可用应该返回错误")
	}
}

func TestAuditWithFormatErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试无效格式
	SetupMockOutput("audit --format=invalid", "", errors.New("Invalid format specified"))
	_, err = composer.AuditWithFormat("invalid")
	if err == nil {
		t.Error("无效格式应该返回错误")
	}

	// 测试JSON格式解析错误
	SetupMockOutput("audit --format=json", "invalid json output", nil)
	_, err = composer.AuditWithFormat("json")
	if err != nil {
		t.Errorf("JSON格式解析错误测试失败: %v", err)
	}

	// 测试空格式参数
	SetupMockOutput("audit --format=", "", errors.New("Format parameter cannot be empty"))
	_, err = composer.AuditWithFormat("")
	if err == nil {
		t.Error("空格式参数应该返回错误")
	}
}

func TestHasVulnerabilitiesWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试audit命令失败
	SetupMockOutput("audit --format=json", "", errors.New("Audit command failed"))
	_, err = composer.HasVulnerabilities()
	if err == nil {
		t.Error("audit命令失败应该返回错误")
	}

	// 测试JSON解析失败
	SetupMockOutput("audit --format=json", "not valid json", nil)
	_, err = composer.HasVulnerabilities()
	if err == nil {
		t.Error("JSON解析失败应该返回错误")
	}

	// 测试空JSON响应
	SetupMockOutput("audit --format=json", "", nil)
	_, err = composer.HasVulnerabilities()
	if err == nil {
		t.Error("空JSON响应应该返回错误")
	}
}

func TestGetHighSeverityVulnerabilitiesWithEmptyResult(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持空漏洞结果
	extendMockScript(t, execPath, "audit --format=json", `{"vulnerabilities":[],"found":0}`)

	highVulns, err := composer.GetHighSeverityVulnerabilities()
	if err != nil {
		t.Errorf("GetHighSeverityVulnerabilities执行失败: %v", err)
	}

	if len(highVulns) != 0 {
		t.Errorf("空结果应该返回0个漏洞，实际为%d", len(highVulns))
	}
}

func TestAuditWithOptionsAndNilOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持无选项的audit命令
	extendMockScript(t, execPath, "audit", "No security vulnerabilities found")

	output, err := composer.AuditWithOptions(nil)
	if err != nil {
		t.Errorf("AuditWithOptions（nil选项）执行失败: %v", err)
	}

	if !contains(output, "No security") {
		t.Errorf("输出应包含审计结果，实际为\"%s\"", output)
	}
}

func TestAuditWithOptionsAndEmptyOptions(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持无选项的audit命令
	extendMockScript(t, execPath, "audit", "No security vulnerabilities found")

	output, err := composer.AuditWithOptions(map[string]string{})
	if err != nil {
		t.Errorf("AuditWithOptions（空选项）执行失败: %v", err)
	}

	if !contains(output, "No security") {
		t.Errorf("输出应包含审计结果，实际为\"%s\"", output)
	}
}
