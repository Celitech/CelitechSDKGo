package esim

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type GetEsimDeviceOkResponse struct {
	Device Device `json:"device" required:"true"`
}

func (g GetEsimDeviceOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimDeviceOkResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimDeviceOkResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimDeviceOkResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimDeviceOkResponse(tmp)
	return nil
}
