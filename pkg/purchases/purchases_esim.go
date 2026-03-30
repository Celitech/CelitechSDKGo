package purchases

import "encoding/json"

type PurchasesEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
}

func (p *PurchasesEsim) GetIccid() *string {
	if p == nil {
		return nil
	}
	return p.Iccid
}

func (p *PurchasesEsim) SetIccid(iccid string) {
	p.Iccid = &iccid
}

func (p PurchasesEsim) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PurchasesEsim to string"
	}
	return string(jsonData)
}
