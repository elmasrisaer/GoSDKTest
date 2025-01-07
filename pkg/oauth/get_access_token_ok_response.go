package oauth

import (
	"encoding/json"
)

type GetAccessTokenOkResponse struct {
	AccessToken *string `json:"access_token,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty"`
	touched     map[string]bool
}

func (g *GetAccessTokenOkResponse) GetAccessToken() *string {
	if g == nil {
		return nil
	}
	return g.AccessToken
}

func (g *GetAccessTokenOkResponse) SetAccessToken(accessToken string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AccessToken"] = true
	g.AccessToken = &accessToken
}

func (g *GetAccessTokenOkResponse) SetAccessTokenNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AccessToken"] = true
	g.AccessToken = nil
}

func (g *GetAccessTokenOkResponse) GetTokenType() *string {
	if g == nil {
		return nil
	}
	return g.TokenType
}

func (g *GetAccessTokenOkResponse) SetTokenType(tokenType string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TokenType"] = true
	g.TokenType = &tokenType
}

func (g *GetAccessTokenOkResponse) SetTokenTypeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TokenType"] = true
	g.TokenType = nil
}

func (g *GetAccessTokenOkResponse) GetExpiresIn() *int64 {
	if g == nil {
		return nil
	}
	return g.ExpiresIn
}

func (g *GetAccessTokenOkResponse) SetExpiresIn(expiresIn int64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresIn"] = true
	g.ExpiresIn = &expiresIn
}

func (g *GetAccessTokenOkResponse) SetExpiresInNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresIn"] = true
	g.ExpiresIn = nil
}
func (g GetAccessTokenOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["AccessToken"] && g.AccessToken == nil {
		data["access_token"] = nil
	} else if g.AccessToken != nil {
		data["access_token"] = g.AccessToken
	}

	if g.touched["TokenType"] && g.TokenType == nil {
		data["token_type"] = nil
	} else if g.TokenType != nil {
		data["token_type"] = g.TokenType
	}

	if g.touched["ExpiresIn"] && g.ExpiresIn == nil {
		data["expires_in"] = nil
	} else if g.ExpiresIn != nil {
		data["expires_in"] = g.ExpiresIn
	}

	return json.Marshal(data)
}
