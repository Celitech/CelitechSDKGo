package oauth

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

func (r *RawClient) GetAccessToken(ctx context.Context, oAuthTokenRequest *api.OAuthTokenRequest, opts ...option.RequestOption) (*core.Response[*api.OAuthTokenResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.celitech.net/v1",
	)
	endpointURL := baseURL + "/oauth2/token"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	var response api.OAuthTokenResponse
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
			Request:         oAuthTokenRequest,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*api.OAuthTokenResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
