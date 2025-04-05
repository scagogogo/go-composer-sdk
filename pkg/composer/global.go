package composer

// GlobalRequire 全局安装包
func (c *Composer) GlobalRequire(packageName string, version string) error {
	args := []string{"global", "require"}

	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// GlobalUpdate 全局更新包
func (c *Composer) GlobalUpdate(packages []string) error {
	args := []string{"global", "update"}
	args = append(args, packages...)

	_, err := c.Run(args...)
	return err
}

// GlobalRemove 全局移除包
func (c *Composer) GlobalRemove(packageName string) error {
	_, err := c.Run("global", "remove", packageName)
	return err
}

// GlobalInstall 全局安装依赖
func (c *Composer) GlobalInstall() error {
	_, err := c.Run("global", "install")
	return err
}

// GlobalList 列出全局安装的包
func (c *Composer) GlobalList() (string, error) {
	return c.Run("global", "show")
}

// GlobalHome 获取全局目录路径
func (c *Composer) GlobalHome() (string, error) {
	return c.Run("global", "home")
}

// GlobalExecute 执行全局安装的包中的二进制文件
func (c *Composer) GlobalExecute(command string, args ...string) (string, error) {
	cmdArgs := append([]string{"global", "exec", command}, args...)
	return c.Run(cmdArgs...)
}

// GlobalStatus 显示全局安装的包的状态
func (c *Composer) GlobalStatus() (string, error) {
	return c.Run("global", "status")
}

// GlobalDumpAutoload 为全局安装生成自动加载文件
func (c *Composer) GlobalDumpAutoload(optimize bool) error {
	args := []string{"global", "dump-autoload"}

	if optimize {
		args = append(args, "--optimize")
	}

	_, err := c.Run(args...)
	return err
}
