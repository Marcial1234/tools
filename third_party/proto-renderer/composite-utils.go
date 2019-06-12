package genrenderer

import (
	"reflect"
	"sync"
	"text/template"

	"github.com/googlecodelabs/tools/third_party"
)

// This will be a template f(x) ~ needs to be accesed by downstream rendering pkgs
func InnerContents(
	contents []devrel_tutorial.ProtoExtractor,
	t *template.Template) []string {
	sz := len(contents)
	renderedChildren := make([]string, sz)

	for i, c := range contents {
		// O(n)
		oneof := c.InnerContents()
		renderedChildren[j] = ExecuteTemplate(oneof, t)
	}

	return renderedChildren
}

func InnerContentsConcurrently(
	contents []devrel_tutorial.ProtoExtractor,
	t *template.Template) []string {
	sz := len(contents)
	renderedChildren := make([]string, sz)
	wg := new(sync.WaitGroup)
	wg.Add(sz)

	for i, c := range contents {
		// O(h) vs default O(n)
		go func(j int) {
			defer wg.Done()

			oneof := c.InnerContents()
			renderedChildren[j] = ExecuteTemplate(oneof, t)
		}(i)
	}

	wg.Wait()
	// will [beautifying/make HTML readable] later
	return renderedChildren
	// return string.Join(renderedChildren, "") => nope!
	// let the templates deal with this
}
