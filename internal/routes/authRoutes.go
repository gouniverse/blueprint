package routes

import (
	"project/controllers/auth"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func authRoutes() []router.Route {
	return []router.Route{
		{
			Name:    "Auth > Auth Controller > Index Page",
			Path:    links.AUTH_AUTH,
			Handler: auth.NewAuthenticationController().AnyIndex,
		},
		{
			Name:    "Auth > Login Controller > Index Page",
			Path:    links.AUTH_LOGIN,
			Handler: auth.NewLoginController().AnyIndex,
		},
		{
			Name:    "Auth > Logout Controller > Index Page",
			Path:    links.AUTH_LOGOUT,
			Handler: auth.NewLogoutController().AnyIndex,
		},
	}
}
