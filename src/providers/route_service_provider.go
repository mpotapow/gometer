package providers

import (
	"gometer/modules/core/contracts"
	httpModule "gometer/modules/http"
	viewContracts "gometer/modules/view/contracts"
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
	router := routeInst.(*httpModule.Router)

	viewInst, _ := a.Get("view")
	view := viewInst.(viewContracts.View)

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
