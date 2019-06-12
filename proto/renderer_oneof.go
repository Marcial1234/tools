package devrel_tutorial

// TODO: How to make the following cleaner/more generic?
//       { [typString] = func(x) { return x.[Type] } ?
//       using: https://github.com/golang/protobuf/issues/261#issuecomment-430496210

func (el *InlineContent) GetInnerContent() interface{} {
	switch x := el.Content.(type) {
	case *InlineContent_Text:
		return x.Text
	}
	return nil
}

func (el *BlockContent) GetInnerContent() interface{} {
	switch x := el.Content.(type) {
	case *BlockContent_Text:
		return x.Text
	}
	return nil
}

// Moving parts:
//		Gateway => Redirects to rendering logic based on type
//		Direct Rendering => Done in templates
//		Packaging => Import swap vs protock pkg interface vs hybrid
// 				what's hybrid? need to have some type of common interface
// 				to be able to group them all...
//		Oneof logic: => Must have type-based-switch
//		Nested/Repeated => Handle them recurrently

// Testing: Init & Stress Testing via other stack overflow thing ...
