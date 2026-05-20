package history

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// GetESimHistoryRequestParams holds the optional parameters for the API request.
type GetESimHistoryRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
