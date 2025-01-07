package esim

import (
	"encoding/json"
)

type GetEsimOkResponse struct {
	Esim    *GetEsimOkResponseEsim `json:"esim,omitempty"`
	touched map[string]bool
}

func (g *GetEsimOkResponse) GetEsim() *GetEsimOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimOkResponse) SetEsim(esim GetEsimOkResponseEsim) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = &esim
}

func (g *GetEsimOkResponse) SetEsimNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = nil
}
func (g GetEsimOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Esim"] && g.Esim == nil {
		data["esim"] = nil
	} else if g.Esim != nil {
		data["esim"] = g.Esim
	}

	return json.Marshal(data)
}

type GetEsimOkResponseEsim struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// SM-DP+ Address
	SmdpAddress *string `json:"smdpAddress,omitempty"`
	// The manual activation code
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
	// Status of the eSIM, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status  *string `json:"status,omitempty"`
	touched map[string]bool
}

func (g *GetEsimOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimOkResponseEsim) SetIccid(iccid string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = &iccid
}

func (g *GetEsimOkResponseEsim) SetIccidNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = nil
}

func (g *GetEsimOkResponseEsim) GetSmdpAddress() *string {
	if g == nil {
		return nil
	}
	return g.SmdpAddress
}

func (g *GetEsimOkResponseEsim) SetSmdpAddress(smdpAddress string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["SmdpAddress"] = true
	g.SmdpAddress = &smdpAddress
}

func (g *GetEsimOkResponseEsim) SetSmdpAddressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["SmdpAddress"] = true
	g.SmdpAddress = nil
}

func (g *GetEsimOkResponseEsim) GetManualActivationCode() *string {
	if g == nil {
		return nil
	}
	return g.ManualActivationCode
}

func (g *GetEsimOkResponseEsim) SetManualActivationCode(manualActivationCode string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ManualActivationCode"] = true
	g.ManualActivationCode = &manualActivationCode
}

func (g *GetEsimOkResponseEsim) SetManualActivationCodeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ManualActivationCode"] = true
	g.ManualActivationCode = nil
}

func (g *GetEsimOkResponseEsim) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetEsimOkResponseEsim) SetStatus(status string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = &status
}

func (g *GetEsimOkResponseEsim) SetStatusNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = nil
}
func (g GetEsimOkResponseEsim) MarshalJSON() ([]byte, error) {
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

	if g.touched["Status"] && g.Status == nil {
		data["status"] = nil
	} else if g.Status != nil {
		data["status"] = g.Status
	}

	return json.Marshal(data)
}
