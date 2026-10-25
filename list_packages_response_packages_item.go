package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListPackagesResponsePackagesItem struct {
	// ID of the package
	ID string `json:"id" xml:"id" required:"true"`
	// ISO3 representation of the package's destination.
	Destination string `json:"destination" xml:"destination" required:"true"`
	// ISO2 representation of the package's destination.
	DestinationIso2 string `json:"destinationISO2" xml:"destinationISO2" required:"true"`
	// Size of the package in Bytes. A value of `-1` indicates an unlimited package.
	DataLimitInBytes float64 `json:"dataLimitInBytes" xml:"dataLimitInBytes" required:"true"`
	// Size of the package in GB. A value of `-1` indicates an unlimited (date-based) package.
	DataLimitInGb float64 `json:"dataLimitInGB" xml:"dataLimitInGB" required:"true"`
	// Min number of days for the package
	MinDays float64 `json:"minDays" xml:"minDays" required:"true"`
	// Max number of days for the package
	MaxDays float64 `json:"maxDays" xml:"maxDays" required:"true"`
	// Price of the package in cents
	PriceInCents    float64 `json:"priceInCents" xml:"priceInCents" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListPackagesResponsePackagesItem) GetID() (value string) {
	if l == nil {
		return
	}
	return l.ID
}

func (l *ListPackagesResponsePackagesItem) GetDestination() (value string) {
	if l == nil {
		return
	}
	return l.Destination
}

func (l *ListPackagesResponsePackagesItem) GetDestinationIso2() (value string) {
	if l == nil {
		return
	}
	return l.DestinationIso2
}

func (l *ListPackagesResponsePackagesItem) GetDataLimitInBytes() (value float64) {
	if l == nil {
		return
	}
	return l.DataLimitInBytes
}

func (l *ListPackagesResponsePackagesItem) GetDataLimitInGb() (value float64) {
	if l == nil {
		return
	}
	return l.DataLimitInGb
}

func (l *ListPackagesResponsePackagesItem) GetMinDays() (value float64) {
	if l == nil {
		return
	}
	return l.MinDays
}

func (l *ListPackagesResponsePackagesItem) GetMaxDays() (value float64) {
	if l == nil {
		return
	}
	return l.MaxDays
}

func (l *ListPackagesResponsePackagesItem) GetPriceInCents() (value float64) {
	if l == nil {
		return
	}
	return l.PriceInCents
}

func (l *ListPackagesResponsePackagesItem) SetID(id string) {
	l.ID = id
}

func (l *ListPackagesResponsePackagesItem) SetDestination(destination string) {
	l.Destination = destination
}

func (l *ListPackagesResponsePackagesItem) SetDestinationIso2(destinationIso2 string) {
	l.DestinationIso2 = destinationIso2
}

func (l *ListPackagesResponsePackagesItem) SetDataLimitInBytes(dataLimitInBytes float64) {
	l.DataLimitInBytes = dataLimitInBytes
}

func (l *ListPackagesResponsePackagesItem) SetDataLimitInGb(dataLimitInGb float64) {
	l.DataLimitInGb = dataLimitInGb
}

func (l *ListPackagesResponsePackagesItem) SetMinDays(minDays float64) {
	l.MinDays = minDays
}

func (l *ListPackagesResponsePackagesItem) SetMaxDays(maxDays float64) {
	l.MaxDays = maxDays
}

func (l *ListPackagesResponsePackagesItem) SetPriceInCents(priceInCents float64) {
	l.PriceInCents = priceInCents
}

func (l *ListPackagesResponsePackagesItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPackagesResponsePackagesItem) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPackagesResponsePackagesItem to string"
	}
	return string(jsonData)
}

func (l *ListPackagesResponsePackagesItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias ListPackagesResponsePackagesItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = ListPackagesResponsePackagesItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListPackagesResponsePackagesItem) MarshalJSON() ([]byte, error) {
	type embed ListPackagesResponsePackagesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
