package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type GetPurchaseConsumptionOkResponse struct {
	// Remaining balance of the package in bytes
	DataUsageRemainingInBytes float64 `json:"dataUsageRemainingInBytes" required:"true"`
	// Remaining balance of the package in GB
	DataUsageRemainingInGb float64 `json:"dataUsageRemainingInGB" required:"true"`
	// Status of the connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'
	Status string `json:"status" required:"true"`
}

func (g GetPurchaseConsumptionOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPurchaseConsumptionOkResponse to string"
	}
	return string(jsonData)
}

func (g *GetPurchaseConsumptionOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetPurchaseConsumptionOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetPurchaseConsumptionOkResponse(tmp)
	return nil
}
