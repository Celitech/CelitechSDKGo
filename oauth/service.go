package oauth

import (
	"context"
	"github.com/Celitech/CelitechSDKGo/celitechconfig"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/internal/oauthtokenmanager"
	"github.com/Celitech/CelitechSDKGo/sharedmodels"
	"time"
)

// Service provides methods to interact with OAuth-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager              *configmanager.ConfigManager
	hook                 hooks.Hook
	getAccessTokenConfig []celitechconfig.RequestOption
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
	return api.manager.GetOAuth()
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

// SetGetAccessTokenConfig sets method-level configuration for GetAccessToken.
// Options are applied to every future call to GetAccessToken and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetAccessTokenConfig(opts ...celitechconfig.RequestOption) *Service {
	api.getAccessTokenConfig = opts
	return api
}

// This endpoint was added by liblab
func (api *Service) GetAccessToken(ctx context.Context, getAccessTokenRequest GetAccessTokenRequest, opts ...celitechconfig.RequestOption) (*GetAccessTokenOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getAccessTokenConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/oauth2/token").
		WithConfig(config).
		WithBody(getAccessTokenRequest).
		AddHeader("CONTENT-TYPE", "application/x-www-form-urlencoded").
		WithContentType(httptransport.ContentTypeFormUrlEncoded).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithScopes(nil).
		Build()

	httpClient := restClient.NewRestClient[GetAccessTokenOkResponse, []byte](config, api.manager, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewCelitechError[[]byte](err)
	}

	return &resp.Data, nil
}

func (api *Service) GetOAuthAccessToken(config celitechconfig.Config, scopes []string) (oauthtokenmanager.GetOAuthAccessTokenResponse, error) {
	oauthBaseURL := config.OAuthBaseURL
	if oauthBaseURL == "" {
		oauthBaseURL = config.BaseURL
	}
	response, err := api.GetAccessToken(
		context.Background(),
		GetAccessTokenRequest{
			GrantType:    getPointer(GrantTypeClientCredentials),
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
		},

		celitechconfig.WithBaseURL(oauthBaseURL),
	)

	if err != nil {
		return nil, err
	}
	return oauthtokenmanager.NewOAuthTokenResponse(response.AccessToken, response.ExpiresIn), nil
}

func getPointer[T any](value T) *T {
	return &value
}
