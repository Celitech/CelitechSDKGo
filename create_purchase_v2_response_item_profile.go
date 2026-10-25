package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type CreatePurchaseV2ResponseItemProfile struct {
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode string `json:"activationCode" xml:"activationCode" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode string `json:"manualActivationCode" xml:"manualActivationCode" required:"true"`
	// iOS Activation Link of the eSIM
	IosActivationLink string `json:"iosActivationLink" xml:"iosActivationLink" required:"true"`
	// Android Activation Link of the eSIM
	AndroidActivationLink string `json:"androidActivationLink" xml:"androidActivationLink" required:"true"`
	extraProperties       map[string]interface{}
}

func (c *CreatePurchaseV2ResponseItemProfile) GetIccid() (value string) {
	if c == nil {
		return
	}
	return c.Iccid
}

func (c *CreatePurchaseV2ResponseItemProfile) GetActivationCode() (value string) {
	if c == nil {
		return
	}
	return c.ActivationCode
}

func (c *CreatePurchaseV2ResponseItemProfile) GetManualActivationCode() (value string) {
	if c == nil {
		return
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseV2ResponseItemProfile) GetIosActivationLink() (value string) {
	if c == nil {
		return
	}
	return c.IosActivationLink
}

func (c *CreatePurchaseV2ResponseItemProfile) GetAndroidActivationLink() (value string) {
	if c == nil {
		return
	}
	return c.AndroidActivationLink
}

func (c *CreatePurchaseV2ResponseItemProfile) SetIccid(iccid string) {
	c.Iccid = iccid
}

func (c *CreatePurchaseV2ResponseItemProfile) SetActivationCode(activationCode string) {
	c.ActivationCode = activationCode
}

func (c *CreatePurchaseV2ResponseItemProfile) SetManualActivationCode(manualActivationCode string) {
	c.ManualActivationCode = manualActivationCode
}

func (c *CreatePurchaseV2ResponseItemProfile) SetIosActivationLink(iosActivationLink string) {
	c.IosActivationLink = iosActivationLink
}

func (c *CreatePurchaseV2ResponseItemProfile) SetAndroidActivationLink(androidActivationLink string) {
	c.AndroidActivationLink = androidActivationLink
}

func (c *CreatePurchaseV2ResponseItemProfile) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseV2ResponseItemProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2ResponseItemProfile to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2ResponseItemProfile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2ResponseItemProfile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2ResponseItemProfile(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseV2ResponseItemProfile) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseV2ResponseItemProfile
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
