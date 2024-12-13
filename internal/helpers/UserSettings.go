package helpers

import (
	"errors"
	"net/http"
	"project/config"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/utils"
)

func UserSettingGet(r *http.Request, key string, defaultValue string) string {
	authUser := GetAuthUser(r)

	if authUser == nil {
		return defaultValue
	}

	session, err := config.SessionStore.SessionFindByKey(key)

	if err != nil {
		return defaultValue
	}

	if session.GetUserID() != authUser.ID() {
		return defaultValue
	}

	if session.GetIPAddress() != utils.IP(r) {
		return defaultValue
	}

	if session.GetUserAgent() != r.UserAgent() {
		return defaultValue
	}

	return session.GetValue()
}

func UserSettingSet(r *http.Request, key string, value string) error {
	authUser := GetAuthUser(r)

	if authUser == nil {
		return errors.New("auth user is nil")
	}

	session, err := config.SessionStore.SessionFindByKey(key)

	if err != nil {
		return err
	}

	if session.GetUserID() != authUser.ID() {
		return errors.New("session user id does not match auth user id")
	}

	if session.GetIPAddress() != utils.IP(r) {
		return errors.New("session ip address does not match request ip address")
	}

	if session.GetUserAgent() != r.UserAgent() {
		return errors.New("session user agent does not match request user agent")
	}

	session.SetValue(value)
	session.SetExpiresAt(carbon.Now(carbon.UTC).AddHours(1).ToDateTimeString(carbon.UTC))

	err = config.SessionStore.SessionUpdate(session)

	return err
}
