package user

import (
	userAccount "project/app/controllers/user/account"
	userHome "project/app/controllers/user/home"

	"project/app/middlewares"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	home := &router.Route{
		Name:        "User > Home",
		Path:        links.USER_HOME,
		HTMLHandler: userHome.NewHomeController().Handler,
	}

	homeCatchAll := &router.Route{
		Name:        "User > Catch All Controller > Index Page",
		Path:        links.USER_HOME + links.CATCHALL,
		HTMLHandler: userHome.NewHomeController().Handler,
	}

	profile := &router.Route{
		Name:        "User > Profile",
		Path:        links.USER_PROFILE,
		HTMLHandler: userAccount.NewProfileController().Handler,
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
