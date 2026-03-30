package packages

import "encoding/json"

type Packages struct {
	// ID of the package
	Id *string `json:"id,omitempty" required:"true"`
	// ISO3 representation of the package's destination.
	Destination *string `json:"destination,omitempty" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 *string `json:"destinationISO2,omitempty" required:"true"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty" required:"true"`
	// Size of the package in GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Min number of days for the package
	MinDays *float64 `json:"minDays,omitempty" required:"true"`
	// Max number of days for the package
	MaxDays *float64 `json:"maxDays,omitempty" required:"true"`
	// Price of the package in cents
	PriceInCents *float64 `json:"priceInCents,omitempty" required:"true"`
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

func (p *Packages) GetDestinationIso2() *string {
	if p == nil {
		return nil
	}
	return p.DestinationIso2
}

func (p *Packages) SetDestinationIso2(destinationIso2 string) {
	p.DestinationIso2 = &destinationIso2
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

func (p *Packages) GetDataLimitInGb() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInGb
}

func (p *Packages) SetDataLimitInGb(dataLimitInGb float64) {
	p.DataLimitInGb = &dataLimitInGb
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
