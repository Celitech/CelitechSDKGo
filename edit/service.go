package edit

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

// Service provides methods to interact with Edit-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager            *configmanager.ConfigManager
	hook               hooks.Hook
	editPurchaseConfig []celitechconfig.RequestOption
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
	return api.manager.GetEdit()
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

// SetEditPurchaseConfig sets method-level configuration for EditPurchase.
// Options are applied to every future call to EditPurchase and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditPurchaseConfig(opts ...celitechconfig.RequestOption) *Service {
	api.editPurchaseConfig = opts
	return api
}

// This endpoint allows you to modify the validity dates of an existing purchase.
//
// **Behavior:**
// - If the purchase has **not yet been activated**, both the start and end dates can be updated.
// - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).
// - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.
//
// The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.
func (api *Service) EditPurchase(ctx context.Context, editPurchaseRequest EditPurchaseRequest, params EditPurchaseRequestParams, opts ...celitechconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editPurchaseConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/edit").
		WithConfig(config).
		WithBody(editPurchaseRequest).
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
