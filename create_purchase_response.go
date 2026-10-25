package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type CreatePurchaseResponse struct {
	Purchase        *CreatePurchaseResponsePurchase `json:"purchase" xml:"purchase" required:"true"`
	Profile         *CreatePurchaseResponseProfile  `json:"profile" xml:"profile" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePurchaseResponse) GetPurchase() (value *CreatePurchaseResponsePurchase) {
	if c == nil {
		return
	}
	return c.Purchase
}

func (c *CreatePurchaseResponse) GetProfile() (value *CreatePurchaseResponseProfile) {
	if c == nil {
		return
	}
	return c.Profile
}

func (c *CreatePurchaseResponse) SetPurchase(purchase *CreatePurchaseResponsePurchase) {
	c.Purchase = purchase
}

func (c *CreatePurchaseResponse) SetProfile(profile *CreatePurchaseResponseProfile) {
	c.Profile = profile
}

func (c *CreatePurchaseResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseResponse to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseResponse) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
