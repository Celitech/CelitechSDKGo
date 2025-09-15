package purchases

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type PurchasesService struct {
	manager *configmanager.ConfigManager
}

func NewPurchasesService() *PurchasesService {
	return &PurchasesService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *PurchasesService) WithConfigManager(manager *configmanager.ConfigManager) *PurchasesService {
	api.manager = manager
	return api
}

func (api *PurchasesService) getConfig() *celitechconfig.Config {
	return api.manager.GetPurchases()
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
func (api *PurchasesService) CreatePurchaseV2(ctx context.Context, createPurchaseV2Request CreatePurchaseV2Request) (*shared.CelitechResponse[[]CreatePurchaseV2OkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[[]CreatePurchaseV2OkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[[]CreatePurchaseV2OkResponse](err)
	}

	return shared.NewCelitechResponse[[]CreatePurchaseV2OkResponse](resp), nil
}

// This endpoint can be used to list all the successful purchases made between a given interval.
func (api *PurchasesService) ListPurchases(ctx context.Context, params ListPurchasesRequestParams) (*shared.CelitechResponse[ListPurchasesOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[ListPurchasesOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[ListPurchasesOkResponse](err)
	}

	return shared.NewCelitechResponse[ListPurchasesOkResponse](resp), nil
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (api *PurchasesService) CreatePurchase(ctx context.Context, createPurchaseRequest CreatePurchaseRequest) (*shared.CelitechResponse[CreatePurchaseOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[CreatePurchaseOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[CreatePurchaseOkResponse](err)
	}

	return shared.NewCelitechResponse[CreatePurchaseOkResponse](resp), nil
}

// This endpoint is used to top-up an eSIM with the previously associated destination by providing an existing ICCID and the package details. The top-up is only feasible for eSIMs in "ENABLED" or "INSTALLED" state. You can check this state using the Get eSIM Status endpoint.
func (api *PurchasesService) TopUpEsim(ctx context.Context, topUpEsimRequest TopUpEsimRequest) (*shared.CelitechResponse[TopUpEsimOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[TopUpEsimOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[TopUpEsimOkResponse](err)
	}

	return shared.NewCelitechResponse[TopUpEsimOkResponse](resp), nil
}

// This endpoint allows you to modify the dates of an existing package with a future activation start time. Editing can only be performed for packages that have not been activated, and it cannot change the package size. The modification must not change the package duration category to ensure pricing consistency.
func (api *PurchasesService) EditPurchase(ctx context.Context, editPurchaseRequest EditPurchaseRequest) (*shared.CelitechResponse[EditPurchaseOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[EditPurchaseOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[EditPurchaseOkResponse](err)
	}

	return shared.NewCelitechResponse[EditPurchaseOkResponse](resp), nil
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (api *PurchasesService) GetPurchaseConsumption(ctx context.Context, purchaseId string) (*shared.CelitechResponse[GetPurchaseConsumptionOkResponse], *shared.CelitechError) {
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

	client := restClient.NewRestClient[GetPurchaseConsumptionOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetPurchaseConsumptionOkResponse](err)
	}

	return shared.NewCelitechResponse[GetPurchaseConsumptionOkResponse](resp), nil
}
