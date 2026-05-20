package oauth

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type OAuthTokenRequest struct {
	GrantType    GrantType `json:"grant_type" xml:"grant_type" required:"true"`
	ClientID     string    `json:"client_id" xml:"client_id" required:"true"`
	ClientSecret string    `json:"client_secret" xml:"client_secret" required:"true"`
	Scope        string    `json:"scope" xml:"scope" required:"true"`
}

func (o OAuthTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OAuthTokenRequest to string"
	}
	return string(jsonData)
}

func (o *OAuthTokenRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OAuthTokenRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OAuthTokenRequest(tmp)
	return nil
}
