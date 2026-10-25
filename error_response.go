package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
)

type ErrorResponse struct {
	// Message of the error
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ErrorResponse) GetMessage() (value *string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *ErrorResponse) SetMessage(message *string) {
	e.Message = message
}

func (e *ErrorResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ErrorResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ErrorResponse to string"
	}
	return string(jsonData)
}

func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ErrorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ErrorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ErrorResponse) MarshalJSON() ([]byte, error) {
	type embed ErrorResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
