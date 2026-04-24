package purchases

// ListPurchasesRequestParams holds the optional parameters for the API request.
type ListPurchasesRequestParams struct {
	PurchaseID  *string  `explode:"true" serializationStyle:"form" queryParam:"purchaseId"`
	Iccid       *string  `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid"`
	AfterDate   *string  `explode:"true" serializationStyle:"form" queryParam:"afterDate"`
	BeforeDate  *string  `explode:"true" serializationStyle:"form" queryParam:"beforeDate"`
	Email       *string  `explode:"true" serializationStyle:"form" queryParam:"email"`
	ReferenceID *string  `explode:"true" serializationStyle:"form" queryParam:"referenceId"`
	AfterCursor *string  `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *float64 `explode:"true" serializationStyle:"form" queryParam:"limit"`
	After       *float64 `explode:"true" serializationStyle:"form" queryParam:"after"`
	Before      *float64 `explode:"true" serializationStyle:"form" queryParam:"before"`
}
