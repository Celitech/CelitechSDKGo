package purchases

import "encoding/json"

type EditPurchaseOkResponse struct {
	// ID of the purchase
	PurchaseId *string `json:"purchaseId,omitempty"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewStartDate *string `json:"newStartDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewEndDate *string `json:"newEndDate,omitempty"`
	// Epoch value representing the new start time of the package's validity
	NewStartTime *float64 `json:"newStartTime,omitempty"`
	// Epoch value representing the new end time of the package's validity
	NewEndTime *float64 `json:"newEndTime,omitempty"`
}

func (e *EditPurchaseOkResponse) GetPurchaseId() *string {
	if e == nil {
		return nil
	}
	return e.PurchaseId
}

func (e *EditPurchaseOkResponse) SetPurchaseId(purchaseId string) {
	e.PurchaseId = &purchaseId
}

func (e *EditPurchaseOkResponse) GetNewStartDate() *string {
	if e == nil {
		return nil
	}
	return e.NewStartDate
}

func (e *EditPurchaseOkResponse) SetNewStartDate(newStartDate string) {
	e.NewStartDate = &newStartDate
}

func (e *EditPurchaseOkResponse) GetNewEndDate() *string {
	if e == nil {
		return nil
	}
	return e.NewEndDate
}

func (e *EditPurchaseOkResponse) SetNewEndDate(newEndDate string) {
	e.NewEndDate = &newEndDate
}

func (e *EditPurchaseOkResponse) GetNewStartTime() *float64 {
	if e == nil {
		return nil
	}
	return e.NewStartTime
}

func (e *EditPurchaseOkResponse) SetNewStartTime(newStartTime float64) {
	e.NewStartTime = &newStartTime
}

func (e *EditPurchaseOkResponse) GetNewEndTime() *float64 {
	if e == nil {
		return nil
	}
	return e.NewEndTime
}

func (e *EditPurchaseOkResponse) SetNewEndTime(newEndTime float64) {
	e.NewEndTime = &newEndTime
}

func (e EditPurchaseOkResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseOkResponse to string"
	}
	return string(jsonData)
}
