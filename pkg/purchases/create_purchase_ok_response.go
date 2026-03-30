package purchases

import "encoding/json"

type CreatePurchaseOkResponse struct {
	Purchase *CreatePurchaseOkResponsePurchase `json:"purchase,omitempty" required:"true"`
	Profile  *CreatePurchaseOkResponseProfile  `json:"profile,omitempty" required:"true"`
}

func (c *CreatePurchaseOkResponse) GetPurchase() *CreatePurchaseOkResponsePurchase {
	if c == nil {
		return nil
	}
	return c.Purchase
}

func (c *CreatePurchaseOkResponse) SetPurchase(purchase CreatePurchaseOkResponsePurchase) {
	c.Purchase = &purchase
}

func (c *CreatePurchaseOkResponse) GetProfile() *CreatePurchaseOkResponseProfile {
	if c == nil {
		return nil
	}
	return c.Profile
}

func (c *CreatePurchaseOkResponse) SetProfile(profile CreatePurchaseOkResponseProfile) {
	c.Profile = &profile
}

func (c CreatePurchaseOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponse to string"
	}
	return string(jsonData)
}
