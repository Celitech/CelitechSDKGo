package purchases

// CreatePurchaseV2RequestParams holds the optional parameters for the API request.
type CreatePurchaseV2RequestParams struct {
	IdempotencyKey *string `explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
}

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

// CreatePurchaseRequestParams holds the optional parameters for the API request.
type CreatePurchaseRequestParams struct {
	IdempotencyKey *string `explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
}

// TopUpEsimRequestParams holds the optional parameters for the API request.
type TopUpEsimRequestParams struct {
	IdempotencyKey *string `explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
}

// EditPurchaseRequestParams holds the optional parameters for the API request.
type EditPurchaseRequestParams struct {
	IdempotencyKey *string `explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
}
