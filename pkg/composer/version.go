package composer

import (
	"fmt"
	"strings"
)

// GetVersion 获取composer版本
func (c *Composer) GetVersion() (string, error) {
	output, err := c.Run("--version")
	if err != nil {
		return "", err
	}

	// 解析版本信息
	// 示例输出: "Composer version 2.1.6 2021-08-19 17:11:08"
	parts := strings.Split(output, " ")
	if len(parts) >= 3 {
		return parts[2], nil
	}

	return "", fmt.Errorf("无法解析版本信息: %s", output)
}

// SelfUpdate 执行composer self-update命令
func (c *Composer) SelfUpdate() error {
	_, err := c.Run("self-update")
	return err
}
