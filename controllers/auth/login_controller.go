package auth

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
)

type loginController struct{}

func NewLoginController() *loginController {
	return &loginController{}
}

func (controller *loginController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	if !config.UserStoreUsed || config.UserStore == nil {
		return helpers.ToFlashError(w, r, `user store is required`, links.NewWebsiteLinks().Home(), 5)
	}

	if !config.VaultStoreUsed || config.VaultStore == nil {
		return helpers.ToFlashError(w, r, `vault store is required`, links.NewWebsiteLinks().Home(), 5)
	}

	http.Redirect(w, r, links.NewAuthLinks().Login(links.NewWebsiteLinks().Home()), http.StatusSeeOther)
	return ""
}
