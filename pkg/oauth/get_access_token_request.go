package oauth

import (
	"encoding/json"
)

type GetAccessTokenRequest struct {
	GrantType    *GrantType `json:"grant_type,omitempty"`
	ClientId     *string    `json:"client_id,omitempty"`
	ClientSecret *string    `json:"client_secret,omitempty"`
	touched      map[string]bool
}

func (g *GetAccessTokenRequest) GetGrantType() *GrantType {
	if g == nil {
		return nil
	}
	return g.GrantType
}

func (g *GetAccessTokenRequest) SetGrantType(grantType GrantType) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["GrantType"] = true
	g.GrantType = &grantType
}

func (g *GetAccessTokenRequest) SetGrantTypeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["GrantType"] = true
	g.GrantType = nil
}

func (g *GetAccessTokenRequest) GetClientId() *string {
	if g == nil {
		return nil
	}
	return g.ClientId
}

func (g *GetAccessTokenRequest) SetClientId(clientId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ClientId"] = true
	g.ClientId = &clientId
}

func (g *GetAccessTokenRequest) SetClientIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ClientId"] = true
	g.ClientId = nil
}

func (g *GetAccessTokenRequest) GetClientSecret() *string {
	if g == nil {
		return nil
	}
	return g.ClientSecret
}

func (g *GetAccessTokenRequest) SetClientSecret(clientSecret string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ClientSecret"] = true
	g.ClientSecret = &clientSecret
}

func (g *GetAccessTokenRequest) SetClientSecretNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ClientSecret"] = true
	g.ClientSecret = nil
}

func (g GetAccessTokenRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["GrantType"] && g.GrantType == nil {
		data["grant_type"] = nil
	} else if g.GrantType != nil {
		data["grant_type"] = g.GrantType
	}

	if g.touched["ClientId"] && g.ClientId == nil {
		data["client_id"] = nil
	} else if g.ClientId != nil {
		data["client_id"] = g.ClientId
	}

	if g.touched["ClientSecret"] && g.ClientSecret == nil {
		data["client_secret"] = nil
	} else if g.ClientSecret != nil {
		data["client_secret"] = g.ClientSecret
	}

	return json.Marshal(data)
}

func (g GetAccessTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAccessTokenRequest to string"
	}
	return string(jsonData)
}

type GrantType string

const (
	GRANT_TYPE_CLIENT_CREDENTIALS GrantType = "client_credentials"
)
