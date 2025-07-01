package esim

import "encoding/json"

type GetEsimHistoryOkResponse struct {
	Esim *GetEsimHistoryOkResponseEsim `json:"esim,omitempty"`
}

func (g *GetEsimHistoryOkResponse) GetEsim() *GetEsimHistoryOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimHistoryOkResponse) SetEsim(esim GetEsimHistoryOkResponseEsim) {
	g.Esim = &esim
}

func (g GetEsimHistoryOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponse to string"
	}
	return string(jsonData)
}

type GetEsimHistoryOkResponseEsim struct {
	// ID of the eSIM
	Iccid   *string   `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	History []History `json:"history,omitempty"`
}

func (g *GetEsimHistoryOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimHistoryOkResponseEsim) SetIccid(iccid string) {
	g.Iccid = &iccid
}

func (g *GetEsimHistoryOkResponseEsim) GetHistory() []History {
	if g == nil {
		return nil
	}
	return g.History
}

func (g *GetEsimHistoryOkResponseEsim) SetHistory(history []History) {
	g.History = history
}

func (g GetEsimHistoryOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponseEsim to string"
	}
	return string(jsonData)
}

type History struct {
	// The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status *string `json:"status,omitempty"`
	// The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StatusDate *string `json:"statusDate,omitempty"`
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
