package contracts

// Option ...
type Option interface {
	GetName() string
	IsFillable() bool
	GetValue() interface{}
	GetDescription() string
	SetValue(val interface{})
}

// Argument ...
type Argument interface {
	GetName() string
	GetValue() string
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
	GetOption(name string) (interface{}, error)
	GetArgument(name string) (interface{}, error)
	Handle(f Formatter)
}
