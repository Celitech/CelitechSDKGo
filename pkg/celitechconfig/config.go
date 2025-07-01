package celitechconfig

import "time"

type Config struct {
	BaseUrl      *string
	Timeout      *time.Duration
	ClientId     *string
	ClientSecret *string
	HookParams   map[string]string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	timeout := time.Second * 10
	newConfig := Config{
		BaseUrl:    &baseUrl,
		Timeout:    &timeout,
		HookParams: make(map[string]string),
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) GetBaseUrl() string {
	return *c.BaseUrl
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = &timeout
}

func (c *Config) GetTimeout() time.Duration {
	return *c.Timeout
}

func (c *Config) SetClientId(clientId string) {
	c.ClientId = &clientId
}

func (c *Config) GetClientId() string {
	return *c.ClientId
}

func (c *Config) SetClientSecret(clientSecret string) {
	c.ClientSecret = &clientSecret
}

func (c *Config) GetClientSecret() string {
	return *c.ClientSecret
}
