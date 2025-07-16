# Composer SDK 测试

本目录包含用于测试 go-composer-sdk 功能的测试代码。

## 测试内容

测试包括：

1. Composer安装功能测试
2. PHP项目依赖管理测试
3. Linux环境自动安装测试（Docker/Vagrant）

## 目录结构

- `install/` - 测试Composer安装功能
- `phpunit/` - 测试PHP单元测试和依赖管理
- `docker/` - 使用Docker测试Linux环境中的功能
- `vagrant/` - 使用Vagrant虚拟机测试Linux环境中的功能

## 运行测试

### 1. 安装测试

测试SDK自动安装Composer的功能：

```bash
cd tests/install
RUN_INSTALL_TEST=1 go run main.go
```

### 2. PHP依赖管理和单元测试

测试使用Composer管理PHP项目依赖并运行PHPUnit测试：

```bash
cd tests/phpunit
RUN_COMPOSER_TESTS=1 go run main.go
```

### 3. Linux环境测试（通过Docker）

测试在Linux环境中安装和使用Composer（需要Docker）：

```bash
cd tests/docker
./run_tests.sh
```

### 4. Linux环境测试（通过Vagrant）

如果没有Docker，可以使用Vagrant虚拟机进行测试：

```bash
cd tests/vagrant
./run_test.sh
```

## 手动测试

也可以进入各个测试目录，查看README文件了解详细的手动测试步骤。

## 测试环境要求

- Go 1.20+
- PHP 8.0+
- Docker或Vagrant+VirtualBox（用于Linux环境测试） 