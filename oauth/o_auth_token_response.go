package oauth

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v2/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/v2/param"
)

type OAuthTokenResponse struct {
	AccessToken *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresIn   *param.Nullable[int64] `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

func (o OAuthTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OAuthTokenResponse to string"
	}
	return string(jsonData)
}

func (o *OAuthTokenResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, o)
}
