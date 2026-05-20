package esim

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// GetESimRequestParams holds the optional parameters for the API request.
type GetESimRequestParams struct {
	Iccid  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"iccid"`
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
