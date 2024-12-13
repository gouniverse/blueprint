package auth

import (
	"project/internal/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	return []router.RouteInterface{
		&router.Route{
			Name:        "Auth > Auth Controller",
			Path:        links.AUTH_AUTH,
			HTMLHandler: NewAuthenticationController().Handler,
		},
		&router.Route{
			Name:        "Auth > Login Controller",
			Path:        links.AUTH_LOGIN,
			HTMLHandler: NewLoginController().AnyIndex,
		},
		&router.Route{
			Name:        "Auth > Logout Controller",
			Path:        links.AUTH_LOGOUT,
			HTMLHandler: NewLogoutController().AnyIndex,
		},
		&router.Route{
			Name:        "Auth > Register Controller",
			Path:        links.AUTH_REGISTER,
			HTMLHandler: NewRegisterController().Handler,
		},
	}
}
