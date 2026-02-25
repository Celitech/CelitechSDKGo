package esim

// GetEsimRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type GetEsimRequestParams struct {
	Iccid *string `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid" required:"true"`
}

// SetIccid sets the Iccid parameter.
func (params *GetEsimRequestParams) SetIccid(iccid string) {
	params.Iccid = &iccid
}
