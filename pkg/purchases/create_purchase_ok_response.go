package purchases

import (
	"encoding/json"
)

type CreatePurchaseOkResponse struct {
	Purchase *CreatePurchaseOkResponsePurchase `json:"purchase,omitempty"`
	Profile  *CreatePurchaseOkResponseProfile  `json:"profile,omitempty"`
	touched  map[string]bool
}

func (c *CreatePurchaseOkResponse) GetPurchase() *CreatePurchaseOkResponsePurchase {
	if c == nil {
		return nil
	}
	return c.Purchase
}

func (c *CreatePurchaseOkResponse) SetPurchase(purchase CreatePurchaseOkResponsePurchase) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Purchase"] = true
	c.Purchase = &purchase
}

func (c *CreatePurchaseOkResponse) SetPurchaseNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Purchase"] = true
	c.Purchase = nil
}

func (c *CreatePurchaseOkResponse) GetProfile() *CreatePurchaseOkResponseProfile {
	if c == nil {
		return nil
	}
	return c.Profile
}

func (c *CreatePurchaseOkResponse) SetProfile(profile CreatePurchaseOkResponseProfile) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Profile"] = true
	c.Profile = &profile
}

func (c *CreatePurchaseOkResponse) SetProfileNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Profile"] = true
	c.Profile = nil
}
func (c CreatePurchaseOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Purchase"] && c.Purchase == nil {
		data["purchase"] = nil
	} else if c.Purchase != nil {
		data["purchase"] = c.Purchase
	}

	if c.touched["Profile"] && c.Profile == nil {
		data["profile"] = nil
	} else if c.Profile != nil {
		data["profile"] = c.Profile
	}

	return json.Marshal(data)
}

type CreatePurchaseOkResponsePurchase struct {
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

func (c *CreatePurchaseOkResponsePurchase) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreatePurchaseOkResponsePurchase) SetId(id string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Id"] = true
	c.Id = &id
}

func (c *CreatePurchaseOkResponsePurchase) SetIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Id"] = true
	c.Id = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetPackageId() *string {
	if c == nil {
		return nil
	}
	return c.PackageId
}

func (c *CreatePurchaseOkResponsePurchase) SetPackageId(packageId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PackageId"] = true
	c.PackageId = &packageId
}

func (c *CreatePurchaseOkResponsePurchase) SetPackageIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PackageId"] = true
	c.PackageId = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetStartDate() *string {
	if c == nil {
		return nil
	}
	return c.StartDate
}

func (c *CreatePurchaseOkResponsePurchase) SetStartDate(startDate string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartDate"] = true
	c.StartDate = &startDate
}

func (c *CreatePurchaseOkResponsePurchase) SetStartDateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartDate"] = true
	c.StartDate = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetEndDate() *string {
	if c == nil {
		return nil
	}
	return c.EndDate
}

func (c *CreatePurchaseOkResponsePurchase) SetEndDate(endDate string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndDate"] = true
	c.EndDate = &endDate
}

func (c *CreatePurchaseOkResponsePurchase) SetEndDateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndDate"] = true
	c.EndDate = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetCreatedDate() *string {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePurchaseOkResponsePurchase) SetCreatedDate(createdDate string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreatedDate"] = true
	c.CreatedDate = &createdDate
}

func (c *CreatePurchaseOkResponsePurchase) SetCreatedDateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreatedDate"] = true
	c.CreatedDate = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetStartTime() *float64 {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *CreatePurchaseOkResponsePurchase) SetStartTime(startTime float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = &startTime
}

func (c *CreatePurchaseOkResponsePurchase) SetStartTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = nil
}

func (c *CreatePurchaseOkResponsePurchase) GetEndTime() *float64 {
	if c == nil {
		return nil
	}
	return c.EndTime
}

func (c *CreatePurchaseOkResponsePurchase) SetEndTime(endTime float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndTime"] = true
	c.EndTime = &endTime
}

func (c *CreatePurchaseOkResponsePurchase) SetEndTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EndTime"] = true
	c.EndTime = nil
}
func (c CreatePurchaseOkResponsePurchase) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Id"] && c.Id == nil {
		data["id"] = nil
	} else if c.Id != nil {
		data["id"] = c.Id
	}

	if c.touched["PackageId"] && c.PackageId == nil {
		data["packageId"] = nil
	} else if c.PackageId != nil {
		data["packageId"] = c.PackageId
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

	if c.touched["CreatedDate"] && c.CreatedDate == nil {
		data["createdDate"] = nil
	} else if c.CreatedDate != nil {
		data["createdDate"] = c.CreatedDate
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

type CreatePurchaseOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	// QR Code of the eSIM as base64
	ActivationCode *string `json:"activationCode,omitempty" maxLength:"8000" minLength:"1000"`
	// Manual Activation Code of the eSIM
	ManualActivationCode *string `json:"manualActivationCode,omitempty"`
	touched              map[string]bool
}

func (c *CreatePurchaseOkResponseProfile) GetIccid() *string {
	if c == nil {
		return nil
	}
	return c.Iccid
}

func (c *CreatePurchaseOkResponseProfile) SetIccid(iccid string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Iccid"] = true
	c.Iccid = &iccid
}

func (c *CreatePurchaseOkResponseProfile) SetIccidNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Iccid"] = true
	c.Iccid = nil
}

func (c *CreatePurchaseOkResponseProfile) GetActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetActivationCode(activationCode string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ActivationCode"] = true
	c.ActivationCode = &activationCode
}

func (c *CreatePurchaseOkResponseProfile) SetActivationCodeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ActivationCode"] = true
	c.ActivationCode = nil
}

func (c *CreatePurchaseOkResponseProfile) GetManualActivationCode() *string {
	if c == nil {
		return nil
	}
	return c.ManualActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetManualActivationCode(manualActivationCode string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ManualActivationCode"] = true
	c.ManualActivationCode = &manualActivationCode
}

func (c *CreatePurchaseOkResponseProfile) SetManualActivationCodeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ManualActivationCode"] = true
	c.ManualActivationCode = nil
}
func (c CreatePurchaseOkResponseProfile) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Iccid"] && c.Iccid == nil {
		data["iccid"] = nil
	} else if c.Iccid != nil {
		data["iccid"] = c.Iccid
	}

	if c.touched["ActivationCode"] && c.ActivationCode == nil {
		data["activationCode"] = nil
	} else if c.ActivationCode != nil {
		data["activationCode"] = c.ActivationCode
	}

	if c.touched["ManualActivationCode"] && c.ManualActivationCode == nil {
		data["manualActivationCode"] = nil
	} else if c.ManualActivationCode != nil {
		data["manualActivationCode"] = c.ManualActivationCode
	}

	return json.Marshal(data)
}
