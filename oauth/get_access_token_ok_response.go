package oauth

import "encoding/json"

type GetAccessTokenOkResponse struct {
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TokenType   *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

func (g GetAccessTokenOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAccessTokenOkResponse to string"
	}
	return string(jsonData)
}
