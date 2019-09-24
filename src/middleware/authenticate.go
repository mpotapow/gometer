package middleware

import (
	gometerHttp "gometer/modules/http"
	"gometer/modules/http/contracts"
	"gometer/src/tools"
	"net/http"
)

// Authenticate ...
func Authenticate(next contracts.Handler) contracts.Handler {
	return gometerHttp.HandlerFunc(func(w contracts.ResponseWriter, r *http.Request) {

		helper := tools.GetSessionInstance(r)
		session := helper.GetStorage()

		if val := session.Get("user_id"); val == nil {
			http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
			return
		}

		next.ServeHTTP(w, r)
	})
}
