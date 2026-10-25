package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimResponseEsim struct {
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
	IsTopUpAllowed  bool `json:"isTopUpAllowed" xml:"isTopUpAllowed" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimResponseEsim) GetIccid() (value string) {
	if g == nil {
		return
	}
	return g.Iccid
}

func (g *GetEsimResponseEsim) GetSmdpAddress() (value string) {
	if g == nil {
		return
	}
	return g.SmdpAddress
}

func (g *GetEsimResponseEsim) GetActivationCode() (value string) {
	if g == nil {
		return
	}
	return g.ActivationCode
}

func (g *GetEsimResponseEsim) GetManualActivationCode() (value string) {
	if g == nil {
		return
	}
	return g.ManualActivationCode
}

func (g *GetEsimResponseEsim) GetStatus() (value string) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetEsimResponseEsim) GetConnectivityStatus() (value string) {
	if g == nil {
		return
	}
	return g.ConnectivityStatus
}

func (g *GetEsimResponseEsim) GetIsTopUpAllowed() (value bool) {
	if g == nil {
		return
	}
	return g.IsTopUpAllowed
}

func (g *GetEsimResponseEsim) SetIccid(iccid string) {
	g.Iccid = iccid
}

func (g *GetEsimResponseEsim) SetSmdpAddress(smdpAddress string) {
	g.SmdpAddress = smdpAddress
}

func (g *GetEsimResponseEsim) SetActivationCode(activationCode string) {
	g.ActivationCode = activationCode
}

func (g *GetEsimResponseEsim) SetManualActivationCode(manualActivationCode string) {
	g.ManualActivationCode = manualActivationCode
}

func (g *GetEsimResponseEsim) SetStatus(status string) {
	g.Status = status
}

func (g *GetEsimResponseEsim) SetConnectivityStatus(connectivityStatus string) {
	g.ConnectivityStatus = connectivityStatus
}

func (g *GetEsimResponseEsim) SetIsTopUpAllowed(isTopUpAllowed bool) {
	g.IsTopUpAllowed = isTopUpAllowed
}

func (g *GetEsimResponseEsim) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimResponseEsim to string"
	}
	return string(jsonData)
}

func (g *GetEsimResponseEsim) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimResponseEsim
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimResponseEsim(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimResponseEsim) MarshalJSON() ([]byte, error) {
	type embed GetEsimResponseEsim
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
