package shared

import (
	"net/http"

	"github.com/gouniverse/router"
)

// == CONTROLLER ==============================================================

type pageNotFoundController struct{}

var _ router.HTMLControllerInterface = (*pageNotFoundController)(nil)

// == CONSTRUCTOR =============================================================

func NewPageNotFoundController() *pageNotFoundController {
	return &pageNotFoundController{}
}

// PUBLIC METHODS =============================================================

func (controller *pageNotFoundController) Handler(w http.ResponseWriter, r *http.Request) string {
	w.WriteHeader(http.StatusNotFound)
	return "Sorry, page not found."
}
