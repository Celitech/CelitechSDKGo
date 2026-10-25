package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimHistoryResponse struct {
	Esim            *GetEsimHistoryResponseEsim `json:"esim" xml:"esim" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimHistoryResponse) GetEsim() (value *GetEsimHistoryResponseEsim) {
	if g == nil {
		return
	}
	return g.Esim
}

func (g *GetEsimHistoryResponse) SetEsim(esim *GetEsimHistoryResponseEsim) {
	g.Esim = esim
}

func (g *GetEsimHistoryResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimHistoryResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimHistoryResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimHistoryResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimHistoryResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimHistoryResponse) MarshalJSON() ([]byte, error) {
	type embed GetEsimHistoryResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
