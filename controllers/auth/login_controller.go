package auth

import (
	"net/http"
	"project/internal/links"
)

type loginController struct{}

func NewLoginController() *loginController {
	return &loginController{}
}

func (controller *loginController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	http.Redirect(w, r, links.NewAuthLinks().Login(links.NewWebsiteLinks().Home()), http.StatusSeeOther)
	return ""
}
