package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type CreatePurchaseV2OkResponsePurchase struct {
	// ID of the purchase
	ID string `json:"id" xml:"id" required:"true"`
	// ID of the package
	PackageID string `json:"packageId" xml:"packageId" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate string `json:"createdDate" xml:"createdDate" required:"true"`
}

func (c CreatePurchaseV2OkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponsePurchase to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2OkResponsePurchase) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2OkResponsePurchase
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2OkResponsePurchase(tmp)
	return nil
}
