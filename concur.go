package main

import (
  "fmt"
  // "io/ioutil"
  // "strings"
  "reflect"

  // "golang.org/x/net/html"
  // "./claat/new-render"
  // "./claat/new-render/html"
  "./third_party"
)

func main() {
  q := &devrel_tutorial.StylizedText{
    Text: `print('Hello World!')`,
  }

  // will have helper f(x)
  h := &devrel_tutorial.TextBlock{
    Content: []*devrel_tutorial.InlineContent{
      {
        Content: &devrel_tutorial.InlineContent_Text{
          q,
        },
      },
    },
  }

  // 2nd type tells you which inline it is... but,?

  // fmt.Println(h)innerFieldValue
  // a.Content => is of some type ~
  // p, _ := reflect.TypeOf(h).Elem().FieldByName("Content")
  // fmt.Println(p.Type.Kind(reflect.ValueOf(el)) == reflect.Slice)
  // type check ^
  // now onto value...
  v := reflect.ValueOf(h).Elem().FieldByName("Content")
  // fmt.Println(v.Type.Elem().Name())
  fmt.Println(v)

  // a := reflect.ValueOf(h).Elem().FieldByName("Content")
  // fmt.Println(reflect.TypeOf(a.Pointer()))
  // fmt.Println(reflect.TypeOf(h).Elem().Elem().Name())
  fmt.Println(devrel_tutorial.Render(h))
}
