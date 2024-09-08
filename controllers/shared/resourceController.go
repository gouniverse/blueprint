package shared

import (
	"net/http"
	"project/internal/resources"
	"strings"
)

func NewResourceController() *resourceController {
	return &resourceController{}
}

type resourceController struct{}

func (controller resourceController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	uri := r.RequestURI

	contentType := ""

	if strings.HasSuffix(uri, ".js") {
		contentType = "application/javascript"
	}

	if strings.HasPrefix(uri, "/resources/") {
		uri = strings.Replace(uri, "/resources/", "", 1)
	}

	resourceContent := resources.Resource(uri)

	if resourceContent == "" {
		return uri
	}

	w.Header().Set("Content-Type", contentType)

	return resourceContent
}
