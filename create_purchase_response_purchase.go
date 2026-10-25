package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type CreatePurchaseResponsePurchase struct {
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

func (c *CreatePurchaseResponsePurchase) GetID() (value string) {
	if c == nil {
		return
	}
	return c.ID
}

func (c *CreatePurchaseResponsePurchase) GetPackageID() (value string) {
	if c == nil {
		return
	}
	return c.PackageID
}

func (c *CreatePurchaseResponsePurchase) GetStartDate() (value *time.Time) {
	if c == nil {
		return
	}
	return c.StartDate
}

func (c *CreatePurchaseResponsePurchase) GetEndDate() (value *time.Time) {
	if c == nil {
		return
	}
	return c.EndDate
}

func (c *CreatePurchaseResponsePurchase) GetCreatedDate() (value time.Time) {
	if c == nil {
		return
	}
	return c.CreatedDate
}

func (c *CreatePurchaseResponsePurchase) GetStartTime() (value *float64) {
	if c == nil {
		return
	}
	return c.StartTime
}

func (c *CreatePurchaseResponsePurchase) GetEndTime() (value *float64) {
	if c == nil {
		return
	}
	return c.EndTime
}

func (c *CreatePurchaseResponsePurchase) SetID(id string) {
	c.ID = id
}

func (c *CreatePurchaseResponsePurchase) SetPackageID(packageID string) {
	c.PackageID = packageID
}

func (c *CreatePurchaseResponsePurchase) SetStartDate(startDate *time.Time) {
	c.StartDate = startDate
}

func (c *CreatePurchaseResponsePurchase) SetEndDate(endDate *time.Time) {
	c.EndDate = endDate
}

func (c *CreatePurchaseResponsePurchase) SetCreatedDate(createdDate time.Time) {
	c.CreatedDate = createdDate
}

func (c *CreatePurchaseResponsePurchase) SetStartTime(startTime *float64) {
	c.StartTime = startTime
}

func (c *CreatePurchaseResponsePurchase) SetEndTime(endTime *float64) {
	c.EndTime = endTime
}

func (c *CreatePurchaseResponsePurchase) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseResponsePurchase to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseResponsePurchase) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseResponsePurchase) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseResponsePurchase
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
