package blogtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) ulToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.UL().HTML(text)
}

func (t *theme) olToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.OL().HTML(text)
}

func (t *theme) liToHtml(block ui.BlockInterface) *hb.Tag {
	text := block.Parameter("content")
	return hb.LI().HTML(text)
}
