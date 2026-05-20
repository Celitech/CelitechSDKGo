package token

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// GenerateTokenRequestParams holds the optional parameters for the API request.
type GenerateTokenRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
