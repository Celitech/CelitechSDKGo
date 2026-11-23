package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type SuspendEsimResponse struct {
	// Confirmation message
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SuspendEsimResponse) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SuspendEsimResponse) SetMessage(message string) {
	s.Message = message
}

func (s *SuspendEsimResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SuspendEsimResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SuspendEsimResponse to string"
	}
	return string(jsonData)
}

func (s *SuspendEsimResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SuspendEsimResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SuspendEsimResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SuspendEsimResponse) MarshalJSON() ([]byte, error) {
	type embed SuspendEsimResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
