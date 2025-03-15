package config_v2

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/blindindexstore"
)

// func init() {
// 	if BlindIndexStoreUsed {
// 		addDatabaseInit(BlindIndexStoreInitialize)
// 		addDatabaseMigration(BlindIndexStoreAutoMigrate)
// 	}
// }

func (c *Config) BlindIndexStoreInitialize(db *sql.DB) error {
	if !c.BlindIndexStoreUsed {
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

	c.BlindIndexStoreEmail = blindIndexStoreEmailInstance

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

	c.BlindIndexStoreFirstName = blindIndexStoreFirstNameInstance

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

	c.BlindIndexStoreLastName = blindIndexStoreLastNameInstance

	return nil
}

func (c *Config) BlindIndexStoreAutoMigrate(_ context.Context) error {
	if !c.BlindIndexStoreUsed {
		return nil
	}

	err := c.BlindIndexStoreEmail.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	err = c.BlindIndexStoreFirstName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	err = c.BlindIndexStoreLastName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstore.AutoMigrate"), err)
	}

	return nil
}
