package handlers

import (
	"gometer/modules/session/contracts"
	"time"
)

// MemoryHandler ...
type MemoryHandler struct {
	ttl  int
	heap map[string]memoryData
}

type memoryData struct {
	lifetime int64
	content  map[string]interface{}
}

// GetMemoryHandlerInstance ...
func GetMemoryHandlerInstance(ttl int) contracts.Handler {
	return &MemoryHandler{
		ttl:  ttl,
		heap: make(map[string]memoryData),
	}
}

func (m *MemoryHandler) getTimestampWithTTL() int64 {
	now := time.Now()
	return now.Unix() + int64(m.ttl)
}

// Read ...
func (m *MemoryHandler) Read(name string) map[string]interface{} {
	data, ok := m.heap[name]
	if ok {
		return data.content
	}
	return nil
}

// Write ...
func (m *MemoryHandler) Write(name string, content map[string]interface{}) {
	data, ok := m.heap[name]
	if ok {
		data.content = content
		data.lifetime = m.getTimestampWithTTL()
	} else {
		data = memoryData{m.getTimestampWithTTL(), content}
		m.heap[name] = data
	}
}

// Destroy ...
func (m *MemoryHandler) Destroy(name string) {
	if _, ok := m.heap[name]; ok {
		delete(m.heap, name)
	}
}

// GC ...
func (m *MemoryHandler) GC(lifetime int) {

}

// Save ...
func (m *MemoryHandler) Save() {

}
