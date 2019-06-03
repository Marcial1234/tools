package devrel_tutorial

import (
	"go/build"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
	funcMap := template.FuncMap{
		"repeatedHeading": repeatedHeading,
	}
	mdTmplsAbsDir := filepath.Join(build.Default.GOPATH, mdTmplsRltvDir)
	md = template.Must(template.New("master").Funcs(funcMap).ParseGlob(mdTmplsAbsDir))
}

func repeatedHeading(level int32) string {
	return strings.Repeat("#", int(level))
}

// TODO if possible: Template names follow its calling struct type,
//                   use `reflect.TypeOf(el).Elem().Name()`
//                   as a generic caller, once a catch-all type is figured out
func (el *Heading) Md() string {
	return executeTemplate(&el, "Heading", md)
}

func (el *StylizedText) Md() string {
	return executeTemplate(&el, "StylizedText", md)
}
