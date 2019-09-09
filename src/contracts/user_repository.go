package contracts

import "gometer/src/models"

// UserRepository ...
type UserRepository interface {
	FindByLogin(login string) (*models.User, error)
}
