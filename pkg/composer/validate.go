package composer

// ValidateStrict 严格验证 composer.json 是否有效
func (c *Composer) ValidateStrict() (string, error) {
	return c.Run("validate", "--strict")
}

// ValidateWithNoCheck 验证 composer.json 而不检查平台需求
func (c *Composer) ValidateWithNoCheck() (string, error) {
	return c.Run("validate", "--no-check-all")
}

// ValidateWithNoCheckPublish 验证 composer.json 而不检查发布要求
func (c *Composer) ValidateWithNoCheckPublish() (string, error) {
	return c.Run("validate", "--no-check-publish")
}

// ValidateWithCheckVersion 验证 composer.json 并检查版本约束
func (c *Composer) ValidateWithCheckVersion() (string, error) {
	return c.Run("validate", "--with-dependencies")
}

// CheckPlatformReqsLock 检查 composer.lock 文件中的平台需求
func (c *Composer) CheckPlatformReqsLock() (string, error) {
	return c.Run("check-platform-reqs", "--lock")
}

// CheckForOutdatedPackages 检查过时的包，并显示可能的更新
func (c *Composer) CheckForOutdatedPackages(direct bool, minor bool, format string) (string, error) {
	args := []string{"outdated"}

	if direct {
		args = append(args, "--direct")
	}

	if minor {
		args = append(args, "--minor-only")
	}

	if format != "" {
		args = append(args, "--format", format)
	}

	return c.Run(args...)
}

// ValidateSchema 验证 composer.json 和 composer.lock 是否符合架构
func (c *Composer) ValidateSchema() (string, error) {
	return c.Run("validate", "--no-check-all", "--no-check-publish", "--no-check-version")
}

// ValidateWithOptions 使用自定义选项验证 composer.json
func (c *Composer) ValidateWithOptions(options map[string]string) (string, error) {
	args := []string{"validate"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}
