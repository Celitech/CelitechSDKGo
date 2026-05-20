package destinations

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

// Service provides methods to interact with Destinations-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                *configmanager.ConfigManager
	hook                   hooks.Hook
	listDestinationsConfig []celitechconfig.RequestOption
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
	return api.manager.GetDestinations()
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

// SetListDestinationsConfig sets method-level configuration for ListDestinations.
// Options are applied to every future call to ListDestinations and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListDestinationsConfig(opts ...celitechconfig.RequestOption) *Service {
	api.listDestinationsConfig = opts
	return api
}

// List Destinations
func (api *Service) ListDestinations(ctx context.Context, opts ...celitechconfig.RequestOption) (*ListDestinationsOkResponse, error) {
	config := *api.config()
	for _, opt := range api.listDestinationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/destinations").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[ListDestinationsOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}
