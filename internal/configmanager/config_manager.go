package configmanager

import (
	"github.com/Celitech/CelitechSDKGo/internal/oauthtokenmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	OAuth             celitechconfig.Config
	Destinations      celitechconfig.Config
	Packages          celitechconfig.Config
	Purchases         celitechconfig.Config
	ESim              celitechconfig.Config
	IFrame            celitechconfig.Config
	oAuthTokenManager *oauthtokenmanager.OAuthTokenManager
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config celitechconfig.Config, tokenService oauthtokenmanager.TokenService) *ConfigManager {
	return &ConfigManager{
		OAuth:             config,
		Destinations:      config,
		Packages:          config,
		Purchases:         config,
		ESim:              config,
		IFrame:            config,
		oAuthTokenManager: oauthtokenmanager.NewOAuthTokenManager(tokenService, 5000),
	}
}

// SetBaseUrl updates the BaseUrl configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.OAuth.SetBaseUrl(baseUrl)
	c.Destinations.SetBaseUrl(baseUrl)
	c.Packages.SetBaseUrl(baseUrl)
	c.Purchases.SetBaseUrl(baseUrl)
	c.ESim.SetBaseUrl(baseUrl)
	c.IFrame.SetBaseUrl(baseUrl)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.OAuth.SetTimeout(timeout)
	c.Destinations.SetTimeout(timeout)
	c.Packages.SetTimeout(timeout)
	c.Purchases.SetTimeout(timeout)
	c.ESim.SetTimeout(timeout)
	c.IFrame.SetTimeout(timeout)
}

// SetClientId updates the ClientId configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetClientId(clientId string) {
	c.OAuth.SetClientId(clientId)
	c.Destinations.SetClientId(clientId)
	c.Packages.SetClientId(clientId)
	c.Purchases.SetClientId(clientId)
	c.ESim.SetClientId(clientId)
	c.IFrame.SetClientId(clientId)
}

// SetClientSecret updates the ClientSecret configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetClientSecret(clientSecret string) {
	c.OAuth.SetClientSecret(clientSecret)
	c.Destinations.SetClientSecret(clientSecret)
	c.Packages.SetClientSecret(clientSecret)
	c.Purchases.SetClientSecret(clientSecret)
	c.ESim.SetClientSecret(clientSecret)
	c.IFrame.SetClientSecret(clientSecret)
}

// SetOAuthBaseUrl updates the OAuthBaseUrl configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetOAuthBaseUrl(oAuthBaseUrl string) {
	c.OAuth.SetOAuthBaseUrl(oAuthBaseUrl)
	c.Destinations.SetOAuthBaseUrl(oAuthBaseUrl)
	c.Packages.SetOAuthBaseUrl(oAuthBaseUrl)
	c.Purchases.SetOAuthBaseUrl(oAuthBaseUrl)
	c.ESim.SetOAuthBaseUrl(oAuthBaseUrl)
	c.IFrame.SetOAuthBaseUrl(oAuthBaseUrl)
}

// GetTokenManager returns the OAuth token manager for handling access token operations.
// Used by OAuth handlers to fetch and refresh tokens.
func (c *ConfigManager) GetTokenManager() *oauthtokenmanager.OAuthTokenManager {
	return c.oAuthTokenManager
}

// GetOAuth returns the configuration for the OAuth service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOAuth() *celitechconfig.Config {
	return &c.OAuth
}

// GetDestinations returns the configuration for the Destinations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDestinations() *celitechconfig.Config {
	return &c.Destinations
}

// GetPackages returns the configuration for the Packages service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPackages() *celitechconfig.Config {
	return &c.Packages
}

// GetPurchases returns the configuration for the Purchases service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPurchases() *celitechconfig.Config {
	return &c.Purchases
}

// GetESim returns the configuration for the ESim service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetESim() *celitechconfig.Config {
	return &c.ESim
}

// GetIFrame returns the configuration for the IFrame service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetIFrame() *celitechconfig.Config {
	return &c.IFrame
}
