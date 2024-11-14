package user

import (
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
)

func breadcrumbs(path []bs.Breadcrumb) *hb.Tag {
	breadcrumbsPath := []bs.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewWebsiteLinks().Home(),
			Icon: icons.Icon("bi-house-fill", 16, 16, "gray").ToHTML(),
		},
	}

	breadcrumbsPath = append(breadcrumbsPath, path...)

	breadcrumbs := bs.Breadcrumbs(breadcrumbsPath)

	return breadcrumbs
}

func breadcrumbsSection(path []bs.Breadcrumb) *hb.Tag {
	return hb.Section().
		ID("SectionBreadcrumbs").
		Child(breadcrumbs(path)).
		Style("margin-bottom:10px;")
}

func breadcrumbsSectionWithContainer(path []bs.Breadcrumb) *hb.Tag {
	return hb.Section().
		ID("SectionBreadcrumbs").
		Children([]hb.TagInterface{
			hb.Div().
				Class("container").
				Child(breadcrumbs(path)),
		})
}
