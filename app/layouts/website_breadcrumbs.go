package layouts

import (
	"project/app/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
)

func websiteBreadcrumbs(path []bs.Breadcrumb) hb.TagInterface {
	breadcrumbsPath := []bs.Breadcrumb{
		{
			Name: "",
			URL:  links.NewWebsiteLinks().Home(),
			Icon: icons.Icon("bi-house", 16, 16, "gray").ToHTML(),
		},
	}

	breadcrumbsPath = append(breadcrumbsPath, path...)

	breadcrumbs := bs.Breadcrumbs(breadcrumbsPath)

	return breadcrumbs
}

func NewWebsiteBreadcrumbsSection(path []bs.Breadcrumb) hb.TagInterface {
	return hb.Section().
		ID("SectionBreadcrumbs").
		Child(websiteBreadcrumbs(path)).
		Style("margin-bottom:10px;")
}

func NewWebsiteBreadcrumbsSectionWithContainer(path []bs.Breadcrumb) hb.TagInterface {
	return hb.Section().
		ID("SectionBreadcrumbs").
		Child(
			hb.Div().
				Class("container").
				Child(websiteBreadcrumbs(path)),
		)
}
