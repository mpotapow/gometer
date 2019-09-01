package console

import "gometer/modules/console/contracts"

// Option ...
type Option struct {
	Name        string
	Description string
	Fillable    bool
	value       interface{}
}

// GetName ...
func (o *Option) GetName() string {
	return o.Name
}

// GetDescription ...
func (o *Option) GetDescription() string {
	return o.Description
}

// IsFillable ...
func (o *Option) IsFillable() bool {
	return o.Fillable
}

// SetValue ...
func (o *Option) SetValue(val interface{}) {
	o.value = val
}

// Argument ...
type Argument struct {
	Name        string
	Description string
	value       string
}

// GetName ...
func (a *Argument) GetName() string {
	return a.Name
}

// GetDescription ...
func (a *Argument) GetDescription() string {
	return a.Description
}

// SetValue ...
func (a *Argument) SetValue(val string) {
	a.value = val
}

// Command ...
type Command struct {
	Name        string
	Description string
	Arguments   []contracts.Argument
	Options     map[string]contracts.Option
}

// GetCommandInstance ...
func GetCommandInstance(name string, description string) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Arguments:   []contracts.Argument{},
		Options:     make(map[string]contracts.Option),
	}
}

// GetName ...
func (c *Command) GetName() string {
	return c.Name
}

// GetDescription ...
func (c *Command) GetDescription() string {
	return c.Description
}

// AddOption ...
func (c *Command) AddOption(option contracts.Option) {
	if !option.IsFillable() {
		option.SetValue(false)
	}
	c.Options[option.GetName()] = option
}

// AddArgument ...
func (c *Command) AddArgument(argument contracts.Argument) {
	c.Arguments = append(c.Arguments, argument)
}

// GetArguments ...
func (c *Command) GetArguments() []contracts.Argument {
	return c.Arguments
}

// GetOptions ...
func (c *Command) GetOptions() map[string]contracts.Option {
	return c.Options
}

// Handle ...
func (c *Command) Handle() {
}
