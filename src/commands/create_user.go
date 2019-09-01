package commands

import (
	"database/sql"
	"fmt"
	"gometer/modules/console"
	"gometer/modules/console/contracts"
	"gometer/modules/core"
	"gometer/src/tools"
)

// CreateUser ...
type CreateUser struct {
	*console.Command
}

// GetCreateUserInstance ...
func GetCreateUserInstance() *CreateUser {

	command := &CreateUser{
		Command: console.GetCommandInstance(
			"app:user-create",
			"Create new user",
		),
	}

	command.AddOption(&console.Option{Name: "name", Description: "User name", Fillable: true})
	command.AddOption(&console.Option{Name: "login", Description: "User login", Fillable: true})
	command.AddOption(&console.Option{Name: "password", Description: "User password", Fillable: true})

	return command
}

// Handle ...
func (c *CreateUser) Handle(f contracts.Formatter) {

	name, _ := c.GetOption("name")
	login, _ := c.GetOption("login")
	password, _ := c.GetOption("password")

	if login == nil || len(login.(string)) <= 0 {
		f.Writeln("<red>login required</>")
		return
	}
	if password == nil || len(password.(string)) <= 0 {
		f.Writeln("<red>password required</>")
		return
	}
	if name == nil {
		name = "Unknown"
	}

	id, err := c.createUser(login.(string), password.(string), name.(string))
	if err != nil {
		f.Writeln("<red>Failed create user</>")
	} else {
		f.Writeln(fmt.Sprintf("<green>Success! user[%d] created</>", id))
	}
}

func (c *CreateUser) createUser(login string, password string, name string) (int64, error) {

	app := core.GetApplicationInstance()

	dbInst, _ := app.Get("db")
	connection := dbInst.(*sql.DB)

	stmt, err := connection.Prepare("INSERT INTO users (login, password, name) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		login,
		tools.GetHashInstance().Make(password),
		name,
	)
	if err != nil {
		panic(err)
	}

	return res.LastInsertId()
}
