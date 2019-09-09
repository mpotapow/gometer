package services

import (
	"errors"
	"gometer/modules/core"
	"gometer/src/contracts"
	"gometer/src/models"
	"gometer/src/tools"
)

// AuthService ...
type AuthService struct {
}

// ErrPasswordHash ...
var ErrPasswordHash = errors.New("auth: wrong password")

// NewAuthService ...
func NewAuthService() contracts.AuthService {
	return &AuthService{}
}

// Authorize ...
func (a *AuthService) Authorize(login string, password string) (*models.User, error) {

	userRepository, _ := core.GetApplicationInstance().Get("user-repository")
	repository := userRepository.(contracts.UserRepository)

	user, err := repository.FindByLogin(login)
	if err != nil {
		return user, err
	}

	hasher := tools.GetHashInstance()
	if hasher.Check(password, user.Password) {
		return user, nil
	}

	return nil, ErrPasswordHash
}
