package packages

type ListPackagesRequestParams struct {
	Destination *string  `queryParam:"destination"`
	StartDate   *string  `queryParam:"startDate"`
	EndDate     *string  `queryParam:"endDate"`
	AfterCursor *string  `queryParam:"afterCursor"`
	Limit       *float64 `queryParam:"limit"`
	StartTime   *int64   `queryParam:"startTime"`
	EndTime     *int64   `queryParam:"endTime"`
	Duration    *float64 `queryParam:"duration"`
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
