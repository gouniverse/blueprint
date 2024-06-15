package routes

import (
	"net/http"
	"project/controllers/shared"
	"project/internal/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/router"
)

func sharedRoutes() []router.Route {
	sharedRoutes := []router.Route{
		{
			Name:    "Shared > Media Controller > Handler",
			Path:    links.MEDIA,
			Methods: []string{http.MethodGet},
			Handler: shared.NewMediaController().Handler,
		},
		{
			Name:    "Shared > Flash Controller > Index Page",
			Path:    links.FLASH,
			Handler: shared.NewFlashController().AnyIndex,
		},
		// {
		// 	Path:    links.MEDIA,
		// 	Handler: websiteControllers.NewMediaController().Index,
		// },
		{
			Name: "Shared > Theme",
			Path: links.THEME,
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				dashboard.ThemeHandler(w, r)
				return ""
			},
		},
	}

	return sharedRoutes
}
