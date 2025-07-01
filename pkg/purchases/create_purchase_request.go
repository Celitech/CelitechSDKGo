package purchases

import "encoding/json"

type CreatePurchaseRequest struct {
	// ISO representation of the package's destination.
	Destination *string `json:"destination,omitempty" required:"true"`
	// Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Email address where the purchase confirmation email will be sent (including QR Code & activation steps)
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters and numbers. This feature is available to platforms with Diamond tier only.
	NetworkBrand *string `json:"networkBrand,omitempty"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
}

func (c *CreatePurchaseRequest) GetDestination() *string {
	if c == nil {
		return nil
	}
	return c.Destination
}

func (c *CreatePurchaseRequest) SetDestination(destination string) {
	c.Destination = &destination
}

func (c *CreatePurchaseRequest) GetDataLimitInGb() *float64 {
	if c == nil {
		return nil
	}
	return c.DataLimitInGb
}

func (c *CreatePurchaseRequest) SetDataLimitInGb(dataLimitInGb float64) {
	c.DataLimitInGb = &dataLimitInGb
}

func (c *CreatePurchaseRequest) GetStartDate() *string {
	if c == nil {
		return nil
	}
	return c.StartDate
}

func (c *CreatePurchaseRequest) SetStartDate(startDate string) {
	c.StartDate = &startDate
}

func (c *CreatePurchaseRequest) GetEndDate() *string {
	if c == nil {
		return nil
	}
	return c.EndDate
}

func (c *CreatePurchaseRequest) SetEndDate(endDate string) {
	c.EndDate = &endDate
}

func (c *CreatePurchaseRequest) GetEmail() *string {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreatePurchaseRequest) SetEmail(email string) {
	c.Email = &email
}

func (c *CreatePurchaseRequest) GetReferenceId() *string {
	if c == nil {
		return nil
	}
	return c.ReferenceId
}

func (c *CreatePurchaseRequest) SetReferenceId(referenceId string) {
	c.ReferenceId = &referenceId
}

func (c *CreatePurchaseRequest) GetNetworkBrand() *string {
	if c == nil {
		return nil
	}
	return c.NetworkBrand
}

func (c *CreatePurchaseRequest) SetNetworkBrand(networkBrand string) {
	c.NetworkBrand = &networkBrand
}

func (c *CreatePurchaseRequest) GetEmailBrand() *string {
	if c == nil {
		return nil
	}
	return c.EmailBrand
}

func (c *CreatePurchaseRequest) SetEmailBrand(emailBrand string) {
	c.EmailBrand = &emailBrand
}

func (c *CreatePurchaseRequest) GetStartTime() *float64 {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *CreatePurchaseRequest) SetStartTime(startTime float64) {
	c.StartTime = &startTime
}

func (c *CreatePurchaseRequest) GetEndTime() *float64 {
	if c == nil {
		return nil
	}
	return c.EndTime
}

func (c *CreatePurchaseRequest) SetEndTime(endTime float64) {
	c.EndTime = &endTime
}

func (c CreatePurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseRequest to string"
	}
	return string(jsonData)
}
