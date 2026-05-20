package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type CreatePurchaseRequest struct {
	Destination   *param.Nullable[string]  `json:"destination,omitempty" xml:"destination,omitempty"`
	DataLimitInGb *param.Nullable[float64] `json:"dataLimitInGB,omitempty" xml:"dataLimitInGB,omitempty"`
	StartDate     *param.Nullable[string]  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	EndDate       *param.Nullable[string]  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Email         *param.Nullable[string]  `json:"email,omitempty" xml:"email,omitempty"`
	ReferenceID   *param.Nullable[string]  `json:"referenceId,omitempty" xml:"referenceId,omitempty"`
	NetworkBrand  *param.Nullable[string]  `json:"networkBrand,omitempty" xml:"networkBrand,omitempty"`
	EmailBrand    *param.Nullable[string]  `json:"emailBrand,omitempty" xml:"emailBrand,omitempty"`
	Language      *param.Nullable[string]  `json:"language,omitempty" xml:"language,omitempty"`
	StartTime     *param.Nullable[float64] `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime       *param.Nullable[float64] `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (c CreatePurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseRequest to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
