package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type GetEsimDeviceResponseDevice struct {
	// Name of the OEM
	Oem string `json:"oem" xml:"oem" required:"true"`
	// Name of the Device
	HardwareName string `json:"hardwareName" xml:"hardwareName" required:"true"`
	// Model of the Device
	HardwareModel string `json:"hardwareModel" xml:"hardwareModel" required:"true"`
	// Serial Number of the eSIM
	Eid             string `json:"eid" xml:"eid" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetEsimDeviceResponseDevice) GetOem() (value string) {
	if g == nil {
		return
	}
	return g.Oem
}

func (g *GetEsimDeviceResponseDevice) GetHardwareName() (value string) {
	if g == nil {
		return
	}
	return g.HardwareName
}

func (g *GetEsimDeviceResponseDevice) GetHardwareModel() (value string) {
	if g == nil {
		return
	}
	return g.HardwareModel
}

func (g *GetEsimDeviceResponseDevice) GetEid() (value string) {
	if g == nil {
		return
	}
	return g.Eid
}

func (g *GetEsimDeviceResponseDevice) SetOem(oem string) {
	g.Oem = oem
}

func (g *GetEsimDeviceResponseDevice) SetHardwareName(hardwareName string) {
	g.HardwareName = hardwareName
}

func (g *GetEsimDeviceResponseDevice) SetHardwareModel(hardwareModel string) {
	g.HardwareModel = hardwareModel
}

func (g *GetEsimDeviceResponseDevice) SetEid(eid string) {
	g.Eid = eid
}

func (g *GetEsimDeviceResponseDevice) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetEsimDeviceResponseDevice) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimDeviceResponseDevice to string"
	}
	return string(jsonData)
}

func (g *GetEsimDeviceResponseDevice) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetEsimDeviceResponseDevice
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetEsimDeviceResponseDevice(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetEsimDeviceResponseDevice) MarshalJSON() ([]byte, error) {
	type embed GetEsimDeviceResponseDevice
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
