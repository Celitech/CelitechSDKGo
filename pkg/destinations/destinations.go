package destinations

import "encoding/json"

type Destinations struct {
	// Name of the destination
	Name *string `json:"name,omitempty" required:"true"`
	// ISO3 representation of the destination
	Destination *string `json:"destination,omitempty" required:"true"`
	// ISO2 representation of the destination
	DestinationIso2 *string `json:"destinationISO2,omitempty" required:"true"`
	// This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region.
	SupportedCountries []string `json:"supportedCountries,omitempty" required:"true"`
}

func (d *Destinations) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Destinations) SetName(name string) {
	d.Name = &name
}

func (d *Destinations) GetDestination() *string {
	if d == nil {
		return nil
	}
	return d.Destination
}

func (d *Destinations) SetDestination(destination string) {
	d.Destination = &destination
}

func (d *Destinations) GetDestinationIso2() *string {
	if d == nil {
		return nil
	}
	return d.DestinationIso2
}

func (d *Destinations) SetDestinationIso2(destinationIso2 string) {
	d.DestinationIso2 = &destinationIso2
}

func (d *Destinations) GetSupportedCountries() []string {
	if d == nil {
		return nil
	}
	return d.SupportedCountries
}

func (d *Destinations) SetSupportedCountries(supportedCountries []string) {
	d.SupportedCountries = supportedCountries
}

func (d Destinations) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Destinations to string"
	}
	return string(jsonData)
}
