package esim

import "encoding/json"

type GetEsimDeviceOkResponse struct {
	Device *Device `json:"device,omitempty"`
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

type Device struct {
	// Name of the OEM
	Oem *string `json:"oem,omitempty"`
	// Name of the Device
	HardwareName *string `json:"hardwareName,omitempty"`
	// Model of the Device
	HardwareModel *string `json:"hardwareModel,omitempty"`
	// Serial Number of the eSIM
	Eid *string `json:"eid,omitempty"`
}

func (d *Device) GetOem() *string {
	if d == nil {
		return nil
	}
	return d.Oem
}

func (d *Device) SetOem(oem string) {
	d.Oem = &oem
}

func (d *Device) GetHardwareName() *string {
	if d == nil {
		return nil
	}
	return d.HardwareName
}

func (d *Device) SetHardwareName(hardwareName string) {
	d.HardwareName = &hardwareName
}

func (d *Device) GetHardwareModel() *string {
	if d == nil {
		return nil
	}
	return d.HardwareModel
}

func (d *Device) SetHardwareModel(hardwareModel string) {
	d.HardwareModel = &hardwareModel
}

func (d *Device) GetEid() *string {
	if d == nil {
		return nil
	}
	return d.Eid
}

func (d *Device) SetEid(eid string) {
	d.Eid = &eid
}

func (d Device) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Device to string"
	}
	return string(jsonData)
}
