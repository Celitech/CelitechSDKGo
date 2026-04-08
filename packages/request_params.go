package packages

// ListPackagesRequestParams holds the optional parameters for the API request.
type ListPackagesRequestParams struct {
	Destination *string  `explode:"true" serializationStyle:"form" queryParam:"destination"`
	StartDate   *string  `explode:"true" serializationStyle:"form" queryParam:"startDate"`
	EndDate     *string  `explode:"true" serializationStyle:"form" queryParam:"endDate"`
	AfterCursor *string  `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *float64 `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartTime   *int64   `explode:"true" serializationStyle:"form" queryParam:"startTime"`
	EndTime     *int64   `explode:"true" serializationStyle:"form" queryParam:"endTime"`
}
