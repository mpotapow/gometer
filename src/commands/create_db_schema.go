package commands

import (
	"database/sql"
	"gometer/modules/console"
	"gometer/modules/console/contracts"
	"gometer/modules/core"
)

// CreateDBSchema ...
type CreateDBSchema struct {
	*console.Command
}

// GetCreateDBSchemaInstance ...
func GetCreateDBSchemaInstance() *CreateDBSchema {

	command := &CreateDBSchema{
		Command: console.GetCommandInstance(
			"db:install",
			"Create database schema",
		),
	}

	return command
}

// Handle ...
func (c *CreateDBSchema) Handle(f contracts.Formatter) {

	app := core.GetApplicationInstance()

	dbInst, _ := app.Get("db")
	connection := dbInst.(*sql.DB)

	if ok := c.createUsersTable(connection); ok {
		f.Writeln("<green>Users table created</>")
	} else {
		f.Writeln("<red>Failed to create users table</>")
	}

	if ok := c.createTestsTable(connection); ok {
		f.Writeln("<green>Tests table created</>")
	} else {
		f.Writeln("<red>Failed to create tests table</>")
	}

	if ok := c.createProjectsTable(connection); ok {
		f.Writeln("<green>Projects table created</>")
	} else {
		f.Writeln("<red>Failed to create projects table</>")
	}
}

func (c *CreateDBSchema) createUsersTable(con *sql.DB) bool {
	_, err := con.Exec(`
		CREATE TABLE users(
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			login TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		return false
	}

	return true
}

func (c *CreateDBSchema) createTestsTable(con *sql.DB) bool {
	_, err := con.Exec(`
		CREATE TABLE tests(
			id INTEGER PRIMARY KEY,
			project_id INTEGER NOT NULL,
			started_at INTEGER NOT NULL,
			duration INTEGER NOT NULL,
			transactions_cnt INTEGER NOT NULL,
			min_response_time INTEGER NOT NULL,
			max_response_time INTEGER NOT NULL,
			avg_response_time INTEGER NOT NULL,
			max_virtual_users INTEGER NOT NULL
		);
	`)
	if err != nil {
		return false
	}

	return true
}

func (c *CreateDBSchema) createProjectsTable(con *sql.DB) bool {
	_, err := con.Exec(`
		CREATE TABLE projects(
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return false
	}

	return true
}
