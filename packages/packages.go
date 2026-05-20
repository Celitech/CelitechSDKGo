package packages

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type Packages struct {
	// ID of the package
	ID string `json:"id" xml:"id" required:"true"`
	// ISO3 representation of the package's destination.
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// Size of the package in Bytes
	DataLimitInBytes float64 `json:"dataLimitInBytes" xml:"dataLimitInBytes" required:"true"`
	// Size of the package in GB
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// Min number of days for the package
	MinDays float64 `json:"minDays" xml:"minDays" required:"true"`
	// Max number of days for the package
	MaxDays float64 `json:"maxDays" xml:"maxDays" required:"true"`
	// Price of the package in cents
	PriceInCents float64 `json:"priceInCents" xml:"priceInCents" required:"true"`
}

func (p Packages) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Packages to string"
	}
	return string(jsonData)
}

func (p *Packages) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias Packages
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = Packages(tmp)
	return nil
}
