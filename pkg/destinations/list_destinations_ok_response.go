package destinations

import (
	"encoding/json"
)

type ListDestinationsOkResponse struct {
	Destinations []Destinations `json:"destinations,omitempty"`
	touched      map[string]bool
}

func (l *ListDestinationsOkResponse) GetDestinations() []Destinations {
	if l == nil {
		return nil
	}
	return l.Destinations
}

func (l *ListDestinationsOkResponse) SetDestinations(destinations []Destinations) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Destinations"] = true
	l.Destinations = destinations
}

func (l *ListDestinationsOkResponse) SetDestinationsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Destinations"] = true
	l.Destinations = nil
}
func (l ListDestinationsOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Destinations"] && l.Destinations == nil {
		data["destinations"] = nil
	} else if l.Destinations != nil {
		data["destinations"] = l.Destinations
	}

	return json.Marshal(data)
}

type Destinations struct {
	// Name of the destination
	Name *string `json:"name,omitempty"`
	// ISO representation of the destination
	Destination *string `json:"destination,omitempty"`
	// This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region.
	SupportedCountries []string `json:"supportedCountries,omitempty"`
	touched            map[string]bool
}

func (d *Destinations) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Destinations) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *Destinations) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *Destinations) GetDestination() *string {
	if d == nil {
		return nil
	}
	return d.Destination
}

func (d *Destinations) SetDestination(destination string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Destination"] = true
	d.Destination = &destination
}

func (d *Destinations) SetDestinationNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Destination"] = true
	d.Destination = nil
}

func (d *Destinations) GetSupportedCountries() []string {
	if d == nil {
		return nil
	}
	return d.SupportedCountries
}

func (d *Destinations) SetSupportedCountries(supportedCountries []string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["SupportedCountries"] = true
	d.SupportedCountries = supportedCountries
}

func (d *Destinations) SetSupportedCountriesNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["SupportedCountries"] = true
	d.SupportedCountries = nil
}
func (d Destinations) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["Destination"] && d.Destination == nil {
		data["destination"] = nil
	} else if d.Destination != nil {
		data["destination"] = d.Destination
	}

	if d.touched["SupportedCountries"] && d.SupportedCountries == nil {
		data["supportedCountries"] = nil
	} else if d.SupportedCountries != nil {
		data["supportedCountries"] = d.SupportedCountries
	}

	return json.Marshal(data)
}
