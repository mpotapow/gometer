package middleware

import (
	gometerHttp "gometer/modules/http"
	httpContracts "gometer/modules/http/contracts"
	"gometer/modules/session/contracts"
	"gometer/src/tools"
	"net/http"
)

// StartSession ...
func StartSession(next httpContracts.Handler) httpContracts.Handler {
	return gometerHttp.HandlerFunc(func(w httpContracts.ResponseWriter, r *http.Request) {

		helper := tools.GetSessionInstance(r)
		session := helper.GetStorage()
		session.Start()

		addCookieResponse(helper.GetManager(), session, w)

		next.ServeHTTP(w, r)

		session.Save()
	})
}

func addCookieResponse(manager contracts.Manager, session contracts.Session, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Value:    session.GetID(),
		Name:     session.GetName(),
		Path:     manager.GetPath(),
		Secure:   manager.IsSecure(),
		Domain:   manager.GetDomain(),
		HttpOnly: manager.IsHttpOnly(),
		MaxAge:   3600 * 24,
	}
	http.SetCookie(w, cookie)
}
