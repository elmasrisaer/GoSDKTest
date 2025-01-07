package configmanager

import (
	"github.com/elmasrisaer/GoSDKTest/internal/oauthtokenmanager"
	"github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
	"time"
)

type ConfigManager struct {
	OAuth             celitechconfig.Config
	Destinations      celitechconfig.Config
	Packages          celitechconfig.Config
	Purchases         celitechconfig.Config
	ESim              celitechconfig.Config
	oAuthTokenManager *oauthtokenmanager.OAuthTokenManager
}

func NewConfigManager(config celitechconfig.Config, tokenService oauthtokenmanager.TokenService) *ConfigManager {
	return &ConfigManager{
		OAuth:             config,
		Destinations:      config,
		Packages:          config,
		Purchases:         config,
		ESim:              config,
		oAuthTokenManager: oauthtokenmanager.NewOAuthTokenManager(tokenService),
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.OAuth.SetBaseUrl(baseUrl)
	c.Destinations.SetBaseUrl(baseUrl)
	c.Packages.SetBaseUrl(baseUrl)
	c.Purchases.SetBaseUrl(baseUrl)
	c.ESim.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.OAuth.SetTimeout(timeout)
	c.Destinations.SetTimeout(timeout)
	c.Packages.SetTimeout(timeout)
	c.Purchases.SetTimeout(timeout)
	c.ESim.SetTimeout(timeout)
}

func (c *ConfigManager) SetClientId(clientId string) {
	c.OAuth.SetClientId(clientId)
	c.Destinations.SetClientId(clientId)
	c.Packages.SetClientId(clientId)
	c.Purchases.SetClientId(clientId)
	c.ESim.SetClientId(clientId)
}

func (c *ConfigManager) SetClientSecret(clientSecret string) {
	c.OAuth.SetClientSecret(clientSecret)
	c.Destinations.SetClientSecret(clientSecret)
	c.Packages.SetClientSecret(clientSecret)
	c.Purchases.SetClientSecret(clientSecret)
	c.ESim.SetClientSecret(clientSecret)
}

func (c *ConfigManager) GetTokenManager() *oauthtokenmanager.OAuthTokenManager {
	return c.oAuthTokenManager
}

func (c *ConfigManager) GetOAuth() *celitechconfig.Config {
	return &c.OAuth
}
func (c *ConfigManager) GetDestinations() *celitechconfig.Config {
	return &c.Destinations
}
func (c *ConfigManager) GetPackages() *celitechconfig.Config {
	return &c.Packages
}
func (c *ConfigManager) GetPurchases() *celitechconfig.Config {
	return &c.Purchases
}
func (c *ConfigManager) GetESim() *celitechconfig.Config {
	return &c.ESim
}
