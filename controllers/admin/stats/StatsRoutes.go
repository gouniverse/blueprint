package admin

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

func StatsRoutes() []router.RouteInterface {
	handler := func(w http.ResponseWriter, r *http.Request) string {
		controller := utils.Req(r, "controller", "")

		if controller == "visitor-activity" {
			return NewVisitorActivityController().Handler(w, r)
		}

		if controller == "visitor-paths" {
			return NewVisitorPathsController().Handler(w, r)
		}

		return NewHomeController().Handler(w, r)
	}

	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Stats > Orders",
			Path:        links.ADMIN_STATS,
			HTMLHandler: handler,
		},
		&router.Route{
			Name:        "Admin > Stats > Catchall",
			Path:        links.ADMIN_STATS + links.CATCHALL,
			HTMLHandler: handler,
		},
	}
}
