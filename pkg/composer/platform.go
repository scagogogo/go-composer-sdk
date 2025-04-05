package composer

import (
	"encoding/json"
	"strings"
)

// PlatformInfo 表示平台信息
type PlatformInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Available bool   `json:"available"`
	Required  string `json:"required,omitempty"`
}

// PlatformRequirements 表示平台需求
type PlatformRequirements struct {
	Platform map[string]PlatformInfo `json:"platform"`
	Lock     map[string]PlatformInfo `json:"lock,omitempty"`
}

// CheckPlatform 检查平台需求
func (c *Composer) CheckPlatform() ([]PlatformInfo, error) {
	output, err := c.Run("check-platform", "--format=json")
	if err != nil {
		return nil, err
	}

	var platformReqs PlatformRequirements
	if err := json.Unmarshal([]byte(output), &platformReqs); err != nil {
		return nil, err
	}

	var result []PlatformInfo
	for _, info := range platformReqs.Platform {
		result = append(result, info)
	}

	return result, nil
}

// CheckPlatformWithLock 检查 composer.lock 文件中的平台需求
func (c *Composer) CheckPlatformWithLock() ([]PlatformInfo, error) {
	output, err := c.Run("check-platform", "--lock", "--format=json")
	if err != nil {
		return nil, err
	}

	var platformReqs PlatformRequirements
	if err := json.Unmarshal([]byte(output), &platformReqs); err != nil {
		return nil, err
	}

	var result []PlatformInfo
	for _, info := range platformReqs.Lock {
		result = append(result, info)
	}

	return result, nil
}

// IsPlatformAvailable 检查指定的平台需求是否满足
func (c *Composer) IsPlatformAvailable(platform string, version string) (bool, error) {
	platformReqs, err := c.CheckPlatform()
	if err != nil {
		return false, err
	}

	for _, info := range platformReqs {
		if info.Name == platform {
			return info.Available, nil
		}
	}

	// 平台未在需求中列出，尝试直接检查
	versionCheck := platform
	if version != "" {
		versionCheck += ":" + version
	}

	output, err := c.Run("check-platform", versionCheck)
	if err != nil {
		return false, err
	}

	// 如果输出包含 "is not available"，则平台不可用
	return !strings.Contains(output, "is not available"), nil
}

// GetPHPVersion 获取当前 PHP 版本
func (c *Composer) GetPHPVersion() (string, error) {
	output, err := c.Run("run", "--php-show-version")
	if err != nil {
		return "", err
	}

	// 解析输出以提取 PHP 版本
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PHP ") {
			parts := strings.Split(line, " ")
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}

	return "", nil
}

// GetExtensions 获取已安装的 PHP 扩展
func (c *Composer) GetExtensions() ([]string, error) {
	output, err := c.Run("run", "--show-extensions")
	if err != nil {
		return nil, err
	}

	// 解析输出以提取扩展名
	extensions := []string{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "Loaded extensions:") {
			extensions = append(extensions, line)
		}
	}

	return extensions, nil
}

// HasExtension 检查是否安装了指定的 PHP 扩展
func (c *Composer) HasExtension(extension string) (bool, error) {
	extensions, err := c.GetExtensions()
	if err != nil {
		return false, err
	}

	for _, ext := range extensions {
		if ext == extension {
			return true, nil
		}
	}

	return false, nil
}
