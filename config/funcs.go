package config

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"

	"github.com/gouniverse/envenc"
)

var databaseInits = []func(db *sql.DB) error{}
var databaseMigrations = []func(ctx context.Context) error{}

func addDatabaseInit(init func(db *sql.DB) error) {
	databaseInits = append(databaseInits, init)
}

func addDatabaseMigration(migration func(ctx context.Context) error) {
	databaseMigrations = append(databaseMigrations, migration)
}

// buildEnvEncKey builds the envenc key
//
// Business logic:
//   - deobfuscates the salt
//   - creates the temp key based on the salt and key
//   - hashes the temp key
//   - returns the hash
//
// Parameters:
// - envEncryptionKey: the env encryption key
//
// Returns:
// - string: the final key
func buildEnvEncKey(envEncryptionKey string) string {
	envEncryptionSalt, _ := envenc.Deobfuscate(ENV_ENCRYPTION_SALT)
	tempKey := envEncryptionSalt + envEncryptionKey

	hash := sha256.Sum256([]byte(tempKey))
	realKey := fmt.Sprintf("%x", hash)

	return realKey
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
