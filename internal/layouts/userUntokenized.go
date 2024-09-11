package layouts

import (
	"project/config"
	"project/pkg/userstore"
)

func userUntokenized(authUser userstore.User) (firstName string, lastName string, err error) {
	firstNameToken := authUser.FirstName()
	lastNameToken := authUser.LastName()
	emailToken := authUser.Email()

	firstName, err = config.VaultStore.TokenRead(firstNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading first name", err.Error())
		return "", "", err
	}

	lastName, err = config.VaultStore.TokenRead(lastNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading last name", err.Error())
		return "", "", err
	}

	if firstName != "" {
		return firstName, lastName, nil
	}

	email, err := config.VaultStore.TokenRead(emailToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading email", err.Error())
		return "", "", err
	}

	if firstName == "" {
		firstName = email
	}

	return firstName, lastName, nil
}
