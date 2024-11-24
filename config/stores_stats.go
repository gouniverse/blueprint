package config

import (
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

func StatsStoreAutoMigrate() error {
	if !StatsStoreUsed {
		return nil
	}

	err := StatsStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("statsstore.AutoMigrate"), err)
	}

	return nil
}
