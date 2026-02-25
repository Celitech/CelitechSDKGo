package purchases

import "encoding/json"

type CreatePurchaseV2OkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty" required:"true"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty" required:"true"`
}

func (c *CreatePurchaseV2OkResponsePurchase) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreatePurchaseV2OkResponsePurchase) SetId(id string) {
	c.Id = &id
}

func (c *CreatePurchaseV2OkResponsePurchase) GetPackageId() *string {
	if c == nil {
		return nil
	}
	return c.PackageId
}

func (c *CreatePurchaseV2OkResponsePurchase) SetPackageId(packageId string) {
	c.PackageId = &packageId
}

func (c *CreatePurchaseV2OkResponsePurchase) GetCreatedDate() *string {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePurchaseV2OkResponsePurchase) SetCreatedDate(createdDate string) {
	c.CreatedDate = &createdDate
}

func (c CreatePurchaseV2OkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponsePurchase to string"
	}
	return string(jsonData)
}
