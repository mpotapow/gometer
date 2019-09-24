package requests

import "errors"

// LoginRequest ...
type LoginRequest struct {
	baseRequest
	Login    string
	Password string
}

// Validate ...
func (r *LoginRequest) Validate() error {
	if len(r.Login) <= 0 {
		return errors.New("Login field is required")
	}
	if len(r.Password) <= 0 {
		return errors.New("Password field is required")
	}
	return nil
}
