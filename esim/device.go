package esim

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
)

type Device struct {
	// Name of the OEM
	Oem string `json:"oem" required:"true"`
	// Name of the Device
	HardwareName string `json:"hardwareName" required:"true"`
	// Model of the Device
	HardwareModel string `json:"hardwareModel" required:"true"`
	// Serial Number of the eSIM
	Eid string `json:"eid" required:"true"`
}

func (d Device) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Device to string"
	}
	return string(jsonData)
}

func (d *Device) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias Device
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = Device(tmp)
	return nil
}
