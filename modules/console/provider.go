package console

import (
	"gometer/modules/console/contracts"
	coreContracts "gometer/modules/core/contracts"
	"os"
)

// Provider ...
type Provider struct {
	cli contracts.Application
}

// GetProvider ...
func GetProvider() coreContracts.ServiceProvider {

	return &Provider{}
}

// Register ...
func (p *Provider) Register(a coreContracts.Application) {

	cli := GetApplicationInstance()
	cli.SetName("GoMeter")
	cli.SetVersion("0.0.1")

	p.cli = cli
	a.Set("console", cli)
}

// Boot ...
func (p *Provider) Boot() {

	p.cli.Run(os.Args)
}
