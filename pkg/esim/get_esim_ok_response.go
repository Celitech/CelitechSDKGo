package esim

import "encoding/json"

type GetEsimOkResponse struct {
	Esim *GetEsimOkResponseEsim `json:"esim,omitempty"`
}

func (g *GetEsimOkResponse) GetEsim() *GetEsimOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimOkResponse) SetEsim(esim GetEsimOkResponseEsim) {
	g.Esim = &esim
}

func (g GetEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimOkResponse to string"
	}
	return string(jsonData)
}

type GetEsimOkResponseEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// SM-DP+ Address
	SmdpAddress *string `json:"smdpAddress,omitempty"`
	// The manual activation code
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
	// Status of the eSIM, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status *string `json:"status,omitempty"`
}

func (g *GetEsimOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimOkResponseEsim) SetIccid(iccid string) {
	g.Iccid = &iccid
}

func (g *GetEsimOkResponseEsim) GetSmdpAddress() *string {
	if g == nil {
		return nil
	}
	return g.SmdpAddress
}

func (g *GetEsimOkResponseEsim) SetSmdpAddress(smdpAddress string) {
	g.SmdpAddress = &smdpAddress
}

func (g *GetEsimOkResponseEsim) GetManualActivationCode() *string {
	if g == nil {
		return nil
	}
	return g.ManualActivationCode
}

func (g *GetEsimOkResponseEsim) SetManualActivationCode(manualActivationCode string) {
	g.ManualActivationCode = &manualActivationCode
}

func (g *GetEsimOkResponseEsim) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetEsimOkResponseEsim) SetStatus(status string) {
	g.Status = &status
}

func (g GetEsimOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimOkResponseEsim to string"
	}
	return string(jsonData)
}
