package esim

import "encoding/json"

type GetEsimHistoryOkResponse struct {
	Esim *GetEsimHistoryOkResponseEsim `json:"esim,omitempty" required:"true"`
}

func (g *GetEsimHistoryOkResponse) GetEsim() *GetEsimHistoryOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimHistoryOkResponse) SetEsim(esim GetEsimHistoryOkResponseEsim) {
	g.Esim = &esim
}

func (g GetEsimHistoryOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponse to string"
	}
	return string(jsonData)
}
