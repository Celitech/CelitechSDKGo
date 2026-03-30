package purchases

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

// PurchasesService provides methods to interact with PurchasesService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type PurchasesService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewPurchasesService() *PurchasesService {
	return &PurchasesService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *PurchasesService) WithConfigManager(manager *configmanager.ConfigManager) *PurchasesService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *PurchasesService) WithHook(hook hooks.Hook) *PurchasesService {
	api.hook = hook
	return api
}

func (api *PurchasesService) getConfig() *celitechconfig.Config {
	return api.manager.GetPurchases()
}

func (api *PurchasesService) getHook() hooks.Hook {
	return api.hook
}

func (api *PurchasesService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *PurchasesService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *PurchasesService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *PurchasesService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

func (api *PurchasesService) SetOAuthBaseUrl(oAuthBaseUrl string) {
	config := api.getConfig()
	config.SetOAuthBaseUrl(oAuthBaseUrl)
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *PurchasesService) CreatePurchaseV2(ctx context.Context, createPurchaseV2Request CreatePurchaseV2Request) (*shared.CelitechResponse[[]CreatePurchaseV2OkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/v2").
		WithConfig(config).
		WithBody(createPurchaseV2Request).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[[]CreatePurchaseV2OkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[[]CreatePurchaseV2OkResponse](resp), nil
}

// This endpoint can be used to list all the successful purchases made between a given interval.
func (api *PurchasesService) ListPurchases(ctx context.Context, params ListPurchasesRequestParams) (*shared.CelitechResponse[ListPurchasesOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[ListPurchasesOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[ListPurchasesOkResponse](resp), nil
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *PurchasesService) CreatePurchase(ctx context.Context, createPurchaseRequest CreatePurchaseRequest) (*shared.CelitechResponse[CreatePurchaseOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases").
		WithConfig(config).
		WithBody(createPurchaseRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[CreatePurchaseOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[CreatePurchaseOkResponse](resp), nil
}

// This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.
func (api *PurchasesService) TopUpEsim(ctx context.Context, topUpEsimRequest TopUpEsimRequest) (*shared.CelitechResponse[TopUpEsimOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/topup").
		WithConfig(config).
		WithBody(topUpEsimRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[TopUpEsimOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[TopUpEsimOkResponse](resp), nil
}

// This endpoint allows you to modify the validity dates of an existing purchase.
//
// **Behavior:**
// - If the purchase has **not yet been activated**, both the start and end dates can be updated.
// - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).
// - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.
//
// The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.
func (api *PurchasesService) EditPurchase(ctx context.Context, editPurchaseRequest EditPurchaseRequest) (*shared.CelitechResponse[EditPurchaseOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/purchases/edit").
		WithConfig(config).
		WithBody(editPurchaseRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[EditPurchaseOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[EditPurchaseOkResponse](resp), nil
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (api *PurchasesService) GetPurchaseConsumption(ctx context.Context, purchaseId string) (*shared.CelitechResponse[GetPurchaseConsumptionOkResponse], *shared.CelitechError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/purchases/{purchaseId}/consumption").
		WithConfig(config).
		AddPathParam("purchaseId", purchaseId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes([]string{}).
		Build()

	client := restClient.NewRestClient[GetPurchaseConsumptionOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]byte](err)
	}

	return shared.NewCelitechResponse[GetPurchaseConsumptionOkResponse](resp), nil
}
