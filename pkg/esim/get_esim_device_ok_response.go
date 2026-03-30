package esim

import "encoding/json"

type GetEsimDeviceOkResponse struct {
	Device *Device `json:"device,omitempty" required:"true"`
}

func (g *GetEsimDeviceOkResponse) GetDevice() *Device {
	if g == nil {
		return nil
	}
	return g.Device
}

func (g *GetEsimDeviceOkResponse) SetDevice(device Device) {
	g.Device = &device
}

func (g GetEsimDeviceOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimDeviceOkResponse to string"
	}
	return string(jsonData)
}
