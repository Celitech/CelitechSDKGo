package esim

type GetEsimRequestParams struct {
	Iccid *string `explode:"true" serializationStyle:"form" maxLength:"22" minLength:"18" queryParam:"iccid" required:"true"`
}

func (params *GetEsimRequestParams) SetIccid(iccid string) {
	params.Iccid = &iccid
}
