package composer

import (
	"encoding/json"
	"strings"
)

// AuditResult 表示安全审计结果
type AuditResult struct {
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Found           int             `json:"found"`
	Advisory        string          `json:"advisory,omitempty"`
	WithoutDev      bool            `json:"without-dev,omitempty"`
}

// Vulnerability 表示一个安全漏洞
type Vulnerability struct {
	Package     string   `json:"package"`
	Version     string   `json:"version"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	CVE         []string `json:"cve,omitempty"`
	Advisory    string   `json:"advisory"`
	Abandoned   bool     `json:"abandoned,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Source      string   `json:"source,omitempty"`
	Affectedver string   `json:"affectedver,omitempty"`
}

// Audit 执行安全审计
func (c *Composer) Audit() (string, error) {
	return c.Run("audit")
}

// AuditWithJSON 执行安全审计并返回 JSON 格式结果
func (c *Composer) AuditWithJSON() (*AuditResult, error) {
	output, err := c.Run("audit", "--format=json")
	if err != nil {
		return nil, err
	}

	var result AuditResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AuditWithoutDev 执行安全审计，不包含开发依赖
func (c *Composer) AuditWithoutDev() (string, error) {
	return c.Run("audit", "--no-dev")
}

// AuditWithFormat 使用指定格式执行安全审计
func (c *Composer) AuditWithFormat(format string) (string, error) {
	return c.Run("audit", "--format="+format)
}

// HasVulnerabilities 检查当前项目是否有漏洞
func (c *Composer) HasVulnerabilities() (bool, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		// 如果命令失败且出错信息包含漏洞，也认为有漏洞
		if strings.Contains(err.Error(), "Found") && strings.Contains(err.Error(), "vulnerabilities") {
			return true, nil
		}
		return false, err
	}

	return result.Found > 0, nil
}

// GetHighSeverityVulnerabilities 获取高严重性的漏洞
func (c *Composer) GetHighSeverityVulnerabilities() ([]Vulnerability, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		return nil, err
	}

	var highSeverity []Vulnerability
	for _, vuln := range result.Vulnerabilities {
		if vuln.Severity == "high" || vuln.Severity == "critical" {
			highSeverity = append(highSeverity, vuln)
		}
	}

	return highSeverity, nil
}

// AuditWithOptions 使用自定义选项执行安全审计
func (c *Composer) AuditWithOptions(options map[string]string) (string, error) {
	args := []string{"audit"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// AuditLock 审计 composer.lock 文件
func (c *Composer) AuditLock(lockFilePath string) (string, error) {
	if lockFilePath == "" {
		return c.Run("audit")
	}
	return c.Run("audit", lockFilePath)
}

// GetAbandonedPackages 获取已放弃的包
func (c *Composer) GetAbandonedPackages() ([]Vulnerability, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		return nil, err
	}

	var abandoned []Vulnerability
	for _, vuln := range result.Vulnerabilities {
		if vuln.Abandoned {
			abandoned = append(abandoned, vuln)
		}
	}

	return abandoned, nil
}
