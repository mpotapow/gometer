package contracts

// Response ...
type Response interface {
	GetStatus() int
	SetStatus(status int) Response

	HasError() bool
	SetHasError(flag bool) Response
}
