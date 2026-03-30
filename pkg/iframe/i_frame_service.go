package iframe

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

// IFrameService provides methods to interact with IFrameService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type IFrameService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewIFrameService() *IFrameService {
	return &IFrameService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *IFrameService) WithConfigManager(manager *configmanager.ConfigManager) *IFrameService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *IFrameService) WithHook(hook hooks.Hook) *IFrameService {
	api.hook = hook
	return api
}

func (api *IFrameService) getConfig() *celitechconfig.Config {
	return api.manager.GetIFrame()
}

func (api *IFrameService) getHook() hooks.Hook {
	return api.hook
}

func (api *IFrameService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *IFrameService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *IFrameService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *IFrameService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

func (api *IFrameService) SetOAuthBaseUrl(oAuthBaseUrl string) {
	config := api.getConfig()
	config.SetOAuthBaseUrl(oAuthBaseUrl)
}

// Generate a new token to be used in the iFrame
func (api *IFrameService) Token(ctx context.Context) (*shared.CelitechResponse[TokenOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/iframe/token").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[TokenOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[TokenOkResponse](resp), nil
}
