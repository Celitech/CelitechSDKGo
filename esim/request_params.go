package esim

// GetEsimRequestParams holds the optional parameters for the API request.
type GetEsimRequestParams struct {
	Iccid *string `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid" required:"true"`
}

// SuspendEsimRequestParams holds the optional parameters for the API request.
type SuspendEsimRequestParams struct {
	IdempotencyKey *string `explode:"false" serializationStyle:"simple" maxLength:"255" minLength:"1" headerParam:"Idempotency-Key"`
}
