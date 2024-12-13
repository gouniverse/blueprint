package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/cachestore"
)

func init() {
	if CacheStoreUsed {
		addDatabaseInit(CacheStoreInitialize)
		addDatabaseMigration(CacheStoreAutoMigrate)
	}
}

func CacheStoreInitialize(db *sql.DB) error {
	if !CacheStoreUsed {
		return nil
	}

	cacheStoreInstance, err := cachestore.NewStore(cachestore.NewStoreOptions{
		DB:             db,
		CacheTableName: "snv_caches_cache",
	})

	if err != nil {
		return errors.Join(errors.New("cachestore.NewStore"), err)
	}

	if cacheStoreInstance == nil {
		return errors.New("cachestore.NewStore: cacheStoreInstance is nil")
	}

	CacheStore = *cacheStoreInstance

	return nil
}

func CacheStoreAutoMigrate(_ context.Context) error {
	if !CacheStoreUsed {
		return nil
	}

	err := CacheStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("cachestore.AutoMigrate"), err)
	}

	return nil
}
