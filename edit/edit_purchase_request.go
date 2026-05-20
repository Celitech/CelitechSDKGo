package edit

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type EditPurchaseRequest struct {
	PurchaseID *param.Nullable[string]  `json:"purchaseId,omitempty" xml:"purchaseId,omitempty"`
	StartDate  *param.Nullable[string]  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	EndDate    *param.Nullable[string]  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	StartTime  *param.Nullable[float64] `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime    *param.Nullable[float64] `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (e EditPurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditPurchaseRequest to string"
	}
	return string(jsonData)
}

func (e *EditPurchaseRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}
