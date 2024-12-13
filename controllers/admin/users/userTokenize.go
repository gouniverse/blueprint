package admin

import (
	"context"
	"project/config"

	"github.com/gouniverse/userstore"
)

func userTokenize(user userstore.UserInterface, firstName string, lastName string, email string) (err error) {
	firstNameToken := user.FirstName()
	lastNameToken := user.LastName()
	emailToken := user.Email()

	ctx := context.Background()

	err = config.VaultStore.TokenUpdate(ctx, firstNameToken, firstName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating first name", err.Error())
		return err
	}

	err = config.VaultStore.TokenUpdate(ctx, lastNameToken, lastName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating last name", err.Error())
		return err
	}

	err = config.VaultStore.TokenUpdate(ctx, emailToken, email, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating email", err.Error())
		return err
	}

	return nil
}
