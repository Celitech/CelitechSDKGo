package esim

import "encoding/json"

type GetEsimHistoryOkResponseEsim struct {
	// ID of the eSIM
	Iccid   *string   `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	History []History `json:"history,omitempty" required:"true"`
}

func (g *GetEsimHistoryOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimHistoryOkResponseEsim) SetIccid(iccid string) {
	g.Iccid = &iccid
}

func (g *GetEsimHistoryOkResponseEsim) GetHistory() []History {
	if g == nil {
		return nil
	}
	return g.History
}

func (g *GetEsimHistoryOkResponseEsim) SetHistory(history []History) {
	g.History = history
}

func (g GetEsimHistoryOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponseEsim to string"
	}
	return string(jsonData)
}
