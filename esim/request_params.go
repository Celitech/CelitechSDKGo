package esim

// GetEsimRequestParams holds the optional parameters for the API request.
type GetEsimRequestParams struct {
	Iccid *string `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid" required:"true"`
}
