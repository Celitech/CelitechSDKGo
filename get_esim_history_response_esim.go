package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimHistoryResponseEsim struct {
	// ID of the eSIM
	Iccid           string                                   `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	History         []*GetEsimHistoryResponseEsimHistoryItem `json:"history" xml:"history" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimHistoryResponseEsim) GetIccid() (value string) {
	if g == nil {
		return
	}
	return g.Iccid
}

func (g *GetEsimHistoryResponseEsim) GetHistory() (value []*GetEsimHistoryResponseEsimHistoryItem) {
	if g == nil {
		return
	}
	return g.History
}

func (g *GetEsimHistoryResponseEsim) SetIccid(iccid string) {
	g.Iccid = iccid
}

func (g *GetEsimHistoryResponseEsim) SetHistory(history []*GetEsimHistoryResponseEsimHistoryItem) {
	g.History = history
}

func (g *GetEsimHistoryResponseEsim) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimHistoryResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryResponseEsim to string"
	}
	return string(jsonData)
}

func (g *GetEsimHistoryResponseEsim) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimHistoryResponseEsim
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimHistoryResponseEsim(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimHistoryResponseEsim) MarshalJSON() ([]byte, error) {
	type embed GetEsimHistoryResponseEsim
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
