package helpers

import (
	"net/http"
	"project/config"
	"project/pkg/userstore"
)

// GetAuthUser returns the authenticated user
func GetAuthUser(r *http.Request) *userstore.User {
	if r == nil {
		return nil
	}

	value := r.Context().Value(config.AuthenticatedUserKey{})
	if value == nil {
		return nil
	}

	user := value.(*userstore.User)
	return user
}
