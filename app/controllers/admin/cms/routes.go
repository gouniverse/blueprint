package admin

import (
	"project/internal/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	cmsOld := &router.Route{
		Name:        "Admin > Cms Manager",
		Path:        links.ADMIN_CMS,
		HTMLHandler: NewCmsController().Handler,
	}
	cmsNew := &router.Route{
		Name:    "Admin > Cms New Manager",
		Path:    links.ADMIN_CMS_NEW,
		Handler: NewCmsNewController().Handler,
	}

	return []router.RouteInterface{
		cmsOld,
		cmsNew,
	}
}
