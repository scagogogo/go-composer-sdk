# 如何使用act在本地测试GitHub Actions

[act](https://github.com/nektos/act) 是一个工具，可以在本地运行GitHub Actions，无需将代码推送到GitHub即可测试工作流。

## 安装act

### macOS
```bash
brew install act
```

### 其他操作系统
请参考[官方安装指南](https://github.com/nektos/act#installation)

## 使用方法

在项目根目录下运行：

```bash
# 列出所有可用的actions
act -l

# 运行所有actions
act

# 只运行测试作业
act -j test

# 使用指定的事件触发工作流
act push

# 在运行act时指定要使用的Docker镜像（必要时）
act -P ubuntu-latest=node:16-buster-slim
```

## 针对本项目的示例命令

在推送代码前，您可以通过以下命令来测试我们的测试工作流：

```bash
# 运行测试工作流
act -j test

# 或者指定使用与Go 1.20兼容的Docker镜像
act -j test -P ubuntu-latest=golang:1.20
```

## 注意事项

1. act使用Docker来模拟GitHub Actions的环境，因此您需要安装Docker。
2. 某些GitHub Actions可能在本地运行时与在GitHub上运行略有不同，但对于基本的测试目的来说应该足够了。
3. act默认使用最小的Docker镜像，如果需要更完整的环境，可以使用 `-P ubuntu-latest=catthehacker/ubuntu:act-latest` 参数。 