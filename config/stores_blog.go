package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/blogstore"
)

func init() {
	if BlogStoreUsed {
		addDatabaseInit(BlogStoreInitialize)
		addDatabaseMigration(BlogStoreAutoMigrate)
	}
}

func BlogStoreInitialize(db *sql.DB) error {
	if !BlogStoreUsed {
		return nil
	}

	blogStoreInstance, err := blogstore.NewStore(blogstore.NewStoreOptions{
		DB:            db,
		PostTableName: "snv_blogs_post",
	})

	if err != nil {
		return errors.Join(errors.New("blogstore.NewStore"), err)
	}

	if blogStoreInstance == nil {
		return errors.New("blogstore.NewStore: blogStoreInstance is nil")
	}

	BlogStore = *blogStoreInstance

	return nil
}

func BlogStoreAutoMigrate(_ context.Context) error {
	if !BlogStoreUsed {
		return nil
	}

	err := BlogStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blogstore.AutoMigrate"), err)
	}

	return nil
}
