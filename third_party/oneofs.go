package devrel_tutorial

// Optional: shared dummy interface that ALL protos have to implement for compile time checks
// type ProtoRenderer interface {
// 	Name () string
// }

// repeated 'Step' and 'Textblock' are dealth with directly in templates!
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
