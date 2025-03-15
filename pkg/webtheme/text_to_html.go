package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) textToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.Span().HTML(text)
}
