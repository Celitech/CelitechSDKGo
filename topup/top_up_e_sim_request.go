package topup

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type TopUpESimRequest struct {
	Iccid         *param.Nullable[string]  `json:"iccid,omitempty" xml:"iccid,omitempty"`
	DataLimitInGb *param.Nullable[float64] `json:"dataLimitInGB,omitempty" xml:"dataLimitInGB,omitempty"`
	StartDate     *param.Nullable[string]  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	EndDate       *param.Nullable[string]  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Duration      *param.Nullable[float64] `json:"duration,omitempty" xml:"duration,omitempty"`
	Email         *param.Nullable[string]  `json:"email,omitempty" xml:"email,omitempty"`
	ReferenceID   *param.Nullable[string]  `json:"referenceId,omitempty" xml:"referenceId,omitempty"`
	EmailBrand    *param.Nullable[string]  `json:"emailBrand,omitempty" xml:"emailBrand,omitempty"`
	StartTime     *param.Nullable[float64] `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime       *param.Nullable[float64] `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (t TopUpESimRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpESimRequest to string"
	}
	return string(jsonData)
}

func (t *TopUpESimRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, t)
}
