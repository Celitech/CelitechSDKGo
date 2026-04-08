package esim

import (
	"context"
	"example.com/celitech/celitechconfig"
	restClient "example.com/celitech/internal/clients/rest"
	"example.com/celitech/internal/clients/rest/hooks"
	"example.com/celitech/internal/clients/rest/httptransport"
	"example.com/celitech/internal/configmanager"
	"example.com/celitech/sharedmodels"
	"time"
)

// Service provides methods to interact with ESim-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager              *configmanager.ConfigManager
	hook                 hooks.Hook
	getEsimConfig        []celitechconfig.RequestOption
	getEsimDeviceConfig  []celitechconfig.RequestOption
	getEsimHistoryConfig []celitechconfig.RequestOption
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
	return api.manager.GetESim()
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

// SetGetEsimConfig sets method-level configuration for GetEsim.
// Options are applied to every future call to GetEsim and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEsimConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getEsimConfig = opts
	return api
}

// SetGetEsimDeviceConfig sets method-level configuration for GetEsimDevice.
// Options are applied to every future call to GetEsimDevice and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEsimDeviceConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getEsimDeviceConfig = opts
	return api
}

// SetGetEsimHistoryConfig sets method-level configuration for GetEsimHistory.
// Options are applied to every future call to GetEsimHistory and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEsimHistoryConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getEsimHistoryConfig = opts
	return api
}

// Get eSIM
func (api *Service) GetEsim(ctx context.Context, params GetEsimRequestParams, opts ...celitechconfig.RequestOption) (*GetEsimOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getEsimConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[GetEsimOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get eSIM Device
func (api *Service) GetEsimDevice(ctx context.Context, iccid string, opts ...celitechconfig.RequestOption) (*GetEsimDeviceOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getEsimDeviceConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim/{iccid}/device").
		WithConfig(config).
		AddPathParam("iccid", iccid).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[GetEsimDeviceOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get eSIM History
func (api *Service) GetEsimHistory(ctx context.Context, iccid string, opts ...celitechconfig.RequestOption) (*GetEsimHistoryOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getEsimHistoryConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/esim/{iccid}/history").
		WithConfig(config).
		AddPathParam("iccid", iccid).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[GetEsimHistoryOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}
