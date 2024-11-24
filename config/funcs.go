package config

import "database/sql"

var databaseInits = []func(db *sql.DB) error{}
var databaseMigrations = []func() error{}

func addDatabaseInit(init func(db *sql.DB) error) {
	databaseInits = append(databaseInits, init)
}

func addDatabaseMigration(migration func() error) {
	databaseMigrations = append(databaseMigrations, migration)
}
