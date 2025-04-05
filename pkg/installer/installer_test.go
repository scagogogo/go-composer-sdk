package installer

import (
	"testing"
)

func TestNewInstaller(t *testing.T) {
	config := Config{
		DownloadURL:     "http://example.com/composer.php",
		InstallPath:     "/test/path",
		UseProxy:        true,
		ProxyURL:        "http://proxy.example.com",
		TimeoutSeconds:  120,
		UseSudo:         true,
		PreferBrewOnMac: false,
	}

	installer := NewInstaller(config)

	// 检查配置是否正确保存
	if installer.config.DownloadURL != config.DownloadURL {
		t.Errorf("NewInstaller.config.DownloadURL = %v, 期望值 %v",
			installer.config.DownloadURL, config.DownloadURL)
	}
	if installer.config.InstallPath != config.InstallPath {
		t.Errorf("NewInstaller.config.InstallPath = %v, 期望值 %v",
			installer.config.InstallPath, config.InstallPath)
	}
	if installer.config.UseProxy != config.UseProxy {
		t.Errorf("NewInstaller.config.UseProxy = %v, 期望值 %v",
			installer.config.UseProxy, config.UseProxy)
	}
	if installer.config.ProxyURL != config.ProxyURL {
		t.Errorf("NewInstaller.config.ProxyURL = %v, 期望值 %v",
			installer.config.ProxyURL, config.ProxyURL)
	}
	if installer.config.TimeoutSeconds != config.TimeoutSeconds {
		t.Errorf("NewInstaller.config.TimeoutSeconds = %v, 期望值 %v",
			installer.config.TimeoutSeconds, config.TimeoutSeconds)
	}
	if installer.config.UseSudo != config.UseSudo {
		t.Errorf("NewInstaller.config.UseSudo = %v, 期望值 %v",
			installer.config.UseSudo, config.UseSudo)
	}
	if installer.config.PreferBrewOnMac != config.PreferBrewOnMac {
		t.Errorf("NewInstaller.config.PreferBrewOnMac = %v, 期望值 %v",
			installer.config.PreferBrewOnMac, config.PreferBrewOnMac)
	}
}

func TestDefaultInstaller(t *testing.T) {
	installer := DefaultInstaller()

	// 检查是否使用了默认配置
	defaultConfig := DefaultConfig()
	if installer.config.DownloadURL != defaultConfig.DownloadURL {
		t.Errorf("DefaultInstaller().config.DownloadURL = %v, 期望值 %v",
			installer.config.DownloadURL, defaultConfig.DownloadURL)
	}
}

func TestGetConfig(t *testing.T) {
	config := Config{
		DownloadURL:     "http://example.com/composer.php",
		InstallPath:     "/test/path",
		UseProxy:        true,
		ProxyURL:        "http://proxy.example.com",
		TimeoutSeconds:  120,
		UseSudo:         true,
		PreferBrewOnMac: false,
	}

	installer := NewInstaller(config)
	gotConfig := installer.GetConfig()

	// 检查获取的配置是否匹配
	if gotConfig.DownloadURL != config.DownloadURL {
		t.Errorf("GetConfig().DownloadURL = %v, 期望值 %v",
			gotConfig.DownloadURL, config.DownloadURL)
	}
	if gotConfig.InstallPath != config.InstallPath {
		t.Errorf("GetConfig().InstallPath = %v, 期望值 %v",
			gotConfig.InstallPath, config.InstallPath)
	}
}

func TestSetConfig(t *testing.T) {
	originalConfig := DefaultConfig()
	installer := NewInstaller(originalConfig)

	// 创建新配置
	newConfig := Config{
		DownloadURL:     "http://example.com/composer.php",
		InstallPath:     "/new/test/path",
		UseProxy:        true,
		ProxyURL:        "http://proxy.example.com",
		TimeoutSeconds:  120,
		UseSudo:         true,
		PreferBrewOnMac: false,
	}

	// 设置新配置
	installer.SetConfig(newConfig)
	gotConfig := installer.GetConfig()

	// 检查新配置是否生效
	if gotConfig.DownloadURL != newConfig.DownloadURL {
		t.Errorf("SetConfig后, GetConfig().DownloadURL = %v, 期望值 %v",
			gotConfig.DownloadURL, newConfig.DownloadURL)
	}
	if gotConfig.InstallPath != newConfig.InstallPath {
		t.Errorf("SetConfig后, GetConfig().InstallPath = %v, 期望值 %v",
			gotConfig.InstallPath, newConfig.InstallPath)
	}
}
