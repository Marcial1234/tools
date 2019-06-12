package devrel_tutorial

import (
	"reflect"
	"sync"
	"text/template"
)

// This will be a template f(x) ~ needs to be accesed by
func InnerContents(contents []ProtoExtractor, outputFormat string) []interface{} {
	sz := len(contents)
	wg := new(sync.WaitGroup)
	renderedChildren := make([]interface{}, sz)
	wg.Add(sz)

	for i, c := range contents {
		// O(h) vs default O(n)
		go func(j int) {
			defer wg.Done()

			oneof := c.InnerContents()
			o := strings.ToLower(outputFormat)
			t := protorenderer.Templates[o]
			tmplName := reflect.Typeof(el).Elem().Name()
			renderedChildren[i] = executeTemplate(oneof, tmplName)
		}(i)
	}

	wg.Wait()
	// will [beautifying/make HTML readable] later
	return renderedChildren
	// return string.Join(renderedChildren, "")
}
