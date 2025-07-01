package oauth

import (
	"context"
	restClient "github.com/Celitech/CelitechSDKGo/internal/clients/rest"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/internal/oauthtokenmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/shared"
	"time"
)

type OAuthService struct {
	manager *configmanager.ConfigManager
}

func NewOAuthService() *OAuthService {
	return &OAuthService{
		manager: configmanager.NewConfigManager(celitechconfig.Config{}, nil),
	}
}

func (api *OAuthService) WithConfigManager(manager *configmanager.ConfigManager) *OAuthService {
	api.manager = manager
	return api
}

func (api *OAuthService) getConfig() *celitechconfig.Config {
	return api.manager.GetOAuth()
}

func (api *OAuthService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *OAuthService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *OAuthService) SetClientId(clientId string) {
	config := api.getConfig()
	config.SetClientId(clientId)
}

func (api *OAuthService) SetClientSecret(clientSecret string) {
	config := api.getConfig()
	config.SetClientSecret(clientSecret)
}

// This endpoint was added by liblab
func (api *OAuthService) GetAccessToken(ctx context.Context, getAccessTokenRequest GetAccessTokenRequest) (*shared.CelitechResponse[GetAccessTokenOkResponse], *shared.CelitechError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/oauth2/token").
		WithConfig(config).
		WithBody(getAccessTokenRequest).
		AddHeader("CONTENT-TYPE", "application/x-www-form-urlencoded").
		WithContentType(httptransport.ContentTypeFormUrlEncoded).
		WithResponseContentType(httptransport.ContentTypeJson).
		WithScopes(nil).
		Build()

	client := restClient.NewRestClient[GetAccessTokenOkResponse](config, api.manager)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewCelitechError[GetAccessTokenOkResponse](err)
	}

	return shared.NewCelitechResponse[GetAccessTokenOkResponse](resp), nil
}

func (api *OAuthService) GetOAuthAccessToken(config celitechconfig.Config, scopes []string) (oauthtokenmanager.GetOAuthAccessTokenResponse, error) {
	response, err := api.GetAccessToken(
		context.Background(),
		GetAccessTokenRequest{
			GrantType:    getPointer(GRANT_TYPE_CLIENT_CREDENTIALS),
			ClientId:     config.ClientId,
			ClientSecret: config.ClientSecret,
		},
	)

	if err != nil {
		return nil, err
	}
	data := response.GetData()
	return &data, nil
}

func getPointer[T any](value T) *T {
	return &value
}
