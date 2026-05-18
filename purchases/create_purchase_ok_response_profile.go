package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type CreatePurchaseOkResponseProfile struct {
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode string `json:"activationCode" xml:"activationCode" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode string `json:"manualActivationCode" xml:"manualActivationCode" required:"true"`
}

func (c CreatePurchaseOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponseProfile to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseOkResponseProfile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseOkResponseProfile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseOkResponseProfile(tmp)
	return nil
}
