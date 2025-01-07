package purchases

import (
	"encoding/json"
)

type TopUpEsimOkResponse struct {
	Purchase *TopUpEsimOkResponsePurchase `json:"purchase,omitempty"`
	Profile  *TopUpEsimOkResponseProfile  `json:"profile,omitempty"`
	touched  map[string]bool
}

func (t *TopUpEsimOkResponse) GetPurchase() *TopUpEsimOkResponsePurchase {
	if t == nil {
		return nil
	}
	return t.Purchase
}

func (t *TopUpEsimOkResponse) SetPurchase(purchase TopUpEsimOkResponsePurchase) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Purchase"] = true
	t.Purchase = &purchase
}

func (t *TopUpEsimOkResponse) SetPurchaseNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Purchase"] = true
	t.Purchase = nil
}

func (t *TopUpEsimOkResponse) GetProfile() *TopUpEsimOkResponseProfile {
	if t == nil {
		return nil
	}
	return t.Profile
}

func (t *TopUpEsimOkResponse) SetProfile(profile TopUpEsimOkResponseProfile) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Profile"] = true
	t.Profile = &profile
}

func (t *TopUpEsimOkResponse) SetProfileNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Profile"] = true
	t.Profile = nil
}
func (t TopUpEsimOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Purchase"] && t.Purchase == nil {
		data["purchase"] = nil
	} else if t.Purchase != nil {
		data["purchase"] = t.Purchase
	}

	if t.touched["Profile"] && t.Profile == nil {
		data["profile"] = nil
	} else if t.Profile != nil {
		data["profile"] = t.Profile
	}

	return json.Marshal(data)
}

type TopUpEsimOkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *string `json:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *string `json:"endDate,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty"`
	// Epoch value representing the start time of the package's validity
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *float64 `json:"endTime,omitempty"`
	touched map[string]bool
}

func (t *TopUpEsimOkResponsePurchase) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TopUpEsimOkResponsePurchase) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TopUpEsimOkResponsePurchase) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TopUpEsimOkResponsePurchase) GetPackageId() *string {
	if t == nil {
		return nil
	}
	return t.PackageId
}

func (t *TopUpEsimOkResponsePurchase) SetPackageId(packageId string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["PackageId"] = true
	t.PackageId = &packageId
}

func (t *TopUpEsimOkResponsePurchase) SetPackageIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["PackageId"] = true
	t.PackageId = nil
}

func (t *TopUpEsimOkResponsePurchase) GetStartDate() *string {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TopUpEsimOkResponsePurchase) SetStartDate(startDate string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartDate"] = true
	t.StartDate = &startDate
}

func (t *TopUpEsimOkResponsePurchase) SetStartDateNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartDate"] = true
	t.StartDate = nil
}

func (t *TopUpEsimOkResponsePurchase) GetEndDate() *string {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TopUpEsimOkResponsePurchase) SetEndDate(endDate string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndDate"] = true
	t.EndDate = &endDate
}

func (t *TopUpEsimOkResponsePurchase) SetEndDateNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndDate"] = true
	t.EndDate = nil
}

func (t *TopUpEsimOkResponsePurchase) GetCreatedDate() *string {
	if t == nil {
		return nil
	}
	return t.CreatedDate
}

func (t *TopUpEsimOkResponsePurchase) SetCreatedDate(createdDate string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedDate"] = true
	t.CreatedDate = &createdDate
}

func (t *TopUpEsimOkResponsePurchase) SetCreatedDateNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedDate"] = true
	t.CreatedDate = nil
}

func (t *TopUpEsimOkResponsePurchase) GetStartTime() *float64 {
	if t == nil {
		return nil
	}
	return t.StartTime
}

func (t *TopUpEsimOkResponsePurchase) SetStartTime(startTime float64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartTime"] = true
	t.StartTime = &startTime
}

func (t *TopUpEsimOkResponsePurchase) SetStartTimeNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["StartTime"] = true
	t.StartTime = nil
}

func (t *TopUpEsimOkResponsePurchase) GetEndTime() *float64 {
	if t == nil {
		return nil
	}
	return t.EndTime
}

func (t *TopUpEsimOkResponsePurchase) SetEndTime(endTime float64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndTime"] = true
	t.EndTime = &endTime
}

func (t *TopUpEsimOkResponsePurchase) SetEndTimeNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["EndTime"] = true
	t.EndTime = nil
}
func (t TopUpEsimOkResponsePurchase) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Id"] && t.Id == nil {
		data["id"] = nil
	} else if t.Id != nil {
		data["id"] = t.Id
	}

	if t.touched["PackageId"] && t.PackageId == nil {
		data["packageId"] = nil
	} else if t.PackageId != nil {
		data["packageId"] = t.PackageId
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

	if t.touched["CreatedDate"] && t.CreatedDate == nil {
		data["createdDate"] = nil
	} else if t.CreatedDate != nil {
		data["createdDate"] = t.CreatedDate
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

type TopUpEsimOkResponseProfile struct {
	// ID of the eSIM
	Iccid   *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	touched map[string]bool
}

func (t *TopUpEsimOkResponseProfile) GetIccid() *string {
	if t == nil {
		return nil
	}
	return t.Iccid
}

func (t *TopUpEsimOkResponseProfile) SetIccid(iccid string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Iccid"] = true
	t.Iccid = &iccid
}

func (t *TopUpEsimOkResponseProfile) SetIccidNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Iccid"] = true
	t.Iccid = nil
}
func (t TopUpEsimOkResponseProfile) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Iccid"] && t.Iccid == nil {
		data["iccid"] = nil
	} else if t.Iccid != nil {
		data["iccid"] = t.Iccid
	}

	return json.Marshal(data)
}
