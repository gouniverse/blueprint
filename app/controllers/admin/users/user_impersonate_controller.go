package admin

import (
	"net/http"

	"project/internal/authentication"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

// == CONTROLLER ==============================================================

// userImpersonateController represents a controller for handling user impersonation.
type userImpersonateController struct{}

// == CONSTRUCTOR =============================================================

func NewUserImpersonateController() *userImpersonateController {
	return &userImpersonateController{}
}

var _ router.HTMLControllerInterface = (*userImpersonateController)(nil)

// == PUBLIC METHODS ==========================================================

func (c *userImpersonateController) Handler(w http.ResponseWriter, r *http.Request) string {
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return helpers.ToFlashError(w, r, "User not found", links.NewAdminLinks().Users(map[string]string{}), 15)
	}

	if !authUser.IsAdministrator() {
		return helpers.ToFlashError(w, r, "Not authorized", links.NewAdminLinks().Users(map[string]string{}), 15)
	}

	userID := utils.Req(r, "user_id", "")

	if userID == "" {
		return helpers.ToFlashError(w, r, "User ID not found", links.NewAdminLinks().Users(map[string]string{}), 15)
	}

	err := authentication.Impersonate(w, r, userID)

	if err != nil {
		return helpers.ToFlashError(w, r, err.Error(), links.NewAdminLinks().Users(map[string]string{}), 15)
	}

	return helpers.ToFlashSuccess(w, r, "Impersonation is successful", links.NewUserLinks().Home(map[string]string{}), 15)

}
