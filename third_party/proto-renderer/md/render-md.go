package protorenderer

import (
	"go/build"
	"path/filepath"
	"text/template"

	"github.com/googlecodelabs/tools/third_party/proto-renderer"
)

const (
	tmplsRltvDir = "src/github.com/googlecodelabs/tools/third_party/proto-renderer/md/templates/*"
)

var (
	t *template.Template
)

func init() {
	tmplsAbsDir := filepath.Join(build.Default.GOPATH, mdTmplsRltvDir)
	funcMap := template.FuncMap{
		InnerContents,
		InnerConcurrentContents,
	}

	t = template.Must(template.New("master").Funcs(funcMap).ParseGlob(tmplsAbsDir))
}

func Render(el inferface) string {
	return genrenderer.ExecuteTemplate(el, t)
}
