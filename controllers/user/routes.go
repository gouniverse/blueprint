package user

import (
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	userRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "User > Home",
			Path:        links.USER_HOME,
			HTMLHandler: NewHomeController().Handler,
		},
		&router.Route{
			Name:        "User > Profile",
			Path:        links.USER_PROFILE,
			HTMLHandler: NewProfileController().Handler,
		},
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	return userRoutes
}
