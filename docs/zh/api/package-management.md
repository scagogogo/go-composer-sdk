# 包管理 API

包管理 API 提供管理 PHP 包和依赖的全面功能。

## 即将推出

本节正在开发中。请稍后查看全面的包管理 API 文档。

目前，请参阅[英文版包管理 API](/api/package-management) 获取完整文档。

## 主要功能

- 安装和更新依赖
- 添加和删除包
- 包信息查询
- 依赖分析
- 过时包检查

## 基本示例

```go
// 安装依赖
err := comp.Install(false, false)

// 添加包
err = comp.RequirePackage("monolog/monolog", "^3.0")

// 更新包
err = comp.Update(false, false)

// 删除包
err = comp.RemovePackage("old-package/deprecated")
```

更多详细信息，请参阅[英文版文档](/api/package-management)。
