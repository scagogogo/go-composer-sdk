package composer

// RequirePackage 添加一个新的包依赖
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
func (c *Composer) ShowPackage(packageName string) (string, error) {
	return c.Run("show", packageName)
}

// Search 搜索包
func (c *Composer) Search(query string) (string, error) {
	return c.Run("search", query)
}

// ShowAllPackages 显示所有已安装的包
func (c *Composer) ShowAllPackages() (string, error) {
	return c.Run("show")
}

// ShowDependencyTree 显示依赖树
func (c *Composer) ShowDependencyTree(packageName string) (string, error) {
	if packageName != "" {
		return c.Run("show", "--tree", packageName)
	}
	return c.Run("show", "--tree")
}

// ShowReverseDependencies 显示哪些包依赖于指定包（depends命令）
func (c *Composer) ShowReverseDependencies(packageName string) (string, error) {
	return c.Run("depends", packageName)
}

// WhyPackage 解释为什么安装了指定的包（why命令）
func (c *Composer) WhyPackage(packageName string) (string, error) {
	return c.Run("why", packageName)
}

// OutdatedPackages 显示所有过时的包
func (c *Composer) OutdatedPackages() (string, error) {
	return c.Run("outdated")
}

// OutdatedPackagesDirect 显示直接依赖中过时的包
func (c *Composer) OutdatedPackagesDirect() (string, error) {
	return c.Run("outdated", "--direct")
}

// RequirePackageWithOptions 添加带有选项的包依赖
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

	// 添加包名和版本
	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// BumpPackages 将包依赖升级到其最新版本
func (c *Composer) BumpPackages(packages []string) error {
	args := []string{"bump"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// BumpPackagesWithOptions 使用更多选项将包依赖升级到其最新版本
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

// Reinstall 重新安装包
func (c *Composer) Reinstall(packageName string) error {
	_, err := c.Run("reinstall", packageName)
	return err
}

// BrowsePackage 打开包的仓库或主页
func (c *Composer) BrowsePackage(packageName string) error {
	_, err := c.Run("browse", packageName)
	return err
}

// BrowsePackageWithOptions 使用更多选项打开包的仓库或主页
func (c *Composer) BrowsePackageWithOptions(packageName string, options map[string]string) error {
	args := []string{"browse"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	args = append(args, packageName)
	_, err := c.Run(args...)
	return err
}

// WhyNotPackage 解释为什么某个包不能被安装（prohibits命令）
func (c *Composer) WhyNotPackage(packageName string, version string) (string, error) {
	if version != "" {
		return c.Run("prohibits", packageName, version)
	}
	return c.Run("prohibits", packageName)
}
