package v2

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type CreatePurchaseV2Request struct {
	Destination   *param.Nullable[string]  `json:"destination,omitempty" xml:"destination,omitempty"`
	DataLimitInGb *param.Nullable[float64] `json:"dataLimitInGB,omitempty" xml:"dataLimitInGB,omitempty"`
	Quantity      *param.Nullable[float64] `json:"quantity,omitempty" xml:"quantity,omitempty"`
	StartDate     *param.Nullable[string]  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	EndDate       *param.Nullable[string]  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Duration      *param.Nullable[float64] `json:"duration,omitempty" xml:"duration,omitempty"`
	Email         *param.Nullable[string]  `json:"email,omitempty" xml:"email,omitempty"`
	ReferenceID   *param.Nullable[string]  `json:"referenceId,omitempty" xml:"referenceId,omitempty"`
	NetworkBrand  *param.Nullable[string]  `json:"networkBrand,omitempty" xml:"networkBrand,omitempty"`
	EmailBrand    *param.Nullable[string]  `json:"emailBrand,omitempty" xml:"emailBrand,omitempty"`
	Language      *param.Nullable[string]  `json:"language,omitempty" xml:"language,omitempty"`
}

func (c CreatePurchaseV2Request) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2Request to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseV2Request) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
