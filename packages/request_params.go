package packages

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// ListPackagesRequestParams holds the optional parameters for the API request.
type ListPackagesRequestParams struct {
	Destination *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"destination"`
	StartDate   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"startDate"`
	EndDate     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"endDate"`
	AfterCursor *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartTime   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"startTime"`
	EndTime     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"endTime"`
	Accept      *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
