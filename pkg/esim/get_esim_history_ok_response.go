package esim

import (
	"encoding/json"
)

type GetEsimHistoryOkResponse struct {
	Esim    *GetEsimHistoryOkResponseEsim `json:"esim,omitempty"`
	touched map[string]bool
}

func (g *GetEsimHistoryOkResponse) GetEsim() *GetEsimHistoryOkResponseEsim {
	if g == nil {
		return nil
	}
	return g.Esim
}

func (g *GetEsimHistoryOkResponse) SetEsim(esim GetEsimHistoryOkResponseEsim) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = &esim
}

func (g *GetEsimHistoryOkResponse) SetEsimNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Esim"] = true
	g.Esim = nil
}

func (g GetEsimHistoryOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Esim"] && g.Esim == nil {
		data["esim"] = nil
	} else if g.Esim != nil {
		data["esim"] = g.Esim
	}

	return json.Marshal(data)
}

func (g GetEsimHistoryOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponse to string"
	}
	return string(jsonData)
}

type GetEsimHistoryOkResponseEsim struct {
	// ID of the eSIM
	Iccid   *string   `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	History []History `json:"history,omitempty"`
	touched map[string]bool
}

func (g *GetEsimHistoryOkResponseEsim) GetIccid() *string {
	if g == nil {
		return nil
	}
	return g.Iccid
}

func (g *GetEsimHistoryOkResponseEsim) SetIccid(iccid string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = &iccid
}

func (g *GetEsimHistoryOkResponseEsim) SetIccidNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Iccid"] = true
	g.Iccid = nil
}

func (g *GetEsimHistoryOkResponseEsim) GetHistory() []History {
	if g == nil {
		return nil
	}
	return g.History
}

func (g *GetEsimHistoryOkResponseEsim) SetHistory(history []History) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["History"] = true
	g.History = history
}

func (g *GetEsimHistoryOkResponseEsim) SetHistoryNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["History"] = true
	g.History = nil
}

func (g GetEsimHistoryOkResponseEsim) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Iccid"] && g.Iccid == nil {
		data["iccid"] = nil
	} else if g.Iccid != nil {
		data["iccid"] = g.Iccid
	}

	if g.touched["History"] && g.History == nil {
		data["history"] = nil
	} else if g.History != nil {
		data["history"] = g.History
	}

	return json.Marshal(data)
}

func (g GetEsimHistoryOkResponseEsim) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetEsimHistoryOkResponseEsim to string"
	}
	return string(jsonData)
}

type History struct {
	// The status of the eSIM at a given time, possible values are 'RELEASED', 'DOWNLOADED', 'INSTALLED', 'ENABLED', 'DELETED', or 'ERROR'
	Status *string `json:"status,omitempty"`
	// The date when the eSIM status changed in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StatusDate *string `json:"statusDate,omitempty"`
	// Epoch value representing the date when the eSIM status changed
	Date    *float64 `json:"date,omitempty"`
	touched map[string]bool
}

func (h *History) GetStatus() *string {
	if h == nil {
		return nil
	}
	return h.Status
}

func (h *History) SetStatus(status string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Status"] = true
	h.Status = &status
}

func (h *History) SetStatusNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Status"] = true
	h.Status = nil
}

func (h *History) GetStatusDate() *string {
	if h == nil {
		return nil
	}
	return h.StatusDate
}

func (h *History) SetStatusDate(statusDate string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["StatusDate"] = true
	h.StatusDate = &statusDate
}

func (h *History) SetStatusDateNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["StatusDate"] = true
	h.StatusDate = nil
}

func (h *History) GetDate() *float64 {
	if h == nil {
		return nil
	}
	return h.Date
}

func (h *History) SetDate(date float64) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Date"] = true
	h.Date = &date
}

func (h *History) SetDateNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Date"] = true
	h.Date = nil
}

func (h History) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if h.touched["Status"] && h.Status == nil {
		data["status"] = nil
	} else if h.Status != nil {
		data["status"] = h.Status
	}

	if h.touched["StatusDate"] && h.StatusDate == nil {
		data["statusDate"] = nil
	} else if h.StatusDate != nil {
		data["statusDate"] = h.StatusDate
	}

	if h.touched["Date"] && h.Date == nil {
		data["date"] = nil
	} else if h.Date != nil {
		data["date"] = h.Date
	}

	return json.Marshal(data)
}

func (h History) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: History to string"
	}
	return string(jsonData)
}
