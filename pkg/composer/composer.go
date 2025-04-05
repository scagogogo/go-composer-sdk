package composer

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/scagogogo/go-composer-sdk/pkg/detector"
	"github.com/scagogogo/go-composer-sdk/pkg/installer"
)

// 常见错误
var (
	ErrComposerNotFound     = errors.New("未找到composer可执行文件")
	ErrComposerInstallation = errors.New("安装composer失败")
	ErrCommandExecution     = errors.New("执行composer命令失败")
	ErrUnsupportedOS        = errors.New("不支持的操作系统")
	ErrInitFailed           = errors.New("初始化Composer失败")
	ErrSelfUpdateFailed     = errors.New("composer自更新失败")
	ErrRequirePackageFailed = errors.New("安装包失败")
	ErrInstallFailed        = errors.New("composer install失败")
	ErrUpdateFailed         = errors.New("composer update失败")
	ErrRemoveFailed         = errors.New("删除包失败")
	ErrValidateFailed       = errors.New("composer validate失败")
	ErrDumpAutoloadFailed   = errors.New("composer dump-autoload失败")
	ErrCreateProjectFailed  = errors.New("创建项目失败")
	ErrRunScriptFailed      = errors.New("运行脚本失败")
	ErrShowPackageFailed    = errors.New("查看包失败")
	ErrSearchFailed         = errors.New("搜索失败")
	ErrInitProjectFailed    = errors.New("初始化项目失败")
	ErrGetComposerHome      = errors.New("获取Composer主目录失败")
)

// mockCommandOutput stores predefined outputs for different commands
var mockCommandOutput = map[string]struct {
	output string
	err    error
}{}

// SetupMockOutput sets up mock output for a specific command
func SetupMockOutput(command string, output string, err error) {
	mockCommandOutput[command] = struct {
		output string
		err    error
	}{
		output: output,
		err:    err,
	}
}

// ClearMockOutputs clears all mock outputs
func ClearMockOutputs() {
	mockCommandOutput = map[string]struct {
		output string
		err    error
	}{}
}

// getMockOutput returns mock output for a command if it exists
func getMockOutput(args ...string) (string, error, bool) {
	if len(args) == 0 {
		return "", nil, false
	}

	// Try to match the exact command
	cmdStr := args[0]
	for i := 1; i < len(args); i++ {
		cmdStr += " " + args[i]
	}

	if mock, ok := mockCommandOutput[cmdStr]; ok {
		return mock.output, mock.err, true
	}

	// If no exact match, try to match just the first argument (the command)
	if mock, ok := mockCommandOutput[args[0]]; ok {
		return mock.output, mock.err, true
	}

	return "", nil, false
}

// Composer 封装PHP Composer的功能
type Composer struct {
	// composer可执行文件的路径
	executablePath string
	// 工作目录
	workingDir string
	// 是否自动安装
	autoInstall bool
	// 自定义安装器
	installer *installer.Installer
	// 自定义检测器
	detector *detector.Detector
	// 环境变量
	env []string
	// 默认超时时间
	defaultTimeout time.Duration
}

// Options 用于自定义Composer实例的选项
type Options struct {
	// composer可执行文件的路径
	ExecutablePath string
	// 工作目录
	WorkingDir string
	// 是否在未找到composer时自动安装
	AutoInstall bool
	// 自定义安装器
	Installer *installer.Installer
	// 自定义检测器
	Detector *detector.Detector
	// 环境变量
	Env []string
	// 默认超时时间
	DefaultTimeout time.Duration
}

// DefaultOptions 返回默认选项
func DefaultOptions() Options {
	return Options{
		WorkingDir:     "",
		AutoInstall:    true,
		DefaultTimeout: 10 * time.Minute,
	}
}

// New 创建一个新的Composer实例
func New(options Options) (*Composer, error) {
	c := &Composer{
		executablePath: options.ExecutablePath,
		workingDir:     options.WorkingDir,
		autoInstall:    options.AutoInstall,
		installer:      options.Installer,
		detector:       options.Detector,
		env:            options.Env,
		defaultTimeout: options.DefaultTimeout,
	}

	// 如果未提供安装器，使用默认安装器
	if c.installer == nil {
		c.installer = installer.NewInstaller(installer.DefaultConfig())
	}

	// 如果未提供检测器，使用默认检测器
	if c.detector == nil {
		c.detector = detector.NewDetector()
	}

	// 如果未指定可执行文件路径，则尝试检测
	if c.executablePath == "" {
		execPath, err := c.detector.Detect()
		if err != nil {
			if c.autoInstall {
				// 尝试安装composer
				if err := c.installer.Install(); err != nil {
					return nil, fmt.Errorf("%w: %v", ErrComposerInstallation, err)
				}

				// 重新尝试检测
				execPath, err = c.detector.Detect()
				if err != nil {
					return nil, fmt.Errorf("%w: %v", ErrComposerNotFound, err)
				}
			} else {
				return nil, fmt.Errorf("%w: %v", ErrComposerNotFound, err)
			}
		}
		c.executablePath = execPath
	}

	return c, nil
}

// SetWorkingDir 设置composer命令的工作目录
func (c *Composer) SetWorkingDir(dir string) {
	c.workingDir = dir
}

// SetEnv 设置环境变量
func (c *Composer) SetEnv(env []string) {
	c.env = env
}

// GetExecutablePath 获取composer可执行文件的路径
func (c *Composer) GetExecutablePath() string {
	return c.executablePath
}

// Run 执行composer命令并返回输出
func (c *Composer) Run(args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.defaultTimeout)
	defer cancel()

	return c.RunWithContext(ctx, args...)
}

// RunWithContext 使用指定的上下文执行composer命令并返回输出
func (c *Composer) RunWithContext(ctx context.Context, args ...string) (string, error) {
	// 当处于测试模式时，检查是否有模拟输出
	if output, err, ok := getMockOutput(args...); ok {
		return output, err
	}

	cmd := exec.CommandContext(ctx, c.executablePath, args...)

	if c.workingDir != "" {
		cmd.Dir = c.workingDir
	}

	if len(c.env) > 0 {
		cmd.Env = c.env
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("%w: %v, 输出: %s", ErrCommandExecution, err, output)
	}

	return string(output), nil
}

// IsInstalled 检查composer是否已安装
func (c *Composer) IsInstalled() bool {
	return c.executablePath != ""
}
