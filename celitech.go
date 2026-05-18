package celitech

import (
	"github.com/Celitech/CelitechSDKGo/destinations"
	"github.com/Celitech/CelitechSDKGo/esim"
	"github.com/Celitech/CelitechSDKGo/iframe"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/oauth"
	"github.com/Celitech/CelitechSDKGo/packages"
	"github.com/Celitech/CelitechSDKGo/purchases"
	"time"
)

// Celitech is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type Celitech struct {
	Destinations *destinations.Service
	Packages     *packages.Service
	Purchases    *purchases.Service
	ESim         *esim.Service
	IFrame       *iframe.Service
	manager      *configmanager.ConfigManager
}

func NewCelitech(config Config) *Celitech {
	oAuth := oauth.NewService()
	oAuth.SetBaseURL("https://auth.celitech.net")
	destinations := destinations.NewService()
	packages := packages.NewService()
	purchases := purchases.NewService()
	eSim := esim.NewService()
	iFrame := iframe.NewService()

	manager := configmanager.NewConfigManager(config, oAuth)
	hook := hooks.NewDefaultHook()
	oAuth.WithConfigManager(manager)
	destinations.WithConfigManager(manager)
	packages.WithConfigManager(manager)
	purchases.WithConfigManager(manager)
	eSim.WithConfigManager(manager)
	iFrame.WithConfigManager(manager)
	oAuth.WithHook(hook)
	destinations.WithHook(hook)
	packages.WithHook(hook)
	purchases.WithHook(hook)
	eSim.WithHook(hook)
	iFrame.WithHook(hook)

	return &Celitech{
		Destinations: destinations,
		Packages:     packages,
		Purchases:    purchases,
		ESim:         eSim,
		IFrame:       iFrame,
		manager:      manager,
	}
}

func (c *Celitech) SetBaseURL(baseURL string) {
	c.manager.SetBaseURL(baseURL)
}

func (c *Celitech) SetTimeout(timeout time.Duration) {
	c.manager.SetTimeout(timeout)
}

func (c *Celitech) SetClientID(clientID string) {
	c.manager.SetClientID(clientID)
}

func (c *Celitech) SetClientSecret(clientSecret string) {
	c.manager.SetClientSecret(clientSecret)
}

func (c *Celitech) SetOAuthBaseURL(oAuthBaseURL string) {
	c.manager.SetOAuthBaseURL(oAuthBaseURL)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (c *Celitech) SetEnvironment(environment Environment) {
	c.manager.SetBaseURL(string(environment))
}

// SetTokenCache attaches a persistent token cache to the SDK.
func (c *Celitech) SetTokenCache(cache TokenCache) {
	c.manager.SetTokenCache(cache)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
