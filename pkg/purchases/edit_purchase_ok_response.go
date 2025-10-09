package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type EditPurchaseOkResponse struct {
	// ID of the purchase
	PurchaseId *string `json:"purchaseId,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewStartDate *util.Nullable[string] `json:"newStartDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewEndDate *util.Nullable[string] `json:"newEndDate,omitempty" required:"true"`
	// Epoch value representing the new start time of the package's validity
	NewStartTime *util.Nullable[float64] `json:"newStartTime,omitempty"`
	// Epoch value representing the new end time of the package's validity
	NewEndTime *util.Nullable[float64] `json:"newEndTime,omitempty"`
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

func (e *EditPurchaseOkResponse) GetNewStartDate() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.NewStartDate
}

func (e *EditPurchaseOkResponse) SetNewStartDate(newStartDate util.Nullable[string]) {
	e.NewStartDate = &newStartDate
}

func (e *EditPurchaseOkResponse) SetNewStartDateNull() {
	e.NewStartDate = &util.Nullable[string]{IsNull: true}
}

func (e *EditPurchaseOkResponse) GetNewEndDate() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.NewEndDate
}

func (e *EditPurchaseOkResponse) SetNewEndDate(newEndDate util.Nullable[string]) {
	e.NewEndDate = &newEndDate
}

func (e *EditPurchaseOkResponse) SetNewEndDateNull() {
	e.NewEndDate = &util.Nullable[string]{IsNull: true}
}

func (e *EditPurchaseOkResponse) GetNewStartTime() *util.Nullable[float64] {
	if e == nil {
		return nil
	}
	return e.NewStartTime
}

func (e *EditPurchaseOkResponse) SetNewStartTime(newStartTime util.Nullable[float64]) {
	e.NewStartTime = &newStartTime
}

func (e *EditPurchaseOkResponse) SetNewStartTimeNull() {
	e.NewStartTime = &util.Nullable[float64]{IsNull: true}
}

func (e *EditPurchaseOkResponse) GetNewEndTime() *util.Nullable[float64] {
	if e == nil {
		return nil
	}
	return e.NewEndTime
}

func (e *EditPurchaseOkResponse) SetNewEndTime(newEndTime util.Nullable[float64]) {
	e.NewEndTime = &newEndTime
}

func (e *EditPurchaseOkResponse) SetNewEndTimeNull() {
	e.NewEndTime = &util.Nullable[float64]{IsNull: true}
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
