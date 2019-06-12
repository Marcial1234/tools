// Copyright 2019 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package html
package main

import (
	"bytes"
	"html/template"
	"fmt"

	"github.com/googlecodelabs/tools/third_party"
	"github.com/googlecodelabs/tools/claat/util"
)

// have a "global" template variable?
// could read from files for complex templates
var (
	t = template.New("master-template-var")
	ts = map[string]string{
		"heading": `<h{{.Level}}> {{.Text}} </h{{.Level}}>`,
	}
)

func main() {
	element := &devrel_tutorial.Heading{
		Text:  "hello",
		Level: 2,
	}

	// # Templates on same package as proto
	// sending to Stdout
	fmt.Printf("AsHTML: ")
	element.AsHTML()
	fmt.Printf("\n")

	// secycling an io.Writer
	var w bytes.Buffer
	fmt.Printf("AsHTMLWithWritter: %s\n", element.AsHTMLWithWritter(&w))

	// # Template logic on its own package with 
	// 'render' switch f(x) - similar to current flow
	// register all templates
	for tmp_name, tmp_str := range ts { 
		registerTemplate(t, tmp_name, tmp_str)
	}

	var ww bytes.Buffer
	fmt.Printf("render: %s\n", render(element, &ww))
}

// # Code below is for "own package" approach
func registerTemplate(t *template.Template, name string, tmp_str string) {
	t.New(name).Parse(tmp_str)
}

// switch :/
func render(data interface{}, w *bytes.Buffer) string {
	var r string

	switch data.(type) {
	case *devrel_tutorial.Heading:
		r = executeTemplate("heading", data, w)
	}

	return r
}

func executeTemplate(tmp_name string, data interface{}, w *bytes.Buffer) string {
	util.LogIfError(t.ExecuteTemplate(w, tmp_name, data))
	return w.String()
}

//  ts = map[string]string{
//    "text": `{{.Text}}`,
//    "emphasized-text": `<em>{{.EmphasizedText}}</em>`,
//    "strong-text": `<strong>{{.StrongText}}</strong>`,
//    "heading": `<h{{.Level}}> {{.Text}} </h{{.Level}}>`,
//    // not allowing for bold &| italics links ~
//    "link": `<a href="{{.Location}}" target="_blank"> {{.Text}} </a>`,
//    // missing: %2f
//    // change "alt_title" to just "alt"
//    "image": `<img alt="{{.Alt}}" title="{{.Alt}}" src={{.Source}} 
//              style="width: {{.Size}}px" />`,
//  }