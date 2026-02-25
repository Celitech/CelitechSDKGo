package esim

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

// ESimService provides methods to interact with ESimService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type ESimService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewESimService() *ESimService {
	return &ESimService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *ESimService) WithConfigManager(manager *configmanager.ConfigManager) *ESimService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *ESimService) WithHook(hook hooks.Hook) *ESimService {
	api.hook = hook
	return api
}

func (api *ESimService) getConfig() *celitechconfig.Config {
	return api.manager.GetESim()
}

func (api *ESimService) getHook() hooks.Hook {
	return api.hook
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

// Get eSIM
func (api *ESimService) GetEsim(ctx context.Context, params GetEsimRequestParams) (*shared.CelitechResponse[GetEsimOkResponse], *shared.CelitechError[[]byte]) {
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

	client := restClient.NewRestClient[GetEsimOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[GetEsimOkResponse](resp), nil
}

// Get eSIM Device
func (api *ESimService) GetEsimDevice(ctx context.Context, iccid string) (*shared.CelitechResponse[GetEsimDeviceOkResponse], *shared.CelitechError[[]byte]) {
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

	client := restClient.NewRestClient[GetEsimDeviceOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[GetEsimDeviceOkResponse](resp), nil
}

// Get eSIM History
func (api *ESimService) GetEsimHistory(ctx context.Context, iccid string) (*shared.CelitechResponse[GetEsimHistoryOkResponse], *shared.CelitechError[[]byte]) {
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

	client := restClient.NewRestClient[GetEsimHistoryOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[GetEsimHistoryOkResponse](resp), nil
}
