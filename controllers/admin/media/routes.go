package admin

import (
	"project/internal/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Media Manager",
			Path:        links.ADMIN_MEDIA,
			HTMLHandler: NewMediaManagerController().AnyIndex,
		},
	}
}
