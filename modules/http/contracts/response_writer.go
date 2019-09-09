package contracts

import "net/http"

// Handler ...
type Handler interface {
	ServeHTTP(ResponseWriter, *http.Request)
}

// ResponseWriter ...
type ResponseWriter interface {
	http.ResponseWriter
	ToJson(v Response)
}
