package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
	"time"
)

type CreatePurchaseV2ResponseItemPurchase struct {
	// ID of the purchase
	ID string `json:"id" xml:"id" required:"true"`
	// ID of the package
	PackageID string `json:"packageId" xml:"packageId" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate     time.Time `json:"createdDate" xml:"createdDate" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePurchaseV2ResponseItemPurchase) GetID() (value string) {
	if c == nil {
		return
	}
	return c.ID
}

func (c *CreatePurchaseV2ResponseItemPurchase) GetPackageID() (value string) {
	if c == nil {
		return
	}
	return c.PackageID
}

func (c *CreatePurchaseV2ResponseItemPurchase) GetCreatedDate() (value time.Time) {
	if c == nil {
		return
	}
	return c.CreatedDate
}

func (c *CreatePurchaseV2ResponseItemPurchase) SetID(id string) {
	c.ID = id
}

func (c *CreatePurchaseV2ResponseItemPurchase) SetPackageID(packageID string) {
	c.PackageID = packageID
}

func (c *CreatePurchaseV2ResponseItemPurchase) SetCreatedDate(createdDate time.Time) {
	c.CreatedDate = createdDate
}

func (c *CreatePurchaseV2ResponseItemPurchase) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePurchaseV2ResponseItemPurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2ResponseItemPurchase to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2ResponseItemPurchase) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2ResponseItemPurchase
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2ResponseItemPurchase(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePurchaseV2ResponseItemPurchase) MarshalJSON() ([]byte, error) {
	type embed CreatePurchaseV2ResponseItemPurchase
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
