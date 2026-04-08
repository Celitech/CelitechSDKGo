package esim

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type GetEsimHistoryOkResponseEsim struct {
	// ID of the eSIM
	Iccid   string    `json:"iccid" required:"true" maxLength:"22" minLength:"18"`
	History []History `json:"history" required:"true"`
}

func (g GetEsimHistoryOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponseEsim to string"
	}
	return string(jsonData)
}

func (g *GetEsimHistoryOkResponseEsim) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimHistoryOkResponseEsim
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimHistoryOkResponseEsim(tmp)
	return nil
}
