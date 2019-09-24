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
	lastSlash     *regexp.Regexp
	namedParam    *regexp.Regexp
	splatParam    *regexp.Regexp
	escapeRegExp  *regexp.Regexp
	optionalParam *regexp.Regexp

	namespace  string
	balancer   *balancer
	middleware []func(contracts.Handler) contracts.Handler
}

type balancer struct {
	partition map[string]map[string]contracts.Handler
}

// HandlerFunc ...
type HandlerFunc func(contracts.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w contracts.ResponseWriter, r *http.Request) {
	f(w, r)
}

// GetRouterInstance ...
func GetRouterInstance() contracts.Router {

	lastSlash, _ := regexp.Compile(`\/$`)
	splatParam, _ := regexp.Compile(`\*\w+`)
	namedParam, _ := regexp.Compile(`(\(\?)?:\w+`)
	optionalParam, _ := regexp.Compile(`\((.*?)\)`)
	escapeRegExp, _ := regexp.Compile(`[\-{}\[\]+?.,\\\^$|#\s]`)

	r := &Router{
		lastSlash:     lastSlash,
		namedParam:    namedParam,
		splatParam:    splatParam,
		escapeRegExp:  escapeRegExp,
		optionalParam: optionalParam,

		middleware: []func(contracts.Handler) contracts.Handler{},

		balancer: &balancer{
			partition: map[string]map[string]contracts.Handler{
				"GET":    map[string]contracts.Handler{},
				"PUT":    map[string]contracts.Handler{},
				"POST":   map[string]contracts.Handler{},
				"PATCH":  map[string]contracts.Handler{},
				"DELETE": map[string]contracts.Handler{},
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
func (r *Router) AddMiddleware(middleware func(contracts.Handler) contracts.Handler) {

	r.middleware = append(r.middleware, middleware)
}

// Get ...
func (r *Router) Get(url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	r.append("GET", url, f)
}

// Post ...
func (r *Router) Post(url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	r.append("POST", url, f)
}

// Patch ...
func (r *Router) Patch(url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	r.append("PATCH", url, f)
}

// Put ...
func (r *Router) Put(url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	r.append("PUT", url, f)
}

// Delete ...
func (r *Router) Delete(url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	r.append("DELETE", url, f)
}

func (r *Router) append(method string, url string, f func(w contracts.ResponseWriter, r *http.Request)) {

	url = r.routeToRegExp(r.applyNamespace(url))
	handler := r.applyMiddleware(HandlerFunc(f))

	r.balancer.append(method, url, handler)
}

func (r *Router) routeToRegExp(route string) string {

	replRoute := r.escapeRegExp.ReplaceAll([]byte(route), []byte("\\$&"))
	replRoute = r.optionalParam.ReplaceAll(replRoute, []byte("(?:$1)?"))
	replRoute = r.namedParam.ReplaceAllFunc(replRoute, func(data []byte) []byte {
		return []byte(`([^\/]+)`)
	})
	replRoute = r.lastSlash.ReplaceAll(replRoute, []byte("[/]?"))
	replRoute = r.splatParam.ReplaceAll(replRoute, []byte("(.*?)"))

	return fmt.Sprintf("^%s$", string(replRoute))
}

func (r *Router) applyMiddleware(next contracts.Handler) contracts.Handler {

	for i := len(r.middleware) - 1; i >= 0; i-- {
		next = r.middleware[i](next)
	}

	return next
}

func (r *Router) applyNamespace(path string) string {

	path = strings.Trim(path, "/")
	namespace := strings.Trim(r.namespace, "/")

	url := "/" + namespace + "/" + path
	url = strings.Trim(url, "/")

	if len(url) <= 0 {
		return "/"
	}

	return "/" + url + "/"
}

func (b *balancer) append(method string, url string, h contracts.Handler) {

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
			http.Error(w, fmt.Sprintf("Method: %s not allow", req.Method), http.StatusMethodNotAllowed)
			return
		}

		for route, handler := range routes {
			routeReg, _ := regexp.Compile(route)
			if routeReg.MatchString(req.URL.Path) {
				params := routeReg.FindStringSubmatch(req.URL.Path)[1:]
				fmt.Println("PARAMS", params)
				handler.ServeHTTP(NewResponse(w), req)
				return
			}
		}

		http.NotFound(w, req)
	})
}
