package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type TopUpEsimRequest struct {
	// ID of the eSIM
	Iccid string `json:"iccid" required:"true" maxLength:"22" minLength:"18"`
	// Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB
	DataLimitInGb float64 `json:"dataLimitInGB" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty"`
	// Duration of the package in days. Available values are 1, 2, 7, 14, 30, or 90. Either provide startDate/endDate or duration.
	Duration *float64 `json:"duration,omitempty"`
	// Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps).
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceID *string `json:"referenceId,omitempty"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
}

func (t TopUpEsimRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimRequest to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimRequest(tmp)
	return nil
}
