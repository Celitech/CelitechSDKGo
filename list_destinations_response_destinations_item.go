package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListDestinationsResponseDestinationsItem struct {
	// Name of the destination
	Name string `json:"name" xml:"name" required:"true"`
	// ISO3 representation of the destination
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the destination
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region.
	SupportedCountries []string `json:"supportedCountries" xml:"supportedCountries" required:"true"`
	extraProperties    map[string]interface{}
}

func (l *ListDestinationsResponseDestinationsItem) GetName() (value string) {
	if l == nil {
		return
	}
	return l.Name
}

func (l *ListDestinationsResponseDestinationsItem) GetDestination() (value string) {
	if l == nil {
		return
	}
	return l.Destination
}

func (l *ListDestinationsResponseDestinationsItem) GetDestinationIso2() (value string) {
	if l == nil {
		return
	}
	return l.DestinationIso2
}

func (l *ListDestinationsResponseDestinationsItem) GetSupportedCountries() (value []string) {
	if l == nil {
		return
	}
	return l.SupportedCountries
}

func (l *ListDestinationsResponseDestinationsItem) SetName(name string) {
	l.Name = name
}

func (l *ListDestinationsResponseDestinationsItem) SetDestination(destination string) {
	l.Destination = destination
}

func (l *ListDestinationsResponseDestinationsItem) SetDestinationIso2(destinationIso2 string) {
	l.DestinationIso2 = destinationIso2
}

func (l *ListDestinationsResponseDestinationsItem) SetSupportedCountries(supportedCountries []string) {
	l.SupportedCountries = supportedCountries
}

func (l *ListDestinationsResponseDestinationsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListDestinationsResponseDestinationsItem) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListDestinationsResponseDestinationsItem to string"
	}
	return string(jsonData)
}

func (l *ListDestinationsResponseDestinationsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias ListDestinationsResponseDestinationsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = ListDestinationsResponseDestinationsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListDestinationsResponseDestinationsItem) MarshalJSON() ([]byte, error) {
	type embed ListDestinationsResponseDestinationsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
