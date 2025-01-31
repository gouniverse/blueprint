package blogtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) hyperlinkToHtml(block ui.BlockInterface) *hb.Tag {
	url := block.Parameter("url")
	text := block.Parameter("content")
	return hb.A().Href(url).HTML(text)
}
