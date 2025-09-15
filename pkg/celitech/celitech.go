package celitech

import (
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/destinations"
	"github.com/Celitech/CelitechSDKGo/pkg/esim"
	"github.com/Celitech/CelitechSDKGo/pkg/iframe"
	"github.com/Celitech/CelitechSDKGo/pkg/oauth"
	"github.com/Celitech/CelitechSDKGo/pkg/packages"
	"github.com/Celitech/CelitechSDKGo/pkg/purchases"
	"time"
)

type Celitech struct {
	Destinations *destinations.DestinationsService
	Packages     *packages.PackagesService
	Purchases    *purchases.PurchasesService
	ESim         *esim.ESimService
	IFrame       *iframe.IFrameService
	manager      *configmanager.ConfigManager
}

func NewCelitech(config celitechconfig.Config) *Celitech {
	oAuth := oauth.NewOAuthService()
	oAuth.SetBaseUrl("https://auth.celitech.net")
	destinations := destinations.NewDestinationsService()
	packages := packages.NewPackagesService()
	purchases := purchases.NewPurchasesService()
	eSim := esim.NewESimService()
	iFrame := iframe.NewIFrameService()

	manager := configmanager.NewConfigManager(config, oAuth)
	oAuth.WithConfigManager(manager)
	destinations.WithConfigManager(manager)
	packages.WithConfigManager(manager)
	purchases.WithConfigManager(manager)
	eSim.WithConfigManager(manager)
	iFrame.WithConfigManager(manager)

	return &Celitech{
		Destinations: destinations,
		Packages:     packages,
		Purchases:    purchases,
		ESim:         eSim,
		IFrame:       iFrame,
		manager:      manager,
	}
}

func (c *Celitech) SetBaseUrl(baseUrl string) {
	c.manager.SetBaseUrl(baseUrl)
}

func (c *Celitech) SetTimeout(timeout time.Duration) {
	c.manager.SetTimeout(timeout)
}

func (c *Celitech) SetClientId(clientId string) {
	c.manager.SetClientId(clientId)
}

func (c *Celitech) SetClientSecret(clientSecret string) {
	c.manager.SetClientSecret(clientSecret)
}

func (c *Celitech) SetOAuthBaseUrl(oAuthBaseUrl string) {
	c.manager.SetOAuthBaseUrl(oAuthBaseUrl)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
