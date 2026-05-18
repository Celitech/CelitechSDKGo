package destinations

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type Destinations struct {
	// Name of the destination
	Name string `json:"name" xml:"name" required:"true"`
	// ISO3 representation of the destination
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the destination
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region.
	SupportedCountries []string `json:"supportedCountries" xml:"supportedCountries" required:"true"`
}

func (d Destinations) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Destinations to string"
	}
	return string(jsonData)
}

func (d *Destinations) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias Destinations
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = Destinations(tmp)
	return nil
}
