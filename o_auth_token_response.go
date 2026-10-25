package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type OAuthTokenResponse struct {
	AccessToken     *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresIn       *int    `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OAuthTokenResponse) GetAccessToken() (value *string) {
	if o == nil {
		return
	}
	return o.AccessToken
}

func (o *OAuthTokenResponse) GetExpiresIn() (value *int) {
	if o == nil {
		return
	}
	return o.ExpiresIn
}

func (o *OAuthTokenResponse) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

func (o *OAuthTokenResponse) SetExpiresIn(expiresIn *int) {
	o.ExpiresIn = expiresIn
}

func (o *OAuthTokenResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OAuthTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OAuthTokenResponse to string"
	}
	return string(jsonData)
}

func (o *OAuthTokenResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OAuthTokenResponse) MarshalJSON() ([]byte, error) {
	type embed OAuthTokenResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
