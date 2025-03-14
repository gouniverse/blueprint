package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/blindindexstore"
)

func init() {
	if BlindIndexStoreUsed {
		addDatabaseInit(BlindIndexStoreInitialize)
		addDatabaseMigration(BlindIndexStoreAutoMigrate)
	}
}

func BlindIndexStoreInitialize(db *sql.DB) error {
	if !BlindIndexStoreUsed {
		return nil
	}

	blindIndexStoreEmailInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          db,
		TableName:   "snv_bindx_email",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreEmailInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreEmailInstance is nil")
	}

	BlindIndexStoreEmail = blindIndexStoreEmailInstance

	blindIndexStoreFirstNameInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          db,
		TableName:   "snv_bindx_first_name",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreFirstNameInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreFirstNameInstance is nil")
	}

	BlindIndexStoreFirstName = blindIndexStoreFirstNameInstance

	blindIndexStoreLastNameInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          db,
		TableName:   "snv_bindx_last_name",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreLastNameInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreLastNameInstance is nil")
	}

	BlindIndexStoreLastName = blindIndexStoreLastNameInstance

	return nil
}

func BlindIndexStoreAutoMigrate(_ context.Context) error {
	if !BlindIndexStoreUsed {
		return nil
	}

	err := BlindIndexStoreEmail.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	err = BlindIndexStoreFirstName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	err = BlindIndexStoreLastName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	return nil
}
