package main

import (
  "fmt"
  // "reflect"

  // "./claat/new-render"
  // "./claat/new-render/html"
  "./third_party"
  "./third_party/proto-renderer/html"
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

  fmt.Println(protorenderer.Render(h))
  // fmt.Println(codelabRender.Render(h))
}
