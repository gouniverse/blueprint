package admin

import (
	"project/config"
	"project/pkg/userstore"
)

func untokenize(tokens []string) (values map[string]string, err error) {
	values, err = config.VaultStore.TokensRead(tokens, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading tokens", err.Error())
		return map[string]string{}, err
	}

	return values, nil
}

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

func userUntokenized(authUser userstore.User) (firstName string, lastName string, email string, err error) {
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
