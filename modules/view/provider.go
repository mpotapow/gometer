package view

import (
	"gometer/modules/core/contracts"
)

// Provider ...
type Provider struct {
}

// GetProvider ...
func GetProvider() contracts.ServiceProvider {

	return &Provider{}
}

// Register ...
func (p *Provider) Register(a contracts.Application) {

	a.Set("view", GetViewInstance(a.GetResourcesPath()+"/view"))
}

// Boot ...
func (p *Provider) Boot() {

}
