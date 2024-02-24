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
			// User > Home Controller > Index Page
			Name:    "User > Home Controller > Index Page",
			Path:    links.USER_HOME,
			Handler: user.NewHomeController().AnyIndex,
		},
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	return userRoutes
}
