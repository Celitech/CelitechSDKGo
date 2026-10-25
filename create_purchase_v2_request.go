package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type CreatePurchaseV2Request struct {
	// Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces.
	IdempotencyKey *string `json:"-" url:"-" explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
	// ISO representation of the package's destination. Supports both ISO2 (e.g., 'FR') and ISO3 (e.g., 'FRA') country codes.
	Destination string `json:"destination" xml:"destination" required:"true"`
	// Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB. Use `-1` to purchase an unlimited (date-based) package — provide `startDate`/`endDate` spanning 3 to 30 days (`duration` is not supported for unlimited packages).
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *time.Time `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *time.Time `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Duration of the package in days. Available values are 1, 2, 7, 14, 30, or 90. Either provide startDate/endDate or duration. Not supported for unlimited packages (`dataLimitInGB` = -1), which are date-based — provide startDate/endDate instead.
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Number of eSIMs to purchase.
	Quantity float64 `json:"quantity" xml:"quantity" required:"true" min:"1" max:"5"`
	// Email address where the purchase confirmation email will be sent (including QR Code & activation steps)
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes. Letters, digits, `-`, `_`, `.`, `@` and `+` are allowed (so an email address works), up to 255 characters.
	ReferenceID *string `json:"referenceId,omitempty" xml:"referenceId,omitempty" maxLength:"255" pattern:"^[a-zA-Z0-9\-_.@+]+$"`
	// Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters, numbers, dots (.), ampersands (&), and spaces. This feature is available to platforms with Diamond tier only.
	NetworkBrand *string `json:"networkBrand,omitempty" xml:"networkBrand,omitempty"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty" xml:"emailBrand,omitempty"`
	// Language of the confirmation email sent to the customer.
	Language *CreatePurchaseV2RequestLanguage `json:"language,omitempty" xml:"language,omitempty"`
}

func (c *CreatePurchaseV2Request) GetIdempotencyKey() (value *string) {
	if c == nil {
		return
	}
	return c.IdempotencyKey
}

func (c *CreatePurchaseV2Request) GetDestination() (value string) {
	if c == nil {
		return
	}
	return c.Destination
}

func (c *CreatePurchaseV2Request) GetDataLimitInGb() (value float64) {
	if c == nil {
		return
	}
	return c.DataLimitInGb
}

func (c *CreatePurchaseV2Request) GetStartDate() (value *time.Time) {
	if c == nil {
		return
	}
	return c.StartDate
}

func (c *CreatePurchaseV2Request) GetEndDate() (value *time.Time) {
	if c == nil {
		return
	}
	return c.EndDate
}

func (c *CreatePurchaseV2Request) GetDuration() (value *float64) {
	if c == nil {
		return
	}
	return c.Duration
}

func (c *CreatePurchaseV2Request) GetQuantity() (value float64) {
	if c == nil {
		return
	}
	return c.Quantity
}

func (c *CreatePurchaseV2Request) GetEmail() (value *string) {
	if c == nil {
		return
	}
	return c.Email
}

func (c *CreatePurchaseV2Request) GetReferenceID() (value *string) {
	if c == nil {
		return
	}
	return c.ReferenceID
}

func (c *CreatePurchaseV2Request) GetNetworkBrand() (value *string) {
	if c == nil {
		return
	}
	return c.NetworkBrand
}

func (c *CreatePurchaseV2Request) GetEmailBrand() (value *string) {
	if c == nil {
		return
	}
	return c.EmailBrand
}

func (c *CreatePurchaseV2Request) GetLanguage() (value *CreatePurchaseV2RequestLanguage) {
	if c == nil {
		return
	}
	return c.Language
}

func (c *CreatePurchaseV2Request) SetIdempotencyKey(idempotencyKey *string) {
	c.IdempotencyKey = idempotencyKey
}

func (c *CreatePurchaseV2Request) SetDestination(destination string) {
	c.Destination = destination
}

func (c *CreatePurchaseV2Request) SetDataLimitInGb(dataLimitInGb float64) {
	c.DataLimitInGb = dataLimitInGb
}

func (c *CreatePurchaseV2Request) SetStartDate(startDate *time.Time) {
	c.StartDate = startDate
}

func (c *CreatePurchaseV2Request) SetEndDate(endDate *time.Time) {
	c.EndDate = endDate
}

func (c *CreatePurchaseV2Request) SetDuration(duration *float64) {
	c.Duration = duration
}

func (c *CreatePurchaseV2Request) SetQuantity(quantity float64) {
	c.Quantity = quantity
}

func (c *CreatePurchaseV2Request) SetEmail(email *string) {
	c.Email = email
}

func (c *CreatePurchaseV2Request) SetReferenceID(referenceID *string) {
	c.ReferenceID = referenceID
}

func (c *CreatePurchaseV2Request) SetNetworkBrand(networkBrand *string) {
	c.NetworkBrand = networkBrand
}

func (c *CreatePurchaseV2Request) SetEmailBrand(emailBrand *string) {
	c.EmailBrand = emailBrand
}

func (c *CreatePurchaseV2Request) SetLanguage(language *CreatePurchaseV2RequestLanguage) {
	c.Language = language
}

func (c *CreatePurchaseV2Request) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreatePurchaseV2Request) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2Request to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2Request) UnmarshalJSON(data []byte) error {
	data = internal.ParseDateFields(data, "startDate", "endDate")
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2Request
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2Request(tmp)
	return nil
}

func (c *CreatePurchaseV2Request) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseV2Request
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return internal.FormatDateFields("startDate", "endDate")(json.Marshal(marshaler))
}
