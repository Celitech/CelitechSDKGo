package purchases

import "encoding/json"

type Package_ struct {
	// ID of the package
	Id *string `json:"id,omitempty" required:"true"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty" required:"true"`
	// Size of the package in GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
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

func (p *Package_) GetDataLimitInGb() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInGb
}

func (p *Package_) SetDataLimitInGb(dataLimitInGb float64) {
	p.DataLimitInGb = &dataLimitInGb
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
