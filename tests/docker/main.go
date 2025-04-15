package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/scagogogo/go-composer-sdk/pkg/composer"
	"github.com/scagogogo/go-composer-sdk/pkg/installer"
)

func main() {
	fmt.Println("===== 开始在Linux环境中测试Composer自动安装 =====")

	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "composer-linux-test-*")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Printf("已创建临时目录: %s\n", tempDir)

	// 测试PHP是否正常安装
	phpPath, err := exec.LookPath("php")
	if err != nil {
		log.Fatalf("未找到PHP: %v", err)
	}
	fmt.Printf("PHP路径: %s\n", phpPath)

	// 获取PHP版本
	cmd := exec.Command(phpPath, "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("获取PHP版本失败: %v", err)
	}
	fmt.Printf("PHP版本: %s\n", string(output))

	// 检查Composer是否已安装
	_, err = exec.LookPath("composer")
	if err == nil {
		fmt.Println("警告: Composer已在系统中安装，将移除以确保测试环境干净")
		removeCmd := exec.Command("rm", "-f", "/usr/local/bin/composer", "/usr/bin/composer")
		if err := removeCmd.Run(); err != nil {
			fmt.Printf("移除现有Composer失败: %v，但将继续测试\n", err)
		}
	}

	fmt.Println("\n===== 测试1: 使用SDK自动安装Composer =====")

	// 创建一个未指定Composer路径的实例
	comp, err := composer.New(composer.Options{
		AutoInstall:    true, // 启用自动安装
		DefaultTimeout: 3 * time.Minute,
	})
	if err != nil {
		log.Fatalf("创建Composer实例失败: %v", err)
	}

	// 通过运行版本命令来触发自动安装
	version, err := comp.GetVersion()
	if err != nil {
		log.Fatalf("获取Composer版本失败: %v", err)
	}
	fmt.Printf("自动安装成功! Composer版本: %s\n", version)

	// 获取安装路径
	path, err := exec.LookPath("composer")
	if err != nil {
		log.Fatalf("安装后未找到Composer: %v", err)
	}
	fmt.Printf("Composer安装路径: %s\n", path)

	fmt.Println("\n===== 测试2: 使用自定义安装目录 =====")

	// 创建自定义安装目录
	customDir := filepath.Join(tempDir, "custom-composer")
	if err := os.MkdirAll(customDir, 0755); err != nil {
		log.Fatalf("创建自定义目录失败: %v", err)
	}

	// 配置自定义安装选项
	installConfig := installer.DefaultConfig()
	installConfig.InstallPath = customDir
	installConfig.UseSudo = false

	inst := installer.NewInstaller(installConfig)

	// 执行安装
	fmt.Println("开始自定义安装...")
	if err := inst.Install(); err != nil {
		log.Fatalf("自定义安装失败: %v", err)
	}

	customPharPath := filepath.Join(customDir, "composer.phar")
	customScriptPath := filepath.Join(customDir, "composer")

	// 检查安装的文件
	if _, err := os.Stat(customPharPath); os.IsNotExist(err) {
		log.Fatalf("自定义安装未创建composer.phar文件: %v", err)
	}
	if _, err := os.Stat(customScriptPath); os.IsNotExist(err) {
		log.Fatalf("自定义安装未创建composer脚本: %v", err)
	}

	fmt.Printf("自定义安装成功! 文件路径:\n- composer.phar: %s\n- composer脚本: %s\n",
		customPharPath, customScriptPath)

	// 使用自定义安装路径创建Composer实例
	customComp, err := composer.New(composer.Options{
		ExecutablePath: customScriptPath,
		DefaultTimeout: 2 * time.Minute,
	})
	if err != nil {
		log.Fatalf("创建自定义Composer实例失败: %v", err)
	}

	// 测试自定义安装的Composer
	customVersion, err := customComp.GetVersion()
	if err != nil {
		log.Fatalf("获取自定义Composer版本失败: %v", err)
	}
	fmt.Printf("自定义Composer版本: %s\n", customVersion)

	fmt.Println("\n===== 测试3: 创建简单PHP项目 =====")

	// 创建项目目录
	projectDir := filepath.Join(tempDir, "test-project")
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		log.Fatalf("创建项目目录失败: %v", err)
	}

	// 创建composer.json
	composerJson := `{
		"name": "test/linux-composer-project",
		"description": "Linux Composer测试项目",
		"type": "project",
		"require": {
			"monolog/monolog": "^2.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(projectDir, "composer.json"), []byte(composerJson), 0644); err != nil {
		log.Fatalf("创建composer.json失败: %v", err)
	}

	// 切换到项目目录
	customComp.SetWorkingDir(projectDir)

	// 安装依赖
	fmt.Println("安装项目依赖...")
	err = customComp.Install(false, false)
	if err != nil {
		log.Fatalf("安装依赖失败: %v", err)
	}

	// 检查vendor目录
	vendorDir := filepath.Join(projectDir, "vendor")
	if _, err := os.Stat(vendorDir); os.IsNotExist(err) {
		log.Fatalf("vendor目录未创建: %v", err)
	}

	monologDir := filepath.Join(vendorDir, "monolog", "monolog")
	if _, err := os.Stat(monologDir); os.IsNotExist(err) {
		log.Fatalf("monolog库未安装: %v", err)
	}

	fmt.Println("项目依赖安装成功!")

	fmt.Println("\n===== 所有测试通过! =====")
}
