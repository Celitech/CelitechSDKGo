package celitech

import (
	"github.com/Celitech/CelitechSDKGo/consumption"
	"github.com/Celitech/CelitechSDKGo/destinations"
	"github.com/Celitech/CelitechSDKGo/device"
	"github.com/Celitech/CelitechSDKGo/edit"
	"github.com/Celitech/CelitechSDKGo/esim"
	"github.com/Celitech/CelitechSDKGo/history"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/oauth"
	"github.com/Celitech/CelitechSDKGo/packages"
	"github.com/Celitech/CelitechSDKGo/purchases"
	"github.com/Celitech/CelitechSDKGo/token"
	"github.com/Celitech/CelitechSDKGo/topup"
	"github.com/Celitech/CelitechSDKGo/v2"
	"time"
)

// Celitech is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type Celitech struct {
	Destinations *destinations.Service
	Packages     *packages.Service
	V2           *v2.Service
	Topup        *topup.Service
	Edit         *edit.Service
	Consumption  *consumption.Service
	Purchases    *purchases.Service
	Device       *device.Service
	History      *history.Service
	Esim         *esim.Service
	Token        *token.Service

	manager *configmanager.ConfigManager
}

func NewCelitech(config Config) *Celitech {
	destinations := destinations.NewService()
	packages := packages.NewService()
	v2 := v2.NewService()
	topup := topup.NewService()
	edit := edit.NewService()
	consumption := consumption.NewService()
	purchases := purchases.NewService()
	device := device.NewService()
	history := history.NewService()
	esim := esim.NewService()
	token := token.NewService()
	oAuth := oauth.NewService()
	oAuth.SetBaseURL("https://auth.celitech.net")

	manager := configmanager.NewConfigManager(config, oAuth)
	hook := hooks.NewDefaultHook()
	destinations.WithConfigManager(manager)
	packages.WithConfigManager(manager)
	v2.WithConfigManager(manager)
	topup.WithConfigManager(manager)
	edit.WithConfigManager(manager)
	consumption.WithConfigManager(manager)
	purchases.WithConfigManager(manager)
	device.WithConfigManager(manager)
	history.WithConfigManager(manager)
	esim.WithConfigManager(manager)
	token.WithConfigManager(manager)
	oAuth.WithConfigManager(manager)
	destinations.WithHook(hook)
	packages.WithHook(hook)
	v2.WithHook(hook)
	topup.WithHook(hook)
	edit.WithHook(hook)
	consumption.WithHook(hook)
	purchases.WithHook(hook)
	device.WithHook(hook)
	history.WithHook(hook)
	esim.WithHook(hook)
	token.WithHook(hook)
	oAuth.WithHook(hook)

	return &Celitech{
		Destinations: destinations,
		Packages:     packages,
		V2:           v2,
		Topup:        topup,
		Edit:         edit,
		Consumption:  consumption,
		Purchases:    purchases,
		Device:       device,
		History:      history,
		Esim:         esim,
		Token:        token,
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
