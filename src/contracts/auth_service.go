package contracts

import "gometer/src/models"

// AuthService ...
type AuthService interface {
	Authorize(login string, password string) (*models.User, error)
}
