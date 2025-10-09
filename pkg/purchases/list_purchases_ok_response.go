package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type ListPurchasesOkResponse struct {
	Purchases []Purchases `json:"purchases,omitempty" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination.
	AfterCursor *util.Nullable[string] `json:"afterCursor,omitempty" required:"true"`
}

func (l *ListPurchasesOkResponse) GetPurchases() []Purchases {
	if l == nil {
		return nil
	}
	return l.Purchases
}

func (l *ListPurchasesOkResponse) SetPurchases(purchases []Purchases) {
	l.Purchases = purchases
}

func (l *ListPurchasesOkResponse) GetAfterCursor() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursor(afterCursor util.Nullable[string]) {
	l.AfterCursor = &afterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursorNull() {
	l.AfterCursor = &util.Nullable[string]{IsNull: true}
}

func (l ListPurchasesOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesOkResponse to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesOkResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

type Purchases struct {
	// ID of the purchase
	Id *string `json:"id,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *util.Nullable[string] `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *util.Nullable[string] `json:"endDate,omitempty" required:"true"`
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

type Package_ struct {
	// ID of the package
	Id *string `json:"id,omitempty" required:"true"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty" required:"true"`
	// ISO3 representation of the package's destination.
	Destination *string `json:"destination,omitempty" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 *string `json:"destinationISO2,omitempty" required:"true"`
	// Name of the package's destination
	DestinationName *string `json:"destinationName,omitempty" required:"true"`
	// Price of the package in cents
	PriceInCents *float64 `json:"priceInCents,omitempty" required:"true"`
}

func (p *Package_) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Package_) SetId(id string) {
	p.Id = &id
}

func (p *Package_) GetDataLimitInBytes() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInBytes
}

func (p *Package_) SetDataLimitInBytes(dataLimitInBytes float64) {
	p.DataLimitInBytes = &dataLimitInBytes
}

func (p *Package_) GetDestination() *string {
	if p == nil {
		return nil
	}
	return p.Destination
}

func (p *Package_) SetDestination(destination string) {
	p.Destination = &destination
}

func (p *Package_) GetDestinationIso2() *string {
	if p == nil {
		return nil
	}
	return p.DestinationIso2
}

func (p *Package_) SetDestinationIso2(destinationIso2 string) {
	p.DestinationIso2 = &destinationIso2
}

func (p *Package_) GetDestinationName() *string {
	if p == nil {
		return nil
	}
	return p.DestinationName
}

func (p *Package_) SetDestinationName(destinationName string) {
	p.DestinationName = &destinationName
}

func (p *Package_) GetPriceInCents() *float64 {
	if p == nil {
		return nil
	}
	return p.PriceInCents
}

func (p *Package_) SetPriceInCents(priceInCents float64) {
	p.PriceInCents = &priceInCents
}

func (p Package_) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Package_ to string"
	}
	return string(jsonData)
}

type PurchasesEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
}

func (p *PurchasesEsim) GetIccid() *string {
	if p == nil {
		return nil
	}
	return p.Iccid
}

func (p *PurchasesEsim) SetIccid(iccid string) {
	p.Iccid = &iccid
}

func (p PurchasesEsim) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PurchasesEsim to string"
	}
	return string(jsonData)
}
