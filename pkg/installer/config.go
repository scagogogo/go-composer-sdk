package installer

import (
	"os"
	"path/filepath"
	"runtime"
)

// Config 保存安装器的配置
type Config struct {
	// 下载URL
	DownloadURL string
	// 安装路径
	InstallPath string
	// 是否使用代理
	UseProxy bool
	// 代理地址
	ProxyURL string
	// 超时时间（秒）
	TimeoutSeconds int
	// 是否使用sudo/管理员权限（类Unix系统）
	UseSudo bool
	// 在Mac系统上是否优先使用Homebrew安装
	PreferBrewOnMac bool
}

// DefaultConfig 返回适合当前操作系统的默认配置
func DefaultConfig() Config {
	config := Config{
		DownloadURL:     "https://getcomposer.org/installer",
		TimeoutSeconds:  300,
		UseProxy:        false,
		PreferBrewOnMac: true, // 默认优先使用brew安装
	}

	// 根据操作系统设置默认安装路径
	switch runtime.GOOS {
	case "windows":
		config.InstallPath = filepath.Join(os.Getenv("ProgramFiles"), "Composer")
	case "darwin", "linux":
		config.InstallPath = "/usr/local/bin"
	}

	return config
}
