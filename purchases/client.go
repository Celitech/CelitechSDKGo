package purchases

import (
	context "context"
	api "github.com/Celitech/CelitechSDKGo/v3"
	core "github.com/Celitech/CelitechSDKGo/v3/core"
	internal "github.com/Celitech/CelitechSDKGo/v3/internal"
	option "github.com/Celitech/CelitechSDKGo/v3/option"
)

type Client struct {
	WithRawResponse *RawClient

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		options:         options,
		baseURL:         options.BaseURL,
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
func (c *Client) CreatePurchaseV2(ctx context.Context, request *api.CreatePurchaseV2Request, opts ...option.RequestOption) ([]*api.CreatePurchaseV2ResponseItem, error) {
	response, err := c.WithRawResponse.CreatePurchaseV2(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// This endpoint can be used to list all the successful purchases made between a given interval.
func (c *Client) ListPurchases(ctx context.Context, request *api.ListPurchasesRequest, opts ...option.RequestOption) (*api.ListPurchasesResponse, error) {
	response, err := c.WithRawResponse.ListPurchases(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// This endpoint is used to purchase a new eSIM by providing the package details.
func (c *Client) CreatePurchase(ctx context.Context, request *api.CreatePurchaseRequest, opts ...option.RequestOption) (*api.CreatePurchaseResponse, error) {
	response, err := c.WithRawResponse.CreatePurchase(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.
func (c *Client) TopUpEsim(ctx context.Context, request *api.TopUpEsimRequest, opts ...option.RequestOption) (*api.TopUpEsimResponse, error) {
	response, err := c.WithRawResponse.TopUpEsim(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// This endpoint allows you to modify the validity dates of an existing purchase.
//
// **Behavior:**
// - If the purchase has **not yet been activated**, both the start and end dates can be updated.
// - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).
// - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.
//
// The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.
func (c *Client) EditPurchase(ctx context.Context, request *api.EditPurchaseRequest, opts ...option.RequestOption) (*api.EditPurchaseResponse, error) {
	response, err := c.WithRawResponse.EditPurchase(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.
func (c *Client) GetPurchaseConsumption(ctx context.Context, request *api.GetPurchaseConsumptionRequest, opts ...option.RequestOption) (*api.GetPurchaseConsumptionResponse, error) {
	response, err := c.WithRawResponse.GetPurchaseConsumption(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
