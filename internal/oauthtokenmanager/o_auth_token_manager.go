package oauthtokenmanager

import (
	"errors"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"time"
)

type OAuthTokenManager struct {
	token         *OAuthToken
	tokenService  TokenService
	refreshBuffer int64
}

type TokenService interface {
	GetOAuthAccessToken(config celitechconfig.Config, scopes []string) (GetOAuthAccessTokenResponse, error)
}

type GetOAuthAccessTokenResponse interface {
	GetExpiresIn() *int64
	GetAccessToken() *string
}

func NewOAuthTokenManager(tokenService TokenService, refreshBuffer int64) *OAuthTokenManager {
	return &OAuthTokenManager{
		tokenService:  tokenService,
		refreshBuffer: refreshBuffer,
	}
}

func (m *OAuthTokenManager) GetToken(scopes []string, config celitechconfig.Config) (*OAuthToken, error) {

	// Check if the token is expired or within the buffer window
	// If so, return it instead of fetching a new one
	if m.token != nil && m.token.HasAllScopes(scopes) {
		// A token without an expiry is considered valid indefinitely.
		if m.token.ExpiresAt == nil {
			return m.token, nil
		}

		// Check if the token is still valid and not within the refresh buffer window.
		bufferedExpiry := m.token.ExpiresAt.Add(-time.Duration(m.refreshBuffer) * time.Second)
		if bufferedExpiry.After(time.Now()) {
			return m.token, nil
		}
	}

	updatedScopesMap := make(map[string]struct{})
	if m.token != nil {
		for scope := range m.token.Scopes {
			updatedScopesMap[scope] = struct{}{}
		}
	}
	for _, scope := range scopes {
		updatedScopesMap[scope] = struct{}{}
	}

	updatedScopes := make([]string, 0, len(updatedScopesMap))
	for scope := range updatedScopesMap {
		updatedScopes = append(updatedScopes, scope)
	}

	response, err := m.tokenService.GetOAuthAccessToken(config, updatedScopes)
	if err != nil || response == nil {
		return nil, errors.New("OAuthError: failed to fetch access token")
	}

	if *response.GetAccessToken() == "" {
		return nil, errors.New("OAuthError: token endpoint response did not return access token")
	}

	var expiresIn *time.Time
	if *response.GetExpiresIn() > 0 {
		exp := time.Now().Add(time.Duration(*response.GetExpiresIn()) * time.Second)
		expiresIn = &exp
	}

	m.token = NewOAuthToken(*response.GetAccessToken(), updatedScopes, expiresIn)

	return m.token, nil
}
