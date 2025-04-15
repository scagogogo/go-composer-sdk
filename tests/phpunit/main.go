package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func runTests() {
	fmt.Println("开始运行单元测试...")

	// 查找已安装的Composer
	composerPath, err := exec.LookPath("composer")
	if err != nil {
		log.Fatalf("无法找到composer: %v", err)
	}
	fmt.Printf("找到composer: %s\n", composerPath)

	// 创建Composer实例
	composerOptions := composer.Options{
		ExecutablePath: composerPath,
		AutoInstall:    false,
		DefaultTimeout: 30 * time.Second, // 设置超时时间
	}

	comp, err := composer.New(composerOptions)
	if err != nil {
		log.Fatalf("创建Composer实例失败: %v", err)
	}

	// 获取Composer版本
	version, err := comp.GetVersion()
	if err != nil {
		log.Fatalf("获取Composer版本失败: %v", err)
	}
	fmt.Printf("Composer版本: %s\n", version)

	// 创建一个临时测试项目
	tempDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Printf("创建临时项目目录: %s\n", tempDir)

	// 初始化composer.json
	composerJson := `{
		"name": "test/composer-unit-test",
		"description": "Test project for Composer unit tests",
		"type": "project",
		"license": "MIT",
		"require": {},
		"require-dev": {
			"phpunit/phpunit": "^9.0"
		}
	}`

	err = os.WriteFile(filepath.Join(tempDir, "composer.json"), []byte(composerJson), 0644)
	if err != nil {
		log.Fatalf("创建composer.json失败: %v", err)
	}
	fmt.Println("已创建composer.json")

	// 设置composer工作目录
	comp.SetWorkingDir(tempDir)

	// 安装依赖
	fmt.Println("正在安装依赖...")
	err = comp.Install(false, false)
	if err != nil {
		fmt.Printf("安装依赖失败: %v\n", err)
		fmt.Println("尝试直接使用composer命令...")

		cmd := exec.Command(composerPath, "install")
		cmd.Dir = tempDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("使用命令行安装依赖失败: %v\n%s", err, output)
		}
		fmt.Printf("依赖安装输出: %s\n", output)
	} else {
		fmt.Println("依赖安装成功")
	}

	// 创建一个简单的PHP测试文件
	testFile := `<?php
class ExampleTest extends \PHPUnit\Framework\TestCase {
	public function testAddition() {
		$this->assertEquals(4, 2 + 2);
	}
	
	public function testSubtraction() {
		$this->assertEquals(1, 3 - 2);
	}
}
`
	err = os.WriteFile(filepath.Join(tempDir, "ExampleTest.php"), []byte(testFile), 0644)
	if err != nil {
		log.Fatalf("创建测试文件失败: %v", err)
	}
	fmt.Println("已创建测试文件")

	// 尝试运行phpunit测试
	fmt.Println("正在运行测试...")
	cmd := exec.Command(filepath.Join(tempDir, "vendor", "bin", "phpunit"), "ExampleTest.php")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("测试运行失败: %v\n%s\n", err, output)

		// 尝试通过composer运行测试
		fmt.Println("尝试通过composer运行测试...")

		// 添加phpunit脚本到composer.json
		composerJsonWithScripts := `{
			"name": "test/composer-unit-test",
			"description": "Test project for Composer unit tests",
			"type": "project",
			"license": "MIT",
			"require": {},
			"require-dev": {
				"phpunit/phpunit": "^9.0"
			},
			"scripts": {
				"test": "phpunit ExampleTest.php"
			}
		}`

		err = os.WriteFile(filepath.Join(tempDir, "composer.json"), []byte(composerJsonWithScripts), 0644)
		if err != nil {
			log.Fatalf("更新composer.json失败: %v", err)
		}

		// 运行测试脚本
		scriptOutput, err := comp.RunScript("test")
		if err != nil {
			log.Fatalf("运行测试脚本失败: %v", err)
		}
		fmt.Printf("测试输出: %s\n", scriptOutput)
	} else {
		fmt.Printf("测试运行成功:\n%s\n", output)
	}

	fmt.Println("单元测试完成！")
}

// 仅当直接运行此文件时，才执行测试
func main() {
	if os.Getenv("RUN_COMPOSER_TESTS") == "1" {
		runTests()
	} else {
		fmt.Println("设置环境变量 RUN_COMPOSER_TESTS=1 来运行此测试")
	}
}
