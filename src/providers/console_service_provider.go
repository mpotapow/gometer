package providers

import (
	"gometer/modules/console"
	cliContracts "gometer/modules/console/contracts"
	"gometer/modules/core/contracts"
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

	cli.Add(console.GetCommandInstance("list", "Lists commands"))

	cli.Add(console.GetCommandInstance("app:load", "Load JMeter log"))

	createUserCommand := console.GetCommandInstance("app:user-create", "Create user")
	createUserCommand.AddArgument(&console.Argument{Name: "test", Description: "For test"})
	createUserCommand.AddOption(&console.Option{Name: "login", Description: "User login", Fillable: true})
	createUserCommand.AddOption(&console.Option{Name: "password", Description: "User password", Fillable: true})
	createUserCommand.AddOption(&console.Option{Name: "flush", Description: "Flush all"})
	cli.Add(createUserCommand)

	cli.Add(console.GetCommandInstance("db:install", "Create database schema"))
}

// Boot ...
func (p *ConsoleServiceProvider) Boot() {

}
