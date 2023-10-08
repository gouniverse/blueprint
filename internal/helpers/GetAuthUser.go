package helpers

import (
	"net/http"
	"project/config"
	"project/models"
)

// GetAuthUser returns the authenticated user
func GetAuthUser(r *http.Request) *models.User {
	if r == nil {
		return nil
	}

	value := r.Context().Value(config.AuthenticatedUserKey{})
	if value == nil {
		return nil
	}
	user := value.(*models.User)
	return user
}
