# 检测器 API

检测器 API 提供检测和验证 Composer 安装的功能。

## 即将推出

本节正在开发中。请稍后查看全面的检测器 API 文档。

目前，请参阅[英文版检测器 API](/api/detector) 获取完整文档。

## 主要功能

- Composer 安装检测
- 版本验证
- 路径发现
- 兼容性检查

## 基本示例

```go
// 检测器示例
detector := detector.New()
isInstalled := detector.IsComposerInstalled()
```

更多详细信息，请参阅[英文版文档](/api/detector)。
