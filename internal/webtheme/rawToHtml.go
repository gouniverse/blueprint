package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) rawToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.Raw(text)
}
