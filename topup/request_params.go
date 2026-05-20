package topup

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// TopUpESimRequestParams holds the optional parameters for the API request.
type TopUpESimRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
