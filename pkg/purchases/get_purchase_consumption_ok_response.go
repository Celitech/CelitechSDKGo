package purchases

import "encoding/json"

type GetPurchaseConsumptionOkResponse struct {
	// Remaining balance of the package in bytes
	DataUsageRemainingInBytes *float64 `json:"dataUsageRemainingInBytes,omitempty"`
	// Status of the connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'
	Status *string `json:"status,omitempty"`
}

func (g *GetPurchaseConsumptionOkResponse) GetDataUsageRemainingInBytes() *float64 {
	if g == nil {
		return nil
	}
	return g.DataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionOkResponse) SetDataUsageRemainingInBytes(dataUsageRemainingInBytes float64) {
	g.DataUsageRemainingInBytes = &dataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionOkResponse) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetPurchaseConsumptionOkResponse) SetStatus(status string) {
	g.Status = &status
}

func (g GetPurchaseConsumptionOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPurchaseConsumptionOkResponse to string"
	}
	return string(jsonData)
}
