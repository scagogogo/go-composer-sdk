package composer

// Install 安装依赖项
//
// 参数：
//   - noDev: 如果为true，则不安装开发依赖
//   - optimize: 如果为true，则优化自动加载器
//
// 返回值：
//   - error: 如果安装依赖项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法安装项目的所有依赖，基于composer.json文件中定义的依赖项。
//	相当于执行`composer install [--no-dev] [--optimize-autoloader]`
//
// 用法示例：
//
//	// 安装所有依赖项（包括开发依赖）
//	err := comp.Install(false, false)
//	if err != nil {
//	    log.Fatalf("安装依赖失败: %v", err)
//	}
//
//	// 只安装生产依赖并优化自动加载
//	err = comp.Install(true, true)
//	if err != nil {
//	    log.Fatalf("安装依赖失败: %v", err)
//	}
func (c *Composer) Install(noDev bool, optimize bool) error {
	args := []string{"install"}

	if noDev {
		args = append(args, "--no-dev")
	}

	if optimize {
		args = append(args, "--optimize-autoloader")
	}

	_, err := c.Run(args...)
	return err
}

// InstallWithOptions 使用更多选项安装依赖项
//
// 参数：
//   - options: 安装选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果安装依赖项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法安装项目的依赖，支持更多自定义选项。
//
// 用法示例：
//
//	// 使用多个选项安装依赖
//	options := map[string]string{
//	    "no-dev": "",
//	    "optimize-autoloader": "",
//	    "prefer-dist": "",
//	    "no-progress": "",
//	}
//	err := comp.InstallWithOptions(options)
//	if err != nil {
//	    log.Fatalf("安装依赖失败: %v", err)
//	}
func (c *Composer) InstallWithOptions(options map[string]string) error {
	args := []string{"install"}

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

// Update 更新依赖项
//
// 参数：
//   - packages: 要更新的包名列表，为空则更新所有包
//   - noDev: 如果为true，则不更新开发依赖
//
// 返回值：
//   - error: 如果更新依赖项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法更新项目的依赖到最新版本，可以指定特定的包进行更新，或者更新所有包。
//	相当于执行`composer update [--no-dev] [package1 package2 ...]`
//
// 用法示例：
//
//	// 更新所有依赖（包括开发依赖）
//	err := comp.Update([]string{}, false)
//	if err != nil {
//	    log.Fatalf("更新依赖失败: %v", err)
//	}
//
//	// 只更新指定的包
//	err = comp.Update([]string{"symfony/console", "symfony/process"}, false)
//	if err != nil {
//	    log.Fatalf("更新依赖失败: %v", err)
//	}
//
//	// 只更新生产依赖
//	err = comp.Update([]string{}, true)
//	if err != nil {
//	    log.Fatalf("更新依赖失败: %v", err)
//	}
func (c *Composer) Update(packages []string, noDev bool) error {
	args := []string{"update"}

	if noDev {
		args = append(args, "--no-dev")
	}

	args = append(args, packages...)

	_, err := c.Run(args...)
	return err
}

// UpdateWithOptions 使用更多选项更新依赖项
//
// 参数：
//   - packages: 要更新的包名列表，为空则更新所有包
//   - options: 更新选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果更新依赖项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法更新项目的依赖，支持更多自定义选项。
//
// 用法示例：
//
//	// 使用多个选项更新依赖
//	options := map[string]string{
//	    "no-dev": "",
//	    "prefer-dist": "",
//	    "with-dependencies": "",
//	    "no-progress": "",
//	}
//	err := comp.UpdateWithOptions([]string{"symfony/console"}, options)
//	if err != nil {
//	    log.Fatalf("更新依赖失败: %v", err)
//	}
func (c *Composer) UpdateWithOptions(packages []string, options map[string]string) error {
	args := []string{"update"}

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

// DumpAutoload 生成自动加载文件
//
// 参数：
//   - optimize: 如果为true，则优化自动加载器，生成类映射
//
// 返回值：
//   - error: 如果生成自动加载文件过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法生成Composer的自动加载文件，可以选择是否优化。
//	相当于执行`composer dump-autoload [--optimize]`
//
// 用法示例：
//
//	// 生成标准自动加载文件
//	err := comp.DumpAutoload(false)
//	if err != nil {
//	    log.Fatalf("生成自动加载文件失败: %v", err)
//	}
//
//	// 生成优化的自动加载文件
//	err = comp.DumpAutoload(true)
//	if err != nil {
//	    log.Fatalf("生成优化的自动加载文件失败: %v", err)
//	}
func (c *Composer) DumpAutoload(optimize bool) error {
	args := []string{"dump-autoload"}

	if optimize {
		args = append(args, "--optimize")
	}

	_, err := c.Run(args...)
	return err
}

// DumpAutoloadWithOptions 使用更多选项生成自动加载文件
//
// 参数：
//   - options: 生成自动加载选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果生成自动加载文件过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法生成Composer的自动加载文件，支持更多自定义选项。
//
// 用法示例：
//
//	// 使用多个选项生成自动加载文件
//	options := map[string]string{
//	    "optimize": "",
//	    "classmap-authoritative": "",
//	    "apcu": "",
//	    "no-dev": "",
//	}
//	err := comp.DumpAutoloadWithOptions(options)
//	if err != nil {
//	    log.Fatalf("生成自动加载文件失败: %v", err)
//	}
func (c *Composer) DumpAutoloadWithOptions(options map[string]string) error {
	args := []string{"dump-autoload"}

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

// CheckDependencies 检查依赖项是否有冲突
//
// 返回值：
//   - string: 检查命令的输出结果
//   - error: 如果检查依赖项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查项目的依赖是否有冲突，composer.json和composer.lock是否同步。
//	相当于执行`composer check`
//
// 用法示例：
//
//	output, err := comp.CheckDependencies()
//	if err != nil {
//	    log.Fatalf("检查依赖失败: %v", err)
//	}
//	fmt.Println("依赖检查结果:", output)
func (c *Composer) CheckDependencies() (string, error) {
	return c.Run("check")
}

// Suggests 安装建议的软件包
//
// 返回值：
//   - error: 如果安装建议的软件包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法查看并安装建议的软件包。
//	相当于执行`composer suggests`
//
// 用法示例：
//
//	err := comp.Suggests()
//	if err != nil {
//	    log.Fatalf("查看建议的软件包失败: %v", err)
//	}
func (c *Composer) Suggests() error {
	_, err := c.Run("suggests")
	return err
}

// FundPackages 列出项目中可以捐赠的软件包
//
// 返回值：
//   - string: 列出可捐赠软件包的输出结果
//   - error: 如果列出可捐赠软件包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出项目中可以捐赠的软件包，显示它们的资金信息。
//	相当于执行`composer fund`
//
// 用法示例：
//
//	output, err := comp.FundPackages()
//	if err != nil {
//	    log.Fatalf("列出可捐赠软件包失败: %v", err)
//	}
//	fmt.Println("可捐赠软件包:", output)
func (c *Composer) FundPackages() (string, error) {
	return c.Run("fund")
}

// RunAudit 查找项目中使用的软件包的已知安全漏洞
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法审计项目的依赖项，查找已知的安全漏洞。
//	相当于执行`composer audit`
//
// 用法示例：
//
//	output, err := comp.RunAudit()
//	if err != nil {
//	    log.Fatalf("安全审计失败: %v", err)
//	}
//	fmt.Println("安全审计结果:", output)
func (c *Composer) RunAudit() (string, error) {
	return c.Run("audit")
}
