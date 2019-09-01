package commands

import (
	"gometer/modules/console"
	"gometer/modules/console/contracts"
)

// LoadJMeter ...
type LoadJMeter struct {
	*console.Command
}

// GetLoadJMeterInstance ...
func GetLoadJMeterInstance() *LoadJMeter {

	command := &LoadJMeter{
		Command: console.GetCommandInstance(
			"app:load",
			"Load JMeter log",
		),
	}

	return command
}

// Handle ...
func (c *LoadJMeter) Handle(f contracts.Formatter) {

}
