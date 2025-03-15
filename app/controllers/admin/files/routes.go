package admin

import (
	"project/app/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > File Manager",
			Path:        links.ADMIN_FILE_MANAGER,
			HTMLHandler: NewFileManagerController().Handler,
		},
	}
}
