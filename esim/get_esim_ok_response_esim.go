package esim

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type GetEsimOkResponseEsim struct {
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// SM-DP+ Address
	SmdpAddress string `json:"smdpAddress" xml:"smdpAddress" required:"true"`
	// QR Code of the eSIM as base64
	ActivationCode string `json:"activationCode" xml:"activationCode" required:"true" maxLength:"8000" minLength:"1000"`
	// The manual activation code
	ManualActivationCode string `json:"manualActivationCode" xml:"manualActivationCode" required:"true"`
	// Status of the eSIM, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status string `json:"status" xml:"status" required:"true"`
	// Status of the eSIM connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'
	ConnectivityStatus string `json:"connectivityStatus" xml:"connectivityStatus" required:"true"`
	// Indicates whether the eSIM is currently eligible for a top-up. This flag should be checked before attempting a top-up request.
	IsTopUpAllowed bool `json:"isTopUpAllowed" xml:"isTopUpAllowed" required:"true"`
}

func (g GetEsimOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimOkResponseEsim to string"
	}
	return string(jsonData)
}

func (g *GetEsimOkResponseEsim) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimOkResponseEsim
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimOkResponseEsim(tmp)
	return nil
}
