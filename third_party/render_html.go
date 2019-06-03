package devrel_tutorial

import (
	// "fmt"
	"go/build"
	"path/filepath"
	"text/template"
)

var (
	html *template.Template
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

func Render(el interface{}) string {
	tName := typeName(el)
	switch {
	case validTemplates[tName]:
		return executeTemplate(el, tName, html)
	case isNestedType(el):
		return executeNestedTemplate(el, html)
	default:
		return "Nope!" + tName
	}
}

// // Leaf Types
// func (el *Heading) Html() string {
// 	return executeTemplate(&el, "Heading", html)
// }

// func (el *StylizedText) Html() string {
// 	return executeTemplate(&el, "StylizedText", html)
// }
