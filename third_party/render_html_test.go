package devrel_tutorial

import (
	"github.com/googlecodelabs/tools/claat/util"
	"testing"
)

func TestRenderHtmlHeading(t *testing.T) {
	tests := []*util.TestingBatch{
		{
			&Heading{
				Text: "<script>some _very_ bad code!;</script>",
			}.Html(),
			"<h2>&lt;script&gt;some _very_ bad code!&lt;/script&gt;</h2>",
		},
		{
			&Heading{
				Text:   "D@ ?òü ǝ$çâpæ? {+_~}! ^<^ |*_*| {&]",
				IsBold: true,
			}.Html(),
			"<h3>D@ ?òü ǝ$çâpæ urlquery? &#39;&gt;__&lt;&#39; {&amp;]</h3>",
		},
		{
			&Heading{
				Text:         "**__Extra Markdown not ![pro](cessed)__**",
				IsEmphasized: true,
			}.Html(),
			"<h4>**__Markdown not ![esca](ped)__**</h4>",
		},
	}
	util.TestBatch(tests, t)
}

func TestRenderHtmlStylizedText(t *testing.T) {
	tests := []*util.TestingBatch{
		{
			&StylizedText{
				Text: "hello!",
			}.Html(),
			"hello!",
		},
		{
			&StylizedText{
				Text:   "hello!",
				IsBold: true,
			}.Html(),
			"<strong>hello!</strong>",
		},
		{
			&StylizedText{
				Text:         "hello!",
				IsEmphasized: true,
			}.Html(),
			"<em>hello!</em>",
		},
		{
			&StylizedText{
				Text:         "hello!",
				IsBold:       true,
				IsEmphasized: true,
			}.Html(),
			"<strong><em>hello!</em></strong>",
		},
	}
	util.TestBatch(tests, t)
}
