package testutils

import (
	"errors"
	"net/http"
	"project/config"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
)

func SeedSession(r *http.Request, user userstore.UserInterface, expiresSeconds int) (sessionstore.SessionInterface, error) {
	if config.SessionStore == nil {
		return nil, errors.New("session store is nil")
	}

	session := sessionstore.NewSession().
		SetUserID(user.ID()).
		SetUserAgent(r.UserAgent()).
		SetIPAddress(utils.IP(r)).
		SetExpiresAt(carbon.Now(carbon.UTC).AddSeconds(expiresSeconds).ToDateTimeString(carbon.UTC))

	err := config.SessionStore.SessionCreate(session)

	if err != nil {
		return nil, err
	}

	return session, nil
}
