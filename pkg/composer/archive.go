package composer

// Archive 创建项目的归档文件（不包含开发依赖）
func (c *Composer) Archive(destination string) (string, error) {
	return c.Run("archive", "--format=zip", "--dir="+destination)
}

// ArchiveWithFormat 使用指定格式创建项目的归档文件
func (c *Composer) ArchiveWithFormat(destination string, format string) (string, error) {
	return c.Run("archive", "--format="+format, "--dir="+destination)
}

// ArchiveWithOptions 使用自定义选项创建项目的归档文件
func (c *Composer) ArchiveWithOptions(destination string, options map[string]string) (string, error) {
	args := []string{"archive", "--dir=" + destination}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ArchivePackage 创建指定包的归档文件
func (c *Composer) ArchivePackage(packageName string, version string, destination string) (string, error) {
	args := []string{"archive"}

	if version != "" {
		args = append(args, packageName+"="+version)
	} else {
		args = append(args, packageName)
	}

	args = append(args, "--dir="+destination)

	return c.Run(args...)
}

// ArchivePackageWithOptions 使用自定义选项创建指定包的归档文件
func (c *Composer) ArchivePackageWithOptions(packageName string, version string, destination string, options map[string]string) (string, error) {
	args := []string{"archive"}

	if version != "" {
		args = append(args, packageName+"="+version)
	} else {
		args = append(args, packageName)
	}

	args = append(args, "--dir="+destination)

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}
