package configmanager

import (
	"example.com/celitech/celitechconfig"
	"example.com/celitech/internal/oauthtokenmanager"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	oAuth             celitechconfig.Config
	destinations      celitechconfig.Config
	packages          celitechconfig.Config
	purchases         celitechconfig.Config
	eSim              celitechconfig.Config
	iFrame            celitechconfig.Config
	oAuthTokenManager *oauthtokenmanager.OAuthTokenManager
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config celitechconfig.Config, tokenService oauthtokenmanager.TokenService) *ConfigManager {
	return &ConfigManager{
		oAuth:             config,
		destinations:      config,
		packages:          config,
		purchases:         config,
		eSim:              config,
		iFrame:            config,
		oAuthTokenManager: oauthtokenmanager.NewOAuthTokenManager(tokenService, 0),
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.oAuth.SetBaseURL(baseURL)
	c.destinations.SetBaseURL(baseURL)
	c.packages.SetBaseURL(baseURL)
	c.purchases.SetBaseURL(baseURL)
	c.eSim.SetBaseURL(baseURL)
	c.iFrame.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.oAuth.SetTimeout(timeout)
	c.destinations.SetTimeout(timeout)
	c.packages.SetTimeout(timeout)
	c.purchases.SetTimeout(timeout)
	c.eSim.SetTimeout(timeout)
	c.iFrame.SetTimeout(timeout)
}

// SetClientID updates the ClientID configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetClientID(clientID string) {
	c.oAuth.SetClientID(clientID)
	c.destinations.SetClientID(clientID)
	c.packages.SetClientID(clientID)
	c.purchases.SetClientID(clientID)
	c.eSim.SetClientID(clientID)
	c.iFrame.SetClientID(clientID)
}

// SetClientSecret updates the ClientSecret configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetClientSecret(clientSecret string) {
	c.oAuth.SetClientSecret(clientSecret)
	c.destinations.SetClientSecret(clientSecret)
	c.packages.SetClientSecret(clientSecret)
	c.purchases.SetClientSecret(clientSecret)
	c.eSim.SetClientSecret(clientSecret)
	c.iFrame.SetClientSecret(clientSecret)
}

// SetOAuthBaseURL updates the OAuthBaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetOAuthBaseURL(oAuthBaseURL string) {
	c.oAuth.SetOAuthBaseURL(oAuthBaseURL)
	c.destinations.SetOAuthBaseURL(oAuthBaseURL)
	c.packages.SetOAuthBaseURL(oAuthBaseURL)
	c.purchases.SetOAuthBaseURL(oAuthBaseURL)
	c.eSim.SetOAuthBaseURL(oAuthBaseURL)
	c.iFrame.SetOAuthBaseURL(oAuthBaseURL)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry celitechconfig.RetryConfig) {
	c.oAuth.SetRetryConfig(retry)
	c.destinations.SetRetryConfig(retry)
	c.packages.SetRetryConfig(retry)
	c.purchases.SetRetryConfig(retry)
	c.eSim.SetRetryConfig(retry)
	c.iFrame.SetRetryConfig(retry)
}

// GetTokenManager returns the OAuth token manager for handling access token operations.
// Used by OAuth handlers to fetch and refresh tokens.
func (c *ConfigManager) GetTokenManager() *oauthtokenmanager.OAuthTokenManager {
	return c.oAuthTokenManager
}

// GetOAuth returns the configuration for the OAuth service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOAuth() *celitechconfig.Config {
	return &c.oAuth
}

// GetDestinations returns the configuration for the Destinations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDestinations() *celitechconfig.Config {
	return &c.destinations
}

// GetPackages returns the configuration for the Packages service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPackages() *celitechconfig.Config {
	return &c.packages
}

// GetPurchases returns the configuration for the Purchases service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPurchases() *celitechconfig.Config {
	return &c.purchases
}

// GetESim returns the configuration for the ESim service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetESim() *celitechconfig.Config {
	return &c.eSim
}

// GetIFrame returns the configuration for the IFrame service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetIFrame() *celitechconfig.Config {
	return &c.iFrame
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.oAuth.BaseURL
}
