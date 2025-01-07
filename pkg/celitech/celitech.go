package celitech

import (
	"github.com/elmasrisaer/GoSDKTest/internal/configmanager"
	"github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
	"github.com/elmasrisaer/GoSDKTest/pkg/destinations"
	"github.com/elmasrisaer/GoSDKTest/pkg/esim"
	"github.com/elmasrisaer/GoSDKTest/pkg/oauth"
	"github.com/elmasrisaer/GoSDKTest/pkg/packages"
	"github.com/elmasrisaer/GoSDKTest/pkg/purchases"
	"time"
)

type Celitech struct {
	OAuth        *oauth.OAuthService
	Destinations *destinations.DestinationsService
	Packages     *packages.PackagesService
	Purchases    *purchases.PurchasesService
	ESim         *esim.ESimService
	manager      *configmanager.ConfigManager
}

func NewCelitech(config celitechconfig.Config) *Celitech {
	oAuth := oauth.NewOAuthService()
	destinations := destinations.NewDestinationsService()
	packages := packages.NewPackagesService()
	purchases := purchases.NewPurchasesService()
	eSim := esim.NewESimService()
	manager := configmanager.NewConfigManager(config, oAuth)
	oAuth.WithConfigManager(manager)
	destinations.WithConfigManager(manager)
	packages.WithConfigManager(manager)
	purchases.WithConfigManager(manager)
	eSim.WithConfigManager(manager)
	return &Celitech{
		OAuth:        oAuth,
		Destinations: destinations,
		Packages:     packages,
		Purchases:    purchases,
		ESim:         eSim,
		manager:      manager,
	}
}

func (c *Celitech) SetBaseUrl(baseUrl string) {
	c.manager.SetBaseUrl(baseUrl)
}

func (c *Celitech) SetTimeout(timeout time.Duration) {
	c.manager.SetTimeout(timeout)
}

func (c *Celitech) SetClientId(clientId string) {
	c.manager.SetClientId(clientId)
}

func (c *Celitech) SetClientSecret(clientSecret string) {
	c.manager.SetClientSecret(clientSecret)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
