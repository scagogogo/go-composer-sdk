package composer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ComposerJSON 表示 composer.json 文件的结构
type ComposerJSON struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Type                string                 `json:"type,omitempty"`
	Keywords            []string               `json:"keywords,omitempty"`
	Homepage            string                 `json:"homepage,omitempty"`
	License             interface{}            `json:"license,omitempty"`
	Authors             []map[string]string    `json:"authors,omitempty"`
	Support             map[string]string      `json:"support,omitempty"`
	Require             map[string]string      `json:"require,omitempty"`
	RequireDev          map[string]string      `json:"require-dev,omitempty"`
	Suggest             map[string]string      `json:"suggest,omitempty"`
	Autoload            map[string]interface{} `json:"autoload,omitempty"`
	AutoloadDev         map[string]interface{} `json:"autoload-dev,omitempty"`
	Repositories        map[string]interface{} `json:"repositories,omitempty"`
	Config              map[string]interface{} `json:"config,omitempty"`
	Scripts             map[string]interface{} `json:"scripts,omitempty"`
	ScriptsDescriptions map[string]string      `json:"scripts-descriptions,omitempty"`
	Extra               map[string]interface{} `json:"extra,omitempty"`
	Bin                 []string               `json:"bin,omitempty"`
	Archive             map[string]interface{} `json:"archive,omitempty"`
	NonFeatureBranches  []string               `json:"non-feature-branches,omitempty"`
	MinimumStability    string                 `json:"minimum-stability,omitempty"`
	PreferStable        bool                   `json:"prefer-stable,omitempty"`
	Replace             map[string]string      `json:"replace,omitempty"`
	Conflict            map[string]string      `json:"conflict,omitempty"`
	Provide             map[string]string      `json:"provide,omitempty"`
}

// ErrComposerJSONNotFound 表示未找到 composer.json 文件
var ErrComposerJSONNotFound = errors.New("composer.json not found")

// mockComposerJSON stores a mock ComposerJSON for testing
var mockComposerJSON *ComposerJSON
var useMockComposerJSON bool

// SetMockComposerJSON sets a mock ComposerJSON for testing
func SetMockComposerJSON(composerJSON *ComposerJSON) {
	mockComposerJSON = composerJSON
	useMockComposerJSON = true
}

// ClearMockComposerJSON clears the mock ComposerJSON
func ClearMockComposerJSON() {
	mockComposerJSON = nil
	useMockComposerJSON = false
}

// ReadComposerJSON 读取并解析 composer.json 文件
func (c *Composer) ReadComposerJSON() (*ComposerJSON, error) {
	// 如果设置了mock，则使用mock
	if useMockComposerJSON && mockComposerJSON != nil {
		return mockComposerJSON, nil
	}

	// 如果指定了工作目录，则从工作目录中读取
	workDir := c.workingDir
	if workDir == "" {
		// 如果未指定工作目录，则使用当前目录
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	// 读取 composer.json 文件
	composerJSONPath := filepath.Join(workDir, "composer.json")
	if _, err := os.Stat(composerJSONPath); os.IsNotExist(err) {
		return nil, ErrComposerJSONNotFound
	}

	content, err := ioutil.ReadFile(composerJSONPath)
	if err != nil {
		return nil, err
	}

	// 解析 JSON
	var composerJSON ComposerJSON
	if err := json.Unmarshal(content, &composerJSON); err != nil {
		return nil, err
	}

	return &composerJSON, nil
}

// WriteComposerJSON 将 ComposerJSON 写入 composer.json 文件
func (c *Composer) WriteComposerJSON(composerJSON *ComposerJSON) error {
	// 如果指定了工作目录，则写入工作目录
	workDir := c.workingDir
	if workDir == "" {
		// 如果未指定工作目录，则使用当前目录
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	// 序列化为 JSON
	content, err := json.MarshalIndent(composerJSON, "", "    ")
	if err != nil {
		return err
	}

	// 写入 composer.json 文件
	composerJSONPath := filepath.Join(workDir, "composer.json")
	return ioutil.WriteFile(composerJSONPath, content, 0644)
}

// AddRequire 添加依赖到 composer.json
func (c *Composer) AddRequire(packageName, version string, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if isDev {
		if composerJSON.RequireDev == nil {
			composerJSON.RequireDev = make(map[string]string)
		}
		composerJSON.RequireDev[packageName] = version
	} else {
		if composerJSON.Require == nil {
			composerJSON.Require = make(map[string]string)
		}
		composerJSON.Require[packageName] = version
	}

	return c.WriteComposerJSON(composerJSON)
}

// RemoveRequire 从 composer.json 中移除依赖
func (c *Composer) RemoveRequire(packageName string, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if isDev {
		if composerJSON.RequireDev != nil {
			delete(composerJSON.RequireDev, packageName)
		}
	} else {
		if composerJSON.Require != nil {
			delete(composerJSON.Require, packageName)
		}
	}

	return c.WriteComposerJSON(composerJSON)
}

// AddScript 添加脚本到 composer.json
func (c *Composer) AddScript(name string, script interface{}, description string) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 添加脚本
	if composerJSON.Scripts == nil {
		composerJSON.Scripts = make(map[string]interface{})
	}
	composerJSON.Scripts[name] = script

	// 添加脚本描述（如果提供）
	if description != "" {
		if composerJSON.ScriptsDescriptions == nil {
			composerJSON.ScriptsDescriptions = make(map[string]string)
		}
		composerJSON.ScriptsDescriptions[name] = description
	}

	return c.WriteComposerJSON(composerJSON)
}

// RemoveScript 从 composer.json 中移除脚本
func (c *Composer) RemoveScript(name string) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 移除脚本
	if composerJSON.Scripts != nil {
		delete(composerJSON.Scripts, name)
	}

	// 移除脚本描述
	if composerJSON.ScriptsDescriptions != nil {
		delete(composerJSON.ScriptsDescriptions, name)
	}

	return c.WriteComposerJSON(composerJSON)
}

