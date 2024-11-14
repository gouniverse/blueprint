package routes

import (
	"net/http"
	"project/controllers/shared"
	"project/internal/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/responses"
	"github.com/gouniverse/router"
)

func sharedRoutes() []router.RouteInterface {
	sharedRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Shared > Files Controller",
			Path:        links.FILES,
			Methods:     []string{http.MethodGet},
			HTMLHandler: shared.NewFileController().Handler,
		},
		&router.Route{
			Name:        "Shared > Media Controller",
			Path:        links.MEDIA,
			Methods:     []string{http.MethodGet},
			HTMLHandler: shared.NewMediaController().Handler,
		},
		&router.Route{
			Name: "Shared > ads.txt",
			Path: "/ads.txt",
			HTMLHandler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
				//return "google.com, pub-8821108004642146, DIRECT, f08c47fec0942fa0"
				return "google.com, pub-YOURNUMBER, DIRECT, YOURSTRING"
			}),
		},
		&router.Route{
			Name:        "Shared > Flash Controller",
			Path:        links.FLASH,
			HTMLHandler: shared.NewFlashController().Handler,
		},
		&router.Route{
			Name:        "Shared > Resources",
			Path:        links.RESOURCES + links.CATCHALL,
			HTMLHandler: shared.NewResourceController().Handler,
		},
		&router.Route{
			Name:    "Shared > Theme",
			Path:    links.THEME,
			Handler: dashboard.ThemeHandler,
		},
		&router.Route{
			Name:        "Shared > Thumbnail",
			Path:        links.THUMB,
			HTMLHandler: shared.NewThumbController().Handler,
		},
	}

	return sharedRoutes
}
