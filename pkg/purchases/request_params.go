package purchases

type ListPurchasesRequestParams struct {
	Iccid       *string  `maxLength:"22" minLength:"18" queryParam:"iccid"`
	AfterDate   *string  `queryParam:"afterDate"`
	BeforeDate  *string  `queryParam:"beforeDate"`
	ReferenceId *string  `queryParam:"referenceId"`
	AfterCursor *string  `queryParam:"afterCursor"`
	Limit       *float64 `queryParam:"limit"`
	After       *float64 `queryParam:"after"`
	Before      *float64 `queryParam:"before"`
}

func (params *ListPurchasesRequestParams) SetIccid(iccid string) {
	params.Iccid = &iccid
}
func (params *ListPurchasesRequestParams) SetAfterDate(afterDate string) {
	params.AfterDate = &afterDate
}
func (params *ListPurchasesRequestParams) SetBeforeDate(beforeDate string) {
	params.BeforeDate = &beforeDate
}
func (params *ListPurchasesRequestParams) SetReferenceId(referenceId string) {
	params.ReferenceId = &referenceId
}
func (params *ListPurchasesRequestParams) SetAfterCursor(afterCursor string) {
	params.AfterCursor = &afterCursor
}
func (params *ListPurchasesRequestParams) SetLimit(limit float64) {
	params.Limit = &limit
}
func (params *ListPurchasesRequestParams) SetAfter(after float64) {
	params.After = &after
}
func (params *ListPurchasesRequestParams) SetBefore(before float64) {
	params.Before = &before
}
