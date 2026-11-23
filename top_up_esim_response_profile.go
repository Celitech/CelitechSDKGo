package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type TopUpEsimResponseProfile struct {
	// ID of the eSIM
	Iccid           string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
	extraProperties map[string]interface{}
}

func (t *TopUpEsimResponseProfile) GetIccid() (value string) {
	if t == nil {
		return
	}
	return t.Iccid
}

func (t *TopUpEsimResponseProfile) SetIccid(iccid string) {
	t.Iccid = iccid
}

func (t *TopUpEsimResponseProfile) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopUpEsimResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimResponseProfile to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimResponseProfile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimResponseProfile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimResponseProfile(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopUpEsimResponseProfile) MarshalJSON() ([]byte, error) {
	type embed TopUpEsimResponseProfile
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
