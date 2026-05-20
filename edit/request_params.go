package edit

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// EditPurchaseRequestParams holds the optional parameters for the API request.
type EditPurchaseRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
