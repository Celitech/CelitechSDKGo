package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimDeviceResponse struct {
	Device          *GetEsimDeviceResponseDevice `json:"device" xml:"device" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimDeviceResponse) GetDevice() (value *GetEsimDeviceResponseDevice) {
	if g == nil {
		return
	}
	return g.Device
}

func (g *GetEsimDeviceResponse) SetDevice(device *GetEsimDeviceResponseDevice) {
	g.Device = device
}

func (g *GetEsimDeviceResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimDeviceResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimDeviceResponse to string"
	}
	return string(jsonData)
}

func (g *GetEsimDeviceResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimDeviceResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimDeviceResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimDeviceResponse) MarshalJSON() ([]byte, error) {
	type embed GetEsimDeviceResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
