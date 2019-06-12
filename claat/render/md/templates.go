package md

import (
	"os"
	"text/template"
)

var (
	templates = map[string]string{
		"text_template":           `{{.Text}}`,
		"emphasizedText_template": `<em>{{.EmphasizedText}}</em>`,
		"strongText_template":     `<strong>{{.StrongText}}</strong>`,
		"heading_template":        `<h{{.Level}}> {{.Text}} </h{{.Level}}>`,
		"link_template":           `<a href="{{.Location}}" target="_blank"> {{.Text}} </a>`,
	}
)

func ExecuteTemplate(s string, data interface{}) {
	t := template.Must(template.New("dummy").Parse(s))
	LogIfError(t.Execute(os.Stdout, data))
}
