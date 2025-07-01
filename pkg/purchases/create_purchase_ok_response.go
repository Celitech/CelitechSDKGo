package purchases

import "encoding/json"

type CreatePurchaseOkResponse struct {
	Purchase *CreatePurchaseOkResponsePurchase `json:"purchase,omitempty"`
	Profile  *CreatePurchaseOkResponseProfile  `json:"profile,omitempty"`
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

type CreatePurchaseOkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *string `json:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *string `json:"endDate,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty"`
	// Epoch value representing the start time of the package's validity
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *float64 `json:"endTime,omitempty"`
}

func (c *CreatePurchaseOkResponsePurchase) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreatePurchaseOkResponsePurchase) SetId(id string) {
	c.Id = &id
}

func (c *CreatePurchaseOkResponsePurchase) GetPackageId() *string {
	if c == nil {
		return nil
	}
	return c.PackageId
}

func (c *CreatePurchaseOkResponsePurchase) SetPackageId(packageId string) {
	c.PackageId = &packageId
}

func (c *CreatePurchaseOkResponsePurchase) GetStartDate() *string {
	if c == nil {
		return nil
	}
	return c.StartDate
}

func (c *CreatePurchaseOkResponsePurchase) SetStartDate(startDate string) {
	c.StartDate = &startDate
}

func (c *CreatePurchaseOkResponsePurchase) GetEndDate() *string {
	if c == nil {
		return nil
	}
	return c.EndDate
}

func (c *CreatePurchaseOkResponsePurchase) SetEndDate(endDate string) {
	c.EndDate = &endDate
}

func (c *CreatePurchaseOkResponsePurchase) GetCreatedDate() *string {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePurchaseOkResponsePurchase) SetCreatedDate(createdDate string) {
	c.CreatedDate = &createdDate
}

func (c *CreatePurchaseOkResponsePurchase) GetStartTime() *float64 {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *CreatePurchaseOkResponsePurchase) SetStartTime(startTime float64) {
	c.StartTime = &startTime
}

func (c *CreatePurchaseOkResponsePurchase) GetEndTime() *float64 {
	if c == nil {
		return nil
	}
	return c.EndTime
}

func (c *CreatePurchaseOkResponsePurchase) SetEndTime(endTime float64) {
	c.EndTime = &endTime
}

func (c CreatePurchaseOkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponsePurchase to string"
	}
	return string(jsonData)
}

type CreatePurchaseOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
}

func (c *CreatePurchaseOkResponseProfile) GetIccid() *string {
	if c == nil {
		return nil
	}
	return c.Iccid
}

func (c *CreatePurchaseOkResponseProfile) SetIccid(iccid string) {
	c.Iccid = &iccid
}

func (c *CreatePurchaseOkResponseProfile) GetActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetActivationCode(activationCode string) {
	c.ActivationCode = &activationCode
}

func (c *CreatePurchaseOkResponseProfile) GetManualActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetManualActivationCode(manualActivationCode string) {
	c.ManualActivationCode = &manualActivationCode
}

func (c CreatePurchaseOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponseProfile to string"
	}
	return string(jsonData)
}
