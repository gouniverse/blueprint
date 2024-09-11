package helpers

import (
	"errors"
	"net/http"
	"project/config"

	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/utils"
)

func UserSettingGet(r *http.Request, key string, defaultValue string) string {
	authUser := GetAuthUser(r)

	if authUser == nil {
		return defaultValue
	}

	hasValue, err := config.SessionStore.Has(key, sessionstore.SessionOptions{
		UserID:    authUser.ID(),
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	if err != nil {
		return defaultValue
	}

	if !hasValue {
		return defaultValue
	}

	value, err := config.SessionStore.Get(key, defaultValue, sessionstore.SessionOptions{
		UserID:    authUser.ID(),
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	if err != nil {
		return defaultValue
	}

	return value
}

func UserSettingSet(r *http.Request, key string, value string) error {
	authUser := GetAuthUser(r)

	if authUser == nil {
		return errors.New("auth user is nil")
	}

	err := config.SessionStore.Set(key, value, 60*60*24, sessionstore.SessionOptions{
		UserID:    authUser.ID(),
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	return err
}
