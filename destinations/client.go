package destinations

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

// List Destinations
func (c *Client) ListDestinations(ctx context.Context, opts ...option.RequestOption) (*api.ListDestinationsResponse, error) {
	response, err := c.WithRawResponse.ListDestinations(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
