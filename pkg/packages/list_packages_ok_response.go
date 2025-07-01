package packages

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type ListPackagesOkResponse struct {
	Packages []Packages `json:"packages,omitempty"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination
	AfterCursor *util.Nullable[string] `json:"afterCursor,omitempty"`
}

func (l *ListPackagesOkResponse) GetPackages() []Packages {
	if l == nil {
		return nil
	}
	return l.Packages
}

func (l *ListPackagesOkResponse) SetPackages(packages []Packages) {
	l.Packages = packages
}

func (l *ListPackagesOkResponse) GetAfterCursor() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursor(afterCursor util.Nullable[string]) {
	l.AfterCursor = &afterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursorNull() {
	l.AfterCursor = &util.Nullable[string]{IsNull: true}
}

func (l ListPackagesOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPackagesOkResponse to string"
	}
	return string(jsonData)
}

func (l *ListPackagesOkResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

type Packages struct {
	// ID of the package
	Id *string `json:"id,omitempty"`
	// ISO representation of the package's destination.
	Destination *string `json:"destination,omitempty"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty"`
	// Min number of days for the package
	MinDays *float64 `json:"minDays,omitempty"`
	// Max number of days for the package
	MaxDays *float64 `json:"maxDays,omitempty"`
	// Price of the package in cents
	PriceInCents *float64 `json:"priceInCents,omitempty"`
}

func (p *Packages) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Packages) SetId(id string) {
	p.Id = &id
}

func (p *Packages) GetDestination() *string {
	if p == nil {
		return nil
	}
	return p.Destination
}

func (p *Packages) SetDestination(destination string) {
	p.Destination = &destination
}

func (p *Packages) GetDataLimitInBytes() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInBytes
}

func (p *Packages) SetDataLimitInBytes(dataLimitInBytes float64) {
	p.DataLimitInBytes = &dataLimitInBytes
}

func (p *Packages) GetMinDays() *float64 {
	if p == nil {
		return nil
	}
	return p.MinDays
}

func (p *Packages) SetMinDays(minDays float64) {
	p.MinDays = &minDays
}

func (p *Packages) GetMaxDays() *float64 {
	if p == nil {
		return nil
	}
	return p.MaxDays
}

func (p *Packages) SetMaxDays(maxDays float64) {
	p.MaxDays = &maxDays
}

func (p *Packages) GetPriceInCents() *float64 {
	if p == nil {
		return nil
	}
	return p.PriceInCents
}

func (p *Packages) SetPriceInCents(priceInCents float64) {
	p.PriceInCents = &priceInCents
}

func (p Packages) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Packages to string"
	}
	return string(jsonData)
}
