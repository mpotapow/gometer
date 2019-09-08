package contracts

// Session ...
type Session interface {
	GetName() string
	GetID() string
	SetID(id string)
	Get(name string) interface{}
	Set(name string, value interface{})
	Forget(name string)
	Flush()
	Start()
	Save()
}
