package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListPurchasesResponsePurchasesItemPackage struct {
	// ID of the package
	ID string `json:"id" xml:"id" required:"true"`
	// Size of the package in Bytes. A value of `-1` indicates an unlimited package.
	DataLimitInBytes float64 `json:"dataLimitInBytes" xml:"dataLimitInBytes" required:"true"`
	// Size of the package in GB. A value of `-1` indicates an unlimited (date-based) package.
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// ISO3 representation of the package's destination.
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// Name of the package's destination
	DestinationName string `json:"destinationName" xml:"destinationName" required:"true"`
	// Price of the package in cents
	PriceInCents    float64 `json:"priceInCents" xml:"priceInCents" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetID() (value string) {
	if l == nil {
		return
	}
	return l.ID
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetDataLimitInBytes() (value float64) {
	if l == nil {
		return
	}
	return l.DataLimitInBytes
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetDataLimitInGb() (value float64) {
	if l == nil {
		return
	}
	return l.DataLimitInGb
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetDestination() (value string) {
	if l == nil {
		return
	}
	return l.Destination
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetDestinationIso2() (value string) {
	if l == nil {
		return
	}
	return l.DestinationIso2
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetDestinationName() (value string) {
	if l == nil {
		return
	}
	return l.DestinationName
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetPriceInCents() (value float64) {
	if l == nil {
		return
	}
	return l.PriceInCents
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetID(id string) {
	l.ID = id
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetDataLimitInBytes(dataLimitInBytes float64) {
	l.DataLimitInBytes = dataLimitInBytes
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetDataLimitInGb(dataLimitInGb float64) {
	l.DataLimitInGb = dataLimitInGb
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetDestination(destination string) {
	l.Destination = destination
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetDestinationIso2(destinationIso2 string) {
	l.DestinationIso2 = destinationIso2
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetDestinationName(destinationName string) {
	l.DestinationName = destinationName
}

func (l *ListPurchasesResponsePurchasesItemPackage) SetPriceInCents(priceInCents float64) {
	l.PriceInCents = priceInCents
}

func (l *ListPurchasesResponsePurchasesItemPackage) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPurchasesResponsePurchasesItemPackage) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesResponsePurchasesItemPackage to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesResponsePurchasesItemPackage) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias ListPurchasesResponsePurchasesItemPackage
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = ListPurchasesResponsePurchasesItemPackage(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListPurchasesResponsePurchasesItemPackage) MarshalJSON() ([]byte, error) {
	type embed ListPurchasesResponsePurchasesItemPackage
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
