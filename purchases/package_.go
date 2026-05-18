package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type Package_ struct {
	// ID of the package
	ID string `json:"id" xml:"id" required:"true"`
	// Size of the package in Bytes
	DataLimitInBytes float64 `json:"dataLimitInBytes" xml:"dataLimitInBytes" required:"true"`
	// Size of the package in GB
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// ISO3 representation of the package's destination.
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// Name of the package's destination
	DestinationName string `json:"destinationName" xml:"destinationName" required:"true"`
	// Price of the package in cents
	PriceInCents float64 `json:"priceInCents" xml:"priceInCents" required:"true"`
}

func (p Package_) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Package_ to string"
	}
	return string(jsonData)
}

func (p *Package_) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias Package_
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = Package_(tmp)
	return nil
}
