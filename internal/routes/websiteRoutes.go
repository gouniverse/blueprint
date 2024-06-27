package routes

import (
	"project/controllers/website"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func websiteRoutes() []router.Route {
	websiteRoutes := []router.Route{
		// {
		// 	Path:    links.HOME,
		// 	Handler: websiteControllers.NewHomeController().AnyIndex,
		// },
		// {
		// 	Path:    links.CATCHALL,
		// 	Handler: sharedControllers.NewPageNotFoundControllerController().AnyIndex,
		// },
		{
			Name:    "Website > Widget Controller > Handler",
			Path:    links.WIDGET,
			Handler: website.NewWidgetController().Handler,
		},
		{
			Name:    "Website > Cms > Home Page",
			Path:    links.HOME,
			Handler: website.NewCmsController().AnyIndex,
		},
		{
			Name:    "Website > Cms > Catch All Pages",
			Path:    links.CATCHALL,
			Handler: website.NewCmsController().AnyIndex,
		},
	}

	return websiteRoutes
}
