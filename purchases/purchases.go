package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type Purchases struct {
	// ID of the purchase
	ID string `json:"id" xml:"id" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *param.Nullable[string] `json:"startDate" xml:"startDate" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *param.Nullable[string] `json:"endDate" xml:"endDate" required:"true"`
	// Duration of the package in days. Possible values are 1, 2, 7, 14, 30, or 90.
	Duration *param.Nullable[float64] `json:"duration,omitempty" xml:"duration,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate string `json:"createdDate" xml:"createdDate" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *param.Nullable[float64] `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *param.Nullable[float64] `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Epoch value representing the date of creation of the purchase
	CreatedAt *float64      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Package   Package_      `json:"package" xml:"package" required:"true"`
	Esim      PurchasesEsim `json:"esim" xml:"esim" required:"true"`
	// The `source` indicates whether the purchase was made from the API, dashboard, landing-page, promo-page or iframe. For purchases made before September 8, 2023, the value will be displayed as 'Not available'.
	Source string `json:"source" xml:"source" required:"true"`
	// The `purchaseType` indicates whether this is the initial purchase that creates the eSIM (First Purchase) or a subsequent top-up on an existing eSIM (Top-up Purchase).
	PurchaseType string `json:"purchaseType" xml:"purchaseType" required:"true"`
	// The `referenceId` that was provided by the partner during the purchase or top-up flow. This identifier can be used for analytics and debugging purposes.
	ReferenceID *param.Nullable[string] `json:"referenceId,omitempty" xml:"referenceId,omitempty"`
}

func (p Purchases) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Purchases to string"
	}
	return string(jsonData)
}

func (p *Purchases) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}
