package core

// Storage ...
type Storage struct {
	values map[string]interface{}
}

// Set ...
func (s *Storage) Set(name string, obj interface{}) {

	s.values[name] = obj
}

// Get ...
func (s *Storage) Get(name string) (interface{}, bool) {

	obj, ok := s.values[name]
	return obj, ok
}
