package contracts

import "net/http"

// Router ...
type Router interface {
	Group(namespace string, callback func(r Router))
	AddMiddleware(middleware func(http.Handler) http.Handler)
	Get(url string, f func(w http.ResponseWriter, r *http.Request))
	Put(url string, f func(w http.ResponseWriter, r *http.Request))
	Post(url string, f func(w http.ResponseWriter, r *http.Request))
	Patch(url string, f func(w http.ResponseWriter, r *http.Request))
	Delete(url string, f func(w http.ResponseWriter, r *http.Request))
}
