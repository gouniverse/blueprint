package user

import "net/http"

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (controller *homeController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return "You are in user dashboard"
}
