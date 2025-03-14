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

func IsEnvDevelopment() bool {
	return AppEnvironment == APP_ENVIRONMENT_DEVELOPMENT
}

func IsEnvLocal() bool {
	return AppEnvironment == APP_ENVIRONMENT_LOCAL
}

func IsEnvProduction() bool {
	return AppEnvironment == APP_ENVIRONMENT_PRODUCTION
}

func IsEnvStaging() bool {
	return AppEnvironment == APP_ENVIRONMENT_STAGING
}

func IsEnvTesting() bool {
	return AppEnvironment == APP_ENVIRONMENT_TESTING
}

func IsDebugEnabled() bool {
	return Debug
}
