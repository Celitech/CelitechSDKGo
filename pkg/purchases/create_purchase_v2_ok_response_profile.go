package purchases

import "encoding/json"

type CreatePurchaseV2OkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty" required:"true"`
	// iOS Activation Link of the eSIM
	IosActivationLink *string `json:"iosActivationLink,omitempty" required:"true"`
	// Android Activation Link of the eSIM
	AndroidActivationLink *string `json:"androidActivationLink,omitempty" required:"true"`
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

func (c *CreatePurchaseV2OkResponseProfile) GetIosActivationLink() *string {
	if c == nil {
		return nil
	}
	return c.IosActivationLink
}

func (c *CreatePurchaseV2OkResponseProfile) SetIosActivationLink(iosActivationLink string) {
	c.IosActivationLink = &iosActivationLink
}

func (c *CreatePurchaseV2OkResponseProfile) GetAndroidActivationLink() *string {
	if c == nil {
		return nil
	}
	return c.AndroidActivationLink
}

func (c *CreatePurchaseV2OkResponseProfile) SetAndroidActivationLink(androidActivationLink string) {
	c.AndroidActivationLink = &androidActivationLink
}

func (c CreatePurchaseV2OkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponseProfile to string"
	}
	return string(jsonData)
}
