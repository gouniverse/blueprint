package auth

import (
	"net/http"
	"project/app/links"
	"project/config"
	"project/internal/helpers"
	"strings"

	"github.com/dracory/base/req"
	"github.com/gouniverse/router"
)

type loginController struct{}

func NewLoginController() router.HTMLControllerInterface {
	return &loginController{}
}

func (controller *loginController) Handler(w http.ResponseWriter, r *http.Request) string {
	if !config.UserStoreUsed || config.UserStore == nil {
		return helpers.ToFlashError(w, r, `user store is required`, links.NewWebsiteLinks().Home(), 5)
	}

	if config.VaultStoreUsed && config.VaultStore == nil {
		return helpers.ToFlashError(w, r, `vault store is required`, links.NewWebsiteLinks().Home(), 5)
	}

	backUrl := req.ValueOr(r, "back_url", links.NewWebsiteLinks().Home())
	if !strings.HasPrefix(backUrl, links.NewWebsiteLinks().Home()) {
		backUrl = links.NewWebsiteLinks().Home()
	}

	loginUrl := links.NewAuthLinks().AuthKnightLogin(backUrl)

	http.Redirect(w, r, loginUrl, http.StatusSeeOther)
	return ""
}
