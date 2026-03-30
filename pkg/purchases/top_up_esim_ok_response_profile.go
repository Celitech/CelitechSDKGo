package purchases

import "encoding/json"

type TopUpEsimOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
}

func (t *TopUpEsimOkResponseProfile) GetIccid() *string {
	if t == nil {
		return nil
	}
	return t.Iccid
}

func (t *TopUpEsimOkResponseProfile) SetIccid(iccid string) {
	t.Iccid = &iccid
}

func (t TopUpEsimOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponseProfile to string"
	}
	return string(jsonData)
}
