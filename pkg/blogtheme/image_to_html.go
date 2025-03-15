package blogtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) imageToHtml(block ui.BlockInterface) *hb.Tag {
	imageUrl := block.Parameter("image_url")
	return hb.Img(imageUrl)
}
