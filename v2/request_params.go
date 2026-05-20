package v2

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// CreatePurchaseV2RequestParams holds the optional parameters for the API request.
type CreatePurchaseV2RequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
