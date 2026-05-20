package destinations

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// ListDestinationsRequestParams holds the optional parameters for the API request.
type ListDestinationsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
