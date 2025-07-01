package oauth

import "encoding/json"

type GetAccessTokenOkResponse struct {
	AccessToken *string `json:"access_token,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty"`
}

func (g *GetAccessTokenOkResponse) GetAccessToken() *string {
	if g == nil {
		return nil
	}
	return g.AccessToken
}

func (g *GetAccessTokenOkResponse) SetAccessToken(accessToken string) {
	g.AccessToken = &accessToken
}

func (g *GetAccessTokenOkResponse) GetTokenType() *string {
	if g == nil {
		return nil
	}
	return g.TokenType
}

func (g *GetAccessTokenOkResponse) SetTokenType(tokenType string) {
	g.TokenType = &tokenType
}

func (g *GetAccessTokenOkResponse) GetExpiresIn() *int64 {
	if g == nil {
		return nil
	}
	return g.ExpiresIn
}

func (g *GetAccessTokenOkResponse) SetExpiresIn(expiresIn int64) {
	g.ExpiresIn = &expiresIn
}

func (g GetAccessTokenOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAccessTokenOkResponse to string"
	}
	return string(jsonData)
}
