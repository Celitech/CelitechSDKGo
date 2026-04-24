package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type TopUpEsimOkResponse struct {
	Purchase TopUpEsimOkResponsePurchase `json:"purchase" xml:"purchase" required:"true"`
	Profile  TopUpEsimOkResponseProfile  `json:"profile" xml:"profile" required:"true"`
}

func (t TopUpEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponse to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopUpEsimOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopUpEsimOkResponse(tmp)
	return nil
}
