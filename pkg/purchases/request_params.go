package purchases

// ListPurchasesRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListPurchasesRequestParams struct {
	PurchaseId  *string  `explode:"true" serializationStyle:"form" queryParam:"purchaseId"`
	Iccid       *string  `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid"`
	AfterDate   *string  `explode:"true" serializationStyle:"form" queryParam:"afterDate"`
	BeforeDate  *string  `explode:"true" serializationStyle:"form" queryParam:"beforeDate"`
	Email       *string  `explode:"true" serializationStyle:"form" queryParam:"email"`
	ReferenceId *string  `explode:"true" serializationStyle:"form" queryParam:"referenceId"`
	AfterCursor *string  `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *float64 `explode:"true" serializationStyle:"form" queryParam:"limit"`
	After       *float64 `explode:"true" serializationStyle:"form" queryParam:"after"`
	Before      *float64 `explode:"true" serializationStyle:"form" queryParam:"before"`
}

// SetPurchaseId sets the PurchaseId parameter.
func (params *ListPurchasesRequestParams) SetPurchaseId(purchaseId string) {
	params.PurchaseId = &purchaseId
}

// SetIccid sets the Iccid parameter.
func (params *ListPurchasesRequestParams) SetIccid(iccid string) {
	params.Iccid = &iccid
}

// SetAfterDate sets the AfterDate parameter.
func (params *ListPurchasesRequestParams) SetAfterDate(afterDate string) {
	params.AfterDate = &afterDate
}

// SetBeforeDate sets the BeforeDate parameter.
func (params *ListPurchasesRequestParams) SetBeforeDate(beforeDate string) {
	params.BeforeDate = &beforeDate
}

// SetEmail sets the Email parameter.
func (params *ListPurchasesRequestParams) SetEmail(email string) {
	params.Email = &email
}

// SetReferenceId sets the ReferenceId parameter.
func (params *ListPurchasesRequestParams) SetReferenceId(referenceId string) {
	params.ReferenceId = &referenceId
}

// SetAfterCursor sets the AfterCursor parameter.
func (params *ListPurchasesRequestParams) SetAfterCursor(afterCursor string) {
	params.AfterCursor = &afterCursor
}

// SetLimit sets the Limit parameter.
func (params *ListPurchasesRequestParams) SetLimit(limit float64) {
	params.Limit = &limit
}

// SetAfter sets the After parameter.
func (params *ListPurchasesRequestParams) SetAfter(after float64) {
	params.After = &after
}

// SetBefore sets the Before parameter.
func (params *ListPurchasesRequestParams) SetBefore(before float64) {
	params.Before = &before
}
