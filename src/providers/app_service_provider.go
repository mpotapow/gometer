package providers

import (
	"database/sql"
	"gometer/modules/core/contracts"
	"gometer/src/repositories"
	"gometer/src/services"
)

// AppServiceProvider ...
type AppServiceProvider struct {
}

// GetAppServiceProvider ...
func GetAppServiceProvider() contracts.ServiceProvider {

	return &AppServiceProvider{}
}

// Register ...
func (p *AppServiceProvider) Register(a contracts.Application) {

	dbInst, _ := a.Get("db")
	connection := dbInst.(*sql.DB)

	a.Set("auth-service", services.NewAuthService())
	a.Set("user-repository", repositories.NewUserRepository(connection))
}

// Boot ...
func (p *AppServiceProvider) Boot() {

}
