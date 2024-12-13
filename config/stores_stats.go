package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/statsstore"
)

func init() {
	if StatsStoreUsed {
		addDatabaseInit(StatsStoreInitialize)
		addDatabaseMigration(StatsStoreAutoMigrate)
	}
}

func StatsStoreInitialize(db *sql.DB) error {
	if !StatsStoreUsed {
		return nil
	}

	statsStoreInstance, err := statsstore.NewStore(statsstore.NewStoreOptions{
		VisitorTableName: "snv_stats_visitor",
		DB:               db,
	})

	if err != nil {
		return errors.Join(errors.New("statsstore.NewStore"), err)
	}

	if statsStoreInstance == nil {
		return errors.Join(errors.New("statsStoreInstance is nil"))
	}

	StatsStore = statsStoreInstance

	return nil
}

func StatsStoreAutoMigrate(_ context.Context) error {
	if !StatsStoreUsed {
		return nil
	}

	if StatsStore == nil {
		return errors.New("statsstore.AutoMigrate: StatsStore is nil")
	}

	err := StatsStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("statsstore.AutoMigrate"), err)
	}

	return nil
}
