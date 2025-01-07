package purchases

import (
	"encoding/json"
)

type EditPurchaseOkResponse struct {
	// ID of the purchase
	PurchaseId *string `json:"purchaseId,omitempty"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewStartDate *string `json:"newStartDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	NewEndDate *string `json:"newEndDate,omitempty"`
	// Epoch value representing the new start time of the package's validity
	NewStartTime *float64 `json:"newStartTime,omitempty"`
	// Epoch value representing the new end time of the package's validity
	NewEndTime *float64 `json:"newEndTime,omitempty"`
	touched    map[string]bool
}

func (e *EditPurchaseOkResponse) GetPurchaseId() *string {
	if e == nil {
		return nil
	}
	return e.PurchaseId
}

func (e *EditPurchaseOkResponse) SetPurchaseId(purchaseId string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["PurchaseId"] = true
	e.PurchaseId = &purchaseId
}

func (e *EditPurchaseOkResponse) SetPurchaseIdNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["PurchaseId"] = true
	e.PurchaseId = nil
}

func (e *EditPurchaseOkResponse) GetNewStartDate() *string {
	if e == nil {
		return nil
	}
	return e.NewStartDate
}

func (e *EditPurchaseOkResponse) SetNewStartDate(newStartDate string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewStartDate"] = true
	e.NewStartDate = &newStartDate
}

func (e *EditPurchaseOkResponse) SetNewStartDateNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewStartDate"] = true
	e.NewStartDate = nil
}

func (e *EditPurchaseOkResponse) GetNewEndDate() *string {
	if e == nil {
		return nil
	}
	return e.NewEndDate
}

func (e *EditPurchaseOkResponse) SetNewEndDate(newEndDate string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewEndDate"] = true
	e.NewEndDate = &newEndDate
}

func (e *EditPurchaseOkResponse) SetNewEndDateNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewEndDate"] = true
	e.NewEndDate = nil
}

func (e *EditPurchaseOkResponse) GetNewStartTime() *float64 {
	if e == nil {
		return nil
	}
	return e.NewStartTime
}

func (e *EditPurchaseOkResponse) SetNewStartTime(newStartTime float64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewStartTime"] = true
	e.NewStartTime = &newStartTime
}

func (e *EditPurchaseOkResponse) SetNewStartTimeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewStartTime"] = true
	e.NewStartTime = nil
}

func (e *EditPurchaseOkResponse) GetNewEndTime() *float64 {
	if e == nil {
		return nil
	}
	return e.NewEndTime
}

func (e *EditPurchaseOkResponse) SetNewEndTime(newEndTime float64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewEndTime"] = true
	e.NewEndTime = &newEndTime
}

func (e *EditPurchaseOkResponse) SetNewEndTimeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NewEndTime"] = true
	e.NewEndTime = nil
}
func (e EditPurchaseOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["PurchaseId"] && e.PurchaseId == nil {
		data["purchaseId"] = nil
	} else if e.PurchaseId != nil {
		data["purchaseId"] = e.PurchaseId
	}

	if e.touched["NewStartDate"] && e.NewStartDate == nil {
		data["newStartDate"] = nil
	} else if e.NewStartDate != nil {
		data["newStartDate"] = e.NewStartDate
	}

	if e.touched["NewEndDate"] && e.NewEndDate == nil {
		data["newEndDate"] = nil
	} else if e.NewEndDate != nil {
		data["newEndDate"] = e.NewEndDate
	}

	if e.touched["NewStartTime"] && e.NewStartTime == nil {
		data["newStartTime"] = nil
	} else if e.NewStartTime != nil {
		data["newStartTime"] = e.NewStartTime
	}

	if e.touched["NewEndTime"] && e.NewEndTime == nil {
		data["newEndTime"] = nil
	} else if e.NewEndTime != nil {
		data["newEndTime"] = e.NewEndTime
	}

	return json.Marshal(data)
}
