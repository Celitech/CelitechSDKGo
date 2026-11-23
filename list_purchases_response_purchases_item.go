package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type ListPurchasesResponsePurchasesItem struct {
	// ID of the purchase
	ID string `json:"id" xml:"id" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *time.Time `json:"startDate" xml:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *time.Time `json:"endDate" xml:"endDate" required:"true"`
	// Duration of the package in days. Possible values are 1, 2, 7, 14, 30, or 90. `null` for unlimited (date-based) packages.
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate time.Time `json:"createdDate" xml:"createdDate" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *float64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *float64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Epoch value representing the date of creation of the purchase
	CreatedAt *float64                                   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Package   *ListPurchasesResponsePurchasesItemPackage `json:"package" xml:"package" required:"true"`
	Esim      *ListPurchasesResponsePurchasesItemEsim    `json:"esim" xml:"esim" required:"true"`
	// The `source` indicates whether the purchase was made from the API, dashboard, landing-page, promo-page or iframe. For purchases made before September 8, 2023, the value will be displayed as 'Not available'.
	Source string `json:"source" xml:"source" required:"true"`
	// The `purchaseType` indicates whether this is the initial purchase that creates the eSIM (First Purchase) or a subsequent top-up on an existing eSIM (Top-up Purchase).
	PurchaseType string `json:"purchaseType" xml:"purchaseType" required:"true"`
	// The `referenceId` that was provided by the partner during the purchase or top-up flow. This identifier can be used for analytics and debugging purposes.
	ReferenceID     *string `json:"referenceId,omitempty" xml:"referenceId,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListPurchasesResponsePurchasesItem) GetID() (value string) {
	if l == nil {
		return
	}
	return l.ID
}

func (l *ListPurchasesResponsePurchasesItem) GetStartDate() (value *time.Time) {
	if l == nil {
		return
	}
	return l.StartDate
}

func (l *ListPurchasesResponsePurchasesItem) GetEndDate() (value *time.Time) {
	if l == nil {
		return
	}
	return l.EndDate
}

func (l *ListPurchasesResponsePurchasesItem) GetDuration() (value *float64) {
	if l == nil {
		return
	}
	return l.Duration
}

func (l *ListPurchasesResponsePurchasesItem) GetCreatedDate() (value time.Time) {
	if l == nil {
		return
	}
	return l.CreatedDate
}

func (l *ListPurchasesResponsePurchasesItem) GetStartTime() (value *float64) {
	if l == nil {
		return
	}
	return l.StartTime
}

func (l *ListPurchasesResponsePurchasesItem) GetEndTime() (value *float64) {
	if l == nil {
		return
	}
	return l.EndTime
}

func (l *ListPurchasesResponsePurchasesItem) GetCreatedAt() (value *float64) {
	if l == nil {
		return
	}
	return l.CreatedAt
}

func (l *ListPurchasesResponsePurchasesItem) GetPackage() (value *ListPurchasesResponsePurchasesItemPackage) {
	if l == nil {
		return
	}
	return l.Package
}

func (l *ListPurchasesResponsePurchasesItem) GetEsim() (value *ListPurchasesResponsePurchasesItemEsim) {
	if l == nil {
		return
	}
	return l.Esim
}

func (l *ListPurchasesResponsePurchasesItem) GetSource() (value string) {
	if l == nil {
		return
	}
	return l.Source
}

func (l *ListPurchasesResponsePurchasesItem) GetPurchaseType() (value string) {
	if l == nil {
		return
	}
	return l.PurchaseType
}

func (l *ListPurchasesResponsePurchasesItem) GetReferenceID() (value *string) {
	if l == nil {
		return
	}
	return l.ReferenceID
}

func (l *ListPurchasesResponsePurchasesItem) SetID(id string) {
	l.ID = id
}

func (l *ListPurchasesResponsePurchasesItem) SetStartDate(startDate *time.Time) {
	l.StartDate = startDate
}

func (l *ListPurchasesResponsePurchasesItem) SetEndDate(endDate *time.Time) {
	l.EndDate = endDate
}

func (l *ListPurchasesResponsePurchasesItem) SetDuration(duration *float64) {
	l.Duration = duration
}

func (l *ListPurchasesResponsePurchasesItem) SetCreatedDate(createdDate time.Time) {
	l.CreatedDate = createdDate
}

func (l *ListPurchasesResponsePurchasesItem) SetStartTime(startTime *float64) {
	l.StartTime = startTime
}

func (l *ListPurchasesResponsePurchasesItem) SetEndTime(endTime *float64) {
	l.EndTime = endTime
}

func (l *ListPurchasesResponsePurchasesItem) SetCreatedAt(createdAt *float64) {
	l.CreatedAt = createdAt
}

func (l *ListPurchasesResponsePurchasesItem) SetPackage(package_ *ListPurchasesResponsePurchasesItemPackage) {
	l.Package = package_
}

func (l *ListPurchasesResponsePurchasesItem) SetEsim(esim *ListPurchasesResponsePurchasesItemEsim) {
	l.Esim = esim
}

func (l *ListPurchasesResponsePurchasesItem) SetSource(source string) {
	l.Source = source
}

func (l *ListPurchasesResponsePurchasesItem) SetPurchaseType(purchaseType string) {
	l.PurchaseType = purchaseType
}

func (l *ListPurchasesResponsePurchasesItem) SetReferenceID(referenceID *string) {
	l.ReferenceID = referenceID
}

func (l *ListPurchasesResponsePurchasesItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPurchasesResponsePurchasesItem) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesResponsePurchasesItem to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesResponsePurchasesItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListPurchasesResponsePurchasesItem) MarshalJSON() ([]byte, error) {
	type embed ListPurchasesResponsePurchasesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
