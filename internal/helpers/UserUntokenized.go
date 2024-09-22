package helpers

import (
	"project/config"
	"project/pkg/userstore"
)

func UserUntokenized(authUser userstore.User) (firstName string, lastName string, email string, err error) {
	firstNameToken := authUser.FirstName()
	lastNameToken := authUser.LastName()
	emailToken := authUser.Email()

	firstName, err = config.VaultStore.TokenRead(firstNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading first name", err.Error())
		return "", "", "", err
	}

	lastName, err = config.VaultStore.TokenRead(lastNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading last name", err.Error())
		return "", "", "", err
	}

	email, err = config.VaultStore.TokenRead(emailToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading email", err.Error())
		return "", "", "", err
	}

	return firstName, lastName, email, nil
}
