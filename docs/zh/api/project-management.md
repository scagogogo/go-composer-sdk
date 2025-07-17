# 项目管理 API

项目管理 API 提供创建、配置和管理 PHP 项目的功能。

## 即将推出

本节正在开发中。请稍后查看全面的项目管理 API 文档。

目前，请参阅[英文版项目管理 API](/api/project-management) 获取完整文档。

## 主要功能

- 创建新项目
- 验证 composer.json
- 运行脚本
- 管理自动加载器
- 项目配置

## 基本示例

```go
// 创建新项目
err := comp.CreateProject("laravel/laravel", "my-app", "")

// 验证 composer.json
err = comp.Validate()

// 运行脚本
err = comp.RunScript("test")
```

更多详细信息，请参阅[英文版文档](/api/project-management)。
