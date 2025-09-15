package esim

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type ESimService struct {
	manager *configmanager.ConfigManager
}

func NewESimService() *ESimService {
	return &ESimService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *ESimService) WithConfigManager(manager *configmanager.ConfigManager) *ESimService {
	api.manager = manager
	return api
}

func (api *ESimService) getConfig() *celitechconfig.Config {
	return api.manager.GetESim()
}

func (api *ESimService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *ESimService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *ESimService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *ESimService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

func (api *ESimService) SetOAuthBaseUrl(oAuthBaseUrl string) {
	config := api.getConfig()
	config.SetOAuthBaseUrl(oAuthBaseUrl)
}

// Get eSIM Status
func (api *ESimService) GetEsim(ctx context.Context, params GetEsimRequestParams) (*shared.CelitechResponse[GetEsimOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[GetEsimOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetEsimOkResponse](err)
	}

	return shared.NewCelitechResponse[GetEsimOkResponse](resp), nil
}

// Get eSIM Device
func (api *ESimService) GetEsimDevice(ctx context.Context, iccid string) (*shared.CelitechResponse[GetEsimDeviceOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim/{iccid}/device").
		WithConfig(config).
		AddPathParam("iccid", iccid).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[GetEsimDeviceOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetEsimDeviceOkResponse](err)
	}

	return shared.NewCelitechResponse[GetEsimDeviceOkResponse](resp), nil
}

// Get eSIM History
func (api *ESimService) GetEsimHistory(ctx context.Context, iccid string) (*shared.CelitechResponse[GetEsimHistoryOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim/{iccid}/history").
		WithConfig(config).
		AddPathParam("iccid", iccid).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[GetEsimHistoryOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetEsimHistoryOkResponse](err)
	}

	return shared.NewCelitechResponse[GetEsimHistoryOkResponse](resp), nil
}

// Get eSIM MAC
func (api *ESimService) GetEsimMac(ctx context.Context, iccid string) (*shared.CelitechResponse[GetEsimMacOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim/{iccid}/mac").
		WithConfig(config).
		AddPathParam("iccid", iccid).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[GetEsimMacOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetEsimMacOkResponse](err)
	}

	return shared.NewCelitechResponse[GetEsimMacOkResponse](resp), nil
}
