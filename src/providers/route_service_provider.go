package providers

import (
	"gometer/modules/core/contracts"
	httpContracts "gometer/modules/http/contracts"
	viewContracts "gometer/modules/view/contracts"
	"gometer/src/middleware"
	"net/http"
)

// RouteServiceProvider ...
type RouteServiceProvider struct {
}

// GetRouteServiceProvider ...
func GetRouteServiceProvider() contracts.ServiceProvider {

	return &RouteServiceProvider{}
}

// Register ...
func (p *RouteServiceProvider) Register(a contracts.Application) {

	routeInst, _ := a.Get("router")
	router := routeInst.(httpContracts.Router)

	viewInst, _ := a.Get("view")
	view := viewInst.(viewContracts.View)

	router.AddMiddleware(middleware.StartSession)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		view.Render(w, "welcome", nil)
	})

	router.Get("/test/:id/test", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("TEST!"))
	})
}

// Boot ...
func (p *RouteServiceProvider) Boot() {

}
