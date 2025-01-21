package helpers

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/sessionstore"
)

// GetAuthSession returns the authenticated session
func GetAuthSession(r *http.Request) sessionstore.SessionInterface {
	if r == nil {
		return nil
	}

	value := r.Context().Value(config.AuthenticatedSessionContextKey{})

	if value == nil {
		return nil
	}

	session := value.(sessionstore.SessionInterface)
	return session
}
