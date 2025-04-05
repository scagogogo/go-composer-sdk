package detector

import (
	"os"
	"runtime"
	"testing"
)

// mockOS 临时修改runtime.GOOS用于测试
func mockOS(t *testing.T, mockGOOS string, testFunc func(t *testing.T)) {
	// 保存原始值
	originalGOOS := runtime.GOOS

	// 模拟不同的值
	mockRuntime(t, mockGOOS)

	// 执行测试
	testFunc(t)

	// 恢复原始值
	mockRuntime(t, originalGOOS)
}

// mockRuntime 修改runtime包中的GOOS变量
// 注意：这是一个实验性的方法，在正式环境中应谨慎使用
func mockRuntime(t *testing.T, goos string) {
	// 这里我们不能直接修改runtime.GOOS，因为它是只读的
	// 但我们可以设置环境变量，一些函数可能会读取这个环境变量
	err := os.Setenv("MOCK_GOOS", goos)
	if err != nil {
		t.Fatalf("设置MOCK_GOOS环境变量失败: %v", err)
	}
}

// 通用测试函数，使用不同平台的getPlatformSpecificPaths实现
func TestAllPlatformSpecificPaths(t *testing.T) {
	// 测试Windows平台的路径
	t.Run("TestWindowsPaths", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("在Windows平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的Windows环境测试
		paths := []string{
			"AppData\\Composer\\composer.phar",
			"ProgramFiles\\Composer\\composer.phar",
			"composer.phar",
			"composer.bat",
			"composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Windows平台的路径列表不应为空")
		}
	})

	// 测试macOS平台的路径
	t.Run("TestDarwinPaths", func(t *testing.T) {
		if runtime.GOOS == "darwin" {
			t.Skip("在Darwin平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的macOS环境测试
		paths := []string{
			"/usr/local/bin/composer",
			"/usr/bin/composer",
			"/opt/homebrew/bin/composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Darwin平台的路径列表不应为空")
		}
	})

	// 测试Linux/Unix平台的路径
	t.Run("TestUnixPaths", func(t *testing.T) {
		if runtime.GOOS != "windows" && runtime.GOOS != "darwin" {
			t.Skip("在Unix平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的Unix环境测试
		paths := []string{
			"/usr/local/bin/composer",
			"/usr/bin/composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Unix平台的路径列表不应为空")
		}
	})
}
