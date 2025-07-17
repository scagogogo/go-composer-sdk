# 安全审计 API

安全审计 API 提供检查依赖漏洞和安全问题的功能。

## 即将推出

本节正在开发中。请稍后查看全面的安全审计 API 文档。

目前，请参阅[英文版安全审计 API](/api/security-audit) 获取完整文档。

## 主要功能

- 安全漏洞扫描
- 依赖审计
- 平台要求检查
- 安全报告生成

## 基本示例

```go
// 执行安全审计
auditResult, err := comp.Audit()

// 检查平台要求
err = comp.CheckPlatformReqs()
```

更多详细信息，请参阅[英文版文档](/api/security-audit)。
