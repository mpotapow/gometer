package contracts

import "net/http"

// Request ...
type Request interface {
	ParseJson(req *http.Request, container Request) error
}
