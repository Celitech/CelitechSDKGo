package destinations

import "encoding/json"

type ListDestinationsOkResponse struct {
	Destinations []Destinations `json:"destinations" xml:"destinations" required:"true"`
}

func (l ListDestinationsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListDestinationsOkResponse to string"
	}
	return string(jsonData)
}
