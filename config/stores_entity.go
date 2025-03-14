package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/entitystore"
)

func init() {
	if EntityStoreUsed {
		addDatabaseInit(EntityStoreInitialize)
		addDatabaseMigration(EntityStoreAutoMigrate)
	}
}

func EntityStoreInitialize(db *sql.DB) error {
	if !EntityStoreUsed {
		return nil
	}

	entityStoreInstance, err := entitystore.NewStore(entitystore.NewStoreOptions{
		DB:                      db,
		EntityTableName:         "snv_entities_entity",
		EntityTrashTableName:    "snv_entities_entity_trash",
		AttributeTableName:      "snv_entities_attribute",
		AttributeTrashTableName: "snv_entities_attribute_trash",
	})

	if err != nil {
		return errors.Join(errors.New("entitystore.NewStore"), err)
	}

	if entityStoreInstance == nil {
		return errors.Join(errors.New("entityStoreInstance is nil"))
	}

	EntityStore = entityStoreInstance

	return nil
}

func EntityStoreAutoMigrate(_ context.Context) error {
	if !EntityStoreUsed {
		return nil
	}

	err := EntityStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("entitystore.AutoMigrate"), err)
	}

	return nil
}
