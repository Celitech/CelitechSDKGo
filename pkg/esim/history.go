package esim

import "encoding/json"

type History struct {
	// The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status *string `json:"status,omitempty" required:"true"`
	// The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StatusDate *string `json:"statusDate,omitempty" required:"true"`
	// Epoch value representing the date when the eSIM status changed
	Date *float64 `json:"date,omitempty"`
}

func (h *History) GetStatus() *string {
	if h == nil {
		return nil
	}
	return h.Status
}

func (h *History) SetStatus(status string) {
	h.Status = &status
}

func (h *History) GetStatusDate() *string {
	if h == nil {
		return nil
	}
	return h.StatusDate
}

func (h *History) SetStatusDate(statusDate string) {
	h.StatusDate = &statusDate
}

func (h *History) GetDate() *float64 {
	if h == nil {
		return nil
	}
	return h.Date
}

func (h *History) SetDate(date float64) {
	h.Date = &date
}

func (h History) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: History to string"
	}
	return string(jsonData)
}
