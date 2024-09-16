package routes

import (
	"project/controllers/auth"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func authRoutes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Auth > Auth Controller",
			Path:        links.AUTH_AUTH,
			HTMLHandler: auth.NewAuthenticationController().Handler,
		},
		&router.Route{
			Name:        "Auth > Login Controller",
			Path:        links.AUTH_LOGIN,
			HTMLHandler: auth.NewLoginController().AnyIndex,
		},
		&router.Route{
			Name:        "Auth > Logout Controller",
			Path:        links.AUTH_LOGOUT,
			HTMLHandler: auth.NewLogoutController().AnyIndex,
		},
		&router.Route{
			Name:        "Auth > Register Controller",
			Path:        links.AUTH_REGISTER,
			HTMLHandler: auth.NewRegisterController().Handler,
		},
	}
}
