package view

import (
	"gometer/modules/view/contracts"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

// View ...
type View struct {
	path     string
	template *template.Template
}

// GetViewInstance ...
func GetViewInstance(path string) contracts.View {

	view := &View{
		path: path,
	}

	if templ, err := view.getTemplateInstance(); err != nil {
		panic(err)
	} else {
		view.template = templ
	}

	return view
}

func (v *View) getTemplateInstance() (*template.Template, error) {

	paths := v.getTemplatesPathFromDir(v.path)
	tmpl := template.New("base").Delims("[[", "]]")

	return tmpl.ParseFiles(paths...)
}

func (v *View) getTemplatesPathFromDir(dir string) []string {

	var paths []string
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	return paths
}

// Render ...
func (v *View) Render(wr io.Writer, tmplName string, data interface{}) {

	err := v.template.ExecuteTemplate(wr, tmplName+".html", data)
	if err != nil {
		panic(err)
	}
}
