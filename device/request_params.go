package device

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// GetESimDeviceRequestParams holds the optional parameters for the API request.
type GetESimDeviceRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
