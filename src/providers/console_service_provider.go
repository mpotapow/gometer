package providers

import (
	cliContracts "gometer/modules/console/contracts"
	"gometer/modules/core/contracts"
	"gometer/src/commands"
)

// ConsoleServiceProvider ...
type ConsoleServiceProvider struct {
}

// GetConsoleServiceProvider ...
func GetConsoleServiceProvider() contracts.ServiceProvider {

	return &ConsoleServiceProvider{}
}

// Register ...
func (p *ConsoleServiceProvider) Register(a contracts.Application) {

	cliInst, _ := a.Get("console")
	cli := cliInst.(cliContracts.Application)

	cli.Add(commands.GetLoadJMeterInstance())
	cli.Add(commands.GetCreateUserInstance())
	cli.Add(commands.GetCreateDBSchemaInstance())
}

// Boot ...
func (p *ConsoleServiceProvider) Boot() {

}
