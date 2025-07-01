package esim

import "encoding/json"

type GetEsimMacOkResponse struct {
	Esim *GetEsimMacOkResponseEsim `json:"esim,omitempty"`
}

func (g *GetEsimMacOkResponse) GetEsim() *GetEsimMacOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimMacOkResponse) SetEsim(esim GetEsimMacOkResponseEsim) {
	g.Esim = &esim
}

func (g GetEsimMacOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimMacOkResponse to string"
	}
	return string(jsonData)
}

type GetEsimMacOkResponseEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// SM-DP+ Address
	SmdpAddress *string `json:"smdpAddress,omitempty"`
	// The manual activation code
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
}

func (g *GetEsimMacOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimMacOkResponseEsim) SetIccid(iccid string) {
	g.Iccid = &iccid
}

func (g *GetEsimMacOkResponseEsim) GetSmdpAddress() *string {
	if g == nil {
		return nil
	}
	return g.SmdpAddress
}

func (g *GetEsimMacOkResponseEsim) SetSmdpAddress(smdpAddress string) {
	g.SmdpAddress = &smdpAddress
}

func (g *GetEsimMacOkResponseEsim) GetManualActivationCode() *string {
	if g == nil {
		return nil
	}
	return g.ManualActivationCode
}

func (g *GetEsimMacOkResponseEsim) SetManualActivationCode(manualActivationCode string) {
	g.ManualActivationCode = &manualActivationCode
}

func (g GetEsimMacOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimMacOkResponseEsim to string"
	}
	return string(jsonData)
}
