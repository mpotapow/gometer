package session

import (
	"gometer/modules/session/contracts"
	"gometer/modules/session/handlers"
)

// Manager ...
type Manager struct {
	config  *Config
	drivers map[string]contracts.Handler
}

// GetManagerInstance ...
func GetManagerInstance(config *Config) contracts.Manager {
	return &Manager{
		config:  config,
		drivers: make(map[string]contracts.Handler),
	}
}

// GetDriver ...
func (m *Manager) GetDriver() contracts.Session {
	if _, ok := m.drivers[m.config.Driver]; !ok {
		m.drivers[m.config.Driver] = m.createDriver(m.config.Driver)
	}

	driver, _ := m.drivers[m.config.Driver]
	return m.buildSession(driver)
}

// GetPath ...
func (m *Manager) GetPath() string {
	return m.config.Path
}

// GetDomain ...
func (m *Manager) GetDomain() string {
	return m.config.Domain
}

// IsSecure ...
func (m *Manager) IsSecure() bool {
	return m.config.Secure
}

// IsHttpOnly ...
func (m *Manager) IsHttpOnly() bool {
	return m.config.HttpOnly
}

func (m *Manager) createDriver(name string) contracts.Handler {
	switch name {
	default:
		return handlers.GetMemoryHandlerInstance(m.config.Lifetime)
	}
}

func (m *Manager) buildSession(handler contracts.Handler) contracts.Session {

	return GetStoreInstance(m.config.Cookie, handler)
}
