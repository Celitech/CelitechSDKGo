package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
)

type ListDestinationsResponse struct {
	Destinations    []*ListDestinationsResponseDestinationsItem `json:"destinations" xml:"destinations" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListDestinationsResponse) GetDestinations() (value []*ListDestinationsResponseDestinationsItem) {
	if l == nil {
		return
	}
	return l.Destinations
}

func (l *ListDestinationsResponse) SetDestinations(destinations []*ListDestinationsResponseDestinationsItem) {
	l.Destinations = destinations
}

func (l *ListDestinationsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListDestinationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListDestinationsResponse to string"
	}
	return string(jsonData)
}

func (l *ListDestinationsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListDestinationsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListDestinationsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListDestinationsResponse) MarshalJSON() ([]byte, error) {
	type embed ListDestinationsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
