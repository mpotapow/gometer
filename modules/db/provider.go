package db

import (
	"database/sql"
	"gometer/modules/core/contracts"
)

// Provider ...
type Provider struct {
	config *Config
}

// GetProvider ...
func GetProvider() contracts.ServiceProvider {

	return &Provider{}
}

// Register ...
func (p *Provider) Register(a contracts.Application) {

	p.config = GetConfig(a.GetConfigPath())

	dbPath := a.GetStoragePath() + "/database/" + p.config.DatabaseName
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	a.Set("db", db)
}

// Boot ...
func (p *Provider) Boot() {

}
