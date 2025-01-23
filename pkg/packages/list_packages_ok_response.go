package packages

import (
	"encoding/json"
)

type ListPackagesOkResponse struct {
	Packages []Packages `json:"packages,omitempty"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination
	AfterCursor *string `json:"afterCursor,omitempty"`
	touched     map[string]bool
}

func (l *ListPackagesOkResponse) GetPackages() []Packages {
	if l == nil {
		return nil
	}
	return l.Packages
}

func (l *ListPackagesOkResponse) SetPackages(packages []Packages) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Packages"] = true
	l.Packages = packages
}

func (l *ListPackagesOkResponse) SetPackagesNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Packages"] = true
	l.Packages = nil
}

func (l *ListPackagesOkResponse) GetAfterCursor() *string {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursor(afterCursor string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["AfterCursor"] = true
	l.AfterCursor = &afterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursorNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["AfterCursor"] = true
	l.AfterCursor = nil
}

func (l ListPackagesOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Packages"] && l.Packages == nil {
		data["packages"] = nil
	} else if l.Packages != nil {
		data["packages"] = l.Packages
	}

	if l.touched["AfterCursor"] && l.AfterCursor == nil {
		data["afterCursor"] = nil
	} else if l.AfterCursor != nil {
		data["afterCursor"] = l.AfterCursor
	}

	return json.Marshal(data)
}

func (l ListPackagesOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPackagesOkResponse to string"
	}
	return string(jsonData)
}

type Packages struct {
	// ID of the package
	Id *string `json:"id,omitempty"`
	// ISO representation of the package's destination
	Destination *string `json:"destination,omitempty"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty"`
	// Min number of days for the package
	MinDays *float64 `json:"minDays,omitempty"`
	// Max number of days for the package
	MaxDays *float64 `json:"maxDays,omitempty"`
	// Price of the package in cents
	PriceInCents *float64 `json:"priceInCents,omitempty"`
	touched      map[string]bool
}

func (p *Packages) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Packages) SetId(id string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = &id
}

func (p *Packages) SetIdNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = nil
}

func (p *Packages) GetDestination() *string {
	if p == nil {
		return nil
	}
	return p.Destination
}

func (p *Packages) SetDestination(destination string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Destination"] = true
	p.Destination = &destination
}

func (p *Packages) SetDestinationNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Destination"] = true
	p.Destination = nil
}

func (p *Packages) GetDataLimitInBytes() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInBytes
}

func (p *Packages) SetDataLimitInBytes(dataLimitInBytes float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DataLimitInBytes"] = true
	p.DataLimitInBytes = &dataLimitInBytes
}

func (p *Packages) SetDataLimitInBytesNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DataLimitInBytes"] = true
	p.DataLimitInBytes = nil
}

func (p *Packages) GetMinDays() *float64 {
	if p == nil {
		return nil
	}
	return p.MinDays
}

func (p *Packages) SetMinDays(minDays float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["MinDays"] = true
	p.MinDays = &minDays
}

func (p *Packages) SetMinDaysNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["MinDays"] = true
	p.MinDays = nil
}

func (p *Packages) GetMaxDays() *float64 {
	if p == nil {
		return nil
	}
	return p.MaxDays
}

func (p *Packages) SetMaxDays(maxDays float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["MaxDays"] = true
	p.MaxDays = &maxDays
}

func (p *Packages) SetMaxDaysNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["MaxDays"] = true
	p.MaxDays = nil
}

func (p *Packages) GetPriceInCents() *float64 {
	if p == nil {
		return nil
	}
	return p.PriceInCents
}

func (p *Packages) SetPriceInCents(priceInCents float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["PriceInCents"] = true
	p.PriceInCents = &priceInCents
}

func (p *Packages) SetPriceInCentsNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["PriceInCents"] = true
	p.PriceInCents = nil
}

func (p Packages) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Id"] && p.Id == nil {
		data["id"] = nil
	} else if p.Id != nil {
		data["id"] = p.Id
	}

	if p.touched["Destination"] && p.Destination == nil {
		data["destination"] = nil
	} else if p.Destination != nil {
		data["destination"] = p.Destination
	}

	if p.touched["DataLimitInBytes"] && p.DataLimitInBytes == nil {
		data["dataLimitInBytes"] = nil
	} else if p.DataLimitInBytes != nil {
		data["dataLimitInBytes"] = p.DataLimitInBytes
	}

	if p.touched["MinDays"] && p.MinDays == nil {
		data["minDays"] = nil
	} else if p.MinDays != nil {
		data["minDays"] = p.MinDays
	}

	if p.touched["MaxDays"] && p.MaxDays == nil {
		data["maxDays"] = nil
	} else if p.MaxDays != nil {
		data["maxDays"] = p.MaxDays
	}

	if p.touched["PriceInCents"] && p.PriceInCents == nil {
		data["priceInCents"] = nil
	} else if p.PriceInCents != nil {
		data["priceInCents"] = p.PriceInCents
	}

	return json.Marshal(data)
}

func (p Packages) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Packages to string"
	}
	return string(jsonData)
}
