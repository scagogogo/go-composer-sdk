package composer

// Install 安装依赖项
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
func (c *Composer) DumpAutoload(optimize bool) error {
	args := []string{"dump-autoload"}

	if optimize {
		args = append(args, "--optimize")
	}

	_, err := c.Run(args...)
	return err
}

// DumpAutoloadWithOptions 使用更多选项生成自动加载文件
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
func (c *Composer) CheckDependencies() (string, error) {
	return c.Run("check")
}

// Suggests 安装建议的软件包
func (c *Composer) Suggests() error {
	_, err := c.Run("suggests")
	return err
}

// FundPackages 列出项目中可以捐赠的软件包
func (c *Composer) FundPackages() (string, error) {
	return c.Run("fund")
}

// RunAudit 查找项目中使用的软件包的已知安全漏洞
func (c *Composer) RunAudit() (string, error) {
	return c.Run("audit")
}
