package contracts

import "net/http"

// Request ...
type Request interface {
	Validate() error
	ParseJson(req *http.Request, container Request) error
}
