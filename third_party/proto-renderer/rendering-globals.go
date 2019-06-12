package genrenderer

import (
	"bytes"
	"reflect"
	"text/template"
)

// Generic Templating Logic
func ExecuteTemplate(el interface{}, t *template.Template) string {
	tmplName := reflect.Typeof(el).Elem().Name()
	// or el.Name() if shared interface

	var w bytes.Buffer
	util.LogIfError(t.ExecuteTemplate(&w, tmplName, el))
	return w.String()
}
