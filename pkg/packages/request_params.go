package packages

type ListPackagesRequestParams struct {
	Destination *string  `explode:"true" serializationStyle:"form" queryParam:"destination"`
	StartDate   *string  `explode:"true" serializationStyle:"form" queryParam:"startDate"`
	EndDate     *string  `explode:"true" serializationStyle:"form" queryParam:"endDate"`
	AfterCursor *string  `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *float64 `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartTime   *int64   `explode:"true" serializationStyle:"form" queryParam:"startTime"`
	EndTime     *int64   `explode:"true" serializationStyle:"form" queryParam:"endTime"`
	Duration    *float64 `explode:"true" serializationStyle:"form" queryParam:"duration"`
}

func (params *ListPackagesRequestParams) SetDestination(destination string) {
	params.Destination = &destination
}
func (params *ListPackagesRequestParams) SetStartDate(startDate string) {
	params.StartDate = &startDate
}
func (params *ListPackagesRequestParams) SetEndDate(endDate string) {
	params.EndDate = &endDate
}
func (params *ListPackagesRequestParams) SetAfterCursor(afterCursor string) {
	params.AfterCursor = &afterCursor
}
func (params *ListPackagesRequestParams) SetLimit(limit float64) {
	params.Limit = &limit
}
func (params *ListPackagesRequestParams) SetStartTime(startTime int64) {
	params.StartTime = &startTime
}
func (params *ListPackagesRequestParams) SetEndTime(endTime int64) {
	params.EndTime = &endTime
}
func (params *ListPackagesRequestParams) SetDuration(duration float64) {
	params.Duration = &duration
}
