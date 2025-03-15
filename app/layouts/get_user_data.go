package layouts

import (
	"errors"
	"net/http"
	"project/config"
	"project/internal/helpers"

	"github.com/gouniverse/userstore"
)

func getUserData(r *http.Request, authUser userstore.UserInterface) (firstName string, lastName string, err error) {
	if authUser == nil {
		return "n/a", "", errors.New("user is nil")
	}

	if !config.VaultStoreUsed {
		firstName = authUser.FirstName()
		lastName = authUser.LastName()

		if firstName == "" && lastName == "" {
			return authUser.Email(), "", nil
		}

		return firstName, lastName, nil
	}

	firtsName, lastName, _, err := helpers.UserUntokenized(r.Context(), authUser)

	return firtsName, lastName, err
}
