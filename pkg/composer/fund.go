package composer

import (
	"encoding/json"
	"strings"
)

// FundingInfo 表示一个包的资金信息
type FundingInfo struct {
	Name    string   `json:"name"`
	URLs    []string `json:"urls"`
	Funding bool     `json:"funding"`
}

// Fund 显示项目的资金信息
func (c *Composer) Fund() (string, error) {
	return c.Run("fund")
}

// FundWithJSON 获取 JSON 格式的资金信息
func (c *Composer) FundWithJSON() ([]FundingInfo, error) {
	output, err := c.Run("fund", "--format=json")
	if err != nil {
		return nil, err
	}

	// 解析 JSON 输出
	var fundingInfo []FundingInfo
	if err := json.Unmarshal([]byte(output), &fundingInfo); err != nil {
		return nil, err
	}

	return fundingInfo, nil
}

// FundWithPackage 显示特定包的资金信息
func (c *Composer) FundWithPackage(packageName string) (string, error) {
	return c.Run("fund", packageName)
}

// GetFundingURLs 获取所有支持资金的包的 URL
func (c *Composer) GetFundingURLs() (map[string][]string, error) {
	fundingInfo, err := c.FundWithJSON()
	if err != nil {
		return nil, err
	}

	// 创建包名到 URL 的映射
	urls := make(map[string][]string)

	for _, info := range fundingInfo {
		if info.Funding && len(info.URLs) > 0 {
			urls[info.Name] = info.URLs
		}
	}

	return urls, nil
}

// HasFunding 检查项目是否有任何资金支持
func (c *Composer) HasFunding() (bool, error) {
	output, err := c.Run("fund", "--format=text")
	if err != nil {
		return false, err
	}

	// 如果输出包含 "No funding"，则没有资金支持
	return !strings.Contains(output, "No funding"), nil
}

// FundWithOptions 使用自定义选项显示资金信息
func (c *Composer) FundWithOptions(options map[string]string) (string, error) {
	args := []string{"fund"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}
