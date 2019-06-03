package devrel_tutorial

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/googlecodelabs/tools/claat/util"
)

// Reviewing comment: Templates can live anywhere within the repo, putting them
// here for now
const (
	mdTmplsRltvDir   = "src/github.com/googlecodelabs/tools/third_party/templates/md/*"
	htmlTmplsRltvDir = "src/github.com/googlecodelabs/tools/third_party/templates/html/*"
)

var (
	validTemplates = map[string]bool{
		"Heading":      true,
		"Link":         true,
		"StylizedText": true,
	}
	nestedTypes = map[string]bool{
		"InlineContent": true,
	}
)

// type (
// 	ProtoRenderer interface {
// 		Md() string
// 		Html() string
// 	}
// )

// Reflect helpers
// Returns the package-stripped type name of the passed argument as string
func typeName(el interface{}) string {
	// Pointer types require 'Elem' derefence to get type name as string
	// el => *devrel_tutorial.elType => devrel_tutorial.elType => elType
	return underlyingType(el).Name()
}

// Returns the package-stripped type name of the passed argument as string
func underlyingType(el interface{}) reflect.Type {
	// Pointer types require 'Elem' derefence to get type name as string
	// el => *devrel_tutorial.elType => devrel_tutorial.elType => elType
	return reflect.TypeOf(el).Elem()
}

func innerFieldType(el interface{}, field string) reflect.Type {
	ty, exists := reflect.TypeOf(el).Elem().FieldByName(field)
	if !exists {
		fmt.Println("error! (%s) of type %T doesn't have field '%s'", el, el, field)
	}
	return ty.Type
}

func innerContentType(el interface{}) reflect.Type {
	// el => el.Content => []*devrel_tutorial.elContentType
	return innerFieldType(el, "Content")
}

func isNestedType(el interface{}) bool {
	return innerContentType(el).Kind() == reflect.Slice
}

// return of '0' means not found!
func innerFieldValue(el interface{}, field string) reflect.Value {
	return reflect.ValueOf(el).Elem().FieldByName(field)
}

func innerContentValue(el interface{}) reflect.Value {
	// el => el.Content => []*devrel_tutorial.elContentType
	return innerFieldValue(el, "Content")
}

func oneofContent(el interface{}) reflect.Value {
	// ugliest thing here...
	// valueNotFound := reflect.ValueOf(nil)
	// switch {
	// case innerFieldValue(el, "Text") != valueNotFound:
	// 	return innerFieldValue(el, "Text")
	// case innerFieldValue(el, "Heading") != valueNotFound:
	// 	return innerFieldValue(el, "Heading")
	// case innerFieldValue(el, "Button") != valueNotFound:
	// 	return innerFieldValue(el, "Button")
	// case innerFieldValue(el, "InfoBox") != valueNotFound:
	// 	return innerFieldValue(el, "InfoBox")
	// case innerFieldValue(el, "CodeBlock") != valueNotFound:
	// 	return innerFieldValue(el, "CodeBlock")
	// case innerFieldValue(el, "List") != valueNotFound:
	// 	return innerFieldValue(el, "List")
	// case innerFieldValue(el, "Survey") != valueNotFound:
	// 	return innerFieldValue(el, "Survey")
	// case innerFieldValue(el, "Image") != valueNotFound:
	// 	return innerFieldValue(el, "Image")
	// }
	// return valueNotFound
	// return innerFieldValue(innerFieldValue(el, "Text"), "Text")
	return innerFieldValue(el, "Text")
}

// Base Templating Logic
func executeTemplate(el interface{}, tmplName string, t *template.Template) string {
	var w bytes.Buffer
	util.LogIfError(t.ExecuteTemplate(&w, tmplName, el))
	return w.String()
}

func executeNestedTemplate(el interface{}, t *template.Template) string {
	// *devrel_tutorial.elType => devrel_tutorial.elType => elType
	childrenVals := innerContentValue(el)
	// childrenType := reflect.ValueOf(el).Elem().FieldByName("Content")

	// call executeTemplate on a syncgroup loop..
	// defer done...
	for i := 0; i < childrenVals.Len(); i++ {
		v := childrenVals.Index(i).Elem() //.Elem() //.Elem().Field(0)
		// ty := underlyingType(innerContentType(el))
		// ty := innerFieldType(innerContentType(el).Elem(), "text")
		ty := innerContentType(el).Elem().Elem().Field(0).Type
		// fmt.Println(executeTemplate(v, ty, t))
		qq := reflect.New(ty)
		qq.Elem().Field(0).Set(v)
		// qw = qq.Elem().Interface().()
		fmt.Println(v, ty, qq)
	}

	return "_"
}
