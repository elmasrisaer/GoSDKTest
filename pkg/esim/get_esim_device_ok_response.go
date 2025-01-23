package esim

import (
	"encoding/json"
)

type GetEsimDeviceOkResponse struct {
	Device  *Device `json:"device,omitempty"`
	touched map[string]bool
}

func (g *GetEsimDeviceOkResponse) GetDevice() *Device {
	if g == nil {
		return nil
	}
	return g.Device
}

func (g *GetEsimDeviceOkResponse) SetDevice(device Device) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Device"] = true
	g.Device = &device
}

func (g *GetEsimDeviceOkResponse) SetDeviceNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Device"] = true
	g.Device = nil
}

func (g GetEsimDeviceOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Device"] && g.Device == nil {
		data["device"] = nil
	} else if g.Device != nil {
		data["device"] = g.Device
	}

	return json.Marshal(data)
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
	Eid     *string `json:"eid,omitempty"`
	touched map[string]bool
}

func (d *Device) GetOem() *string {
	if d == nil {
		return nil
	}
	return d.Oem
}

func (d *Device) SetOem(oem string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Oem"] = true
	d.Oem = &oem
}

func (d *Device) SetOemNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Oem"] = true
	d.Oem = nil
}

func (d *Device) GetHardwareName() *string {
	if d == nil {
		return nil
	}
	return d.HardwareName
}

func (d *Device) SetHardwareName(hardwareName string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["HardwareName"] = true
	d.HardwareName = &hardwareName
}

func (d *Device) SetHardwareNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["HardwareName"] = true
	d.HardwareName = nil
}

func (d *Device) GetHardwareModel() *string {
	if d == nil {
		return nil
	}
	return d.HardwareModel
}

func (d *Device) SetHardwareModel(hardwareModel string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["HardwareModel"] = true
	d.HardwareModel = &hardwareModel
}

func (d *Device) SetHardwareModelNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["HardwareModel"] = true
	d.HardwareModel = nil
}

func (d *Device) GetEid() *string {
	if d == nil {
		return nil
	}
	return d.Eid
}

func (d *Device) SetEid(eid string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Eid"] = true
	d.Eid = &eid
}

func (d *Device) SetEidNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Eid"] = true
	d.Eid = nil
}

func (d Device) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Oem"] && d.Oem == nil {
		data["oem"] = nil
	} else if d.Oem != nil {
		data["oem"] = d.Oem
	}

	if d.touched["HardwareName"] && d.HardwareName == nil {
		data["hardwareName"] = nil
	} else if d.HardwareName != nil {
		data["hardwareName"] = d.HardwareName
	}

	if d.touched["HardwareModel"] && d.HardwareModel == nil {
		data["hardwareModel"] = nil
	} else if d.HardwareModel != nil {
		data["hardwareModel"] = d.HardwareModel
	}

	if d.touched["Eid"] && d.Eid == nil {
		data["eid"] = nil
	} else if d.Eid != nil {
		data["eid"] = d.Eid
	}

	return json.Marshal(data)
}

func (d Device) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Device to string"
	}
	return string(jsonData)
}
