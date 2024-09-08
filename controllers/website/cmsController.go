package website

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/router"
)

const CMS_ENABLE_CACHE = false

// == CONTROLLER ===============================================================

type cmsController struct{}

var _ router.ControllerInterface = (*cmsController)(nil)

// == CONSTRUCTOR ==============================================================

func NewCmsController() *cmsController {
	return &cmsController{}
}

// == PUBLIC METHODS ===========================================================

func (controller cmsController) Handler(w http.ResponseWriter, r *http.Request) string {
	return config.Cms.FrontendHandlerRenderAsString(w, r)
}
