package admin

import (
	"project/app/links"

	"github.com/gouniverse/router"
)

func StatsRoutes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Stats > Orders",
			Path:        links.ADMIN_STATS,
			HTMLHandler: StatsController().Handler,
		},
		&router.Route{
			Name:        "Admin > Stats > Catchall",
			Path:        links.ADMIN_STATS + links.CATCHALL,
			HTMLHandler: StatsController().Handler,
		},
	}
}
