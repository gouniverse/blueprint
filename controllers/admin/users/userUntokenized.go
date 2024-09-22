package admin

import (
	"project/config"
	"project/pkg/userstore"
)

func userTokenize(authUser userstore.User, firstName string, lastName string, email string) (err error) {
	firstNameToken := authUser.FirstName()
	lastNameToken := authUser.LastName()
	emailToken := authUser.Email()

	err = config.VaultStore.TokenUpdate(firstNameToken, firstName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating first name", err.Error())
		return err
	}

	err = config.VaultStore.TokenUpdate(lastNameToken, lastName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating last name", err.Error())
		return err
	}

	err = config.VaultStore.TokenUpdate(emailToken, email, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating email", err.Error())
		return err
	}

	return nil
}
