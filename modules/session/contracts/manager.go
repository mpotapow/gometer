package contracts

// Manager ...
type Manager interface {
	GetDriver() Session
	GetPath() string
	GetDomain() string
	IsSecure() bool
	IsHttpOnly() bool
}
