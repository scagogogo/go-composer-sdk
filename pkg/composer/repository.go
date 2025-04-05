package composer

import (
	"encoding/json"
	"fmt"
)

// RepositoryType 表示 Composer 仓库类型
type RepositoryType string

const (
	// VcsRepository 版本控制系统仓库类型
	VcsRepository RepositoryType = "vcs"
	// ComposerRepository Composer 仓库类型
	ComposerRepository RepositoryType = "composer"
	// PackagistRepository Packagist 仓库类型
	PackagistRepository RepositoryType = "packagist"
	// PathRepository 路径仓库类型
	PathRepository RepositoryType = "path"
	// ArtifactRepository artifact 仓库类型
	ArtifactRepository RepositoryType = "artifact"
	// PearRepository PEAR 仓库类型
	PearRepository RepositoryType = "pear"
)

// Repository 表示一个 Composer 仓库
type Repository struct {
	Type    RepositoryType         `json:"type"`
	URL     string                 `json:"url,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// AddRepository 添加一个仓库到 composer.json
func (c *Composer) AddRepository(name string, repo Repository) error {
	repoJSON, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	args := []string{"config", "repositories." + name, string(repoJSON)}
	_, err = c.Run(args...)
	return err
}

// RemoveRepository 从 composer.json 中移除仓库
func (c *Composer) RemoveRepository(name string) error {
	args := []string{"config", "--unset", "repositories." + name}
	_, err := c.Run(args...)
	return err
}

// ListRepositories 列出当前项目中配置的所有仓库
func (c *Composer) ListRepositories() (string, error) {
	return c.Run("config", "repositories")
}

// AddPackagistRepository 添加 Packagist.org 仓库
func (c *Composer) AddPackagistRepository(url string) error {
	repo := Repository{
		Type: PackagistRepository,
		URL:  url,
	}
	return c.AddRepository("packagist.org", repo)
}

// DisablePackagistRepository 禁用 Packagist.org 仓库
func (c *Composer) DisablePackagistRepository() error {
	args := []string{"config", "repositories.packagist.org.url", "false"}
	_, err := c.Run(args...)
	return err
}

// EnablePackagistRepository 启用 Packagist.org 仓库
func (c *Composer) EnablePackagistRepository() error {
	args := []string{"config", "repositories.packagist.org.url", "https://repo.packagist.org"}
	_, err := c.Run(args...)
	return err
}

// AddVcsRepository 添加版本控制系统仓库
func (c *Composer) AddVcsRepository(name string, url string) error {
	repo := Repository{
		Type: VcsRepository,
		URL:  url,
	}
	return c.AddRepository(name, repo)
}

// AddPathRepository 添加本地路径仓库
func (c *Composer) AddPathRepository(name string, path string, options map[string]interface{}) error {
	repo := Repository{
		Type:    PathRepository,
		URL:     path,
		Options: options,
	}
	return c.AddRepository(name, repo)
}

// AddComposerRepository 添加 Composer 仓库
func (c *Composer) AddComposerRepository(name string, url string) error {
	repo := Repository{
		Type: ComposerRepository,
		URL:  url,
	}
	return c.AddRepository(name, repo)
}

// GetPreferredInstall 获取 preferred-install 配置
func (c *Composer) GetPreferredInstall() (string, error) {
	return c.Run("config", "preferred-install")
}

// SetPreferredInstall 设置 preferred-install 配置
func (c *Composer) SetPreferredInstall(value string) error {
	if value != "dist" && value != "source" && value != "auto" {
		return fmt.Errorf("invalid preferred-install value: %s, must be 'dist', 'source' or 'auto'", value)
	}

	args := []string{"config", "preferred-install", value}
	_, err := c.Run(args...)
	return err
}

// SetMinimumStability 设置最小稳定性配置
func (c *Composer) SetMinimumStability(stability string) error {
	args := []string{"config", "minimum-stability", stability}
	_, err := c.Run(args...)
	return err
}

// GetMinimumStability 获取最小稳定性配置
func (c *Composer) GetMinimumStability() (string, error) {
	return c.Run("config", "minimum-stability")
}

// SetPreferStable 设置是否优先使用稳定版本
func (c *Composer) SetPreferStable(preferStable bool) error {
	value := "0"
	if preferStable {
		value = "1"
	}

	args := []string{"config", "prefer-stable", value}
	_, err := c.Run(args...)
	return err
}
