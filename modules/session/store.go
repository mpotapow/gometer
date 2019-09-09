package session

import (
	"crypto/rand"
	"encoding/base64"
	"gometer/modules/session/contracts"
	"io"
	"sync"
)

// Store ...
type Store struct {
	id      string
	name    string
	lock    sync.Mutex
	handler contracts.Handler
}

// GetStoreInstance ...
func GetStoreInstance(name string, handler contracts.Handler) contracts.Session {
	store := &Store{
		name:    name,
		handler: handler,
	}
	store.SetID(store.generateID())
	return store
}

// SetID ...
func (s *Store) SetID(id string) {
	s.id = id
}

// GetName ...
func (s *Store) GetName() string {
	return s.name
}

// GetID ...
func (s *Store) GetID() string {
	return s.id
}

func (s *Store) generateID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Get ...
func (s *Store) Get(name string) interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	data := s.handler.Read(s.GetID())
	if data == nil {
		return nil
	}
	if val, ok := data[name]; ok {
		return val
	}
	return nil
}

// Set ...
func (s *Store) Set(name string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	data := s.handler.Read(s.GetID())
	if data == nil {
		data = map[string]interface{}{}
		data[name] = value
	} else {
		data[name] = value
	}

	s.handler.Write(s.GetID(), data)
}

// Forget ...
func (s *Store) Forget(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	data := s.handler.Read(s.GetID())
	if data != nil {
		if _, ok := data[name]; ok {
			delete(data, name)
			s.handler.Write(s.GetID(), data)
		}
	}
}

// Flush ...
func (s *Store) Flush() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.handler.Destroy(s.GetID())
}

// Start ...
func (s *Store) Start() {
	s.lock.Lock()
	defer s.lock.Unlock()

	sessionData := s.handler.Read(s.GetID())
	if sessionData == nil {
		s.handler.Write(s.GetID(), map[string]interface{}{})
	} else {
		s.handler.Write(s.GetID(), sessionData)
	}
}

// Save ...
func (s *Store) Save() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.handler.Save()
}
