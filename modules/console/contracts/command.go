package contracts

// Option ...
type Option interface {
	GetName() string
	IsFillable() bool
	GetDescription() string
	SetValue(val interface{})
}

// Argument ...
type Argument interface {
	GetName() string
	SetValue(val string)
	GetDescription() string
}

// Command ...
type Command interface {
	GetName() string
	GetDescription() string
	AddOption(option Option)
	AddArgument(argument Argument)
	GetArguments() []Argument
	GetOptions() map[string]Option
	Handle()
}
