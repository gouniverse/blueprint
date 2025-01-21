package testutils

import (
	"errors"
	"net/http"

	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/userstore"
)

func SeedUserAndSession(userID string, r *http.Request, expiresSeconds int) (user userstore.UserInterface, session sessionstore.SessionInterface, err error) {
	user, err = SeedUser(userID)

	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, errors.New("user should not be nil")
	}

	session, err = SeedSession(r, user, expiresSeconds)

	if err != nil {
		return nil, nil, err
	}

	if session == nil {
		return nil, nil, errors.New("session should not be nil")
	}

	return user, session, nil
}
