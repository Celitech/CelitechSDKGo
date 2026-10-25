package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type EditPurchaseResponse struct {
	// ID of the purchase
	PurchaseID string `json:"purchaseId" xml:"purchaseId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewStartDate *time.Time `json:"newStartDate" xml:"newStartDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewEndDate *time.Time `json:"newEndDate" xml:"newEndDate" required:"true"`
	// Epoch value representing the new start time of the package's validity
	NewStartTime *float64 `json:"newStartTime,omitempty" xml:"newStartTime,omitempty"`
	// Epoch value representing the new end time of the package's validity
	NewEndTime      *float64 `json:"newEndTime,omitempty" xml:"newEndTime,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EditPurchaseResponse) GetPurchaseID() (value string) {
	if e == nil {
		return
	}
	return e.PurchaseID
}

func (e *EditPurchaseResponse) GetNewStartDate() (value *time.Time) {
	if e == nil {
		return
	}
	return e.NewStartDate
}

func (e *EditPurchaseResponse) GetNewEndDate() (value *time.Time) {
	if e == nil {
		return
	}
	return e.NewEndDate
}

func (e *EditPurchaseResponse) GetNewStartTime() (value *float64) {
	if e == nil {
		return
	}
	return e.NewStartTime
}

func (e *EditPurchaseResponse) GetNewEndTime() (value *float64) {
	if e == nil {
		return
	}
	return e.NewEndTime
}

func (e *EditPurchaseResponse) SetPurchaseID(purchaseID string) {
	e.PurchaseID = purchaseID
}

func (e *EditPurchaseResponse) SetNewStartDate(newStartDate *time.Time) {
	e.NewStartDate = newStartDate
}

func (e *EditPurchaseResponse) SetNewEndDate(newEndDate *time.Time) {
	e.NewEndDate = newEndDate
}

func (e *EditPurchaseResponse) SetNewStartTime(newStartTime *float64) {
	e.NewStartTime = newStartTime
}

func (e *EditPurchaseResponse) SetNewEndTime(newEndTime *float64) {
	e.NewEndTime = newEndTime
}

func (e *EditPurchaseResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditPurchaseResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseResponse to string"
	}
	return string(jsonData)
}

func (e *EditPurchaseResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditPurchaseResponse) MarshalJSON() ([]byte, error) {
	type embed EditPurchaseResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
