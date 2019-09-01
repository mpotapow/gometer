package contracts

// Application ...
type Application interface {
	SetName(name string)
	SetVersion(version string)
	Add(command Command)
	Run(args []string)
}
