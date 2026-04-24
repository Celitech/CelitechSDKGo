package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type TopUpEsimOkResponseProfile struct {
	// ID of the eSIM
	Iccid string `json:"iccid" xml:"iccid" required:"true" maxLength:"22" minLength:"18"`
}

func (t TopUpEsimOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponseProfile to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimOkResponseProfile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimOkResponseProfile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimOkResponseProfile(tmp)
	return nil
}
