package packages

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

// PackagesService provides methods to interact with PackagesService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type PackagesService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewPackagesService() *PackagesService {
	return &PackagesService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *PackagesService) WithConfigManager(manager *configmanager.ConfigManager) *PackagesService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *PackagesService) WithHook(hook hooks.Hook) *PackagesService {
	api.hook = hook
	return api
}

func (api *PackagesService) getConfig() *celitechconfig.Config {
	return api.manager.GetPackages()
}

func (api *PackagesService) getHook() hooks.Hook {
	return api.hook
}

func (api *PackagesService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *PackagesService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *PackagesService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *PackagesService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

func (api *PackagesService) SetOAuthBaseUrl(oAuthBaseUrl string) {
	config := api.getConfig()
	config.SetOAuthBaseUrl(oAuthBaseUrl)
}

// List Packages
func (api *PackagesService) ListPackages(ctx context.Context, params ListPackagesRequestParams) (*shared.CelitechResponse[ListPackagesOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/packages").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[ListPackagesOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[ListPackagesOkResponse](resp), nil
}
