package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) iconToHtml(block ui.BlockInterface) *hb.Tag {
	icon := block.Parameter("icon")
	return hb.I().Class(icon)
}
