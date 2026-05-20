package consumption

import (
	"context"
	"github.com/Celitech/CelitechSDKGo/celitechconfig"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/sharedmodels"
	"time"
)

// Service provides methods to interact with Consumption-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                      *configmanager.ConfigManager
	hook                         hooks.Hook
	getPurchaseConsumptionConfig []celitechconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *celitechconfig.Config {
	return api.manager.GetConsumption()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetClientID(clientID string) {
	config := api.config()
	config.SetClientID(clientID)
}

func (api *Service) SetClientSecret(clientSecret string) {
	config := api.config()
	config.SetClientSecret(clientSecret)
}

func (api *Service) SetOAuthBaseURL(oAuthBaseURL string) {
	config := api.config()
	config.SetOAuthBaseURL(oAuthBaseURL)
}

// SetGetPurchaseConsumptionConfig sets method-level configuration for GetPurchaseConsumption.
// Options are applied to every future call to GetPurchaseConsumption and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetPurchaseConsumptionConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getPurchaseConsumptionConfig = opts
	return api
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (api *Service) GetPurchaseConsumption(ctx context.Context, purchaseID string, params GetPurchaseConsumptionRequestParams, opts ...celitechconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getPurchaseConsumptionConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases/{purchaseId}/consumption").
		WithConfig(config).
		AddPathParam("purchaseId", purchaseID).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return resp.Data, nil
}
