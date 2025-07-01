package oauthtokenmanager

import "time"

type OAuthToken struct {
	AccessToken string
	Scopes      map[string]struct{}
	ExpiresAt   *time.Time
}

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

func (t *OAuthToken) HasAllScopes(scopes []string) bool {
	for _, scope := range scopes {
		if _, exists := t.Scopes[scope]; !exists {
			return false
		}
	}
	return true
}
