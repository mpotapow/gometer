package contracts

import "io"

// View ...
type View interface {
	Render(wr io.Writer, tmplName string, data interface{})
}
