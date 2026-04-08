package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type EditPurchaseRequest struct {
	// ID of the purchase
	PurchaseID string `json:"purchaseId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate string `json:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate string `json:"endDate" required:"true"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
}

func (e EditPurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseRequest to string"
	}
	return string(jsonData)
}

func (e *EditPurchaseRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditPurchaseRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditPurchaseRequest(tmp)
	return nil
}
