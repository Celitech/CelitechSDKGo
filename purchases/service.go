package purchases

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

// Service provides methods to interact with Purchases-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager              *configmanager.ConfigManager
	hook                 hooks.Hook
	createPurchaseConfig []celitechconfig.RequestOption
	listPurchasesConfig  []celitechconfig.RequestOption
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
	return api.manager.GetPurchases()
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

// SetCreatePurchaseConfig sets method-level configuration for CreatePurchase.
// Options are applied to every future call to CreatePurchase and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreatePurchaseConfig(opts ...celitechconfig.RequestOption) *Service {
	api.createPurchaseConfig = opts
	return api
}

// SetListPurchasesConfig sets method-level configuration for ListPurchases.
// Options are applied to every future call to ListPurchases and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListPurchasesConfig(opts ...celitechconfig.RequestOption) *Service {
	api.listPurchasesConfig = opts
	return api
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *Service) CreatePurchase(ctx context.Context, createPurchaseRequest CreatePurchaseRequest, params CreatePurchaseRequestParams, opts ...celitechconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createPurchaseConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases").
		WithConfig(config).
		WithBody(createPurchaseRequest).
		AddHeader("CONTENT-TYPE", "application/json").
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

// This endpoint can be used to list all the successful purchases made between a given interval.
func (api *Service) ListPurchases(ctx context.Context, params ListPurchasesRequestParams, opts ...celitechconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.listPurchasesConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases").
		WithConfig(config).
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
