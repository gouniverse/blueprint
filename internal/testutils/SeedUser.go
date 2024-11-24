package testutils

import (
	"project/config"

	"github.com/gouniverse/userstore"
)

// SeedUser find existing or generates a new user with the given ID
func SeedUser(userID string) (userstore.UserInterface, error) {
	user, err := config.UserStore.UserFindByID(userID)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	user = userstore.NewUser().
		SetID(userID).
		SetStatus(userstore.USER_STATUS_ACTIVE)

	if userID == USER_01 {
		user.SetRole(userstore.USER_ROLE_USER)
	}

	if userID == ADMIN_01 {
		user.SetRole(userstore.USER_ROLE_ADMINISTRATOR)
	}

	err = config.UserStore.UserCreate(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
