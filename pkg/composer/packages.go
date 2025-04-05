package composer

// RequirePackage 添加一个新的包依赖
//
// 参数：
//   - packageName: 要添加的包名，例如"symfony/console"
//   - version: 版本约束，例如"^5.0"，如果为空则使用最新版本
//   - dev: 是否作为开发依赖添加
//
// 返回值：
//   - error: 如果添加包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向项目添加一个新的依赖包。相当于执行
//	`composer require [--dev] package/name:version`
//
// 用法示例：
//
//	// 添加生产依赖
//	err := comp.RequirePackage("symfony/console", "^5.0", false)
//	if err != nil {
//	    log.Fatalf("添加依赖失败: %v", err)
//	}
//
//	// 添加开发依赖
//	err = comp.RequirePackage("phpunit/phpunit", "^9.0", true)
func (c *Composer) RequirePackage(packageName string, version string, dev bool) error {
	args := []string{"require"}

	if dev {
		args = append(args, "--dev")
	}

	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// Remove 移除包
//
// 参数：
//   - packageName: 要移除的包名
//   - dev: 是否从开发依赖中移除
//
// 返回值：
//   - error: 如果移除包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法从项目中移除指定的依赖包。相当于执行
//	`composer remove [--dev] package/name`
//
// 用法示例：
//
//	// 移除生产依赖
//	err := comp.Remove("symfony/console", false)
//	if err != nil {
//	    log.Fatalf("移除依赖失败: %v", err)
//	}
//
//	// 移除开发依赖
//	err = comp.Remove("phpunit/phpunit", true)
func (c *Composer) Remove(packageName string, dev bool) error {
	args := []string{"remove"}

	if dev {
		args = append(args, "--dev")
	}

	args = append(args, packageName)

	_, err := c.Run(args...)
	return err
}

// ShowPackage 显示包的详细信息
//
// 参数：
//   - packageName: 要显示信息的包名
//
// 返回值：
//   - string: 包信息的输出结果
//   - error: 如果获取包信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示指定包的详细信息，包括版本、依赖、安装位置等。
//	相当于执行`composer show package/name`
//
// 用法示例：
//
//	output, err := comp.ShowPackage("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取包信息失败: %v", err)
//	}
//	fmt.Println("包信息:", output)
func (c *Composer) ShowPackage(packageName string) (string, error) {
	return c.Run("show", packageName)
}

// Search 搜索包
//
// 参数：
//   - query: 搜索关键词
//
// 返回值：
//   - string: 搜索结果的输出
//   - error: 如果搜索过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法在Packagist上搜索符合关键词的包。
//	相当于执行`composer search query`
//
// 用法示例：
//
//	output, err := comp.Search("logger")
//	if err != nil {
//	    log.Fatalf("搜索包失败: %v", err)
//	}
//	fmt.Println("搜索结果:", output)
func (c *Composer) Search(query string) (string, error) {
	return c.Run("search", query)
}

// ShowAllPackages 显示所有已安装的包
//
// 返回值：
//   - string: 已安装包列表的输出
//   - error: 如果获取包列表过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示项目中所有已安装的包。
//	相当于执行`composer show`
//
// 用法示例：
//
//	output, err := comp.ShowAllPackages()
//	if err != nil {
//	    log.Fatalf("获取包列表失败: %v", err)
//	}
//	fmt.Println("已安装的包:", output)
func (c *Composer) ShowAllPackages() (string, error) {
	return c.Run("show")
}

// ShowDependencyTree 显示依赖树
//
// 参数：
//   - packageName: 要显示依赖树的包名，如果为空则显示整个项目的依赖树
//
// 返回值：
//   - string: 依赖树的输出
//   - error: 如果获取依赖树过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法以树形结构显示包的依赖关系。
//	相当于执行`composer show --tree [package/name]`
//
// 用法示例：
//
//	// 显示整个项目的依赖树
//	output, err := comp.ShowDependencyTree("")
//	if err != nil {
//	    log.Fatalf("获取依赖树失败: %v", err)
//	}
//	fmt.Println("依赖树:", output)
//
//	// 显示特定包的依赖树
//	output, err = comp.ShowDependencyTree("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取依赖树失败: %v", err)
//	}
//	fmt.Println("依赖树:", output)
func (c *Composer) ShowDependencyTree(packageName string) (string, error) {
	if packageName != "" {
		return c.Run("show", "--tree", packageName)
	}
	return c.Run("show", "--tree")
}

// ShowReverseDependencies 显示哪些包依赖于指定包（depends命令）
//
// 参数：
//   - packageName: 要查询反向依赖的包名
//
// 返回值：
//   - string: 反向依赖的输出结果
//   - error: 如果获取反向依赖过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示哪些已安装的包依赖于指定的包。
//	相当于执行`composer depends package/name`
//
// 用法示例：
//
//	output, err := comp.ShowReverseDependencies("symfony/polyfill-mbstring")
//	if err != nil {
//	    log.Fatalf("获取反向依赖失败: %v", err)
//	}
//	fmt.Println("依赖于此包的其他包:", output)
func (c *Composer) ShowReverseDependencies(packageName string) (string, error) {
	return c.Run("depends", packageName)
}

// WhyPackage 解释为什么安装了指定的包（why命令）
//
// 参数：
//   - packageName: 要查询原因的包名
//
// 返回值：
//   - string: 安装原因的输出结果
//   - error: 如果查询安装原因过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法解释为什么项目中安装了指定的包，通常是因为它是某个直接依赖的间接依赖。
//	相当于执行`composer why package/name`
//
// 用法示例：
//
//	output, err := comp.WhyPackage("symfony/polyfill-mbstring")
//	if err != nil {
//	    log.Fatalf("查询安装原因失败: %v", err)
//	}
//	fmt.Println("安装原因:", output)
func (c *Composer) WhyPackage(packageName string) (string, error) {
	return c.Run("why", packageName)
}

// OutdatedPackages 显示所有过时的包
//
// 返回值：
//   - string: 过时包列表的输出
//   - error: 如果获取过时包列表过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示项目中所有过时的包及可用的更新。
//	相当于执行`composer outdated`
//
// 用法示例：
//
//	output, err := comp.OutdatedPackages()
//	if err != nil {
//	    log.Fatalf("获取过时包列表失败: %v", err)
//	}
//	fmt.Println("过时的包:", output)
func (c *Composer) OutdatedPackages() (string, error) {
	return c.Run("outdated")
}

// OutdatedPackagesDirect 只显示直接依赖中过时的包
//
// 返回值：
//   - string: 过时直接依赖的输出
//   - error: 如果获取过时直接依赖过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法只显示在composer.json中直接声明的依赖中过时的包，不包括间接依赖。
//	相当于执行`composer outdated --direct`
//
// 用法示例：
//
//	output, err := comp.OutdatedPackagesDirect()
//	if err != nil {
//	    log.Fatalf("获取过时直接依赖失败: %v", err)
//	}
//	fmt.Println("过时的直接依赖:", output)
func (c *Composer) OutdatedPackagesDirect() (string, error) {
	return c.Run("outdated", "--direct")
}

// RequirePackageWithOptions 使用更多选项添加包依赖
//
// 参数：
//   - packageName: 要添加的包名，例如"symfony/console"
//   - version: 版本约束，例如"^5.0"，如果为空则使用最新版本
//   - options: 添加依赖时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果添加包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向项目添加一个新的依赖包，支持更多自定义选项。
//
// 用法示例：
//
//	// 添加依赖并指定多个选项
//	options := map[string]string{
//	    "dev": "",
//	    "prefer-source": "",
//	    "no-update": "",
//	}
//	err := comp.RequirePackageWithOptions("symfony/console", "^5.0", options)
//	if err != nil {
//	    log.Fatalf("添加依赖失败: %v", err)
//	}
func (c *Composer) RequirePackageWithOptions(packageName string, version string, options map[string]string) error {
	args := []string{"require"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// BumpPackages 升级指定的包至最新兼容版本
//
// 参数：
//   - packages: 要升级的包名列表
//
// 返回值：
//   - error: 如果升级包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法将指定的包升级到符合其在composer.json中版本约束的最新版本。
//	相当于执行`composer bump package1 package2 ...`
//
// 用法示例：
//
//	// 升级多个包
//	err := comp.BumpPackages([]string{"symfony/console", "symfony/process"})
//	if err != nil {
//	    log.Fatalf("升级包失败: %v", err)
//	}
//
//	// 如果packages为空切片，则会升级所有包
//	err = comp.BumpPackages([]string{})
//	if err != nil {
//	    log.Fatalf("升级所有包失败: %v", err)
//	}
func (c *Composer) BumpPackages(packages []string) error {
	args := []string{"bump"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// BumpPackagesWithOptions 使用更多选项升级包
//
// 参数：
//   - packages: 要升级的包名列表
//   - options: 升级包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果升级包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法将指定的包升级到符合其在composer.json中版本约束的最新版本，支持更多自定义选项。
//
// 用法示例：
//
//	// 升级包并指定多个选项
//	options := map[string]string{
//	    "dev-only": "",
//	    "prefer-stable": "",
//	    "dry-run": "",
//	}
//	err := comp.BumpPackagesWithOptions([]string{"symfony/console"}, options)
//	if err != nil {
//	    log.Fatalf("升级包失败: %v", err)
//	}
func (c *Composer) BumpPackagesWithOptions(packages []string, options map[string]string) error {
	args := []string{"bump"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// Reinstall 重新安装指定的包
//
// 参数：
//   - packageName: 要重新安装的包名
//
// 返回值：
//   - error: 如果重新安装包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法重新安装指定的包，先移除然后再安装。
//	相当于执行`composer remove packageName && composer require packageName`
//
// 用法示例：
//
//	err := comp.Reinstall("symfony/console")
//	if err != nil {
//	    log.Fatalf("重新安装包失败: %v", err)
//	}
func (c *Composer) Reinstall(packageName string) error {
	// 先移除包
	err := c.Remove(packageName, false)
	if err != nil {
		return err
	}

	// 然后重新安装
	return c.RequirePackage(packageName, "", false)
}

// BrowsePackage 打开包的项目页面
//
// 参数：
//   - packageName: 要浏览的包名
//
// 返回值：
//   - error: 如果打开包页面过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用默认浏览器打开指定包的项目页面（通常是其GitHub仓库）。
//	相当于执行`composer browse packageName`
//	注意：此功能需要操作系统支持打开浏览器，并且包必须在composer.json中声明了项目URL。
//
// 用法示例：
//
//	err := comp.BrowsePackage("symfony/console")
//	if err != nil {
//	    log.Fatalf("打开包页面失败: %v", err)
//	}
func (c *Composer) BrowsePackage(packageName string) error {
	_, err := c.Run("browse", packageName)
	return err
}

// BrowsePackageWithOptions 使用更多选项打开包的项目页面
//
// 参数：
//   - packageName: 要浏览的包名
//   - options: 浏览包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果打开包页面过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用默认浏览器打开指定包的项目页面，支持更多自定义选项。
//	例如可以指定打开文档页面而非主页（使用--homepage或--docs选项）。
//
// 用法示例：
//
//	// 打开包的文档页面
//	options := map[string]string{
//	    "docs": "",
//	}
//	err := comp.BrowsePackageWithOptions("symfony/console", options)
//	if err != nil {
//	    log.Fatalf("打开包文档页面失败: %v", err)
//	}
//
//	// 打开包的行为页面
//	options = map[string]string{
//	    "issues": "",
//	}
//	err = comp.BrowsePackageWithOptions("symfony/console", options)
//	if err != nil {
//	    log.Fatalf("打开包问题跟踪页面失败: %v", err)
//	}
func (c *Composer) BrowsePackageWithOptions(packageName string, options map[string]string) error {
	args := []string{"browse", packageName}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	_, err := c.Run(args...)
	return err
}

// WhyNotPackage 解释为什么不能安装指定版本的包
//
// 参数：
//   - packageName: 要查询的包名
//   - version: 要查询的版本
//
// 返回值：
//   - string: 解释结果的输出
//   - error: 如果查询过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法解释为什么不能安装指定版本的包，通常是因为版本冲突或其他约束问题。
//	相当于执行`composer why-not packageName version`
//
// 用法示例：
//
//	output, err := comp.WhyNotPackage("symfony/console", "v4.0.0")
//	if err != nil {
//	    log.Fatalf("查询版本约束失败: %v", err)
//	}
//	fmt.Println("无法安装的原因:", output)
func (c *Composer) WhyNotPackage(packageName string, version string) (string, error) {
	return c.Run("why-not", packageName, version)
}
