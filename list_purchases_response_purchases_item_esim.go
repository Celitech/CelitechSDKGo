package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListPurchasesResponsePurchasesItemEsim struct {
	// ID of the eSIM
	Iccid           string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	extraProperties map[string]interface{}
}

func (l *ListPurchasesResponsePurchasesItemEsim) GetIccid() (value string) {
	if l == nil {
		return
	}
	return l.Iccid
}

func (l *ListPurchasesResponsePurchasesItemEsim) SetIccid(iccid string) {
	l.Iccid = iccid
}

func (l *ListPurchasesResponsePurchasesItemEsim) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPurchasesResponsePurchasesItemEsim) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesResponsePurchasesItemEsim to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesResponsePurchasesItemEsim) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias ListPurchasesResponsePurchasesItemEsim
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = ListPurchasesResponsePurchasesItemEsim(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListPurchasesResponsePurchasesItemEsim) MarshalJSON() ([]byte, error) {
	type embed ListPurchasesResponsePurchasesItemEsim
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
