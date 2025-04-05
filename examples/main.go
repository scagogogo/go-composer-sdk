package main

import (
	"fmt"
	"os"

	basic_usage "github.com/scagogogo/go-composer-sdk/examples/01_basic_usage"
	package_management "github.com/scagogogo/go-composer-sdk/examples/02_package_management"
	project_management "github.com/scagogogo/go-composer-sdk/examples/03_project_management"
)

func main() {
	examples := []struct {
		name        string
		description string
		run         func()
	}{
		{"1.1", "基础用法 - 创建Composer实例", basic_usage.Example01NewComposer},
		{"1.2", "基础用法 - 运行命令", basic_usage.Example02RunCommands},
		{"2.1", "包管理 - 安装和更新包", package_management.Example01InstallUpdate},
		{"2.2", "包管理 - 添加和移除包", package_management.Example02RequireRemove},
		{"2.3", "包管理 - 查看包信息", package_management.Example03ShowPackage},
		{"2.4", "包管理 - 搜索包", package_management.Example04SearchPackage},
		{"3.1", "项目管理 - 创建项目", project_management.Example01CreateProject},
		{"3.2", "项目管理 - 运行脚本", project_management.Example02RunScript},
		{"3.3", "项目管理 - 平台需求检查", project_management.Example03PlatformCheck},
		{"3.4", "项目管理 - 依赖分析", project_management.Example04DependencyAnalysis},
		{"3.5", "项目管理 - 完整性检查和诊断", project_management.Example05IntegrityCheck},
	}

	if len(os.Args) < 2 {
		fmt.Println("请指定要运行的示例编号:")
		for _, example := range examples {
			fmt.Printf("  %s: %s\n", example.name, example.description)
		}
		return
	}

	exampleID := os.Args[1]
	for _, example := range examples {
		if example.name == exampleID {
			fmt.Printf("运行示例 %s: %s\n", example.name, example.description)
			fmt.Println("========================================")
			example.run()
			return
		}
	}

	fmt.Printf("未找到示例编号: %s\n", exampleID)
}
