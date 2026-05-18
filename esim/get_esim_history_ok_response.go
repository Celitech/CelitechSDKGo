package esim

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type GetEsimHistoryOkResponse struct {
	Esim GetEsimHistoryOkResponseEsim `json:"esim" xml:"esim" required:"true"`
}

func (g GetEsimHistoryOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimHistoryOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimHistoryOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimHistoryOkResponse(tmp)
	return nil
}
