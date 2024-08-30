package routes

import (
	"project/controllers/auth"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func authRoutes() []router.Route {
	return []router.Route{
		{
			Name:    "Auth > Auth Controller",
			Path:    links.AUTH_AUTH,
			Handler: auth.NewAuthenticationController().Handler,
		},
		{
			Name:    "Auth > Login Controller",
			Path:    links.AUTH_LOGIN,
			Handler: auth.NewLoginController().AnyIndex,
		},
		{
			Name:    "Auth > Logout Controller",
			Path:    links.AUTH_LOGOUT,
			Handler: auth.NewLogoutController().AnyIndex,
		},
		{
			Name:    "Auth > Register Controller",
			Path:    links.AUTH_REGISTER,
			Handler: auth.NewRegisterController().Handler,
		},
	}
}
