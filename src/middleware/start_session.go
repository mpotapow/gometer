package middleware

import (
	"gometer/modules/core"
	"gometer/modules/session/contracts"
	"net/http"
)

// StartSession ...
func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		manager := getSessionManager()
		session := getSession(manager, r)
		session.Start()

		addCookieResponse(manager, session, w)

		next.ServeHTTP(w, r)

		session.Save()
	})
}

func getSessionManager() contracts.Manager {
	managerInst, _ := core.GetApplicationInstance().Get("session")
	return managerInst.(contracts.Manager)
}

func getSession(manager contracts.Manager, r *http.Request) contracts.Session {
	session := manager.GetDriver()
	if cookie, err := r.Cookie(session.GetName()); err == nil {
		session.SetID(cookie.Value)
	}
	return session
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
