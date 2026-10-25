package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetPurchaseConsumptionResponse struct {
	// Remaining balance of the package in bytes. Returns `-1` for unlimited packages.
	DataUsageRemainingInBytes float64 `json:"dataUsageRemainingInBytes" xml:"dataUsageRemainingInBytes" required:"true"`
	// Remaining balance of the package in GB. Returns `-1` for unlimited packages.
	DataUsageRemainingInGb float64 `json:"dataUsageRemainingInGB" xml:"dataUsageRemainingInGB" required:"true"`
	// Status of the connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetPurchaseConsumptionResponse) GetDataUsageRemainingInBytes() (value float64) {
	if g == nil {
		return
	}
	return g.DataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionResponse) GetDataUsageRemainingInGb() (value float64) {
	if g == nil {
		return
	}
	return g.DataUsageRemainingInGb
}

func (g *GetPurchaseConsumptionResponse) GetStatus() (value string) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetPurchaseConsumptionResponse) SetDataUsageRemainingInBytes(dataUsageRemainingInBytes float64) {
	g.DataUsageRemainingInBytes = dataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionResponse) SetDataUsageRemainingInGb(dataUsageRemainingInGb float64) {
	g.DataUsageRemainingInGb = dataUsageRemainingInGb
}

func (g *GetPurchaseConsumptionResponse) SetStatus(status string) {
	g.Status = status
}

func (g *GetPurchaseConsumptionResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPurchaseConsumptionResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPurchaseConsumptionResponse to string"
	}
	return string(jsonData)
}

func (g *GetPurchaseConsumptionResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetPurchaseConsumptionResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetPurchaseConsumptionResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetPurchaseConsumptionResponse) MarshalJSON() ([]byte, error) {
	type embed GetPurchaseConsumptionResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
