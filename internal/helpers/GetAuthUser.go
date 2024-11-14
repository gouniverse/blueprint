package helpers

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/userstore"
)

// GetAuthUser returns the authenticated user
func GetAuthUser(r *http.Request) userstore.UserInterface {
	if r == nil {
		return nil
	}

	value := r.Context().Value(config.AuthenticatedUserKey{})
	if value == nil {
		return nil
	}

	user := value.(userstore.UserInterface)
	return user
}
