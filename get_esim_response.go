package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimResponse struct {
	Esim            *GetEsimResponseEsim `json:"esim" xml:"esim" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimResponse) GetEsim() (value *GetEsimResponseEsim) {
	if g == nil {
		return
	}
	return g.Esim
}

func (g *GetEsimResponse) SetEsim(esim *GetEsimResponseEsim) {
	g.Esim = esim
}

func (g *GetEsimResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimResponse) MarshalJSON() ([]byte, error) {
	type embed GetEsimResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
