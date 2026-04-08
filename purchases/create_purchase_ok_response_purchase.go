package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
	"example.com/celitech/param"
)

type CreatePurchaseOkResponsePurchase struct {
	// ID of the purchase
	ID string `json:"id" required:"true"`
	// ID of the package
	PackageID string `json:"packageId" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *param.Nullable[string] `json:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *param.Nullable[string] `json:"endDate" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate string `json:"createdDate" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *param.Nullable[float64] `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *param.Nullable[float64] `json:"endTime,omitempty"`
}

func (c CreatePurchaseOkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseOkResponsePurchase to string"
	}
	return string(jsonData)
}

func (c *CreatePurchaseOkResponsePurchase) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
