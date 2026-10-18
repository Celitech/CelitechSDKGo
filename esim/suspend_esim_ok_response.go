package esim

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v2/internal/unmarshal"
)

type SuspendEsimOkResponse struct {
	// Confirmation message
	Message string `json:"message" xml:"message" required:"true"`
}

func (s SuspendEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SuspendEsimOkResponse to string"
	}
	return string(jsonData)
}

func (s *SuspendEsimOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SuspendEsimOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SuspendEsimOkResponse(tmp)
	return nil
}
