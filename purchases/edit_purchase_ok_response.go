package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
	"example.com/celitech/param"
)

type EditPurchaseOkResponse struct {
	// ID of the purchase
	PurchaseID string `json:"purchaseId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewStartDate *param.Nullable[string] `json:"newStartDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewEndDate *param.Nullable[string] `json:"newEndDate" required:"true"`
	// Epoch value representing the new start time of the package's validity
	NewStartTime *param.Nullable[float64] `json:"newStartTime,omitempty"`
	// Epoch value representing the new end time of the package's validity
	NewEndTime *param.Nullable[float64] `json:"newEndTime,omitempty"`
}

func (e EditPurchaseOkResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseOkResponse to string"
	}
	return string(jsonData)
}

func (e *EditPurchaseOkResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}
