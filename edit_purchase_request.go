package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type EditPurchaseRequest struct {
	// Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces.
	IdempotencyKey *string `json:"-" url:"-" explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
	// ID of the purchase
	PurchaseID string `json:"purchaseId" xml:"purchaseId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate time.Time `json:"startDate" xml:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate time.Time `json:"endDate" xml:"endDate" required:"true"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (e *EditPurchaseRequest) GetIdempotencyKey() (value *string) {
	if e == nil {
		return
	}
	return e.IdempotencyKey
}

func (e *EditPurchaseRequest) GetPurchaseID() (value string) {
	if e == nil {
		return
	}
	return e.PurchaseID
}

func (e *EditPurchaseRequest) GetStartDate() (value time.Time) {
	if e == nil {
		return
	}
	return e.StartDate
}

func (e *EditPurchaseRequest) GetEndDate() (value time.Time) {
	if e == nil {
		return
	}
	return e.EndDate
}

func (e *EditPurchaseRequest) GetStartTime() (value *float64) {
	if e == nil {
		return
	}
	return e.StartTime
}

func (e *EditPurchaseRequest) GetEndTime() (value *float64) {
	if e == nil {
		return
	}
	return e.EndTime
}

func (e *EditPurchaseRequest) SetIdempotencyKey(idempotencyKey *string) {
	e.IdempotencyKey = idempotencyKey
}

func (e *EditPurchaseRequest) SetPurchaseID(purchaseID string) {
	e.PurchaseID = purchaseID
}

func (e *EditPurchaseRequest) SetStartDate(startDate time.Time) {
	e.StartDate = startDate
}

func (e *EditPurchaseRequest) SetEndDate(endDate time.Time) {
	e.EndDate = endDate
}

func (e *EditPurchaseRequest) SetStartTime(startTime *float64) {
	e.StartTime = startTime
}

func (e *EditPurchaseRequest) SetEndTime(endTime *float64) {
	e.EndTime = endTime
}

func (e *EditPurchaseRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (e EditPurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseRequest to string"
	}
	return string(jsonData)
}

func (e *EditPurchaseRequest) UnmarshalJSON(data []byte) error {
	data = internal.ParseDateFields(data, "startDate", "endDate")
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

func (e *EditPurchaseRequest) MarshalJSON() ([]byte, error) {
	type embed EditPurchaseRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return internal.FormatDateFields("startDate", "endDate")(json.Marshal(marshaler))
}
