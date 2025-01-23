package purchases

import (
	"encoding/json"
)

type GetPurchaseConsumptionOkResponse struct {
	// Remaining balance of the package in bytes
	DataUsageRemainingInBytes *float64 `json:"dataUsageRemainingInBytes,omitempty"`
	// Status of the connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'
	Status  *string `json:"status,omitempty"`
	touched map[string]bool
}

func (g *GetPurchaseConsumptionOkResponse) GetDataUsageRemainingInBytes() *float64 {
	if g == nil {
		return nil
	}
	return g.DataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionOkResponse) SetDataUsageRemainingInBytes(dataUsageRemainingInBytes float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DataUsageRemainingInBytes"] = true
	g.DataUsageRemainingInBytes = &dataUsageRemainingInBytes
}

func (g *GetPurchaseConsumptionOkResponse) SetDataUsageRemainingInBytesNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DataUsageRemainingInBytes"] = true
	g.DataUsageRemainingInBytes = nil
}

func (g *GetPurchaseConsumptionOkResponse) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetPurchaseConsumptionOkResponse) SetStatus(status string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = &status
}

func (g *GetPurchaseConsumptionOkResponse) SetStatusNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = nil
}

func (g GetPurchaseConsumptionOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["DataUsageRemainingInBytes"] && g.DataUsageRemainingInBytes == nil {
		data["dataUsageRemainingInBytes"] = nil
	} else if g.DataUsageRemainingInBytes != nil {
		data["dataUsageRemainingInBytes"] = g.DataUsageRemainingInBytes
	}

	if g.touched["Status"] && g.Status == nil {
		data["status"] = nil
	} else if g.Status != nil {
		data["status"] = g.Status
	}

	return json.Marshal(data)
}

func (g GetPurchaseConsumptionOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPurchaseConsumptionOkResponse to string"
	}
	return string(jsonData)
}
