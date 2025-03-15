package shared

import (
	"net/http"
	"project/app/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/responses"
	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	adsTxt := &router.Route{
		Name: "Shared > ads.txt",
		Path: "/ads.txt",
		HTMLHandler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
			return "google.com, pub-8821108004642146, DIRECT, f08c47fec0942fa0"
		}),
	}

	files := &router.Route{
		Name:        "Shared > Files Controller",
		Path:        links.FILES,
		Methods:     []string{http.MethodGet},
		HTMLHandler: NewFileController().Handler,
	}

	flash := &router.Route{
		Name:        "Shared > Flash Controller",
		Path:        links.FLASH,
		HTMLHandler: NewFlashController().Handler,
	}

	media := &router.Route{
		Name:        "Shared > Media Controller",
		Path:        links.MEDIA,
		Methods:     []string{http.MethodGet},
		HTMLHandler: NewMediaController().Handler,
	}

	resources := &router.Route{
		Name:        "Shared > Resources Controller",
		Path:        links.RESOURCES,
		HTMLHandler: NewResourceController().Handler,
	}

	theme := &router.Route{
		Name:    "Shared > Theme Controller",
		Path:    links.THEME,
		Handler: dashboard.ThemeHandler,
	}

	thumb := &router.Route{
		Name:        "Shared > Thumb Controller",
		Path:        links.THUMB,
		HTMLHandler: NewThumbController().Handler,
	}

	return []router.RouteInterface{
		adsTxt,
		files,
		flash,
		media,
		resources,
		theme,
		thumb,
	}
}
