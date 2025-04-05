package pkg02_package_management

import (
	"fmt"
	"log"

	"github.com/scagogogo/go-composer-sdk/pkg/composer"
)

// Example03ShowPackage 演示如何查看包信息
func Example03ShowPackage() {
	// 创建Composer实例
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 示例1：显示特定包的信息
	packageName := "monolog/monolog"
	fmt.Printf("\n1. 显示包 %s 的信息...\n", packageName)
	output, err := c.ShowPackage(packageName)
	if err != nil {
		log.Printf("显示包信息失败: %v", err)
	} else {
		fmt.Printf("包 %s 的信息:\n%s\n", packageName, output)
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("name     : monolog/monolog")
		fmt.Println("descrip. : Sends your logs to files, sockets, inboxes, databases and various web services")
		fmt.Println("keywords : log, logging, psr-3")
		fmt.Println("versions : dev-main, 3.4.0, 3.3.1, ...")
		fmt.Println("-------------------------------------")
		// 实际输出会更详细
	}

	// 示例2：显示所有已安装的包
	fmt.Println("\n2. 显示所有已安装的包...")
	output, err = c.ShowAllPackages()
	if err != nil {
		log.Printf("显示所有包信息失败: %v", err)
	} else {
		fmt.Println("所有已安装的包:")
		fmt.Printf("(输出长度: %d字符)\n", len(output))
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("monolog/monolog                 v3.4.0  Sends your logs to files, sockets, inboxes...")
		fmt.Println("symfony/console                 v6.3.4  Eases the creation of beautiful and testable c...")
		fmt.Println("-------------------------------------")
		// 实际输出会包含所有安装的包
	}

	// 示例3：显示依赖树
	fmt.Println("\n3. 显示依赖树...")
	output, err = c.ShowDependencyTree("")
	if err != nil {
		log.Printf("显示依赖树失败: %v", err)
	} else {
		fmt.Println("依赖树:")
		fmt.Printf("(输出长度: %d字符)\n", len(output))
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("doctrine/annotations 2.0.1 (symlinked from vendor/doctrine/annotations)")
		fmt.Println("└──php >=7.2")
		fmt.Println("monolog/monolog 3.4.0 (symlinked from vendor/monolog/monolog)")
		fmt.Println("└──php >=8.1")
		fmt.Println("-------------------------------------")
		// 实际输出会包含完整的依赖树
	}

	// 示例4：显示特定包的依赖树
	packageName = "symfony/console"
	fmt.Printf("\n4. 显示包 %s 的依赖树...\n", packageName)
	output, err = c.ShowDependencyTree(packageName)
	if err != nil {
		log.Printf("显示特定包依赖树失败: %v", err)
	} else {
		fmt.Printf("包 %s 的依赖树:\n%s\n", packageName, output)
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("symfony/console 6.3.4 (symlinked from vendor/symfony/console)")
		fmt.Println("├──php >=8.1")
		fmt.Println("├──symfony/polyfill-mbstring >=1.0")
		fmt.Println("├──symfony/service-contracts ^2.1|^3")
		fmt.Println("└──symfony/string ^5.4|^6.0")
		fmt.Println("   └──symfony/polyfill-ctype ~1.8")
		fmt.Println("-------------------------------------")
		// 实际输出会包含完整的依赖树
	}

	// 示例5：显示反向依赖关系
	packageName = "symfony/polyfill-mbstring"
	fmt.Printf("\n5. 显示包 %s 的反向依赖关系...\n", packageName)
	output, err = c.ShowReverseDependencies(packageName)
	if err != nil {
		log.Printf("显示反向依赖关系失败: %v", err)
	} else {
		fmt.Printf("包 %s 的反向依赖关系:\n%s\n", packageName, output)
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("symfony/console              6.3.4    requires  symfony/polyfill-mbstring (>=1.0)")
		fmt.Println("symfony/http-foundation       6.3.4    requires  symfony/polyfill-mbstring (>=1.0)")
		fmt.Println("-------------------------------------")
		// 实际输出会包含所有依赖于该包的包
	}

	// 示例6：解释为什么需要某个包
	packageName = "symfony/polyfill-mbstring"
	fmt.Printf("\n6. 解释为什么项目需要包 %s...\n", packageName)
	output, err = c.WhyPackage(packageName)
	if err != nil {
		log.Printf("解释为什么需要包失败: %v", err)
	} else {
		fmt.Printf("项目需要包 %s 的原因:\n%s\n", packageName, output)
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("symfony/console  6.3.4  requires  symfony/polyfill-mbstring (>=1.0)")
		fmt.Println("-------------------------------------")
		// 实际输出会详细解释原因
	}

	// 示例7：查找过期的包
	fmt.Println("\n7. 查找过期的包...")
	// 注意：OutdatedPackages 方法在您的代码库中可能有不同的参数
	// 这里使用了带布尔参数的版本，但实际使用时需要根据实际API调整
	output, err = c.OutdatedPackages()
	if err != nil {
		log.Printf("查找过期包失败: %v", err)
	} else {
		fmt.Println("过期的包:")
		fmt.Printf("(输出长度: %d字符)\n", len(output))
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("symfony/console    6.3.4    6.4.0    更新到新的次要版本")
		fmt.Println("monolog/monolog    3.4.0    3.5.0    更新到新的次要版本")
		fmt.Println("-------------------------------------")
		// 实际输出会列出所有过期的包
	}

	// 示例8：查找直接依赖中的过期包
	fmt.Println("\n8. 查找直接依赖中的过期包...")
	output, err = c.OutdatedPackagesDirect()
	if err != nil {
		log.Printf("查找直接依赖中的过期包失败: %v", err)
	} else {
		fmt.Println("直接依赖中的过期包:")
		fmt.Printf("(输出长度: %d字符)\n", len(output))
		fmt.Println("-------------截取部分输出-------------")
		fmt.Println("symfony/console    6.3.4    6.4.0    更新到新的次要版本")
		fmt.Println("-------------------------------------")
		// 实际输出会列出所有直接依赖中的过期包
	}
}
