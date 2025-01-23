package purchases

import (
	"encoding/json"
)

type CreatePurchaseRequest struct {
	// ISO representation of the package's destination
	Destination *string `json:"destination,omitempty" required:"true"`
	// Size of the package in GB. The available options are 1, 2, 3, 5, 8, 20GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Email address where the purchase confirmation email will be sent (including QR Code & activation steps)
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Customize the network brand of the issued eSIM. This parameter is accessible to platforms with Diamond tier and requires an alphanumeric string of up to 15 characters
	NetworkBrand *string `json:"networkBrand,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
	touched map[string]bool
}

func (c *CreatePurchaseRequest) GetDestination() *string {
	if c == nil {
		return nil
	}
	return c.Destination
}

func (c *CreatePurchaseRequest) SetDestination(destination string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Destination"] = true
	c.Destination = &destination
}

func (c *CreatePurchaseRequest) SetDestinationNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Destination"] = true
	c.Destination = nil
}

func (c *CreatePurchaseRequest) GetDataLimitInGb() *float64 {
	if c == nil {
		return nil
	}
	return c.DataLimitInGb
}

func (c *CreatePurchaseRequest) SetDataLimitInGb(dataLimitInGb float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DataLimitInGb"] = true
	c.DataLimitInGb = &dataLimitInGb
}

func (c *CreatePurchaseRequest) SetDataLimitInGbNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DataLimitInGb"] = true
	c.DataLimitInGb = nil
}

func (c *CreatePurchaseRequest) GetStartDate() *string {
	if c == nil {
		return nil
	}
	return c.StartDate
}

func (c *CreatePurchaseRequest) SetStartDate(startDate string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartDate"] = true
	c.StartDate = &startDate
}

func (c *CreatePurchaseRequest) SetStartDateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartDate"] = true
	c.StartDate = nil
}

func (c *CreatePurchaseRequest) GetEndDate() *string {
	if c == nil {
		return nil
	}
	return c.EndDate
}

func (c *CreatePurchaseRequest) SetEndDate(endDate string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndDate"] = true
	c.EndDate = &endDate
}

func (c *CreatePurchaseRequest) SetEndDateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndDate"] = true
	c.EndDate = nil
}

func (c *CreatePurchaseRequest) GetEmail() *string {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreatePurchaseRequest) SetEmail(email string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Email"] = true
	c.Email = &email
}

func (c *CreatePurchaseRequest) SetEmailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Email"] = true
	c.Email = nil
}

func (c *CreatePurchaseRequest) GetReferenceId() *string {
	if c == nil {
		return nil
	}
	return c.ReferenceId
}

func (c *CreatePurchaseRequest) SetReferenceId(referenceId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReferenceId"] = true
	c.ReferenceId = &referenceId
}

func (c *CreatePurchaseRequest) SetReferenceIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReferenceId"] = true
	c.ReferenceId = nil
}

func (c *CreatePurchaseRequest) GetNetworkBrand() *string {
	if c == nil {
		return nil
	}
	return c.NetworkBrand
}

func (c *CreatePurchaseRequest) SetNetworkBrand(networkBrand string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NetworkBrand"] = true
	c.NetworkBrand = &networkBrand
}

func (c *CreatePurchaseRequest) SetNetworkBrandNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NetworkBrand"] = true
	c.NetworkBrand = nil
}

func (c *CreatePurchaseRequest) GetStartTime() *float64 {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *CreatePurchaseRequest) SetStartTime(startTime float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = &startTime
}

func (c *CreatePurchaseRequest) SetStartTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = nil
}

func (c *CreatePurchaseRequest) GetEndTime() *float64 {
	if c == nil {
		return nil
	}
	return c.EndTime
}

func (c *CreatePurchaseRequest) SetEndTime(endTime float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndTime"] = true
	c.EndTime = &endTime
}

func (c *CreatePurchaseRequest) SetEndTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndTime"] = true
	c.EndTime = nil
}

func (c CreatePurchaseRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Destination"] && c.Destination == nil {
		data["destination"] = nil
	} else if c.Destination != nil {
		data["destination"] = c.Destination
	}

	if c.touched["DataLimitInGb"] && c.DataLimitInGb == nil {
		data["dataLimitInGB"] = nil
	} else if c.DataLimitInGb != nil {
		data["dataLimitInGB"] = c.DataLimitInGb
	}

	if c.touched["StartDate"] && c.StartDate == nil {
		data["startDate"] = nil
	} else if c.StartDate != nil {
		data["startDate"] = c.StartDate
	}

	if c.touched["EndDate"] && c.EndDate == nil {
		data["endDate"] = nil
	} else if c.EndDate != nil {
		data["endDate"] = c.EndDate
	}

	if c.touched["Email"] && c.Email == nil {
		data["email"] = nil
	} else if c.Email != nil {
		data["email"] = c.Email
	}

	if c.touched["ReferenceId"] && c.ReferenceId == nil {
		data["referenceId"] = nil
	} else if c.ReferenceId != nil {
		data["referenceId"] = c.ReferenceId
	}

	if c.touched["NetworkBrand"] && c.NetworkBrand == nil {
		data["networkBrand"] = nil
	} else if c.NetworkBrand != nil {
		data["networkBrand"] = c.NetworkBrand
	}

	if c.touched["StartTime"] && c.StartTime == nil {
		data["startTime"] = nil
	} else if c.StartTime != nil {
		data["startTime"] = c.StartTime
	}

	if c.touched["EndTime"] && c.EndTime == nil {
		data["endTime"] = nil
	} else if c.EndTime != nil {
		data["endTime"] = c.EndTime
	}

	return json.Marshal(data)
}

func (c CreatePurchaseRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseRequest to string"
	}
	return string(jsonData)
}
