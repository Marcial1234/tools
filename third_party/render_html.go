package devrel_tutorial

import (
	"go/build"
	"path/filepath"
	"text/template"
)

func init() {
	htmlTmplsDir := filepath.Join(build.Default.GOPATH, htmlTmplsRltvDir)
	html = template.Must(template.New("master").ParseGlob(htmlTmplsDir))
}

// TODO if possible:
//     Template names are named after their struct type,
//     use `reflect.TypeOf(el).Elem().Name()`
//     as a generic caller once a catch-all type, or
//     field-persistent inferace-passing process is figured out

// Leaf Types
func (el *Heading) Html() string {
	return executeTemplate(&el, "Heading", html)
}

func (el *StylizedText) Html() string {
	return executeTemplate(&el, "StylizedText", html)
}
