package oauthtokenmanager

import "time"

// OAuthToken represents an OAuth access token with associated scopes and expiration time.
// Used for caching and validating access tokens before making API requests.
type OAuthToken struct {
	AccessToken string
	Scopes      map[string]struct{}
	ExpiresAt   *time.Time
}

// NewOAuthToken creates a new OAuth token with the specified access token, scopes, and expiration.
// Converts the scopes array to a map for efficient lookup. Expiration can be nil for tokens that never expire.
func NewOAuthToken(accessToken string, scopes []string, expiresAt *time.Time) *OAuthToken {
	scopeMap := make(map[string]struct{}, len(scopes))
	for _, scope := range scopes {
		scopeMap[scope] = struct{}{}
	}
	return &OAuthToken{
		AccessToken: accessToken,
		Scopes:      scopeMap,
		ExpiresAt:   expiresAt,
	}
}

// HasAllScopes checks if the token contains all of the specified scopes.
// Returns true only if all requested scopes are present in the token.
func (t *OAuthToken) HasAllScopes(scopes []string) bool {
	for _, scope := range scopes {
		if _, exists := t.Scopes[scope]; !exists {
			return false
		}
	}
	return true
}
