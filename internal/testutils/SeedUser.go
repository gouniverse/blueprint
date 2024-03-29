package testutils

import (
	"project/config"
	"project/pkg/userstore"
)

func SeedUser(userID string) (*userstore.User, error) {
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
