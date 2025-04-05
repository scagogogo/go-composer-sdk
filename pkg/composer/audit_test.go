package composer

import (
	"fmt"
	"testing"
)

func TestAudit(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持audit命令
	extendMockScript(t, execPath, "audit", "No security vulnerabilities found")

	output, err := composer.Audit()
	if err != nil {
		t.Errorf("Audit执行失败: %v", err)
	}

	if output == "" || !contains(output, "No security vulnerabilities found") {
		t.Errorf("输出应包含审计结果，实际为\"%s\"", output)
	}
}

func TestAuditWithJSON(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持audit --format=json命令
	jsonOutput := `{"vulnerabilities":[{"package":"vendor/package","version":"1.0.0","title":"Security Issue","link":"https://example.com/advisory","advisory":"CVE-2023-1234","severity":"high"}],"found":1}`
	extendMockScript(t, execPath, "audit --format=json", jsonOutput)

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
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 扩展模拟可执行文件以支持audit --no-dev命令
	extendMockScript(t, execPath, "audit --no-dev", "No security vulnerabilities found in production dependencies")

	output, err := composer.AuditWithoutDev()
	if err != nil {
		t.Errorf("AuditWithoutDev执行失败: %v", err)
	}

	if output == "" || !contains(output, "production dependencies") {
		t.Errorf("输出应包含生产依赖的审计结果，实际为\"%s\"", output)
	}
}

func TestAuditWithFormat(t *testing.T) {
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试不同的格式
	formats := []string{"text", "json", "summary"}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			// 扩展模拟可执行文件以支持不同格式的audit命令
			output := fmt.Sprintf("Audit output in %s format", format)
			if format == "json" {
				output = `{"vulnerabilities":[],"found":0}`
			}

			extendMockScript(t, execPath, "audit --format="+format, output)

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
	execPath := createMockExecutable(t)

	composer, err := New(Options{ExecutablePath: execPath})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试有漏洞的情况
	jsonWithVuln := `{"vulnerabilities":[{"package":"vendor/package","version":"1.0.0","title":"Security Issue","advisory":"CVE-2023-1234"}],"found":1}`
	extendMockScript(t, execPath, "audit --format=json", jsonWithVuln)

	hasVuln, err := composer.HasVulnerabilities()
	if err != nil {
		t.Errorf("HasVulnerabilities执行失败: %v", err)
	}

	if !hasVuln {
		t.Error("应检测到漏洞")
	}

	// 测试无漏洞的情况
	jsonNoVuln := `{"vulnerabilities":[],"found":0}`
	extendMockScript(t, execPath, "audit --format=json", jsonNoVuln)

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
