package esim

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type GetEsimOkResponse struct {
	Esim GetEsimOkResponseEsim `json:"esim" required:"true"`
}

func (g GetEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimOkResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimOkResponse(tmp)
	return nil
}
