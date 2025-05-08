package purchases

import "encoding/json"

type CreatePurchaseV2Request struct {
	// ISO representation of the package's destination.
	Destination *string `json:"destination,omitempty" required:"true"`
	// Size of the package in GB.
	//
	// For **limited packages**, the available options are: **0.5, 1, 2, 3, 5, 8, 20GB** (supports `duration` or `startDate` / `endDate`).
	//
	// For **unlimited packages** (available to Region-3), please use **-1** as an identifier (supports `duration` only).
	//
	DataLimitInGb *float64 `json:"dataLimitInGB,omitempty" required:"true"`
	// Number of eSIMs to purchase.
	Quantity *float64 `json:"quantity,omitempty" required:"true" min:"1" max:"5"`
	// Email address where the purchase confirmation email will be sent (including QR Code & activation steps)
	Email *string `json:"email,omitempty"`
	// An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters and numbers. This feature is available to platforms with Diamond tier only.
	NetworkBrand *string `json:"networkBrand,omitempty"`
	// Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.
	EmailBrand *string `json:"emailBrand,omitempty"`
}

func (c *CreatePurchaseV2Request) GetDestination() *string {
	if c == nil {
		return nil
	}
	return c.Destination
}

func (c *CreatePurchaseV2Request) SetDestination(destination string) {
	c.Destination = &destination
}

func (c *CreatePurchaseV2Request) GetDataLimitInGb() *float64 {
	if c == nil {
		return nil
	}
	return c.DataLimitInGb
}

func (c *CreatePurchaseV2Request) SetDataLimitInGb(dataLimitInGb float64) {
	c.DataLimitInGb = &dataLimitInGb
}

func (c *CreatePurchaseV2Request) GetQuantity() *float64 {
	if c == nil {
		return nil
	}
	return c.Quantity
}

func (c *CreatePurchaseV2Request) SetQuantity(quantity float64) {
	c.Quantity = &quantity
}

func (c *CreatePurchaseV2Request) GetEmail() *string {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreatePurchaseV2Request) SetEmail(email string) {
	c.Email = &email
}

func (c *CreatePurchaseV2Request) GetReferenceId() *string {
	if c == nil {
		return nil
	}
	return c.ReferenceId
}

func (c *CreatePurchaseV2Request) SetReferenceId(referenceId string) {
	c.ReferenceId = &referenceId
}

func (c *CreatePurchaseV2Request) GetNetworkBrand() *string {
	if c == nil {
		return nil
	}
	return c.NetworkBrand
}

func (c *CreatePurchaseV2Request) SetNetworkBrand(networkBrand string) {
	c.NetworkBrand = &networkBrand
}

func (c *CreatePurchaseV2Request) GetEmailBrand() *string {
	if c == nil {
		return nil
	}
	return c.EmailBrand
}

func (c *CreatePurchaseV2Request) SetEmailBrand(emailBrand string) {
	c.EmailBrand = &emailBrand
}

func (c CreatePurchaseV2Request) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePurchaseV2Request to string"
	}
	return string(jsonData)
}
