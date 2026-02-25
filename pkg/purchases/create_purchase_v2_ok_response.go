package purchases

import "encoding/json"

type CreatePurchaseV2OkResponse struct {
	Purchase *CreatePurchaseV2OkResponsePurchase `json:"purchase,omitempty" required:"true"`
	Profile  *CreatePurchaseV2OkResponseProfile  `json:"profile,omitempty" required:"true"`
}

func (c *CreatePurchaseV2OkResponse) GetPurchase() *CreatePurchaseV2OkResponsePurchase {
	if c == nil {
		return nil
	}
	return c.Purchase
}

func (c *CreatePurchaseV2OkResponse) SetPurchase(purchase CreatePurchaseV2OkResponsePurchase) {
	c.Purchase = &purchase
}

func (c *CreatePurchaseV2OkResponse) GetProfile() *CreatePurchaseV2OkResponseProfile {
	if c == nil {
		return nil
	}
	return c.Profile
}

func (c *CreatePurchaseV2OkResponse) SetProfile(profile CreatePurchaseV2OkResponseProfile) {
	c.Profile = &profile
}

func (c CreatePurchaseV2OkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponse to string"
	}
	return string(jsonData)
}
