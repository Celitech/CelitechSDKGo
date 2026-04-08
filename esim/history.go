package esim

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type History struct {
	// The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status string `json:"status" required:"true"`
	// The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StatusDate string `json:"statusDate" required:"true"`
	// Epoch value representing the date when the eSIM status changed
	Date *float64 `json:"date,omitempty"`
}

func (h History) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: History to string"
	}
	return string(jsonData)
}

func (h *History) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, h); err != nil {
		return err
	}
	type alias History
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*h = History(tmp)
	return nil
}
