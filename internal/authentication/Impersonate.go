package authentication

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/strutils"
	"github.com/gouniverse/utils"
	"github.com/spf13/cast"
)

func Impersonate(w http.ResponseWriter, r *http.Request, userID string) error {
	sessionTimeout := 2 * 60 * 60 // 2 hours

	if config.IsEnvDevelopment() {
		sessionTimeout = 4 * 60 * 60 // 4 hours
	}

	sessionKey := strutils.RandomFromGamma(64, "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz")
	errSession := config.SessionStore.Set(sessionKey, userID, cast.ToInt64(sessionTimeout), sessionstore.SessionOptions{
		UserID:    userID,
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	if errSession != nil {
		config.LogStore.ErrorWithContext("At Impersonate Error: ", errSession.Error())
		return errSession
	}

	auth.AuthCookieSet(w, r, sessionKey)

	return nil
}
