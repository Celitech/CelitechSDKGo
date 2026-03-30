package destinations

import "encoding/json"

type ListDestinationsOkResponse struct {
	Destinations []Destinations `json:"destinations,omitempty" required:"true"`
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
