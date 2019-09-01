package contracts

// Application ...
type Application interface {
	Storage
	GetRootPath() string
	GetConfigPath() string
	GetStoragePath() string
	GetResourcesPath() string
	AddProvider(p ServiceProvider)
	Register()
	Boot()
}
