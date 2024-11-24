package config

import (
	"database/sql"
	"errors"

	"github.com/gouniverse/userstore"
)

func init() {
	if UserStoreUsed {
		addDatabaseInit(UserStoreInitialize)
		addDatabaseMigration(UserStoreAutoMigrate)
	}
}

func UserStoreInitialize(db *sql.DB) error {
	if !UserStoreUsed {
		return nil
	}

	userStoreInstance, err := userstore.NewStore(userstore.NewStoreOptions{
		DB:            db,
		UserTableName: "snv_users_user",
	})

	if err != nil {
		return errors.Join(errors.New("userstore.NewStore"), err)
	}

	if userStoreInstance == nil {
		panic("UserStore is nil")
	}

	UserStore = userStoreInstance

	return nil
}

func UserStoreAutoMigrate() error {
	if !UserStoreUsed {
		return nil
	}

	err := UserStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("userstore.AutoMigrate"), err)
	}

	return nil
}
