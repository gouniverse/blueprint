package config

import (
	"context"
	"database/sql"
)

var databaseInits = []func(db *sql.DB) error{}
var databaseMigrations = []func(ctx context.Context) error{}

func addDatabaseInit(init func(db *sql.DB) error) {
	databaseInits = append(databaseInits, init)
}

func addDatabaseMigration(migration func(ctx context.Context) error) {
	databaseMigrations = append(databaseMigrations, migration)
}
