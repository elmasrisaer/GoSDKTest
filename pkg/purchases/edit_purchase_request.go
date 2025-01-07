package purchases

import (
	"encoding/json"
)

type EditPurchaseRequest struct {
	// ID of the purchase
	PurchaseId *string `json:"purchaseId,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.
	StartDate *string `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.
	EndDate *string `json:"endDate,omitempty" required:"true"`
	// Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.
	EndTime *float64 `json:"endTime,omitempty"`
	touched map[string]bool
}

func (e *EditPurchaseRequest) GetPurchaseId() *string {
	if e == nil {
		return nil
	}
	return e.PurchaseId
}

func (e *EditPurchaseRequest) SetPurchaseId(purchaseId string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["PurchaseId"] = true
	e.PurchaseId = &purchaseId
}

func (e *EditPurchaseRequest) SetPurchaseIdNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["PurchaseId"] = true
	e.PurchaseId = nil
}

func (e *EditPurchaseRequest) GetStartDate() *string {
	if e == nil {
		return nil
	}
	return e.StartDate
}

func (e *EditPurchaseRequest) SetStartDate(startDate string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["StartDate"] = true
	e.StartDate = &startDate
}

func (e *EditPurchaseRequest) SetStartDateNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["StartDate"] = true
	e.StartDate = nil
}

func (e *EditPurchaseRequest) GetEndDate() *string {
	if e == nil {
		return nil
	}
	return e.EndDate
}

func (e *EditPurchaseRequest) SetEndDate(endDate string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["EndDate"] = true
	e.EndDate = &endDate
}

func (e *EditPurchaseRequest) SetEndDateNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["EndDate"] = true
	e.EndDate = nil
}

func (e *EditPurchaseRequest) GetStartTime() *float64 {
	if e == nil {
		return nil
	}
	return e.StartTime
}

func (e *EditPurchaseRequest) SetStartTime(startTime float64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["StartTime"] = true
	e.StartTime = &startTime
}

func (e *EditPurchaseRequest) SetStartTimeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["StartTime"] = true
	e.StartTime = nil
}

func (e *EditPurchaseRequest) GetEndTime() *float64 {
	if e == nil {
		return nil
	}
	return e.EndTime
}

func (e *EditPurchaseRequest) SetEndTime(endTime float64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["EndTime"] = true
	e.EndTime = &endTime
}

func (e *EditPurchaseRequest) SetEndTimeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["EndTime"] = true
	e.EndTime = nil
}
func (e EditPurchaseRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["PurchaseId"] && e.PurchaseId == nil {
		data["purchaseId"] = nil
	} else if e.PurchaseId != nil {
		data["purchaseId"] = e.PurchaseId
	}

	if e.touched["StartDate"] && e.StartDate == nil {
		data["startDate"] = nil
	} else if e.StartDate != nil {
		data["startDate"] = e.StartDate
	}

	if e.touched["EndDate"] && e.EndDate == nil {
		data["endDate"] = nil
	} else if e.EndDate != nil {
		data["endDate"] = e.EndDate
	}

	if e.touched["StartTime"] && e.StartTime == nil {
		data["startTime"] = nil
	} else if e.StartTime != nil {
		data["startTime"] = e.StartTime
	}

	if e.touched["EndTime"] && e.EndTime == nil {
		data["endTime"] = nil
	} else if e.EndTime != nil {
		data["endTime"] = e.EndTime
	}

	return json.Marshal(data)
}
