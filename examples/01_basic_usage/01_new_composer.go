package pkg01_basic_usage

import (
	"fmt"
	"log"
	"time"

	"github.com/scagogogo/go-composer-sdk/pkg/composer"
)

// Example01NewComposer 演示如何创建和初始化Composer实例
func Example01NewComposer() {
	// 方法1：使用默认选项创建Composer实例
	// 这将自动检测Composer可执行文件路径，如果未找到将尝试安装
	defaultComposer, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建默认Composer实例: %v", err)
	}
	fmt.Printf("成功创建默认Composer实例，可执行文件路径: %s\n", defaultComposer.GetExecutablePath())
	// 输出示例：成功创建默认Composer实例，可执行文件路径: /usr/local/bin/composer

	// 方法2：使用自定义选项创建Composer实例
	customOptions := composer.Options{
		ExecutablePath: "/usr/local/bin/composer", // 指定Composer可执行文件路径
		WorkingDir:     "/path/to/project",        // 指定工作目录
		AutoInstall:    true,                      // 未找到Composer时自动安装
		DefaultTimeout: 5 * time.Minute,           // 设置命令执行超时时间
	}

	customComposer, err := composer.New(customOptions)
	if err != nil {
		log.Fatalf("无法创建自定义Composer实例: %v", err)
	}
	fmt.Printf("成功创建自定义Composer实例，工作目录: %s\n", "/path/to/project")
	// 输出示例：成功创建自定义Composer实例，工作目录: /path/to/project

	// 检查Composer是否已安装
	if customComposer.IsInstalled() {
		fmt.Println("Composer已安装")
		// 输出：Composer已安装
	} else {
		fmt.Println("Composer未安装")
	}

	// 获取Composer版本
	version, err := customComposer.GetVersion()
	if err != nil {
		log.Printf("获取版本信息失败: %v", err)
	} else {
		fmt.Printf("Composer版本: %s\n", version)
		// 输出示例：Composer版本: 2.5.7
	}
}
