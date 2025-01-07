package purchases

import (
	"encoding/json"
)

type TopUpEsimRequest struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" required:"true" maxLength:"22" minLength:"18"`
	// Size of the package in GB. The available options are 1, 2, 3, 5, 8, 20GB
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps)
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
	touched map[string]bool
}

func (t *TopUpEsimRequest) GetIccid() *string {
	if t == nil {
		return nil
	}
	return t.Iccid
}

func (t *TopUpEsimRequest) SetIccid(iccid string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Iccid"] = true
	t.Iccid = &iccid
}

func (t *TopUpEsimRequest) SetIccidNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Iccid"] = true
	t.Iccid = nil
}

func (t *TopUpEsimRequest) GetDataLimitInGb() *float64 {
	if t == nil {
		return nil
	}
	return t.DataLimitInGb
}

func (t *TopUpEsimRequest) SetDataLimitInGb(dataLimitInGb float64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DataLimitInGb"] = true
	t.DataLimitInGb = &dataLimitInGb
}

func (t *TopUpEsimRequest) SetDataLimitInGbNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DataLimitInGb"] = true
	t.DataLimitInGb = nil
}

func (t *TopUpEsimRequest) GetStartDate() *string {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TopUpEsimRequest) SetStartDate(startDate string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartDate"] = true
	t.StartDate = &startDate
}

func (t *TopUpEsimRequest) SetStartDateNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartDate"] = true
	t.StartDate = nil
}

func (t *TopUpEsimRequest) GetEndDate() *string {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TopUpEsimRequest) SetEndDate(endDate string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndDate"] = true
	t.EndDate = &endDate
}

func (t *TopUpEsimRequest) SetEndDateNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndDate"] = true
	t.EndDate = nil
}

func (t *TopUpEsimRequest) GetEmail() *string {
	if t == nil {
		return nil
	}
	return t.Email
}

func (t *TopUpEsimRequest) SetEmail(email string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Email"] = true
	t.Email = &email
}

func (t *TopUpEsimRequest) SetEmailNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Email"] = true
	t.Email = nil
}

func (t *TopUpEsimRequest) GetReferenceId() *string {
	if t == nil {
		return nil
	}
	return t.ReferenceId
}

func (t *TopUpEsimRequest) SetReferenceId(referenceId string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ReferenceId"] = true
	t.ReferenceId = &referenceId
}

func (t *TopUpEsimRequest) SetReferenceIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ReferenceId"] = true
	t.ReferenceId = nil
}

func (t *TopUpEsimRequest) GetStartTime() *float64 {
	if t == nil {
		return nil
	}
	return t.StartTime
}

func (t *TopUpEsimRequest) SetStartTime(startTime float64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartTime"] = true
	t.StartTime = &startTime
}

func (t *TopUpEsimRequest) SetStartTimeNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartTime"] = true
	t.StartTime = nil
}

func (t *TopUpEsimRequest) GetEndTime() *float64 {
	if t == nil {
		return nil
	}
	return t.EndTime
}

func (t *TopUpEsimRequest) SetEndTime(endTime float64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndTime"] = true
	t.EndTime = &endTime
}

func (t *TopUpEsimRequest) SetEndTimeNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndTime"] = true
	t.EndTime = nil
}
func (t TopUpEsimRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Iccid"] && t.Iccid == nil {
		data["iccid"] = nil
	} else if t.Iccid != nil {
		data["iccid"] = t.Iccid
	}

	if t.touched["DataLimitInGb"] && t.DataLimitInGb == nil {
		data["dataLimitInGB"] = nil
	} else if t.DataLimitInGb != nil {
		data["dataLimitInGB"] = t.DataLimitInGb
	}

	if t.touched["StartDate"] && t.StartDate == nil {
		data["startDate"] = nil
	} else if t.StartDate != nil {
		data["startDate"] = t.StartDate
	}

	if t.touched["EndDate"] && t.EndDate == nil {
		data["endDate"] = nil
	} else if t.EndDate != nil {
		data["endDate"] = t.EndDate
	}

	if t.touched["Email"] && t.Email == nil {
		data["email"] = nil
	} else if t.Email != nil {
		data["email"] = t.Email
	}

	if t.touched["ReferenceId"] && t.ReferenceId == nil {
		data["referenceId"] = nil
	} else if t.ReferenceId != nil {
		data["referenceId"] = t.ReferenceId
	}

	if t.touched["StartTime"] && t.StartTime == nil {
		data["startTime"] = nil
	} else if t.StartTime != nil {
		data["startTime"] = t.StartTime
	}

	if t.touched["EndTime"] && t.EndTime == nil {
		data["endTime"] = nil
	} else if t.EndTime != nil {
		data["endTime"] = t.EndTime
	}

	return json.Marshal(data)
}
