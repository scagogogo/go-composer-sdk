package composer

import (
	"strings"
)

// Validate 验证composer.json是否有效
func (c *Composer) Validate() error {
	_, err := c.Run("validate")
	return err
}

// GetComposerHome 获取Composer主目录
func (c *Composer) GetComposerHome() (string, error) {
	output, err := c.Run("config", "--global", "home")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// ClearCache 清除Composer缓存
func (c *Composer) ClearCache() error {
	_, err := c.Run("clear-cache")
	return err
}

// GetConfigWithGlobal 获取Composer配置项的值，可选择是否获取全局配置
func (c *Composer) GetConfigWithGlobal(setting string, global bool) (string, error) {
	args := []string{"config"}

	if global {
		args = append(args, "--global")
	}

	args = append(args, setting)

	output, err := c.Run(args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// SetConfigWithGlobal 设置Composer配置项的值，可选择是否设置全局配置
func (c *Composer) SetConfigWithGlobal(setting string, value string, global bool) error {
	args := []string{"config"}

	if global {
		args = append(args, "--global")
	}

	args = append(args, setting, value)

	_, err := c.Run(args...)
	return err
}

// ValidateComposerJson 验证composer.json，可选参数：strict和with-dependencies
func (c *Composer) ValidateComposerJson(strict bool, withDependencies bool) error {
	args := []string{"validate"}

	if strict {
		args = append(args, "--strict")
	}

	if withDependencies {
		args = append(args, "--with-dependencies")
	}

	_, err := c.Run(args...)
	return err
}

// CheckPlatformReqs 检查平台要求
func (c *Composer) CheckPlatformReqs() (string, error) {
	return c.Run("check-platform-reqs")
}
