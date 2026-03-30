package packages

// ListPackagesRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListPackagesRequestParams struct {
	Destination *string  `explode:"true" serializationStyle:"form" queryParam:"destination"`
	StartDate   *string  `explode:"true" serializationStyle:"form" queryParam:"startDate"`
	EndDate     *string  `explode:"true" serializationStyle:"form" queryParam:"endDate"`
	AfterCursor *string  `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *float64 `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartTime   *int64   `explode:"true" serializationStyle:"form" queryParam:"startTime"`
	EndTime     *int64   `explode:"true" serializationStyle:"form" queryParam:"endTime"`
}

// SetDestination sets the Destination parameter.
func (params *ListPackagesRequestParams) SetDestination(destination string) {
	params.Destination = &destination
}

// SetStartDate sets the StartDate parameter.
func (params *ListPackagesRequestParams) SetStartDate(startDate string) {
	params.StartDate = &startDate
}

// SetEndDate sets the EndDate parameter.
func (params *ListPackagesRequestParams) SetEndDate(endDate string) {
	params.EndDate = &endDate
}

// SetAfterCursor sets the AfterCursor parameter.
func (params *ListPackagesRequestParams) SetAfterCursor(afterCursor string) {
	params.AfterCursor = &afterCursor
}

// SetLimit sets the Limit parameter.
func (params *ListPackagesRequestParams) SetLimit(limit float64) {
	params.Limit = &limit
}

// SetStartTime sets the StartTime parameter.
func (params *ListPackagesRequestParams) SetStartTime(startTime int64) {
	params.StartTime = &startTime
}

// SetEndTime sets the EndTime parameter.
func (params *ListPackagesRequestParams) SetEndTime(endTime int64) {
	params.EndTime = &endTime
}
