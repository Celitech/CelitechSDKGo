package iframe

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type TokenOkResponse struct {
	// The generated token
	Token string `json:"token" required:"true"`
}

func (t TokenOkResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TokenOkResponse to string"
	}
	return string(jsonData)
}

func (t *TokenOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TokenOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TokenOkResponse(tmp)
	return nil
}
