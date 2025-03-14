package shared

import (
	"net/http"

	"github.com/gouniverse/router"
)

// == CONSTRUCTOR =============================================================

func PageNotFoundController() router.HTMLControllerInterface {
	return &pageNotFoundController{}
}

// == CONTROLLER ==============================================================

type pageNotFoundController struct{}

// PUBLIC METHODS =============================================================

func (controller *pageNotFoundController) Handler(w http.ResponseWriter, r *http.Request) string {
	w.WriteHeader(http.StatusNotFound)
	return "Sorry, page not found."
}
