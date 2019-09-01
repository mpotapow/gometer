package core

import (
	"gometer/modules/core/contracts"
	"os"
)

// Application ...
type Application struct {
	Storage
	rootPath  string
	providers []contracts.ServiceProvider
}

// GetApplicationInstance ...
func GetApplicationInstance() contracts.Application {

	dir, _ := os.Getwd()

	return &Application{
		Storage: Storage{
			values: map[string]interface{}{},
		},
		rootPath:  dir,
		providers: []contracts.ServiceProvider{},
	}
}

// GetRootPath ...
func (a *Application) GetRootPath() string {

	return a.rootPath
}

// GetConfigPath ...
func (a *Application) GetConfigPath() string {

	return a.GetRootPath() + "/config"
}

// GetStoragePath ...
func (a *Application) GetStoragePath() string {

	return a.GetRootPath() + "/storage"
}

// GetResourcesPath ...
func (a *Application) GetResourcesPath() string {

	return a.GetRootPath() + "/resources"
}

// AddProvider ...
func (a *Application) AddProvider(p contracts.ServiceProvider) {

	a.providers = append(a.providers, p)
}

// Register ...
func (a *Application) Register() {

	for _, p := range a.providers {
		p.Register(a)
	}
}

// Boot ...
func (a *Application) Boot() {

	for _, p := range a.providers {
		p.Boot()
	}
}
