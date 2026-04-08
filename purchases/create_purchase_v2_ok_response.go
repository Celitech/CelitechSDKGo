package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type CreatePurchaseV2OkResponse struct {
	Purchase CreatePurchaseV2OkResponsePurchase `json:"purchase" required:"true"`
	Profile  CreatePurchaseV2OkResponseProfile  `json:"profile" required:"true"`
}

func (c CreatePurchaseV2OkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2OkResponse to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2OkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseV2OkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseV2OkResponse(tmp)
	return nil
}
