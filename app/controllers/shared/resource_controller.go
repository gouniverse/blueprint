package shared

import (
	"net/http"
	"project/internal/resources"
	"strings"

	"github.com/gouniverse/router"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type resourceController struct{}

var _ router.HTMLControllerInterface = (*resourceController)(nil)

// == CONSTRUCTOR =============================================================

func NewResourceController() *resourceController {
	return &resourceController{}
}

// PUBLIC METHODS =============================================================

func (controller resourceController) Handler(w http.ResponseWriter, r *http.Request) string {
	uri := r.RequestURI

	// Remove leading /resources path
	if strings.HasPrefix(uri, "/resources/") {
		uri = strings.Replace(uri, "/resources/", "", 1)
	}

	// Is resource private?
	if strings.HasPrefix(uri, ".") {
		w.WriteHeader(http.StatusNotFound)
		return PageNotFoundController().Handler(w, r)
	}

	contentType := lo.If(strings.HasSuffix(uri, ".css"), "text/css").
		ElseIf(strings.HasSuffix(uri, ".gif"), "image/gif").
		ElseIf(strings.HasSuffix(uri, ".html"), "text/html").
		ElseIf(strings.HasSuffix(uri, ".ico"), "image/x-icon").
		ElseIf(strings.HasSuffix(uri, ".jpg"), "image/jpeg").
		ElseIf(strings.HasSuffix(uri, ".jpeg"), "image/jpeg").
		ElseIf(strings.HasSuffix(uri, ".js"), "application/javascript").
		ElseIf(strings.HasSuffix(uri, ".png"), "image/png").
		ElseIf(strings.HasSuffix(uri, ".ttf"), "font/ttf").
		ElseIf(strings.HasSuffix(uri, ".svg"), "image/svg+xml").
		ElseIf(strings.HasSuffix(uri, ".webp"), "image/webp").
		ElseIf(strings.HasSuffix(uri, ".woff"), "font/woff").
		ElseIf(strings.HasSuffix(uri, ".woff2"), "font/woff2").
		Else("text/plain")

	resourceContent := resources.Resource(uri)

	if resourceContent == "" {
		return PageNotFoundController().Handler(w, r)
	}

	w.Header().Set("Content-Type", contentType)
	return resourceContent
}
