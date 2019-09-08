package contracts

// Handler ...
type Handler interface {
	Save()
	Read(name string) map[string]interface{}
	Write(name string, content map[string]interface{})
	Destroy(name string)
	GC(lifetime int)
}
