package iframe

import "encoding/json"

type TokenOkResponse struct {
	// The generated token
	Token *string `json:"token,omitempty"`
}

func (t *TokenOkResponse) GetToken() *string {
	if t == nil {
		return nil
	}
	return t.Token
}

func (t *TokenOkResponse) SetToken(token string) {
	t.Token = &token
}

func (t TokenOkResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TokenOkResponse to string"
	}
	return string(jsonData)
}
