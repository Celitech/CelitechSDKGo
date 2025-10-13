package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type CreatePurchaseOkResponse struct {
	Purchase *CreatePurchaseOkResponsePurchase `json:"purchase,omitempty" required:"true"`
	Profile  *CreatePurchaseOkResponseProfile  `json:"profile,omitempty" required:"true"`
}

func (c *CreatePurchaseOkResponse) GetPurchase() *CreatePurchaseOkResponsePurchase {
	if c == nil {
		return nil
	}
	return c.Purchase
}

func (c *CreatePurchaseOkResponse) SetPurchase(purchase CreatePurchaseOkResponsePurchase) {
	c.Purchase = &purchase
}

func (c *CreatePurchaseOkResponse) GetProfile() *CreatePurchaseOkResponseProfile {
	if c == nil {
		return nil
	}
	return c.Profile
}

func (c *CreatePurchaseOkResponse) SetProfile(profile CreatePurchaseOkResponseProfile) {
	c.Profile = &profile
}

func (c CreatePurchaseOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponse to string"
	}
	return string(jsonData)
}

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

type CreatePurchaseOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" required:"true" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty" required:"true"`
}

func (c *CreatePurchaseOkResponseProfile) GetIccid() *string {
	if c == nil {
		return nil
	}
	return c.Iccid
}

func (c *CreatePurchaseOkResponseProfile) SetIccid(iccid string) {
	c.Iccid = &iccid
}

func (c *CreatePurchaseOkResponseProfile) GetActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetActivationCode(activationCode string) {
	c.ActivationCode = &activationCode
}

func (c *CreatePurchaseOkResponseProfile) GetManualActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetManualActivationCode(manualActivationCode string) {
	c.ManualActivationCode = &manualActivationCode
}

func (c CreatePurchaseOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponseProfile to string"
	}
	return string(jsonData)
}
