package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type TopUpEsimRequest struct {
	// Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces.
	IdempotencyKey *string `json:"-" url:"-" explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB. Use `-1` to top up with an unlimited (date-based) package — provide `startDate`/`endDate` spanning 3 to 30 days (`duration` is not supported for unlimited packages).
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *time.Time `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *time.Time `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Duration of the package in days. Available values are 1, 2, 7, 14, 30, or 90. Either provide startDate/endDate or duration. Not supported for unlimited packages (`dataLimitInGB` = -1), which are date-based — provide startDate/endDate instead.
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps).
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes. Letters, digits, `-`, `_`, `.`, `@` and `+` are allowed (so an email address works), up to 255 characters.
	ReferenceID *string `json:"referenceId,omitempty" xml:"referenceId,omitempty" maxLength:"255" pattern:"^[a-zA-Z0-9\-_.@+]+$"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty" xml:"emailBrand,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (t *TopUpEsimRequest) GetIdempotencyKey() (value *string) {
	if t == nil {
		return
	}
	return t.IdempotencyKey
}

func (t *TopUpEsimRequest) GetIccid() (value string) {
	if t == nil {
		return
	}
	return t.Iccid
}

func (t *TopUpEsimRequest) GetDataLimitInGb() (value float64) {
	if t == nil {
		return
	}
	return t.DataLimitInGb
}

func (t *TopUpEsimRequest) GetStartDate() (value *time.Time) {
	if t == nil {
		return
	}
	return t.StartDate
}

func (t *TopUpEsimRequest) GetEndDate() (value *time.Time) {
	if t == nil {
		return
	}
	return t.EndDate
}

func (t *TopUpEsimRequest) GetDuration() (value *float64) {
	if t == nil {
		return
	}
	return t.Duration
}

func (t *TopUpEsimRequest) GetEmail() (value *string) {
	if t == nil {
		return
	}
	return t.Email
}

func (t *TopUpEsimRequest) GetReferenceID() (value *string) {
	if t == nil {
		return
	}
	return t.ReferenceID
}

func (t *TopUpEsimRequest) GetEmailBrand() (value *string) {
	if t == nil {
		return
	}
	return t.EmailBrand
}

func (t *TopUpEsimRequest) GetStartTime() (value *float64) {
	if t == nil {
		return
	}
	return t.StartTime
}

func (t *TopUpEsimRequest) GetEndTime() (value *float64) {
	if t == nil {
		return
	}
	return t.EndTime
}

func (t *TopUpEsimRequest) SetIdempotencyKey(idempotencyKey *string) {
	t.IdempotencyKey = idempotencyKey
}

func (t *TopUpEsimRequest) SetIccid(iccid string) {
	t.Iccid = iccid
}

func (t *TopUpEsimRequest) SetDataLimitInGb(dataLimitInGb float64) {
	t.DataLimitInGb = dataLimitInGb
}

func (t *TopUpEsimRequest) SetStartDate(startDate *time.Time) {
	t.StartDate = startDate
}

func (t *TopUpEsimRequest) SetEndDate(endDate *time.Time) {
	t.EndDate = endDate
}

func (t *TopUpEsimRequest) SetDuration(duration *float64) {
	t.Duration = duration
}

func (t *TopUpEsimRequest) SetEmail(email *string) {
	t.Email = email
}

func (t *TopUpEsimRequest) SetReferenceID(referenceID *string) {
	t.ReferenceID = referenceID
}

func (t *TopUpEsimRequest) SetEmailBrand(emailBrand *string) {
	t.EmailBrand = emailBrand
}

func (t *TopUpEsimRequest) SetStartTime(startTime *float64) {
	t.StartTime = startTime
}

func (t *TopUpEsimRequest) SetEndTime(endTime *float64) {
	t.EndTime = endTime
}

func (t *TopUpEsimRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (t TopUpEsimRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimRequest to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimRequest) UnmarshalJSON(data []byte) error {
	data = internal.ParseDateFields(data, "startDate", "endDate")
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimRequest(tmp)
	return nil
}

func (t *TopUpEsimRequest) MarshalJSON() ([]byte, error) {
	type embed TopUpEsimRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return internal.FormatDateFields("startDate", "endDate")(json.Marshal(marshaler))
}
