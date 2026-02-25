package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type CreatePurchaseOkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty" required:"true"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty" required:"true"`
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
}

func (c *CreatePurchaseOkResponsePurchase) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreatePurchaseOkResponsePurchase) SetId(id string) {
	c.Id = &id
}

func (c *CreatePurchaseOkResponsePurchase) GetPackageId() *string {
	if c == nil {
		return nil
	}
	return c.PackageId
}

func (c *CreatePurchaseOkResponsePurchase) SetPackageId(packageId string) {
	c.PackageId = &packageId
}

func (c *CreatePurchaseOkResponsePurchase) GetStartDate() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.StartDate
}

func (c *CreatePurchaseOkResponsePurchase) SetStartDate(startDate util.Nullable[string]) {
	c.StartDate = &startDate
}

func (c *CreatePurchaseOkResponsePurchase) SetStartDateNull() {
	c.StartDate = &util.Nullable[string]{IsNull: true}
}

func (c *CreatePurchaseOkResponsePurchase) GetEndDate() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.EndDate
}

func (c *CreatePurchaseOkResponsePurchase) SetEndDate(endDate util.Nullable[string]) {
	c.EndDate = &endDate
}

func (c *CreatePurchaseOkResponsePurchase) SetEndDateNull() {
	c.EndDate = &util.Nullable[string]{IsNull: true}
}

func (c *CreatePurchaseOkResponsePurchase) GetCreatedDate() *string {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePurchaseOkResponsePurchase) SetCreatedDate(createdDate string) {
	c.CreatedDate = &createdDate
}

func (c *CreatePurchaseOkResponsePurchase) GetStartTime() *util.Nullable[float64] {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *CreatePurchaseOkResponsePurchase) SetStartTime(startTime util.Nullable[float64]) {
	c.StartTime = &startTime
}

func (c *CreatePurchaseOkResponsePurchase) SetStartTimeNull() {
	c.StartTime = &util.Nullable[float64]{IsNull: true}
}

func (c *CreatePurchaseOkResponsePurchase) GetEndTime() *util.Nullable[float64] {
	if c == nil {
		return nil
	}
	return c.EndTime
}

func (c *CreatePurchaseOkResponsePurchase) SetEndTime(endTime util.Nullable[float64]) {
	c.EndTime = &endTime
}

func (c *CreatePurchaseOkResponsePurchase) SetEndTimeNull() {
	c.EndTime = &util.Nullable[float64]{IsNull: true}
}

func (c CreatePurchaseOkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponsePurchase to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseOkResponsePurchase) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
