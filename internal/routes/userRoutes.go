package routes

import (
	"project/controllers/user"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

func userRoutes() []router.RouteInterface {
	userRoutes := []router.RouteInterface{
		&router.Route{
			// User > Home Page
			Name:        "User > Home Controller > Handle",
			Path:        links.USER_HOME,
			HTMLHandler: user.NewHomeController().Handler,
		},
		&router.Route{
			// User > Profile Page
			Name:        "User > Profile Controller > Handle",
			Path:        links.USER_PROFILE,
			HTMLHandler: user.NewProfileController().Handler,
		},
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	return userRoutes
}
