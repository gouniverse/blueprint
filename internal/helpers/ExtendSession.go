package helpers

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/utils"
)

func ExtendSession(r *http.Request, seconds int64) error {
	sessionKey := r.Context().Value(config.AuthenticatedSessionKey{}).(string)

	err := config.SessionStore.Extend(sessionKey, 3600, sessionstore.SessionOptions{
		IPAddress: utils.IP(r),
		UserAgent: r.UserAgent(),
	})

	return err
}
