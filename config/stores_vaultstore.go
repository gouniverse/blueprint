package config

import (
	"database/sql"
	"errors"

	"github.com/gouniverse/vaultstore"
)

func init() {
	if VaultStoreUsed {
		addDatabaseInit(VaultStoreInitialize)
		addDatabaseMigration(VaultStoreAutoMigrate)
	}
}

func VaultStoreInitialize(db *sql.DB) error {
	if !VaultStoreUsed {
		return nil
	}

	vaultStoreInstance, err := vaultstore.NewStore(vaultstore.NewStoreOptions{
		DB:             db,
		VaultTableName: "snv_vault_vault",
	})

	if err != nil {
		return errors.Join(errors.New("vaultstore.NewStore"), err)
	}

	if vaultStoreInstance == nil {
		panic("VaultStore is nil")
	}

	VaultStore = *vaultStoreInstance

	return nil
}

func VaultStoreAutoMigrate() error {
	if !VaultStoreUsed {
		return nil
	}

	err := VaultStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("vaultstore.AutoMigrate"), err)
	}

	return nil
}
