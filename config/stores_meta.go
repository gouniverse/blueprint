package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/metastore"
)

func init() {
	if MetaStoreUsed {
		addDatabaseInit(MetaStoreInitialize)
		addDatabaseMigration(MetaStoreAutoMigrate)
	}
}

func MetaStoreInitialize(db *sql.DB) error {
	if !MetaStoreUsed {
		return nil
	}

	metaStoreInstance, err := metastore.NewStore(metastore.NewStoreOptions{
		DB:            db,
		MetaTableName: "snv_metas_meta",
	})

	if err != nil {
		return errors.Join(errors.New("metastore.NewStore"), err)
	}

	if metaStoreInstance == nil {
		return errors.Join(errors.New("metaStoreInstance is nil"))
	}

	MetaStore = metaStoreInstance

	return nil
}

func MetaStoreAutoMigrate(_ context.Context) error {
	if !MetaStoreUsed {
		return nil
	}

	err := MetaStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("metastore.AutoMigrate"), err)
	}

	return nil
}
