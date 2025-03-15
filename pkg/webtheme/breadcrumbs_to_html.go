package webtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
)

func (t *theme) breadcrumbsToHtml(block ui.BlockInterface) *hb.Tag {
	breadcrumb1Text := block.Parameter("breadcrumb1_text")
	breadcrumb1Url := block.Parameter("breadcrumb1_url")

	breadcrumb2Text := block.Parameter("breadcrumb2_text")
	breadcrumb2Url := block.Parameter("breadcrumb2_url")

	breadcrumb3Text := block.Parameter("breadcrumb3_text")
	breadcrumb3Url := block.Parameter("breadcrumb3_url")

	breadcrumb1 := lo.If(breadcrumb1Text != "", hb.LI().
		Class("breadcrumb-item").
		Child(hb.Hyperlink().
			Href(lo.Ternary(breadcrumb1Url != "", breadcrumb1Url, "#")).
			Text(breadcrumb1Text))).
		Else(nil)
	breadcrumb2 := lo.If(breadcrumb2Text != "", hb.LI().
		Class("breadcrumb-item").
		Child(hb.Hyperlink().
			Href(lo.Ternary(breadcrumb2Url != "", breadcrumb2Url, "#")).
			Text(breadcrumb2Text))).
		Else(nil)
	breadcrumb3 := lo.If(breadcrumb3Text != "", hb.LI().
		Class("breadcrumb-item").
		Child(hb.Hyperlink().
			Href(lo.Ternary(breadcrumb3Url != "", breadcrumb3Url, "#")).
			Text(breadcrumb3Text))).
		Else(nil)

	return hb.Nav().Attr("aria-label", "breadcrumb").
		Child(hb.OL().
			Class("breadcrumb").
			Style("margin-bottom: 0px;").
			ChildIf(breadcrumb1Text != "", hb.LI().
				Class("breadcrumb-item").
				Child(breadcrumb1).
				Child(breadcrumb2).
				Child(breadcrumb3)),
		)

}
