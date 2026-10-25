package client

import (
	"context"
	"errors"
	api "github.com/Celitech/CelitechSDKGo/v3"
	"github.com/Celitech/CelitechSDKGo/v3/core"
	"github.com/Celitech/CelitechSDKGo/v3/destinations"
	"github.com/Celitech/CelitechSDKGo/v3/esim"
	"github.com/Celitech/CelitechSDKGo/v3/iframe"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/oauth"
	"github.com/Celitech/CelitechSDKGo/v3/option"
	"github.com/Celitech/CelitechSDKGo/v3/packages"
	"github.com/Celitech/CelitechSDKGo/v3/purchases"
)

type Celitech struct {
	Destinations *destinations.Client
	Packages     *packages.Client
	Purchases    *purchases.Client
	ESim         *esim.Client
	IFrame       *iframe.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewCelitech(opts ...option.RequestOption) *Celitech {
	options := core.NewRequestOptions(opts...)
	oauthTokenProvider := core.NewTokenProvider(0)
	authOptions := *options
	authClient := oauth.NewClient(&authOptions)
	options.SetTokenGetter(func() (string, error) {
		return oauthTokenProvider.GetOrFetch(func() (string, int64, error) {
			response, err := authClient.GetAccessToken(context.Background(), &api.OAuthTokenRequest{
				ClientID:     options.ClientID,
				ClientSecret: options.ClientSecret,
				GrantType:    "client_credentials",
			})
			if err != nil {
				return "", 0, err
			}
			if response.AccessToken == nil {
				return "", 0, errors.New("oauth response missing access token")
			}
			expiresIn := int64(core.DefaultExpirySeconds)
			if response.ExpiresIn != nil {
				expiresIn = int64(*response.ExpiresIn)
			}
			return *response.AccessToken, expiresIn, nil
		})
	})
	return &Celitech{
		Destinations: destinations.NewClient(options),
		Packages:     packages.NewClient(options),
		Purchases:    purchases.NewClient(options),
		ESim:         esim.NewClient(options),
		IFrame:       iframe.NewClient(options),
		options:      options,
		baseURL:      options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
