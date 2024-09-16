package routes

import (
	"net/http"
	"project/controllers/shared"
	"project/internal/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/router"
)

func sharedRoutes() []router.RouteInterface {
	sharedRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Shared > Media Controller > Handler",
			Path:        links.MEDIA,
			Methods:     []string{http.MethodGet},
			HTMLHandler: shared.NewMediaController().Handler,
		},
		&router.Route{
			Name:        "Shared > Flash Controller",
			Path:        links.FLASH,
			HTMLHandler: shared.NewFlashController().Handler,
		},
		&router.Route{
			Name:        "Resources",
			Path:        links.RESOURCES + links.CATCHALL,
			HTMLHandler: shared.NewResourceController().Handler,
		},
		&router.Route{
			Name:    "Shared > Theme",
			Path:    links.THEME,
			Handler: dashboard.ThemeHandler,
		},
	}

	return sharedRoutes
}
