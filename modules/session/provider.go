package session

import (
	"gometer/modules/core/contracts"
)

// Provider ...
type Provider struct {
	config *Config
}

// GetProvider ...
func GetProvider() contracts.ServiceProvider {

	return &Provider{}
}

// Register ...
func (p *Provider) Register(a contracts.Application) {

	a.Set("session", GetManagerInstance(GetConfig(a.GetConfigPath())))
}

// Boot ...
func (p *Provider) Boot() {

}
