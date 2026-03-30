package esim

import "encoding/json"

type GetEsimOkResponse struct {
	Esim *GetEsimOkResponseEsim `json:"esim,omitempty" required:"true"`
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
