package http

import (
	"fmt"
	"net/http"
	"regexp"
)

// Router ...
type Router struct {
	namedParam    *regexp.Regexp
	splatParam    *regexp.Regexp
	escapeRegExp  *regexp.Regexp
	optionalParam *regexp.Regexp

	partition map[string]map[string]func(w http.ResponseWriter, r *http.Request)
}

// GetRouterInstance ...
func GetRouterInstance() *Router {

	splatParam, _ := regexp.Compile(`\*\w+`)
	namedParam, _ := regexp.Compile(`(\(\?)?:\w+`)
	optionalParam, _ := regexp.Compile(`\((.*?)\)`)
	escapeRegExp, _ := regexp.Compile(`[\-{}\[\]+?.,\\\^$|#\s]`)

	r := &Router{
		namedParam:    namedParam,
		splatParam:    splatParam,
		escapeRegExp:  escapeRegExp,
		optionalParam: optionalParam,

		partition: map[string]map[string]func(w http.ResponseWriter, r *http.Request){
			"GET":    map[string]func(w http.ResponseWriter, r *http.Request){},
			"PUT":    map[string]func(w http.ResponseWriter, r *http.Request){},
			"POST":   map[string]func(w http.ResponseWriter, r *http.Request){},
			"PATCH":  map[string]func(w http.ResponseWriter, r *http.Request){},
			"DELETE": map[string]func(w http.ResponseWriter, r *http.Request){},
		},
	}

	r.registerStaticRoute()
	r.registerInitialRoute()

	return r
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

	r.partition[method][r.routeToRegExp(url)] = f
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

func (r *Router) registerStaticRoute() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
}

func (r *Router) registerInitialRoute() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		routes, ok := r.partition[req.Method]
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
				handler(w, req)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found \n"))
		w.Write([]byte(fmt.Sprintf("Method: %s\nUri: %s", req.Method, req.RequestURI)))
	})
}
