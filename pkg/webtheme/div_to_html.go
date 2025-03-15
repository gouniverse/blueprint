package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) divToHtml(_ ui.BlockInterface) *hb.Tag {
	return hb.Div()
}
