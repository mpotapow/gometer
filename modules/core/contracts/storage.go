package contracts

// Storage ...
type Storage interface {
	Set(name string, obj interface{})
	Get(name string) (interface{}, bool)
}
