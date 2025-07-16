package composer

import (
	"errors"
	"testing"
)

func TestCreateProject(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for create-project command
	SetupMockOutput("create-project vendor/package my-project", "Installing vendor/package (v1.0.0)\nCreated project in my-project", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("vendor/package", "my-project", "")
	if err != nil {
		t.Errorf("CreateProject执行失败: %v", err)
	}
}

func TestCreateProjectWithEmptyPackage(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("", "my-project", "")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestCreateProjectWithEmptyDirectory(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("vendor/package", "", "")
	if err == nil {
		t.Error("空目录名应该返回错误")
	}
}

func TestCreateProjectWithVersion(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for create-project command with version
	SetupMockOutput("create-project vendor/package:^2.0 my-project", "Installing vendor/package (v2.0.0)\nCreated project in my-project", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("vendor/package", "my-project", "^2.0")
	if err != nil {
		t.Errorf("CreateProject执行失败: %v", err)
	}
}

func TestCreateProjectWithVersionAndEmptyVersion(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for create-project command without version
	SetupMockOutput("create-project vendor/package my-project", "Installing vendor/package (v1.0.0)\nCreated project in my-project", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("vendor/package", "my-project", "")
	if err != nil {
		t.Errorf("CreateProject（空版本）执行失败: %v", err)
	}
}

func TestCreateProjectWithOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for create-project command with options
	SetupMockOutput("create-project --prefer-dist --no-scripts --no-dev vendor/package my-project", "Installing vendor/package (v1.0.0)\nCreated project in my-project", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"prefer-dist": "",
		"no-scripts":  "",
		"no-dev":      "",
	}

	err = composer.CreateProjectWithOptions("vendor/package", "my-project", "", options)
	if err != nil {
		t.Errorf("CreateProjectWithOptions执行失败: %v", err)
	}
}

func TestCreateProjectWithOptionsAndNilOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for create-project command without options
	SetupMockOutput("create-project vendor/package my-project", "Installing vendor/package (v1.0.0)\nCreated project in my-project", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProjectWithOptions("vendor/package", "my-project", "", nil)
	if err != nil {
		t.Errorf("CreateProjectWithOptions（nil选项）执行失败: %v", err)
	}
}

func TestCreateProjectWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("create-project nonexistent/package my-project", "", errors.New("package not found"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.CreateProject("nonexistent/package", "my-project", "")
	if err == nil {
		t.Error("不存在的包应该返回错误")
	}
}

func TestInitProject(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for init command
	SetupMockOutput("init", "Welcome to the Composer config generator\n\nThis command will guide you through creating your composer.json config.\n\nPackage name (<vendor>/<name>) [user/project]: my-vendor/my-project\nDescription []: My awesome project\nAuthor [User <user@example.com>, n to skip]: \nMinimum Stability []: \nPackage Type (e.g. library, project, metapackage, composer-plugin) []: \nLicense []: MIT\n\nDefine your dependencies.\n\nWould you like to define your dependencies (require) interactively [yes]? no\nWould you like to define your dev dependencies (require-dev) interactively [yes]? no\n\n{\n    \"name\": \"my-vendor/my-project\",\n    \"description\": \"My awesome project\",\n    \"license\": \"MIT\",\n    \"require\": {}\n}\n\nDo you confirm generation [yes]? yes", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProject()
	if err != nil {
		t.Errorf("InitProject执行失败: %v", err)
	}
}

func TestInitProjectWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("init", "", errors.New("init failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProject()
	if err == nil {
		t.Error("初始化失败时应该返回错误")
	}
}

func TestInitProjectWithName(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for init command with name
	SetupMockOutput("init --name=my-vendor/my-project", "{\n    \"name\": \"my-vendor/my-project\",\n    \"require\": {}\n}\n\nDo you confirm generation [yes]? yes", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProjectWithOptions("my-vendor/my-project", "", "", nil)
	if err != nil {
		t.Errorf("InitProjectWithOptions执行失败: %v", err)
	}
}

func TestInitProjectWithNameAndEmptyName(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProjectWithOptions("", "", "", nil)
	if err == nil {
		t.Error("空项目名应该返回错误")
	}
}

func TestInitProjectWithOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for init command with options
	SetupMockOutput("init --name=my-vendor/my-project --description=\"My awesome project\" --author=\"John Doe <john@example.com>\" --type=library --license=MIT", "{\n    \"name\": \"my-vendor/my-project\",\n    \"description\": \"My awesome project\",\n    \"type\": \"library\",\n    \"license\": \"MIT\",\n    \"authors\": [\n        {\n            \"name\": \"John Doe\",\n            \"email\": \"john@example.com\"\n        }\n    ],\n    \"require\": {}\n}\n\nDo you confirm generation [yes]? yes", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"name":        "my-vendor/my-project",
		"description": "My awesome project",
		"author":      "John Doe <john@example.com>",
		"type":        "library",
		"license":     "MIT",
	}

	err = composer.InitProjectWithOptions("my-vendor/my-project", "My awesome project", "John Doe <john@example.com>", options)
	if err != nil {
		t.Errorf("InitProjectWithOptions执行失败: %v", err)
	}
}

func TestInitProjectWithOptionsAndNilOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for init command without options
	SetupMockOutput("init", "Welcome to the Composer config generator", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProjectWithOptions("", "", "", nil)
	if err != nil {
		t.Errorf("InitProjectWithOptions（nil选项）执行失败: %v", err)
	}
}

func TestInitProjectWithOptionsAndEmptyOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for init command without options
	SetupMockOutput("init", "Welcome to the Composer config generator", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.InitProjectWithOptions("", "", "", map[string]string{})
	if err != nil {
		t.Errorf("InitProjectWithOptions（空选项）执行失败: %v", err)
	}
}
