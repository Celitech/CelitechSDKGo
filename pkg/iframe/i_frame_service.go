package iframe

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type IFrameService struct {
	manager *configmanager.ConfigManager
}

func NewIFrameService() *IFrameService {
	return &IFrameService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *IFrameService) WithConfigManager(manager *configmanager.ConfigManager) *IFrameService {
	api.manager = manager
	return api
}

func (api *IFrameService) getConfig() *celitechconfig.Config {
	return api.manager.GetIFrame()
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

// Generate a new token to be used in the iFrame
func (api *IFrameService) Token(ctx context.Context) (*shared.CelitechResponse[TokenOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/iframe/token").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[TokenOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[TokenOkResponse](err)
	}

	return shared.NewCelitechResponse[TokenOkResponse](resp), nil
}
