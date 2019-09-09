package requests

// LoginRequest ...
type LoginRequest struct {
	baseRequest
	Login    string
	Password string
}
