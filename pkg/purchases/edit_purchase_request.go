package purchases

import "encoding/json"

type EditPurchaseRequest struct {
	// ID of the purchase
	PurchaseId *string `json:"purchaseId,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
}

func (e *EditPurchaseRequest) GetPurchaseId() *string {
	if e == nil {
		return nil
	}
	return e.PurchaseId
}

func (e *EditPurchaseRequest) SetPurchaseId(purchaseId string) {
	e.PurchaseId = &purchaseId
}

func (e *EditPurchaseRequest) GetStartDate() *string {
	if e == nil {
		return nil
	}
	return e.StartDate
}

func (e *EditPurchaseRequest) SetStartDate(startDate string) {
	e.StartDate = &startDate
}

func (e *EditPurchaseRequest) GetEndDate() *string {
	if e == nil {
		return nil
	}
	return e.EndDate
}

func (e *EditPurchaseRequest) SetEndDate(endDate string) {
	e.EndDate = &endDate
}

func (e *EditPurchaseRequest) GetStartTime() *float64 {
	if e == nil {
		return nil
	}
	return e.StartTime
}

func (e *EditPurchaseRequest) SetStartTime(startTime float64) {
	e.StartTime = &startTime
}

func (e *EditPurchaseRequest) GetEndTime() *float64 {
	if e == nil {
		return nil
	}
	return e.EndTime
}

func (e *EditPurchaseRequest) SetEndTime(endTime float64) {
	e.EndTime = &endTime
}

func (e EditPurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseRequest to string"
	}
	return string(jsonData)
}
