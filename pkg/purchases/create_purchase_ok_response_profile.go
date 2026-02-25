package purchases

import "encoding/json"

type CreatePurchaseOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty" required:"true"`
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
