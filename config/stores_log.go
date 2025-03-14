package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/logstore"
)

func init() {
	if LogStoreUsed {
		addDatabaseInit(LogStoreInitialize)
		addDatabaseMigration(LogStoreAutoMigrate)
	}
}

func LogStoreInitialize(db *sql.DB) error {
	if !LogStoreUsed {
		return nil
	}

	logStoreInstance, err := logstore.NewStore(logstore.NewStoreOptions{
		DB:           db,
		LogTableName: "snv_logs_log",
	})

	if err != nil {
		return errors.Join(errors.New("logstore.NewStore"), err)
	}

	if logStoreInstance == nil {
		return errors.Join(errors.New("logStoreInstance is nil"))
	}

	LogStore = logStoreInstance

	return nil
}

func LogStoreAutoMigrate(_ context.Context) error {
	if !LogStoreUsed {
		return nil
	}

	err := LogStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("logstore.AutoMigrate"), err)
	}

	return nil
}
