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
	storage contracts.UserRepository
}

// ErrPasswordHash ...
var ErrPasswordHash = errors.New("auth: wrong password")

// NewAuthService ...
func NewAuthService() contracts.AuthService {

	userRepository, _ := core.GetApplicationInstance().Get("user-repository")

	return &AuthService{
		storage: userRepository.(contracts.UserRepository),
	}
}

// Authorize ...
func (a *AuthService) Authorize(login string, password string) (*models.User, error) {

	user, err := a.storage.FindByLogin(login)
	if err != nil {
		return user, err
	}

	hasher := tools.GetHashInstance()
	if hasher.Check(password, user.Password) {
		return user, nil
	}

	return nil, ErrPasswordHash
}
