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

// testComposerInstall 测试安装Composer
func testComposerInstall() {
	fmt.Println("开始测试 Composer 安装功能...")

	// 创建一个临时目录用于安装
	tempDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Printf("已创建临时安装目录: %s\n", tempDir)

	// 检查PHP是否可用
	phpPath, err := exec.LookPath("php")
	if err != nil {
		log.Fatalf("PHP 命令不可用: %v", err)
	}
	fmt.Printf("检测到PHP命令可用: %s\n", phpPath)

	// 检查PHP版本
	cmd := exec.Command(phpPath, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("获取PHP版本失败: %v", err)
	}
	fmt.Printf("PHP版本信息: %s\n", output)

	// 创建自定义安装配置
	config := installer.DefaultConfig()
	config.InstallPath = tempDir
	config.UseSudo = false         // 不使用 sudo，因为我们安装到临时目录
	config.PreferBrewOnMac = false // 不使用Homebrew，强制使用php脚本安装
	config.TimeoutSeconds = 60     // 设置较短的超时时间

	// 创建安装器
	inst := installer.NewInstaller(config)
	fmt.Println("正在安装 Composer...")

	// 执行安装
	err = inst.Install()
	if err != nil {
		log.Fatalf("安装 Composer 失败: %v", err)
	}

	fmt.Println("Composer 安装成功！")

	// 检查安装的文件
	composerPath := filepath.Join(tempDir, "composer")
	composerPharPath := filepath.Join(tempDir, "composer.phar")

	// 检查文件是否存在
	if _, err := os.Stat(composerPharPath); os.IsNotExist(err) {
		fmt.Printf("警告: composer.phar 文件不存在: %s\n", composerPharPath)

		// 尝试查找系统中已安装的composer
		systemComposer, err := exec.LookPath("composer")
		if err != nil {
			log.Fatalf("无法找到系统中的composer: %v", err)
		}
		fmt.Printf("使用系统安装的composer: %s\n", systemComposer)

		// 使用系统composer
		composerOptions := composer.Options{
			ExecutablePath: systemComposer,
			AutoInstall:    false,
			DefaultTimeout: 30 * time.Second, // 设置更短的超时时间
		}

		comp, err := composer.New(composerOptions)
		if err != nil {
			log.Fatalf("创建 Composer 实例失败: %v", err)
		}

		// 获取 Composer 版本
		fmt.Println("正在检查 Composer 版本...")
		version, err := comp.GetVersion()
		if err != nil {
			log.Fatalf("获取 Composer 版本失败: %v", err)
		}
		fmt.Printf("Composer 版本: %s\n", version)

		// 运行一些基本的 Composer 命令
		fmt.Println("正在测试 Composer 命令...")
		output, err := comp.Run("--version")
		if err != nil {
			log.Fatalf("执行 composer --version 失败: %v", err)
		}
		fmt.Printf("命令输出: %s\n", output)

		fmt.Println("Composer 安装和测试完成！")
		return
	}

	// 如果composer.phar存在
	fmt.Printf("composer.phar 文件已创建: %s\n", composerPharPath)

	if _, err := os.Stat(composerPath); os.IsNotExist(err) {
		fmt.Printf("警告: composer 执行脚本不存在: %s\n", composerPath)
	} else {
		fmt.Printf("composer 执行脚本已创建: %s\n", composerPath)

		// 检查脚本内容
		scriptContent, err := os.ReadFile(composerPath)
		if err != nil {
			fmt.Printf("读取脚本内容失败: %v\n", err)
		} else {
			fmt.Printf("脚本内容: %s\n", scriptContent)
		}

		// 检查执行权限
		info, err := os.Stat(composerPath)
		if err != nil {
			fmt.Printf("获取脚本权限失败: %v\n", err)
		} else {
			fmt.Printf("脚本权限: %v\n", info.Mode())
		}
	}

	// 直接使用php执行composer.phar
	fmt.Println("尝试直接使用PHP执行composer.phar...")
	phpCmd := exec.Command(phpPath, composerPharPath, "--version")
	phpOutput, err := phpCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("使用PHP直接执行composer.phar失败: %v, 输出: %s\n", err, phpOutput)
	} else {
		fmt.Printf("PHP执行composer.phar成功: %s\n", phpOutput)
	}

	// 创建 Composer 实例，设置更短的超时时间
	composerOptions := composer.Options{
		ExecutablePath: composerPharPath,
		AutoInstall:    false,
		DefaultTimeout: 30 * time.Second, // 设置更短的超时时间
	}

	comp, err := composer.New(composerOptions)
	if err != nil {
		log.Fatalf("创建 Composer 实例失败: %v", err)
	}

	// 获取 Composer 版本
	fmt.Println("正在检查 Composer 版本...")
	version, err := comp.GetVersion()
	if err != nil {
		fmt.Printf("获取 Composer 版本失败: %v\n", err)
		fmt.Println("尝试使用composer脚本代替...")

		// 尝试使用composer脚本
		composerOptions.ExecutablePath = composerPath
		comp, err = composer.New(composerOptions)
		if err != nil {
			log.Fatalf("使用composer脚本创建实例失败: %v", err)
		}

		version, err = comp.GetVersion()
		if err != nil {
			log.Fatalf("使用composer脚本获取版本失败: %v", err)
		}
	}
	fmt.Printf("Composer 版本: %s\n", version)

	// 运行一些基本的 Composer 命令
	fmt.Println("正在测试 Composer 命令...")
	cmdOutput, err := comp.Run("--version")
	if err != nil {
		log.Fatalf("执行 composer --version 失败: %v", err)
	}
	fmt.Printf("命令输出: %s\n", cmdOutput)

	fmt.Println("Composer 安装和测试完成！")
}

// 添加main函数调用测试函数
func main() {
	if os.Getenv("RUN_INSTALL_TEST") == "1" {
		testComposerInstall()
	} else {
		fmt.Println("设置环境变量 RUN_INSTALL_TEST=1 来运行安装测试")
	}
}
