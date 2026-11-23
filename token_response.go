package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type TokenResponse struct {
	// The generated token
	Token           string `json:"token" xml:"token" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TokenResponse) GetToken() (value string) {
	if t == nil {
		return
	}
	return t.Token
}

func (t *TokenResponse) SetToken(token string) {
	t.Token = token
}

func (t *TokenResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TokenResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TokenResponse to string"
	}
	return string(jsonData)
}

func (t *TokenResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TokenResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TokenResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TokenResponse) MarshalJSON() ([]byte, error) {
	type embed TokenResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
