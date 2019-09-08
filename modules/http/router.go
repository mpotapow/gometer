package http

import (
	"fmt"
	"gometer/modules/http/contracts"
	"net/http"
	"regexp"
	"strings"
)

// Router ...
type Router struct {
	namedParam    *regexp.Regexp
	splatParam    *regexp.Regexp
	escapeRegExp  *regexp.Regexp
	optionalParam *regexp.Regexp

	namespace  string
	balancer   *balancer
	middleware []func(http.Handler) http.Handler
}

type balancer struct {
	partition map[string]map[string]http.Handler
}

// GetRouterInstance ...
func GetRouterInstance() contracts.Router {

	splatParam, _ := regexp.Compile(`\*\w+`)
	namedParam, _ := regexp.Compile(`(\(\?)?:\w+`)
	optionalParam, _ := regexp.Compile(`\((.*?)\)`)
	escapeRegExp, _ := regexp.Compile(`[\-{}\[\]+?.,\\\^$|#\s]`)

	r := &Router{
		namedParam:    namedParam,
		splatParam:    splatParam,
		escapeRegExp:  escapeRegExp,
		optionalParam: optionalParam,

		middleware: []func(http.Handler) http.Handler{},

		balancer: &balancer{
			partition: map[string]map[string]http.Handler{
				"GET":    map[string]http.Handler{},
				"PUT":    map[string]http.Handler{},
				"POST":   map[string]http.Handler{},
				"PATCH":  map[string]http.Handler{},
				"DELETE": map[string]http.Handler{},
			},
		},
	}

	r.balancer.registerStaticRoute()
	r.balancer.registerInitialRoute()

	return r
}

func (r *Router) clone() *Router {
	clone := *r
	return &clone
}

// Group ...
func (r *Router) Group(namespace string, callback func(r contracts.Router)) {

	router := r.clone()
	router.namespace = router.applyNamespace(namespace)

	callback(router)
}

// AddMiddleware ...
func (r *Router) AddMiddleware(middleware func(http.Handler) http.Handler) {

	r.middleware = append(r.middleware, middleware)
}

// Get ...
func (r *Router) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {

	r.append("GET", url, f)
}

// Post ...
func (r *Router) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {

	r.append("POST", url, f)
}

// Patch ...
func (r *Router) Patch(url string, f func(w http.ResponseWriter, r *http.Request)) {

	r.append("PATCH", url, f)
}

// Put ...
func (r *Router) Put(url string, f func(w http.ResponseWriter, r *http.Request)) {

	r.append("PUT", url, f)
}

// Delete ...
func (r *Router) Delete(url string, f func(w http.ResponseWriter, r *http.Request)) {

	r.append("DELETE", url, f)
}

func (r *Router) append(method string, url string, f func(w http.ResponseWriter, r *http.Request)) {

	url = r.routeToRegExp(r.applyNamespace(url))
	handler := r.applyMiddleware(http.HandlerFunc(f))

	r.balancer.append(method, url, handler)
}

func (r *Router) routeToRegExp(route string) string {

	replRoute := r.escapeRegExp.ReplaceAll([]byte(route), []byte("\\$&"))
	replRoute = r.optionalParam.ReplaceAll(replRoute, []byte("(?:$1)?"))
	replRoute = r.namedParam.ReplaceAllFunc(replRoute, func(data []byte) []byte {
		return []byte(`([^\/]+)`)
	})
	replRoute = r.splatParam.ReplaceAll(replRoute, []byte("(.*?)"))

	return fmt.Sprintf("^%s$", string(replRoute))
}

func (r *Router) applyMiddleware(next http.Handler) http.Handler {

	for i := len(r.middleware) - 1; i >= 0; i-- {
		next = r.middleware[i](next)
	}

	return next
}

func (r *Router) applyNamespace(path string) string {

	path = strings.Trim(path, "/")
	namespace := strings.Trim(r.namespace, "/")

	if len(namespace) > 0 {
		namespace = "/" + namespace + "/"
	}

	return namespace + path + "/"
}

func (b *balancer) append(method string, url string, h http.Handler) {

	b.partition[method][url] = h
}

func (b *balancer) registerStaticRoute() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
}

func (b *balancer) registerInitialRoute() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		routes, ok := b.partition[req.Method]
		if !ok {
			w.WriteHeader(405)
			w.Write([]byte(fmt.Sprintf("Method: %s not allow", req.Method)))
			return
		}

		for route, handler := range routes {
			routeReg, _ := regexp.Compile(route)
			if routeReg.MatchString(req.RequestURI) {
				params := routeReg.FindStringSubmatch(req.RequestURI)[1:]
				fmt.Println("PARAMS", params)
				handler.ServeHTTP(w, req)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found \n"))
		w.Write([]byte(fmt.Sprintf("Method: %s\nUri: %s", req.Method, req.RequestURI)))
	})
}
