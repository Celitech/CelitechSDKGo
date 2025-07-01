package purchases

import "encoding/json"

type CreatePurchaseV2OkResponse struct {
	Purchase *CreatePurchaseV2OkResponsePurchase `json:"purchase,omitempty"`
	Profile  *CreatePurchaseV2OkResponseProfile  `json:"profile,omitempty"`
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

type CreatePurchaseV2OkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty"`
}

func (c *CreatePurchaseV2OkResponsePurchase) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreatePurchaseV2OkResponsePurchase) SetId(id string) {
	c.Id = &id
}

func (c *CreatePurchaseV2OkResponsePurchase) GetPackageId() *string {
	if c == nil {
		return nil
	}
	return c.PackageId
}

func (c *CreatePurchaseV2OkResponsePurchase) SetPackageId(packageId string) {
	c.PackageId = &packageId
}

func (c *CreatePurchaseV2OkResponsePurchase) GetCreatedDate() *string {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePurchaseV2OkResponsePurchase) SetCreatedDate(createdDate string) {
	c.CreatedDate = &createdDate
}

func (c CreatePurchaseV2OkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponsePurchase to string"
	}
	return string(jsonData)
}

type CreatePurchaseV2OkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
}

func (c *CreatePurchaseV2OkResponseProfile) GetIccid() *string {
	if c == nil {
		return nil
	}
	return c.Iccid
}

func (c *CreatePurchaseV2OkResponseProfile) SetIccid(iccid string) {
	c.Iccid = &iccid
}

func (c *CreatePurchaseV2OkResponseProfile) GetActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ActivationCode
}

func (c *CreatePurchaseV2OkResponseProfile) SetActivationCode(activationCode string) {
	c.ActivationCode = &activationCode
}

func (c *CreatePurchaseV2OkResponseProfile) GetManualActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseV2OkResponseProfile) SetManualActivationCode(manualActivationCode string) {
	c.ManualActivationCode = &manualActivationCode
}

func (c CreatePurchaseV2OkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponseProfile to string"
	}
	return string(jsonData)
}
