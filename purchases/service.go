package purchases

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

// Service provides methods to interact with Purchases-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                      *configmanager.ConfigManager
	hook                         hooks.Hook
	createPurchaseV2Config       []celitechconfig.RequestOption
	listPurchasesConfig          []celitechconfig.RequestOption
	createPurchaseConfig         []celitechconfig.RequestOption
	topUpEsimConfig              []celitechconfig.RequestOption
	editPurchaseConfig           []celitechconfig.RequestOption
	getPurchaseConsumptionConfig []celitechconfig.RequestOption
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
	return api.manager.GetPurchases()
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

// SetCreatePurchaseV2Config sets method-level configuration for CreatePurchaseV2.
// Options are applied to every future call to CreatePurchaseV2 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreatePurchaseV2Config(opts ...celitechconfig.RequestOption) *Service {
	api.createPurchaseV2Config = opts
	return api
}

// SetListPurchasesConfig sets method-level configuration for ListPurchases.
// Options are applied to every future call to ListPurchases and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListPurchasesConfig(opts ...celitechconfig.RequestOption) *Service {
	api.listPurchasesConfig = opts
	return api
}

// SetCreatePurchaseConfig sets method-level configuration for CreatePurchase.
// Options are applied to every future call to CreatePurchase and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreatePurchaseConfig(opts ...celitechconfig.RequestOption) *Service {
	api.createPurchaseConfig = opts
	return api
}

// SetTopUpEsimConfig sets method-level configuration for TopUpEsim.
// Options are applied to every future call to TopUpEsim and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetTopUpEsimConfig(opts ...celitechconfig.RequestOption) *Service {
	api.topUpEsimConfig = opts
	return api
}

// SetEditPurchaseConfig sets method-level configuration for EditPurchase.
// Options are applied to every future call to EditPurchase and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditPurchaseConfig(opts ...celitechconfig.RequestOption) *Service {
	api.editPurchaseConfig = opts
	return api
}

// SetGetPurchaseConsumptionConfig sets method-level configuration for GetPurchaseConsumption.
// Options are applied to every future call to GetPurchaseConsumption and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetPurchaseConsumptionConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getPurchaseConsumptionConfig = opts
	return api
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *Service) CreatePurchaseV2(ctx context.Context, createPurchaseV2Request CreatePurchaseV2Request, opts ...celitechconfig.RequestOption) ([]CreatePurchaseV2OkResponse, error) {
	config := *api.config()
	for _, opt := range api.createPurchaseV2Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/v2").
		WithConfig(config).
		WithBody(createPurchaseV2Request).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[[]CreatePurchaseV2OkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return resp.Data, nil
}

// This endpoint can be used to list all the successful purchases made between a given interval.
func (api *Service) ListPurchases(ctx context.Context, params ListPurchasesRequestParams, opts ...celitechconfig.RequestOption) (*ListPurchasesOkResponse, error) {
	config := *api.config()
	for _, opt := range api.listPurchasesConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[ListPurchasesOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *Service) CreatePurchase(ctx context.Context, createPurchaseRequest CreatePurchaseRequest, opts ...celitechconfig.RequestOption) (*CreatePurchaseOkResponse, error) {
	config := *api.config()
	for _, opt := range api.createPurchaseConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases").
		WithConfig(config).
		WithBody(createPurchaseRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[CreatePurchaseOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.
func (api *Service) TopUpEsim(ctx context.Context, topUpEsimRequest TopUpEsimRequest, opts ...celitechconfig.RequestOption) (*TopUpEsimOkResponse, error) {
	config := *api.config()
	for _, opt := range api.topUpEsimConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/topup").
		WithConfig(config).
		WithBody(topUpEsimRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[TopUpEsimOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// This endpoint allows you to modify the validity dates of an existing purchase.
//
// **Behavior:**
// - If the purchase has **not yet been activated**, both the start and end dates can be updated.
// - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).
// - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.
//
// The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.
func (api *Service) EditPurchase(ctx context.Context, editPurchaseRequest EditPurchaseRequest, opts ...celitechconfig.RequestOption) (*EditPurchaseOkResponse, error) {
	config := *api.config()
	for _, opt := range api.editPurchaseConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/edit").
		WithConfig(config).
		WithBody(editPurchaseRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[EditPurchaseOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (api *Service) GetPurchaseConsumption(ctx context.Context, purchaseID string, opts ...celitechconfig.RequestOption) (*GetPurchaseConsumptionOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getPurchaseConsumptionConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases/{purchaseId}/consumption").
		WithConfig(config).
		AddPathParam("purchaseId", purchaseID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes([]string{}).
		Build()

	httpClient := restClient.NewRestClient[GetPurchaseConsumptionOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}
