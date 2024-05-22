package routes

import (
	"project/controllers/user"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

func userRoutes() []router.Route {
	userRoutes := []router.Route{
		{
			// User > Home Page
			Name:    "User > Home Controller > Handle",
			Path:    links.USER_HOME,
			Handler: user.NewHomeController().Handle,
		},
		{
			// User > Profile Page
			Name:    "User > Profile Controller > Handle",
			Path:    links.USER_PROFILE,
			Handler: user.NewProfileController().Handle,
		},
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	return userRoutes
}
