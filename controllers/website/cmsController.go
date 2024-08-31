package website

import (
	"net/http"
	"project/config"
	"strings"
)

const CMS_ENABLE_CACHE = false

// == CONTROLLER ===============================================================

type cmsController struct{}

// == CONSTRUCTOR ==============================================================

func NewCmsController() *cmsController {
	return &cmsController{}
}

// == PUBLIC METHODS ===========================================================

func (controller cmsController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	uri := r.RequestURI

	if strings.HasSuffix(uri, ".ico") {
		return ""
	}

	config.Cms.FrontendHandler(w, r)
	return ""
}
