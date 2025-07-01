package destinations

import "encoding/json"

type ListDestinationsOkResponse struct {
	Destinations []Destinations `json:"destinations,omitempty"`
}

func (l *ListDestinationsOkResponse) GetDestinations() []Destinations {
	if l == nil {
		return nil
	}
	return l.Destinations
}

func (l *ListDestinationsOkResponse) SetDestinations(destinations []Destinations) {
	l.Destinations = destinations
}

func (l ListDestinationsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListDestinationsOkResponse to string"
	}
	return string(jsonData)
}

type Destinations struct {
	// Name of the destination
	Name *string `json:"name,omitempty"`
	// ISO representation of the destination
	Destination *string `json:"destination,omitempty"`
	// This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region.
	SupportedCountries []string `json:"supportedCountries,omitempty"`
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
