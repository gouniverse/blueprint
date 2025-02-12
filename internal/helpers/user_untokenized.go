package helpers

import (
	"context"
	"errors"
	"project/config"

	"github.com/gouniverse/userstore"
)

func UserUntokenized(ctx context.Context, authUser userstore.UserInterface) (firstName string, lastName string, email string, err error) {
	if config.VaultStore == nil {
		return "", "", "", errors.New("vaultstore is nil")
	}

	firstNameToken := authUser.FirstName()
	lastNameToken := authUser.LastName()
	emailToken := authUser.Email()

	firstName, err = config.VaultStore.TokenRead(ctx, firstNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading first name", err.Error())
		return "", "", "", err
	}

	lastName, err = config.VaultStore.TokenRead(ctx, lastNameToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading last name", err.Error())
		return "", "", "", err
	}

	email, err = config.VaultStore.TokenRead(ctx, emailToken, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading email", err.Error())
		return "", "", "", err
	}

	return firstName, lastName, email, nil
}
