package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type TopUpEsimResponsePurchase struct {
	// ID of the purchase
	ID string `json:"id" xml:"id" required:"true"`
	// ID of the package
	PackageID string `json:"packageId" xml:"packageId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *time.Time `json:"startDate" xml:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *time.Time `json:"endDate" xml:"endDate" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate time.Time `json:"createdDate" xml:"createdDate" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *float64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime         *float64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TopUpEsimResponsePurchase) GetID() (value string) {
	if t == nil {
		return
	}
	return t.ID
}

func (t *TopUpEsimResponsePurchase) GetPackageID() (value string) {
	if t == nil {
		return
	}
	return t.PackageID
}

func (t *TopUpEsimResponsePurchase) GetStartDate() (value *time.Time) {
	if t == nil {
		return
	}
	return t.StartDate
}

func (t *TopUpEsimResponsePurchase) GetEndDate() (value *time.Time) {
	if t == nil {
		return
	}
	return t.EndDate
}

func (t *TopUpEsimResponsePurchase) GetCreatedDate() (value time.Time) {
	if t == nil {
		return
	}
	return t.CreatedDate
}

func (t *TopUpEsimResponsePurchase) GetStartTime() (value *float64) {
	if t == nil {
		return
	}
	return t.StartTime
}

func (t *TopUpEsimResponsePurchase) GetEndTime() (value *float64) {
	if t == nil {
		return
	}
	return t.EndTime
}

func (t *TopUpEsimResponsePurchase) SetID(id string) {
	t.ID = id
}

func (t *TopUpEsimResponsePurchase) SetPackageID(packageID string) {
	t.PackageID = packageID
}

func (t *TopUpEsimResponsePurchase) SetStartDate(startDate *time.Time) {
	t.StartDate = startDate
}

func (t *TopUpEsimResponsePurchase) SetEndDate(endDate *time.Time) {
	t.EndDate = endDate
}

func (t *TopUpEsimResponsePurchase) SetCreatedDate(createdDate time.Time) {
	t.CreatedDate = createdDate
}

func (t *TopUpEsimResponsePurchase) SetStartTime(startTime *float64) {
	t.StartTime = startTime
}

func (t *TopUpEsimResponsePurchase) SetEndTime(endTime *float64) {
	t.EndTime = endTime
}

func (t *TopUpEsimResponsePurchase) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopUpEsimResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimResponsePurchase to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimResponsePurchase) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopUpEsimResponsePurchase) MarshalJSON() ([]byte, error) {
	type embed TopUpEsimResponsePurchase
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
