package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/customstore"
)

func init() {
	if CustomStoreUsed {
		addDatabaseInit(CustomStoreInitialize)
		addDatabaseMigration(CustomStoreAutoMigrate)
	}
}

func CustomStoreInitialize(db *sql.DB) error {
	if !CustomStoreUsed {
		return nil
	}

	customStoreInstance, err := customstore.NewStore(customstore.NewStoreOptions{
		DB:        db,
		TableName: "snv_custom_record",
	})

	if err != nil {
		return errors.Join(errors.New("customstore.NewStore"), err)
	}

	if customStoreInstance == nil {
		return errors.Join(errors.New("customStoreInstance is nil"))
	}

	CustomStore = *customStoreInstance

	return nil
}

func CustomStoreAutoMigrate(_ context.Context) error {
	if !CustomStoreUsed {
		return nil
	}

	err := CustomStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("customstore.AutoMigrate"), err)
	}

	return nil
}
