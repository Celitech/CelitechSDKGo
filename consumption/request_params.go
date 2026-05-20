package consumption

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// GetPurchaseConsumptionRequestParams holds the optional parameters for the API request.
type GetPurchaseConsumptionRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
