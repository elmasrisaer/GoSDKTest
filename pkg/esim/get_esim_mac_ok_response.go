package esim

import (
	"encoding/json"
)

type GetEsimMacOkResponse struct {
	Esim    *GetEsimMacOkResponseEsim `json:"esim,omitempty"`
	touched map[string]bool
}

func (g *GetEsimMacOkResponse) GetEsim() *GetEsimMacOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimMacOkResponse) SetEsim(esim GetEsimMacOkResponseEsim) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = &esim
}

func (g *GetEsimMacOkResponse) SetEsimNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = nil
}

func (g GetEsimMacOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Esim"] && g.Esim == nil {
		data["esim"] = nil
	} else if g.Esim != nil {
		data["esim"] = g.Esim
	}

	return json.Marshal(data)
}

func (g GetEsimMacOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimMacOkResponse to string"
	}
	return string(jsonData)
}

type GetEsimMacOkResponseEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// SM-DP+ Address
	SmdpAddress *string `json:"smdpAddress,omitempty"`
	// The manual activation code
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
	touched              map[string]bool
}

func (g *GetEsimMacOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimMacOkResponseEsim) SetIccid(iccid string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = &iccid
}

func (g *GetEsimMacOkResponseEsim) SetIccidNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = nil
}

func (g *GetEsimMacOkResponseEsim) GetSmdpAddress() *string {
	if g == nil {
		return nil
	}
	return g.SmdpAddress
}

func (g *GetEsimMacOkResponseEsim) SetSmdpAddress(smdpAddress string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["SmdpAddress"] = true
	g.SmdpAddress = &smdpAddress
}

func (g *GetEsimMacOkResponseEsim) SetSmdpAddressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["SmdpAddress"] = true
	g.SmdpAddress = nil
}

func (g *GetEsimMacOkResponseEsim) GetManualActivationCode() *string {
	if g == nil {
		return nil
	}
	return g.ManualActivationCode
}

func (g *GetEsimMacOkResponseEsim) SetManualActivationCode(manualActivationCode string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ManualActivationCode"] = true
	g.ManualActivationCode = &manualActivationCode
}

func (g *GetEsimMacOkResponseEsim) SetManualActivationCodeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ManualActivationCode"] = true
	g.ManualActivationCode = nil
}

func (g GetEsimMacOkResponseEsim) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Iccid"] && g.Iccid == nil {
		data["iccid"] = nil
	} else if g.Iccid != nil {
		data["iccid"] = g.Iccid
	}

	if g.touched["SmdpAddress"] && g.SmdpAddress == nil {
		data["smdpAddress"] = nil
	} else if g.SmdpAddress != nil {
		data["smdpAddress"] = g.SmdpAddress
	}

	if g.touched["ManualActivationCode"] && g.ManualActivationCode == nil {
		data["manualActivationCode"] = nil
	} else if g.ManualActivationCode != nil {
		data["manualActivationCode"] = g.ManualActivationCode
	}

	return json.Marshal(data)
}

func (g GetEsimMacOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimMacOkResponseEsim to string"
	}
	return string(jsonData)
}
