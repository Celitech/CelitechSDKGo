package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type TopUpEsimResponse struct {
	Purchase        *TopUpEsimResponsePurchase `json:"purchase" xml:"purchase" required:"true"`
	Profile         *TopUpEsimResponseProfile  `json:"profile" xml:"profile" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TopUpEsimResponse) GetPurchase() (value *TopUpEsimResponsePurchase) {
	if t == nil {
		return
	}
	return t.Purchase
}

func (t *TopUpEsimResponse) GetProfile() (value *TopUpEsimResponseProfile) {
	if t == nil {
		return
	}
	return t.Profile
}

func (t *TopUpEsimResponse) SetPurchase(purchase *TopUpEsimResponsePurchase) {
	t.Purchase = purchase
}

func (t *TopUpEsimResponse) SetProfile(profile *TopUpEsimResponseProfile) {
	t.Profile = profile
}

func (t *TopUpEsimResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopUpEsimResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimResponse to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopUpEsimResponse) MarshalJSON() ([]byte, error) {
	type embed TopUpEsimResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
