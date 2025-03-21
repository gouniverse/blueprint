package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/sessionstore"
)

func init() {
	if SessionStoreUsed {
		addDatabaseInit(SessionStoreInitialize)
		addDatabaseMigration(SessionStoreAutoMigrate)
	}
}

func SessionStoreInitialize(db *sql.DB) error {
	if !SessionStoreUsed {
		return nil
	}

	sessionStoreInstance, err := sessionstore.NewStore(sessionstore.NewStoreOptions{
		DB:               db,
		SessionTableName: "snv_sessions_session",
		TimeoutSeconds:   7200,
	})

	if err != nil {
		return errors.Join(errors.New("sessionstore.NewStore"), err)
	}

	if sessionStoreInstance == nil {
		return errors.Join(errors.New("sessionStoreInstance is nil"))
	}

	SessionStore = sessionStoreInstance

	return nil
}

func SessionStoreAutoMigrate(ctx context.Context) error {
	if !SessionStoreUsed {
		return nil
	}

	if SessionStore == nil {
		return errors.New("sessionstore.AutoMigrate: SessionStore is nil")
	}

	err := SessionStore.AutoMigrate(ctx)

	if err != nil {
		return errors.Join(errors.New("sessionstore.AutoMigrate"), err)
	}

	return nil
}
