package purchases

import "encoding/json"

type TopUpEsimRequest struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	// Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps).
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
}

func (t *TopUpEsimRequest) GetIccid() *string {
	if t == nil {
		return nil
	}
	return t.Iccid
}

func (t *TopUpEsimRequest) SetIccid(iccid string) {
	t.Iccid = &iccid
}

func (t *TopUpEsimRequest) GetDataLimitInGb() *float64 {
	if t == nil {
		return nil
	}
	return t.DataLimitInGb
}

func (t *TopUpEsimRequest) SetDataLimitInGb(dataLimitInGb float64) {
	t.DataLimitInGb = &dataLimitInGb
}

func (t *TopUpEsimRequest) GetStartDate() *string {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TopUpEsimRequest) SetStartDate(startDate string) {
	t.StartDate = &startDate
}

func (t *TopUpEsimRequest) GetEndDate() *string {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TopUpEsimRequest) SetEndDate(endDate string) {
	t.EndDate = &endDate
}

func (t *TopUpEsimRequest) GetEmail() *string {
	if t == nil {
		return nil
	}
	return t.Email
}

func (t *TopUpEsimRequest) SetEmail(email string) {
	t.Email = &email
}

func (t *TopUpEsimRequest) GetReferenceId() *string {
	if t == nil {
		return nil
	}
	return t.ReferenceId
}

func (t *TopUpEsimRequest) SetReferenceId(referenceId string) {
	t.ReferenceId = &referenceId
}

func (t *TopUpEsimRequest) GetEmailBrand() *string {
	if t == nil {
		return nil
	}
	return t.EmailBrand
}

func (t *TopUpEsimRequest) SetEmailBrand(emailBrand string) {
	t.EmailBrand = &emailBrand
}

func (t *TopUpEsimRequest) GetStartTime() *float64 {
	if t == nil {
		return nil
	}
	return t.StartTime
}

func (t *TopUpEsimRequest) SetStartTime(startTime float64) {
	t.StartTime = &startTime
}

func (t *TopUpEsimRequest) GetEndTime() *float64 {
	if t == nil {
		return nil
	}
	return t.EndTime
}

func (t *TopUpEsimRequest) SetEndTime(endTime float64) {
	t.EndTime = &endTime
}

func (t TopUpEsimRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimRequest to string"
	}
	return string(jsonData)
}
