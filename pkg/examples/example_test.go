package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试示例
func TestSum(t *testing.T) {
	// 使用testify/assert简化测试断言
	assert.Equal(t, 5, Sum(2, 3), "Sum(2, 3) should equal 5")
	assert.NotEqual(t, 10, Sum(2, 3), "Sum(2, 3) should not equal 10")
}

// 表格驱动测试示例
func TestSumTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Sum(%d, %d) should equal %d", tt.a, tt.b, tt.expected)
		})
	}
}
