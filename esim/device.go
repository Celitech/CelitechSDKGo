package esim

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
)

type Device struct {
	// Name of the OEM
	Oem string `json:"oem" xml:"oem" required:"true"`
	// Name of the Device
	HardwareName string `json:"hardwareName" xml:"hardwareName" required:"true"`
	// Model of the Device
	HardwareModel string `json:"hardwareModel" xml:"hardwareModel" required:"true"`
	// Serial Number of the eSIM
	Eid string `json:"eid" xml:"eid" required:"true"`
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
