package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) paragraphToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.P().HTML(text)
}
