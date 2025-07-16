package composer

import (
	"testing"
)

func TestAddGitHubToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitHubToken("github.com", "token123")
	if err != nil {
		t.Errorf("AddGitHubToken执行失败: %v", err)
	}
}

func TestAddGitHubTokenWithEmptyToken(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitHubToken("github.com", "")
	if err == nil {
		t.Error("空token应该返回错误")
	}
}

func TestAddGitHubTokenWithEmptyDomain(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitHubToken("", "token123")
	if err == nil {
		t.Error("空域名应该返回错误")
	}
}

func TestAddGitLabToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitLabToken("gitlab.com", "token456")
	if err != nil {
		t.Errorf("AddGitLabToken执行失败: %v", err)
	}
}

func TestAddGitLabTokenWithEmptyDomain(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitLabToken("", "token456")
	if err == nil {
		t.Error("空域名应该返回错误")
	}
}

func TestAddGitLabTokenWithEmptyToken(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddGitLabToken("gitlab.com", "")
	if err == nil {
		t.Error("空token应该返回错误")
	}
}

func TestAddBitbucketToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddBitbucketToken("bitbucket.org", "consumer-key", "consumer-secret")
	if err != nil {
		t.Errorf("AddBitbucketToken执行失败: %v", err)
	}
}

func TestAddBitbucketTokenWithEmptyConsumer(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddBitbucketToken("bitbucket.org", "", "consumer-secret")
	if err == nil {
		t.Error("空consumer应该返回错误")
	}
}

func TestAddBitbucketTokenWithEmptyToken(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddBitbucketToken("bitbucket.org", "consumer-key", "")
	if err == nil {
		t.Error("空token应该返回错误")
	}
}

func TestAddBearerToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddBearerToken("api.example.com", "bearer-token-123")
	if err != nil {
		t.Errorf("AddBearerToken执行失败: %v", err)
	}
}

func TestAddHTTPBasicAuth(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.AddHTTPBasicAuth("repo.example.com", "username", "password")
	if err != nil {
		t.Errorf("AddHTTPBasicAuth执行失败: %v", err)
	}
}

func TestRemoveToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.RemoveToken("github-oauth", "github.com")
	if err != nil {
		t.Errorf("RemoveToken执行失败: %v", err)
	}
}

func TestGetToken(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetToken("github-oauth", "github.com")
	if err != nil {
		t.Errorf("GetToken执行失败: %v", err)
	}
}

func TestGetAuthConfig(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for config command to get composer home
	SetupMockOutput("config --global home", "/tmp/composer", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.GetAuthConfig()
	if err != nil {
		t.Errorf("GetAuthConfig执行失败: %v", err)
	}
}
