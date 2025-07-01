package packages

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type PackagesService struct {
	manager *configmanager.ConfigManager
}

func NewPackagesService() *PackagesService {
	return &PackagesService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *PackagesService) WithConfigManager(manager *configmanager.ConfigManager) *PackagesService {
	api.manager = manager
	return api
}

func (api *PackagesService) getConfig() *celitechconfig.Config {
	return api.manager.GetPackages()
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

// List Packages
func (api *PackagesService) ListPackages(ctx context.Context, params ListPackagesRequestParams) (*shared.CelitechResponse[ListPackagesOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[ListPackagesOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[ListPackagesOkResponse](err)
	}

	return shared.NewCelitechResponse[ListPackagesOkResponse](resp), nil
}
