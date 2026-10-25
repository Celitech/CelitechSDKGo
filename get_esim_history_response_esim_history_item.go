package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type GetEsimHistoryResponseEsimHistoryItem struct {
	// The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status string `json:"status" xml:"status" required:"true"`
	// The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StatusDate time.Time `json:"statusDate" xml:"statusDate" required:"true"`
	// Epoch value representing the date when the eSIM status changed
	Date            *float64 `json:"date,omitempty" xml:"date,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetEsimHistoryResponseEsimHistoryItem) GetStatus() (value string) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetEsimHistoryResponseEsimHistoryItem) GetStatusDate() (value time.Time) {
	if g == nil {
		return
	}
	return g.StatusDate
}

func (g *GetEsimHistoryResponseEsimHistoryItem) GetDate() (value *float64) {
	if g == nil {
		return
	}
	return g.Date
}

func (g *GetEsimHistoryResponseEsimHistoryItem) SetStatus(status string) {
	g.Status = status
}

func (g *GetEsimHistoryResponseEsimHistoryItem) SetStatusDate(statusDate time.Time) {
	g.StatusDate = statusDate
}

func (g *GetEsimHistoryResponseEsimHistoryItem) SetDate(date *float64) {
	g.Date = date
}

func (g *GetEsimHistoryResponseEsimHistoryItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimHistoryResponseEsimHistoryItem) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryResponseEsimHistoryItem to string"
	}
	return string(jsonData)
}

func (g *GetEsimHistoryResponseEsimHistoryItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimHistoryResponseEsimHistoryItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimHistoryResponseEsimHistoryItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimHistoryResponseEsimHistoryItem) MarshalJSON() ([]byte, error) {
	type embed GetEsimHistoryResponseEsimHistoryItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
