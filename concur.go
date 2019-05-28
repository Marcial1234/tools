// // Master types: Tutorial, Step, BlockContent ('sandwich' composition)
// // Intermediate types: Paragraph, Heading, Link, InfoBlock, [table types {List, FAQ, CheckList}]
// // Leaf Types: InlineCode, CodeBlock, Image, YT, StylizedText
// // all intermediaries 'oneof' will have 'content', all repeated named 'children', all final strings as 'text'

// InfoBlock: turn into <dl?>! byee
// Heading: HTML-1, MD-2

package main

import (
  "fmt"
  // "io/ioutil"
  // "strings"
  // "reflect"

  // "golang.org/x/net/html"
  // "./claat/new-render"
  // "./claat/new-render/html"
  "./third_party"
)

func main () {
  q := &devrel_tutorial.StylizedText{
    Text: "why you gotta be so extra",
  }
  // w := &devrel_tutorial.InlineContent{
  //       Content: &devrel_tutorial.InlineContent_Code{
  //         q,
  //   },
  // }
  // !!!!!
  h := &devrel_tutorial.TextBlock{
    Children: []*devrel_tutorial.InlineContent{
      {
        Content: &devrel_tutorial.InlineContent_Text{
          q,
        },
      },
    },
  }

  // fmt.Println(h.Html())
  // i := reflect.TypeOf(&h).Elem().Elem()
  // e := i.Name()
  // fmt.Printf("%T , %+v\n", i, i)
  // fmt.Println(i == reflect.TypeOf(devrel_tutorial.TextBlock{}))
  // fmt.Println()
  // fmt.Printf("%T , %+v\n", e, e)
  fmt.Printf(h.Html())
  // fmt.Println(render_base.IsOneof(i))
  // fmt.Println(i.FieldByName("Children"))

  // fmt.Println(i.FieldByName("Content"))
  // fmt.Println(i.FieldByName("Text"))
  // fmt.Println(h.Children[0].Content.(*devrel_tutorial.InlineContent_Code))

  // s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
  // // turn above into io.Reader? ufg...
  // z := html.NewTokenizer(strings.NewReader(h.Html()))
  // // var tree = {
  // //  {"e", "o"},
  // //  {"e", "o"}
  // // }
  // // fmt.Printf("%+v", z)
  // // fmt.Printf("%v", z)

  // woop!
  // for {
  //   tt := z.Next()

  //   switch {
  //   case tt == html.ErrorToken:
  //       // End of the document, we're done
  //       return
  //   default:
  //       fmt.Println(z.Token().String())
  //   }
  // }
}