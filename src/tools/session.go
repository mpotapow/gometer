package tools

import (
	"gometer/modules/core"
	"gometer/modules/session/contracts"
	"net/http"
)

// Session ...
type Session struct {
	request *http.Request
	manager contracts.Manager
	store   contracts.Session
}

// GetSessionInstance ...
func GetSessionInstance(r *http.Request) *Session {
	sessionInst := &Session{
		request: r,
	}

	managerInst, _ := core.GetApplicationInstance().Get("session")
	sessionInst.manager = managerInst.(contracts.Manager)

	store := sessionInst.manager.GetDriver()
	if cookie, err := r.Cookie(store.GetName()); err == nil {
		store.SetID(cookie.Value)
	}
	sessionInst.store = store

	return sessionInst
}

// GetManager ...
func (s *Session) GetManager() contracts.Manager {
	return s.manager
}

// GetStorage ..
func (s *Session) GetStorage() contracts.Session {
	return s.store
}
