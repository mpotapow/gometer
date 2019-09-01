package main

import (
	"gometer/modules/console"
	"gometer/modules/core"
	"gometer/modules/core/contracts"
	"gometer/modules/db"
	"gometer/src/providers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	app := core.GetApplicationInstance()

	registerModulesProvider(app)
	registerApplicationProvider(app)

	app.Register()
	app.Boot()
}

func registerModulesProvider(app contracts.Application) {

	app.AddProvider(db.GetProvider())
	app.AddProvider(console.GetProvider())
}

func registerApplicationProvider(app contracts.Application) {

	app.AddProvider(providers.GetConsoleServiceProvider())
}
