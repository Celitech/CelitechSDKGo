package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type CreatePurchaseResponseProfile struct {
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode string `json:"activationCode" xml:"activationCode" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode string `json:"manualActivationCode" xml:"manualActivationCode" required:"true"`
	extraProperties      map[string]interface{}
}

func (c *CreatePurchaseResponseProfile) GetIccid() (value string) {
	if c == nil {
		return
	}
	return c.Iccid
}

func (c *CreatePurchaseResponseProfile) GetActivationCode() (value string) {
	if c == nil {
		return
	}
	return c.ActivationCode
}

func (c *CreatePurchaseResponseProfile) GetManualActivationCode() (value string) {
	if c == nil {
		return
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseResponseProfile) SetIccid(iccid string) {
	c.Iccid = iccid
}

func (c *CreatePurchaseResponseProfile) SetActivationCode(activationCode string) {
	c.ActivationCode = activationCode
}

func (c *CreatePurchaseResponseProfile) SetManualActivationCode(manualActivationCode string) {
	c.ManualActivationCode = manualActivationCode
}

func (c *CreatePurchaseResponseProfile) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseResponseProfile to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseResponseProfile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseResponseProfile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseResponseProfile(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseResponseProfile) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseResponseProfile
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
