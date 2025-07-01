package destinations

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type DestinationsService struct {
	manager *configmanager.ConfigManager
}

func NewDestinationsService() *DestinationsService {
	return &DestinationsService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *DestinationsService) WithConfigManager(manager *configmanager.ConfigManager) *DestinationsService {
	api.manager = manager
	return api
}

func (api *DestinationsService) getConfig() *celitechconfig.Config {
	return api.manager.GetDestinations()
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

// List Destinations
func (api *DestinationsService) ListDestinations(ctx context.Context) (*shared.CelitechResponse[ListDestinationsOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/destinations").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[ListDestinationsOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[ListDestinationsOkResponse](err)
	}

	return shared.NewCelitechResponse[ListDestinationsOkResponse](resp), nil
}
