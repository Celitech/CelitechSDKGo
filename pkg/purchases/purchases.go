package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type Purchases struct {
	// ID of the purchase
	Id *string `json:"id,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *util.Nullable[string] `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *util.Nullable[string] `json:"endDate,omitempty" required:"true"`
	// Duration of the package in days. Possible values are 1, 2, 7, 14, 30, or 90.
	Duration *util.Nullable[float64] `json:"duration,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *util.Nullable[float64] `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *util.Nullable[float64] `json:"endTime,omitempty"`
	// Epoch value representing the date of creation of the purchase
	CreatedAt *float64       `json:"createdAt,omitempty"`
	Package_  *Package_      `json:"package,omitempty" required:"true"`
	Esim      *PurchasesEsim `json:"esim,omitempty" required:"true"`
	// The `source` indicates whether the purchase was made from the API, dashboard, landing-page, promo-page or iframe. For purchases made before September 8, 2023, the value will be displayed as 'Not available'.
	Source *string `json:"source,omitempty" required:"true"`
	// The `purchaseType` indicates whether this is the initial purchase that creates the eSIM (First Purchase) or a subsequent top-up on an existing eSIM (Top-up Purchase).
	PurchaseType *string `json:"purchaseType,omitempty" required:"true"`
	// The `referenceId` that was provided by the partner during the purchase or top-up flow. This identifier can be used for analytics and debugging purposes.
	ReferenceId *util.Nullable[string] `json:"referenceId,omitempty"`
}

func (p *Purchases) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Purchases) SetId(id string) {
	p.Id = &id
}

func (p *Purchases) GetStartDate() *util.Nullable[string] {
	if p == nil {
		return nil
	}
	return p.StartDate
}

func (p *Purchases) SetStartDate(startDate util.Nullable[string]) {
	p.StartDate = &startDate
}

func (p *Purchases) SetStartDateNull() {
	p.StartDate = &util.Nullable[string]{IsNull: true}
}

func (p *Purchases) GetEndDate() *util.Nullable[string] {
	if p == nil {
		return nil
	}
	return p.EndDate
}

func (p *Purchases) SetEndDate(endDate util.Nullable[string]) {
	p.EndDate = &endDate
}

func (p *Purchases) SetEndDateNull() {
	p.EndDate = &util.Nullable[string]{IsNull: true}
}

func (p *Purchases) GetDuration() *util.Nullable[float64] {
	if p == nil {
		return nil
	}
	return p.Duration
}

func (p *Purchases) SetDuration(duration util.Nullable[float64]) {
	p.Duration = &duration
}

func (p *Purchases) SetDurationNull() {
	p.Duration = &util.Nullable[float64]{IsNull: true}
}

func (p *Purchases) GetCreatedDate() *string {
	if p == nil {
		return nil
	}
	return p.CreatedDate
}

func (p *Purchases) SetCreatedDate(createdDate string) {
	p.CreatedDate = &createdDate
}

func (p *Purchases) GetStartTime() *util.Nullable[float64] {
	if p == nil {
		return nil
	}
	return p.StartTime
}

func (p *Purchases) SetStartTime(startTime util.Nullable[float64]) {
	p.StartTime = &startTime
}

func (p *Purchases) SetStartTimeNull() {
	p.StartTime = &util.Nullable[float64]{IsNull: true}
}

func (p *Purchases) GetEndTime() *util.Nullable[float64] {
	if p == nil {
		return nil
	}
	return p.EndTime
}

func (p *Purchases) SetEndTime(endTime util.Nullable[float64]) {
	p.EndTime = &endTime
}

func (p *Purchases) SetEndTimeNull() {
	p.EndTime = &util.Nullable[float64]{IsNull: true}
}

func (p *Purchases) GetCreatedAt() *float64 {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *Purchases) SetCreatedAt(createdAt float64) {
	p.CreatedAt = &createdAt
}

func (p *Purchases) GetPackage_() *Package_ {
	if p == nil {
		return nil
	}
	return p.Package_
}

func (p *Purchases) SetPackage_(package_ Package_) {
	p.Package_ = &package_
}

func (p *Purchases) GetEsim() *PurchasesEsim {
	if p == nil {
		return nil
	}
	return p.Esim
}

func (p *Purchases) SetEsim(esim PurchasesEsim) {
	p.Esim = &esim
}

func (p *Purchases) GetSource() *string {
	if p == nil {
		return nil
	}
	return p.Source
}

func (p *Purchases) SetSource(source string) {
	p.Source = &source
}

func (p *Purchases) GetPurchaseType() *string {
	if p == nil {
		return nil
	}
	return p.PurchaseType
}

func (p *Purchases) SetPurchaseType(purchaseType string) {
	p.PurchaseType = &purchaseType
}

func (p *Purchases) GetReferenceId() *util.Nullable[string] {
	if p == nil {
		return nil
	}
	return p.ReferenceId
}

func (p *Purchases) SetReferenceId(referenceId util.Nullable[string]) {
	p.ReferenceId = &referenceId
}

func (p *Purchases) SetReferenceIdNull() {
	p.ReferenceId = &util.Nullable[string]{IsNull: true}
}

func (p Purchases) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Purchases to string"
	}
	return string(jsonData)
}

func (p *Purchases) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}
