package destinations

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

// DestinationsService provides methods to interact with DestinationsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type DestinationsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewDestinationsService() *DestinationsService {
	return &DestinationsService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *DestinationsService) WithConfigManager(manager *configmanager.ConfigManager) *DestinationsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *DestinationsService) WithHook(hook hooks.Hook) *DestinationsService {
	api.hook = hook
	return api
}

func (api *DestinationsService) getConfig() *celitechconfig.Config {
	return api.manager.GetDestinations()
}

func (api *DestinationsService) getHook() hooks.Hook {
	return api.hook
}

func (api *DestinationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *DestinationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *DestinationsService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *DestinationsService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

func (api *DestinationsService) SetOAuthBaseUrl(oAuthBaseUrl string) {
	config := api.getConfig()
	config.SetOAuthBaseUrl(oAuthBaseUrl)
}

// List Destinations
func (api *DestinationsService) ListDestinations(ctx context.Context) (*shared.CelitechResponse[ListDestinationsOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/destinations").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[ListDestinationsOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[ListDestinationsOkResponse](resp), nil
}
