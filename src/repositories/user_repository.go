package repositories

import (
	"database/sql"
	"gometer/src/contracts"
	"gometer/src/models"
)

// UserRepository ...
type UserRepository struct {
	connection *sql.DB
}

// NewUserRepository ...
func NewUserRepository(conn *sql.DB) contracts.UserRepository {
	return &UserRepository{
		connection: conn,
	}
}

// FindByLogin ...
func (u *UserRepository) FindByLogin(login string) (*models.User, error) {

	row := u.connection.QueryRow("Select id, login, password from users where login=?", login)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Login, &user.Password)

	return user, err
}
