package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type OAuthTokenRequest struct {
	GrantType    GrantType `json:"grant_type" xml:"grant_type" required:"true"`
	ClientID     string    `json:"client_id" xml:"client_id" required:"true"`
	ClientSecret string    `json:"client_secret" xml:"client_secret" required:"true"`
	Scope        string    `json:"scope" xml:"scope" required:"true"`
}

func (o *OAuthTokenRequest) GetGrantType() (value GrantType) {
	if o == nil {
		return
	}
	return o.GrantType
}

func (o *OAuthTokenRequest) GetClientID() (value string) {
	if o == nil {
		return
	}
	return o.ClientID
}

func (o *OAuthTokenRequest) GetClientSecret() (value string) {
	if o == nil {
		return
	}
	return o.ClientSecret
}

func (o *OAuthTokenRequest) GetScope() (value string) {
	if o == nil {
		return
	}
	return o.Scope
}

func (o *OAuthTokenRequest) SetGrantType(grantType GrantType) {
	o.GrantType = grantType
}

func (o *OAuthTokenRequest) SetClientID(clientID string) {
	o.ClientID = clientID
}

func (o *OAuthTokenRequest) SetClientSecret(clientSecret string) {
	o.ClientSecret = clientSecret
}

func (o *OAuthTokenRequest) SetScope(scope string) {
	o.Scope = scope
}

func (o *OAuthTokenRequest) GetExtraProperties() map[string]interface{} {
	return nil
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

func (o *OAuthTokenRequest) MarshalJSON() ([]byte, error) {
	type embed OAuthTokenRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
