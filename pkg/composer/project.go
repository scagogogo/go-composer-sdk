package composer

import (
	"encoding/json"
)

// CreateProject 创建一个新的项目
func (c *Composer) CreateProject(packageName string, directory string, version string) error {
	args := []string{"create-project"}

	if version != "" {
		args = append(args, packageName+":"+version, directory)
	} else {
		args = append(args, packageName, directory)
	}

	_, err := c.Run(args...)
	return err
}

// CreateProjectWithOptions 使用更多选项创建一个新的项目
func (c *Composer) CreateProjectWithOptions(packageName string, directory string, version string, options map[string]string) error {
	args := []string{"create-project"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	if version != "" {
		args = append(args, packageName+":"+version, directory)
	} else {
		args = append(args, packageName, directory)
	}

	_, err := c.Run(args...)
	return err
}

// InitProject 初始化一个新的项目
func (c *Composer) InitProject() error {
	_, err := c.Run("init")
	return err
}

// InitProjectWithOptions 使用更多选项初始化一个新的项目
func (c *Composer) InitProjectWithOptions(name string, description string, author string, options map[string]string) error {
	args := []string{"init"}

	// 添加命名选项
	if name != "" {
		args = append(args, "--name="+name)
	}

	if description != "" {
		args = append(args, "--description="+description)
	}

	if author != "" {
		args = append(args, "--author="+author)
	}

	// 添加其他选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	_, err := c.Run(args...)
	return err
}

// RunScript 执行composer脚本
func (c *Composer) RunScript(scriptName string, args ...string) (string, error) {
	cmdArgs := append([]string{"run-script", scriptName}, args...)
	return c.Run(cmdArgs...)
}

// ExecuteScript 执行composer.json中定义的自定义脚本
func (c *Composer) ExecuteScript(scriptName string) (string, error) {
	return c.Run("run", scriptName)
}

// Archive 创建项目的存档
func (c *Composer) ArchiveProject(directory string, format string) error {
	args := []string{"archive"}

	if directory != "" {
		args = append(args, "--dir="+directory)
	}

	if format != "" {
		args = append(args, "--format="+format)
	}

	_, err := c.Run(args...)
	return err
}

// ComposerJsonInfo 表示composer.json文件的部分信息
type ComposerJsonInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Require     map[string]string `json:"require"`
	RequireDev  map[string]string `json:"require-dev"`
}

// GetProjectInfo 获取项目信息
func (c *Composer) GetProjectInfo() (*ComposerJsonInfo, error) {
	output, err := c.Run("config", "--list", "--json")
	if err != nil {
		return nil, err
	}

	var info ComposerJsonInfo
	err = json.Unmarshal([]byte(output), &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// ListScripts 列出composer.json中定义的所有脚本
func (c *Composer) ListScripts() (string, error) {
	return c.Run("run-script", "--list")
}
