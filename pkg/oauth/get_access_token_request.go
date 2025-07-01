package oauth

import "encoding/json"

type GetAccessTokenRequest struct {
	GrantType    *GrantType `json:"grant_type,omitempty"`
	ClientId     *string    `json:"client_id,omitempty"`
	ClientSecret *string    `json:"client_secret,omitempty"`
}

func (g *GetAccessTokenRequest) GetGrantType() *GrantType {
	if g == nil {
		return nil
	}
	return g.GrantType
}

func (g *GetAccessTokenRequest) SetGrantType(grantType GrantType) {
	g.GrantType = &grantType
}

func (g *GetAccessTokenRequest) GetClientId() *string {
	if g == nil {
		return nil
	}
	return g.ClientId
}

func (g *GetAccessTokenRequest) SetClientId(clientId string) {
	g.ClientId = &clientId
}

func (g *GetAccessTokenRequest) GetClientSecret() *string {
	if g == nil {
		return nil
	}
	return g.ClientSecret
}

func (g *GetAccessTokenRequest) SetClientSecret(clientSecret string) {
	g.ClientSecret = &clientSecret
}

func (g GetAccessTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAccessTokenRequest to string"
	}
	return string(jsonData)
}

type GrantType string

const (
	GRANT_TYPE_CLIENT_CREDENTIALS GrantType = "client_credentials"
)
