package purchases

import (
	context "context"
	api "github.com/Celitech/CelitechSDKGo/v3"
	core "github.com/Celitech/CelitechSDKGo/v3/core"
	internal "github.com/Celitech/CelitechSDKGo/v3/internal"
	option "github.com/Celitech/CelitechSDKGo/v3/option"
	http "net/http"
)

type RawClient struct {
	baseURL string
	caller  *internal.Caller
	options *core.RequestOptions
}

func NewRawClient(options *core.RequestOptions) *RawClient {
	return &RawClient{
		options: options,
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (r *RawClient) CreatePurchaseV2(ctx context.Context, request *api.CreatePurchaseV2Request, opts ...option.RequestOption) (*core.Response[[]*api.CreatePurchaseV2ResponseItem], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/purchases/v2"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &api.CreatePurchaseV2Request{}
	}
	internal.SetHeader(headers, "Idempotency-Key", request.IdempotencyKey)
	var response []*api.CreatePurchaseV2ResponseItem
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[[]*api.CreatePurchaseV2ResponseItem]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// This endpoint can be used to list all the successful purchases made between a given interval.
func (r *RawClient) ListPurchases(ctx context.Context, request *api.ListPurchasesRequest, opts ...option.RequestOption) (*core.Response[*api.ListPurchasesResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/purchases"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &api.ListPurchasesRequest{}
	}
	var response api.ListPurchasesResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.ListPurchasesResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (r *RawClient) CreatePurchase(ctx context.Context, request *api.CreatePurchaseRequest, opts ...option.RequestOption) (*core.Response[*api.CreatePurchaseResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/purchases"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &api.CreatePurchaseRequest{}
	}
	internal.SetHeader(headers, "Idempotency-Key", request.IdempotencyKey)
	var response api.CreatePurchaseResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.CreatePurchaseResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.
func (r *RawClient) TopUpEsim(ctx context.Context, request *api.TopUpEsimRequest, opts ...option.RequestOption) (*core.Response[*api.TopUpEsimResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/purchases/topup"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &api.TopUpEsimRequest{}
	}
	internal.SetHeader(headers, "Idempotency-Key", request.IdempotencyKey)
	var response api.TopUpEsimResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.TopUpEsimResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// This endpoint allows you to modify the validity dates of an existing purchase.
//
// **Behavior:**
// - If the purchase has **not yet been activated**, both the start and end dates can be updated.
// - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).
// - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.
//
// The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.
func (r *RawClient) EditPurchase(ctx context.Context, request *api.EditPurchaseRequest, opts ...option.RequestOption) (*core.Response[*api.EditPurchaseResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/purchases/edit"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &api.EditPurchaseRequest{}
	}
	internal.SetHeader(headers, "Idempotency-Key", request.IdempotencyKey)
	var response api.EditPurchaseResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.EditPurchaseResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (r *RawClient) GetPurchaseConsumption(ctx context.Context, request *api.GetPurchaseConsumptionRequest, opts ...option.RequestOption) (*core.Response[*api.GetPurchaseConsumptionResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := internal.EncodeURL(baseURL+"/purchases/%v/consumption", request.PurchaseID)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &api.GetPurchaseConsumptionRequest{}
	}
	var response api.GetPurchaseConsumptionResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.GetPurchaseConsumptionResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
