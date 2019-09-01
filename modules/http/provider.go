package http

import (
	"fmt"
	"gometer/modules/core/contracts"
	"net/http"
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

	p.config = GetConfig(a.GetConfigPath())

	a.Set("router", GetRouterInstance())
}

// Boot ...
func (p *Provider) Boot() {

	fmt.Println(fmt.Sprintf("Start http server %s:%d", p.config.Host, p.config.Port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", p.config.Host, p.config.Port), nil); err != nil {
		panic(err)
	}
}
