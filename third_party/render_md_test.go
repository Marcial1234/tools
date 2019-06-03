package devrel_tutorial

import (
	"github.com/googlecodelabs/tools/claat/util"
	"testing"
)

func TestRenderMdHeading(t *testing.T) {
	tests := []*util.TestingBatch{
		{
			&Heading{
				Text:  "<script>some _very_ bad code!;</script>",
				Level: 2,
			}.Md(),
			"#### <script>some _very_ bad code!;</script>"},
		{
			&Heading{
				Text:  "D@ ?òü ǝ$çâpæ? {+_~}! ^<^ |*_*| {&]",
				Level: 3,
			}.Md(),
			"##### D@ ?òü ǝ$çâpæ? {+_~}! ^<^ |*_*| {&]"},
		{
			&Heading{
				Text:  "**__Extra Markdown not ![pro](cessed)__**",
				Level: 4,
			}.Md(),
			"###### **__Extra Markdown not ![pro](cessed)__**"},
	}
	util.TestBatch(tests, t)
}

func TestRenderMdStylizedText(t *testing.T) {
	tests := []*util.TestingBatch{
		{
			&StylizedText{
				Text: "hello!",
			}.Md(),
			"hello!"},
		{
			&StylizedText{
				Text:   "hello!",
				IsBold: true,
			}.Md(),
			"**hello!**"},
		{
			&StylizedText{
				Text:         "hello!",
				IsEmphasized: true,
			}.Md(),
			"__hello!__"},
		{
			&StylizedText{
				Text:         "hello!",
				IsBold:       true,
				IsEmphasized: true,
			}.Md(),
			"**__hello!__**"},
	}
	util.TestBatch(tests, t)
}
