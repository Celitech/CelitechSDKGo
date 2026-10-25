package esim

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

// Get eSIM
func (c *Client) GetEsim(ctx context.Context, request *api.GetEsimRequest, opts ...option.RequestOption) (*api.GetEsimResponse, error) {
	response, err := c.WithRawResponse.GetEsim(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get eSIM Device
func (c *Client) GetEsimDevice(ctx context.Context, request *api.GetEsimDeviceRequest, opts ...option.RequestOption) (*api.GetEsimDeviceResponse, error) {
	response, err := c.WithRawResponse.GetEsimDevice(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get eSIM History
func (c *Client) GetEsimHistory(ctx context.Context, request *api.GetEsimHistoryRequest, opts ...option.RequestOption) (*api.GetEsimHistoryResponse, error) {
	response, err := c.WithRawResponse.GetEsimHistory(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Suspends an eSIM to prevent abuse. All of the eSIM's packages are terminated and its connectivity is switched off. The eSIM is not refunded: the charge stays.
func (c *Client) SuspendEsim(ctx context.Context, request *api.SuspendEsimRequest, opts ...option.RequestOption) (*api.SuspendEsimResponse, error) {
	response, err := c.WithRawResponse.SuspendEsim(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
