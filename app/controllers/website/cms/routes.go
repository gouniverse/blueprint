package website

import (
	"project/app/links"
	"project/app/middlewares"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Website > Widget Controller > Handler",
			Path:        links.WIDGET,
			HTMLHandler: NewWidgetController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Home Page",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.HOME,
			HTMLHandler: NewCmsController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Catch All Pages",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.CATCHALL,
			HTMLHandler: NewCmsController().Handler,
		},
	}
}
