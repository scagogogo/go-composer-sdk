// Package installer provides the functionality to install and manage PHP Composer.
package installer

import (
	"errors"
)

var (
	// ErrInstallationFailed 表示安装过程失败
	ErrInstallationFailed = errors.New("安装失败")
	// ErrInsufficientRights 表示权限不足
	ErrInsufficientRights = errors.New("权限不足，请使用管理员/sudo权限")
	// ErrUnsupportedPlatform 表示不支持的平台
	ErrUnsupportedPlatform = errors.New("不支持的操作系统平台")
	// ErrDownloadFailed 表示下载失败
	ErrDownloadFailed = errors.New("下载失败")
)

// Installer 负责安装Composer
type Installer struct {
	config Config
}

// NewInstaller 创建一个新的安装器实例
func NewInstaller(config Config) *Installer {
	return &Installer{
		config: config,
	}
}

// DefaultInstaller 创建使用默认配置的安装器实例
func DefaultInstaller() *Installer {
	return &Installer{
		config: DefaultConfig(),
	}
}

// GetConfig 获取安装器配置
func (i *Installer) GetConfig() Config {
	return i.config
}

// SetConfig 设置安装器配置
func (i *Installer) SetConfig(config Config) {
	i.config = config
}

// Install 安装Composer
func (i *Installer) Install() error {
	// 获取平台特定的安装器
	platformInstaller, err := GetPlatformInstaller(i.config)
	if err != nil {
		return err
	}

	// 执行安装
	return platformInstaller.Install()
}
