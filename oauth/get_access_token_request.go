package oauth

import "encoding/json"

type GetAccessTokenRequest struct {
	GrantType    *GrantType `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	ClientID     *string    `json:"client_id,omitempty" xml:"client_id,omitempty"`
	ClientSecret *string    `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
}

func (g GetAccessTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAccessTokenRequest to string"
	}
	return string(jsonData)
}
