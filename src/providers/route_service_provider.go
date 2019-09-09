package providers

import (
	"gometer/modules/core/contracts"
	httpContracts "gometer/modules/http/contracts"
	viewContracts "gometer/modules/view/contracts"
	"gometer/src/controllers"
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

	router.Group("/api/v1", func(r httpContracts.Router) {

		r.AddMiddleware(middleware.Authenticate)
	})

	router.Group("/api/v1", func(r httpContracts.Router) {

		r.Post("/login", controllers.AuthController)
	})

	router.Get("/login", func(w httpContracts.ResponseWriter, r *http.Request) {
		view.Render(w, "main", nil)
	})
}

// Boot ...
func (p *RouteServiceProvider) Boot() {

}
