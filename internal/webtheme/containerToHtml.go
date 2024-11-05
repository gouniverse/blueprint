package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) containerToHtml(block ui.BlockInterface) *hb.Tag {
	return hb.Div().Class("container")
}
