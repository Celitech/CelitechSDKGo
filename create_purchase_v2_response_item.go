package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type CreatePurchaseV2ResponseItem struct {
	Purchase        *CreatePurchaseV2ResponseItemPurchase `json:"purchase" xml:"purchase" required:"true"`
	Profile         *CreatePurchaseV2ResponseItemProfile  `json:"profile" xml:"profile" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePurchaseV2ResponseItem) GetPurchase() (value *CreatePurchaseV2ResponseItemPurchase) {
	if c == nil {
		return
	}
	return c.Purchase
}

func (c *CreatePurchaseV2ResponseItem) GetProfile() (value *CreatePurchaseV2ResponseItemProfile) {
	if c == nil {
		return
	}
	return c.Profile
}

func (c *CreatePurchaseV2ResponseItem) SetPurchase(purchase *CreatePurchaseV2ResponseItemPurchase) {
	c.Purchase = purchase
}

func (c *CreatePurchaseV2ResponseItem) SetProfile(profile *CreatePurchaseV2ResponseItemProfile) {
	c.Profile = profile
}

func (c *CreatePurchaseV2ResponseItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseV2ResponseItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2ResponseItem to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2ResponseItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2ResponseItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2ResponseItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseV2ResponseItem) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseV2ResponseItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