// AddAutoload 添加自动加载配置到 composer.json
func (c *Composer) AddAutoload(type_ string, namespace string, paths interface{}, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 选择正确的自动加载配置
	var autoload map[string]interface{}
	if isDev {
		if composerJSON.AutoloadDev == nil {
			composerJSON.AutoloadDev = make(map[string]interface{})
		}
		autoload = composerJSON.AutoloadDev
	} else {
		if composerJSON.Autoload == nil {
			composerJSON.Autoload = make(map[string]interface{})
		}
		autoload = composerJSON.Autoload
	}

	// 确保类型的配置存在
	if autoload[type_] == nil {
		autoload[type_] = make(map[string]interface{})
	}

	// 添加命名空间和路径
	typeConfig, ok := autoload[type_].(map[string]interface{})
	if !ok {
		return errors.New("invalid autoload configuration")
	}
	typeConfig[namespace] = paths

	return c.WriteComposerJSON(composerJSON)
}

// SetConfig 设置 composer.json 中的配置选项
func (c *Composer) SetConfig(key string, value interface{}) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if composerJSON.Config == nil {
		composerJSON.Config = make(map[string]interface{})
	}
	composerJSON.Config[key] = value

	return c.WriteComposerJSON(composerJSON)
}

// GetConfig 获取 composer.json 中的配置选项
func (c *Composer) GetConfig(key string) (interface{}, error) {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return nil, err
	}

	if composerJSON.Config == nil {
		return nil, nil
	}
	return composerJSON.Config[key], nil
}

// SetProperty 设置 composer.json 中的顶级属性
func (c *Composer) SetProperty(property string, value interface{}) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 使用反射来设置属性可能更好，但为了简单起见，我们这里使用硬编码
	switch property {
	case "name":
		if strValue, ok := value.(string); ok {
			composerJSON.Name = strValue
		}
	case "description":
		if strValue, ok := value.(string); ok {
			composerJSON.Description = strValue
		}
	case "type":
		if strValue, ok := value.(string); ok {
			composerJSON.Type = strValue
		}
	case "keywords":
		if strArray, ok := value.([]string); ok {
			composerJSON.Keywords = strArray
		}
	case "homepage":
		if strValue, ok := value.(string); ok {
			composerJSON.Homepage = strValue
		}
	case "license":
		composerJSON.License = value
	case "minimum-stability":
		if strValue, ok := value.(string); ok {
			composerJSON.MinimumStability = strValue
		}
	case "prefer-stable":
		if boolValue, ok := value.(bool); ok {
			composerJSON.PreferStable = boolValue
		}
	default:
		return errors.New("unsupported property")
	}

	return c.WriteComposerJSON(composerJSON)
}
