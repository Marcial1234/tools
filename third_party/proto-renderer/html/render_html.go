package protorenderer

import (
	"go/build"
	"path/filepath"
	"text/template"

	"github.com/googlecodelabs/tools/third_party/proto-renderer"
)

const (
	tmplsRltvDir = "src/github.com/googlecodelabs/tools/third_party/proto-renderer/html/templates/*"
)

var (
	t *template.Template
)

func init() {
	tmplsAbsDir := filepath.Join(build.Default.GOPATH, tmplsRltvDir)
	funcMap := template.FuncMap{
		"renderInnerContents":             InnerContents,
		"renderInnerContentsConcurrently": InnerContentsConcurrently,
	}
	t = template.Must(template.New("master").Funcs(funcMap).ParseGlob(tmplsAbsDir))
}

// Actual rendering happening on Templates
func Render(el inferface) string {
	// Type check if happens at runtime, not compile time!
	// Need a shared interface (dummy one Name() string?)
	// to have compile time safety ~
	// TODO: chart of the types that'd need this?
	return protorenderer.ExecuteTemplate(el, t)
}

func InnerContents(contents []devrel_tutorial.ProtoExtractor) []string {
	return genrenderer.InnerContents(el, t)
}

func InnerContentsConcurrently(contents []devrel_tutorial.ProtoExtractor) []string {
	return genrenderer.InnerContentsConcurrently(el, t)
}
