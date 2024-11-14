package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) rowToHtml(_ ui.BlockInterface) *hb.Tag {
	return hb.Div().Class("row")
}
