package devrel_tutorial

import (
	"bytes"
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
	md   *template.Template
	html *template.Template
)

type (
	ProtoExtractor interface {
		GetInnerContent() interface{}
	}
)

// Base Templating Logic
func executeTemplate(el interface{}, tmplName string, t *template.Template) string {
	var w bytes.Buffer
	util.LogIfError(t.ExecuteTemplate(&w, tmplName, el))
	return w.String()
}
