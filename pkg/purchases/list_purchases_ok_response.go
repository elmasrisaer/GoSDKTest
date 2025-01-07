package purchases

import (
	"encoding/json"
)

type ListPurchasesOkResponse struct {
	Purchases []Purchases `json:"purchases,omitempty"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination.
	AfterCursor *string `json:"afterCursor,omitempty"`
	touched     map[string]bool
}

func (l *ListPurchasesOkResponse) GetPurchases() []Purchases {
	if l == nil {
		return nil
	}
	return l.Purchases
}

func (l *ListPurchasesOkResponse) SetPurchases(purchases []Purchases) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Purchases"] = true
	l.Purchases = purchases
}

func (l *ListPurchasesOkResponse) SetPurchasesNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Purchases"] = true
	l.Purchases = nil
}

func (l *ListPurchasesOkResponse) GetAfterCursor() *string {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursor(afterCursor string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["AfterCursor"] = true
	l.AfterCursor = &afterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursorNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["AfterCursor"] = true
	l.AfterCursor = nil
}
func (l ListPurchasesOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Purchases"] && l.Purchases == nil {
		data["purchases"] = nil
	} else if l.Purchases != nil {
		data["purchases"] = l.Purchases
	}

	if l.touched["AfterCursor"] && l.AfterCursor == nil {
		data["afterCursor"] = nil
	} else if l.AfterCursor != nil {
		data["afterCursor"] = l.AfterCursor
	}

	return json.Marshal(data)
}

type Purchases struct {
	// ID of the purchase
	Id *string `json:"id,omitempty"`
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
	// Epoch value representing the date of creation of the purchase
	CreatedAt *float64       `json:"createdAt,omitempty"`
	Package_  *Package_      `json:"package,omitempty"`
	Esim      *PurchasesEsim `json:"esim,omitempty"`
	// The source indicates where the eSIM was purchased, which can be from the API, dashboard, landing-page or promo-page. For purchases made before September 8, 2023, the value will be displayed as 'Not available'.
	Source *string `json:"source,omitempty"`
	// The referenceId that was provided by the partner during the purchase or topup flow. This identifier can be used for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	touched     map[string]bool
}

func (p *Purchases) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Purchases) SetId(id string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = &id
}

func (p *Purchases) SetIdNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = nil
}

func (p *Purchases) GetStartDate() *string {
	if p == nil {
		return nil
	}
	return p.StartDate
}

func (p *Purchases) SetStartDate(startDate string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["StartDate"] = true
	p.StartDate = &startDate
}

func (p *Purchases) SetStartDateNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["StartDate"] = true
	p.StartDate = nil
}

func (p *Purchases) GetEndDate() *string {
	if p == nil {
		return nil
	}
	return p.EndDate
}

func (p *Purchases) SetEndDate(endDate string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["EndDate"] = true
	p.EndDate = &endDate
}

func (p *Purchases) SetEndDateNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["EndDate"] = true
	p.EndDate = nil
}

func (p *Purchases) GetCreatedDate() *string {
	if p == nil {
		return nil
	}
	return p.CreatedDate
}

func (p *Purchases) SetCreatedDate(createdDate string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["CreatedDate"] = true
	p.CreatedDate = &createdDate
}

func (p *Purchases) SetCreatedDateNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["CreatedDate"] = true
	p.CreatedDate = nil
}

func (p *Purchases) GetStartTime() *float64 {
	if p == nil {
		return nil
	}
	return p.StartTime
}

func (p *Purchases) SetStartTime(startTime float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["StartTime"] = true
	p.StartTime = &startTime
}

func (p *Purchases) SetStartTimeNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["StartTime"] = true
	p.StartTime = nil
}

func (p *Purchases) GetEndTime() *float64 {
	if p == nil {
		return nil
	}
	return p.EndTime
}

func (p *Purchases) SetEndTime(endTime float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["EndTime"] = true
	p.EndTime = &endTime
}

func (p *Purchases) SetEndTimeNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["EndTime"] = true
	p.EndTime = nil
}

func (p *Purchases) GetCreatedAt() *float64 {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *Purchases) SetCreatedAt(createdAt float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["CreatedAt"] = true
	p.CreatedAt = &createdAt
}

func (p *Purchases) SetCreatedAtNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["CreatedAt"] = true
	p.CreatedAt = nil
}

func (p *Purchases) GetPackage_() *Package_ {
	if p == nil {
		return nil
	}
	return p.Package_
}

func (p *Purchases) SetPackage_(package_ Package_) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Package_"] = true
	p.Package_ = &package_
}

func (p *Purchases) SetPackage_Nil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Package_"] = true
	p.Package_ = nil
}

func (p *Purchases) GetEsim() *PurchasesEsim {
	if p == nil {
		return nil
	}
	return p.Esim
}

func (p *Purchases) SetEsim(esim PurchasesEsim) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Esim"] = true
	p.Esim = &esim
}

func (p *Purchases) SetEsimNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Esim"] = true
	p.Esim = nil
}

func (p *Purchases) GetSource() *string {
	if p == nil {
		return nil
	}
	return p.Source
}

func (p *Purchases) SetSource(source string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Source"] = true
	p.Source = &source
}

func (p *Purchases) SetSourceNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Source"] = true
	p.Source = nil
}

func (p *Purchases) GetReferenceId() *string {
	if p == nil {
		return nil
	}
	return p.ReferenceId
}

func (p *Purchases) SetReferenceId(referenceId string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["ReferenceId"] = true
	p.ReferenceId = &referenceId
}

func (p *Purchases) SetReferenceIdNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["ReferenceId"] = true
	p.ReferenceId = nil
}
func (p Purchases) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Id"] && p.Id == nil {
		data["id"] = nil
	} else if p.Id != nil {
		data["id"] = p.Id
	}

	if p.touched["StartDate"] && p.StartDate == nil {
		data["startDate"] = nil
	} else if p.StartDate != nil {
		data["startDate"] = p.StartDate
	}

	if p.touched["EndDate"] && p.EndDate == nil {
		data["endDate"] = nil
	} else if p.EndDate != nil {
		data["endDate"] = p.EndDate
	}

	if p.touched["CreatedDate"] && p.CreatedDate == nil {
		data["createdDate"] = nil
	} else if p.CreatedDate != nil {
		data["createdDate"] = p.CreatedDate
	}

	if p.touched["StartTime"] && p.StartTime == nil {
		data["startTime"] = nil
	} else if p.StartTime != nil {
		data["startTime"] = p.StartTime
	}

	if p.touched["EndTime"] && p.EndTime == nil {
		data["endTime"] = nil
	} else if p.EndTime != nil {
		data["endTime"] = p.EndTime
	}

	if p.touched["CreatedAt"] && p.CreatedAt == nil {
		data["createdAt"] = nil
	} else if p.CreatedAt != nil {
		data["createdAt"] = p.CreatedAt
	}

	if p.touched["Package_"] && p.Package_ == nil {
		data["package"] = nil
	} else if p.Package_ != nil {
		data["package"] = p.Package_
	}

	if p.touched["Esim"] && p.Esim == nil {
		data["esim"] = nil
	} else if p.Esim != nil {
		data["esim"] = p.Esim
	}

	if p.touched["Source"] && p.Source == nil {
		data["source"] = nil
	} else if p.Source != nil {
		data["source"] = p.Source
	}

	if p.touched["ReferenceId"] && p.ReferenceId == nil {
		data["referenceId"] = nil
	} else if p.ReferenceId != nil {
		data["referenceId"] = p.ReferenceId
	}

	return json.Marshal(data)
}

type Package_ struct {
	// ID of the package
	Id *string `json:"id,omitempty"`
	// Size of the package in Bytes
	DataLimitInBytes *float64 `json:"dataLimitInBytes,omitempty"`
	// ISO representation of the package's destination
	Destination *string `json:"destination,omitempty"`
	// Name of the package's destination
	DestinationName *string `json:"destinationName,omitempty"`
	// Price of the package in cents
	PriceInCents *float64 `json:"priceInCents,omitempty"`
	touched      map[string]bool
}

func (p *Package_) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Package_) SetId(id string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = &id
}

func (p *Package_) SetIdNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = nil
}

func (p *Package_) GetDataLimitInBytes() *float64 {
	if p == nil {
		return nil
	}
	return p.DataLimitInBytes
}

func (p *Package_) SetDataLimitInBytes(dataLimitInBytes float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DataLimitInBytes"] = true
	p.DataLimitInBytes = &dataLimitInBytes
}

func (p *Package_) SetDataLimitInBytesNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DataLimitInBytes"] = true
	p.DataLimitInBytes = nil
}

func (p *Package_) GetDestination() *string {
	if p == nil {
		return nil
	}
	return p.Destination
}

func (p *Package_) SetDestination(destination string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Destination"] = true
	p.Destination = &destination
}

func (p *Package_) SetDestinationNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Destination"] = true
	p.Destination = nil
}

func (p *Package_) GetDestinationName() *string {
	if p == nil {
		return nil
	}
	return p.DestinationName
}

func (p *Package_) SetDestinationName(destinationName string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DestinationName"] = true
	p.DestinationName = &destinationName
}

func (p *Package_) SetDestinationNameNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["DestinationName"] = true
	p.DestinationName = nil
}

func (p *Package_) GetPriceInCents() *float64 {
	if p == nil {
		return nil
	}
	return p.PriceInCents
}

func (p *Package_) SetPriceInCents(priceInCents float64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["PriceInCents"] = true
	p.PriceInCents = &priceInCents
}

func (p *Package_) SetPriceInCentsNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["PriceInCents"] = true
	p.PriceInCents = nil
}
func (p Package_) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Id"] && p.Id == nil {
		data["id"] = nil
	} else if p.Id != nil {
		data["id"] = p.Id
	}

	if p.touched["DataLimitInBytes"] && p.DataLimitInBytes == nil {
		data["dataLimitInBytes"] = nil
	} else if p.DataLimitInBytes != nil {
		data["dataLimitInBytes"] = p.DataLimitInBytes
	}

	if p.touched["Destination"] && p.Destination == nil {
		data["destination"] = nil
	} else if p.Destination != nil {
		data["destination"] = p.Destination
	}

	if p.touched["DestinationName"] && p.DestinationName == nil {
		data["destinationName"] = nil
	} else if p.DestinationName != nil {
		data["destinationName"] = p.DestinationName
	}

	if p.touched["PriceInCents"] && p.PriceInCents == nil {
		data["priceInCents"] = nil
	} else if p.PriceInCents != nil {
		data["priceInCents"] = p.PriceInCents
	}

	return json.Marshal(data)
}

type PurchasesEsim struct {
	// ID of the eSIM
	Iccid   *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
	touched map[string]bool
}

func (p *PurchasesEsim) GetIccid() *string {
	if p == nil {
		return nil
	}
	return p.Iccid
}

func (p *PurchasesEsim) SetIccid(iccid string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Iccid"] = true
	p.Iccid = &iccid
}

func (p *PurchasesEsim) SetIccidNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Iccid"] = true
	p.Iccid = nil
}
func (p PurchasesEsim) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Iccid"] && p.Iccid == nil {
		data["iccid"] = nil
	} else if p.Iccid != nil {
		data["iccid"] = p.Iccid
	}

	return json.Marshal(data)
}
