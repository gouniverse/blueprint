package website

import "net/http"

// == CONTROLLER ===============================================================

type homeController struct{}

// == CONSTRUCTOR ==============================================================

func NewHomeController() *homeController {
	return &homeController{}
}

// == PUBLIC METHODS ===========================================================

func (controller *homeController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return "You are at the website home page"
}
