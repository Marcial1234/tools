package codelab_renderer

import (
	"testing"
)

func testMdBatch(tests []TestingBatch, t *testing.T) {
	for _, test := range tests {
		r := string(test)
		if test.o != r {
			t.Errorf("Expecting:\n\t'%s', but got \n\t'%s'", test.o, r)
			continue
		}
	}
}

func TestRenderHeadingMd(t *testing.T) {
	h2 := &Heading{
    Level: 2,
    Text: "<script>some _very_ bad code!;</script>",
	}
	h3 := &Heading{
    Level: 3,
    Text: "D@ ?òü ǝ$çâpæ? {+_~}! ^<^ |*_*| {&]",
	}
	h4 := &Heading{
    Level: 4,
    Text: "**__Extra Markdown not ![pro](cessed)__**",
	}
	tests := []TestingBatch {
		{&h2, "#### &lt;script&gt;some _very_ bad code!;&lt;/script&gt;"},
		{&h3, "##### D@ ?òü ǝ$çâpæ? {&#43;_~}! ^&lt;^ |*_*| {&amp;]"},
		{&h4, "###### **__Extra Markdown not ![pro](cessed)__**"},
	}
	testMdBatch(tests, t)
}