package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type CreatePurchaseOkResponse struct {
	Purchase CreatePurchaseOkResponsePurchase `json:"purchase" required:"true"`
	Profile  CreatePurchaseOkResponseProfile  `json:"profile" required:"true"`
}

func (c CreatePurchaseOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponse to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePurchaseOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePurchaseOkResponse(tmp)
	return nil
}
