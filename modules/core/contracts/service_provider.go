package contracts

// ServiceProvider ...
type ServiceProvider interface {
	Boot()
	Register(a Application)
}
