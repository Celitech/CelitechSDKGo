package configmanager

import (
	"github.com/Celitech/CelitechSDKGo/internal/oauthtokenmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"time"
)

type ConfigManager struct {
	OAuth             celitechconfig.Config
	Destinations      celitechconfig.Config
	Packages          celitechconfig.Config
	Purchases         celitechconfig.Config
	ESim              celitechconfig.Config
	IFrame            celitechconfig.Config
	oAuthTokenManager *oauthtokenmanager.OAuthTokenManager
}

func NewConfigManager(config celitechconfig.Config, tokenService oauthtokenmanager.TokenService) *ConfigManager {
	return &ConfigManager{
		OAuth:             config,
		Destinations:      config,
		Packages:          config,
		Purchases:         config,
		ESim:              config,
		IFrame:            config,
		oAuthTokenManager: oauthtokenmanager.NewOAuthTokenManager(tokenService),
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.OAuth.SetBaseUrl(baseUrl)
	c.Destinations.SetBaseUrl(baseUrl)
	c.Packages.SetBaseUrl(baseUrl)
	c.Purchases.SetBaseUrl(baseUrl)
	c.ESim.SetBaseUrl(baseUrl)
	c.IFrame.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.OAuth.SetTimeout(timeout)
	c.Destinations.SetTimeout(timeout)
	c.Packages.SetTimeout(timeout)
	c.Purchases.SetTimeout(timeout)
	c.ESim.SetTimeout(timeout)
	c.IFrame.SetTimeout(timeout)
}

func (c *ConfigManager) SetClientId(clientId string) {
	c.OAuth.SetClientId(clientId)
	c.Destinations.SetClientId(clientId)
	c.Packages.SetClientId(clientId)
	c.Purchases.SetClientId(clientId)
	c.ESim.SetClientId(clientId)
	c.IFrame.SetClientId(clientId)
}

func (c *ConfigManager) SetClientSecret(clientSecret string) {
	c.OAuth.SetClientSecret(clientSecret)
	c.Destinations.SetClientSecret(clientSecret)
	c.Packages.SetClientSecret(clientSecret)
	c.Purchases.SetClientSecret(clientSecret)
	c.ESim.SetClientSecret(clientSecret)
	c.IFrame.SetClientSecret(clientSecret)
}

func (c *ConfigManager) GetTokenManager() *oauthtokenmanager.OAuthTokenManager {
	return c.oAuthTokenManager
}

func (c *ConfigManager) GetOAuth() *celitechconfig.Config {
	return &c.OAuth
}
func (c *ConfigManager) GetDestinations() *celitechconfig.Config {
	return &c.Destinations
}
func (c *ConfigManager) GetPackages() *celitechconfig.Config {
	return &c.Packages
}
func (c *ConfigManager) GetPurchases() *celitechconfig.Config {
	return &c.Purchases
}
func (c *ConfigManager) GetESim() *celitechconfig.Config {
	return &c.ESim
}
func (c *ConfigManager) GetIFrame() *celitechconfig.Config {
	return &c.IFrame
}
