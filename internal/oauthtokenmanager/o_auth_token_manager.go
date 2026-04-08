package oauthtokenmanager

import (
	"errors"
	"example.com/celitech/celitechconfig"
	"fmt"
	"time"
)

// OAuthTokenManager manages OAuth access tokens with caching and automatic refresh.
// It caches tokens and refreshes them before expiration using a configurable buffer window.
type OAuthTokenManager struct {
	token         *OAuthToken
	tokenService  TokenService
	refreshBuffer int64
}

// TokenService defines the interface for fetching OAuth access tokens from the API.
type TokenService interface {
	GetOAuthAccessToken(config celitechconfig.Config, scopes []string) (GetOAuthAccessTokenResponse, error)
}

// GetOAuthAccessTokenResponse defines the interface for accessing token response data.
type GetOAuthAccessTokenResponse interface {
	GetExpiresIn() *int64
	GetAccessToken() *string
}

// OAuthTokenResponse is a spec-agnostic adapter that always satisfies GetOAuthAccessTokenResponse.
// It is returned by GetOAuthAccessToken so the token manager never depends on the generated
// model's getter signatures, which vary based on whether fields are required in the spec.
type OAuthTokenResponse struct {
	accessToken *string
	expiresIn   *int64
}

func NewOAuthTokenResponse(accessToken *string, expiresIn *int64) *OAuthTokenResponse {
	return &OAuthTokenResponse{accessToken: accessToken, expiresIn: expiresIn}
}

func (r *OAuthTokenResponse) GetAccessToken() *string { return r.accessToken }
func (r *OAuthTokenResponse) GetExpiresIn() *int64    { return r.expiresIn }

// NewOAuthTokenManager creates a new OAuth token manager with the specified token service and refresh buffer.
// The refresh buffer determines how many seconds before expiration the token should be refreshed.
func NewOAuthTokenManager(tokenService TokenService, refreshBuffer int64) *OAuthTokenManager {
	return &OAuthTokenManager{
		tokenService:  tokenService,
		refreshBuffer: refreshBuffer,
	}
}

// GetToken retrieves a valid OAuth access token for the requested scopes.
// Returns a cached token if valid and contains all requested scopes, otherwise fetches a new token.
// Automatically refreshes tokens that are expired or within the refresh buffer window.
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
	if err != nil {
		return nil, fmt.Errorf("OAuthError: failed to fetch access token: %w", err)
	}

	if response == nil || response.GetAccessToken() == nil || *response.GetAccessToken() == "" {
		return nil, errors.New("OAuthError: token endpoint response did not return access token")
	}

	var expiresIn *time.Time
	if response.GetExpiresIn() != nil && *response.GetExpiresIn() > 0 {
		exp := time.Now().Add(time.Duration(*response.GetExpiresIn()) * time.Second)
		expiresIn = &exp
	}

	m.token = NewOAuthToken(*response.GetAccessToken(), updatedScopes, expiresIn)

	return m.token, nil
}
