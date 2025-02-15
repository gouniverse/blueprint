package user

import (
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"

	account "project/controllers/user/account"
)

func Routes() []router.RouteInterface {
	home := &router.Route{
		Name:        "User > Home",
		Path:        links.USER_HOME,
		HTMLHandler: NewHomeController().Handler,
	}

	homeCatchAll := &router.Route{
		Name:        "User > Catch All Controller > Index Page",
		Path:        links.USER_HOME + links.CATCHALL,
		HTMLHandler: NewHomeController().Handler,
	}

	profile := &router.Route{
		Name:        "User > Profile",
		Path:        links.USER_PROFILE,
		HTMLHandler: account.NewProfileController().Handler,
	}

	userRoutes := []router.RouteInterface{
		profile,
		home,
		homeCatchAll,
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	return userRoutes
}
