package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func (t *theme) columnToHtml(block ui.BlockInterface) *hb.Tag {
	widthXs := block.Parameter("width_xs")
	widthSm := block.Parameter("width_sm")
	widthMd := block.Parameter("width_md")
	widthLg := block.Parameter("width_lg")
	widthXl := block.Parameter("width_xl")
	widthXxl := block.Parameter("width_xxl")

	if widthXs == "" && widthSm == "" && widthMd == "" && widthLg == "" && widthXl == "" && widthXxl == "" {
		widthSm = "12"
	}

	return hb.Div().ClassIf(widthXs != "", "col-xs-"+widthXs).
		ClassIf(widthSm != "", "col-sm-"+widthSm).
		ClassIf(widthMd != "", "col-md-"+widthMd).
		ClassIf(widthLg != "", "col-lg-"+widthLg).
		ClassIf(widthXl != "", "col-xl-"+widthXl).
		ClassIf(widthXxl != "", "col-xxl-"+widthXxl)
}
