package helpers

import (
	"errors"
	"net/http"
	"project/config"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/utils"
	"github.com/spf13/cast"
)

func ExtendSession(r *http.Request, seconds int64) error {
	if config.SessionStore == nil {
		return errors.New("session store is nil")
	}

	session := GetAuthSession(r)

	if session == nil {
		return errors.New("session not found")
	}

	if session.GetIPAddress() != utils.IP(r) {
		return errors.New("session ip address does not match request ip address")
	}

	if session.GetUserAgent() != r.UserAgent() {
		return errors.New("session user agent does not match request user agent")
	}

	session.SetExpiresAt(carbon.Now(carbon.UTC).AddSeconds(cast.ToInt(seconds)).ToDateTimeString(carbon.UTC))

	err := config.SessionStore.SessionUpdate(r.Context(), session)

	return err
}
