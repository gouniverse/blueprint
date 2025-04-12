package config

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"project/resources"
	"strings"

	"github.com/faabiosr/cachego/file"

	"github.com/dracory/base/database"
	"github.com/dracory/base/env"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/jellydator/ttlcache/v3"
	"github.com/samber/lo"

	// _ "github.com/go-sql-driver/mysql" // Enable MySQL driver if needed
	// _ "github.com/lib/pq" // Enable PostgreSQL driver if needed
	_ "modernc.org/sqlite" // Enable SQLite driver if needed
)

// Initialize initializes the application
//
// Business logic:
//   - initializes the environment variables
//   - initializes the database
//   - migrates the database
//   - initializes the in-memory cache
//   - initializes the logger
//
// Parameters:
// - none
//
// Returns:
// - none
func Initialize() error {
	err := initializeEnvVariables()

	if err != nil {
		return err
	}

	os.Setenv("TZ", "UTC")

	err = initializeDatabase()

	if err != nil {
		return err
	}

	err = migrateDatabase()

	if err != nil {
		return err
	}

	initializeCache()

	Logger = *slog.New(logstore.NewSlogHandler(LogStore))

	return nil
}

// initializeEnvVariables initializes the env variables
//
// Business logic:
//   - initializes the environment variables from the .env file
//   - initializes envenc variables based on the app environment
//   - checks all the required env variables
//   - panics if any of the required variable is missing
//
// Parameters:
// - none
//
// Returns:
// - none
func initializeEnvVariables() error {
	env.Initialize(".env")

	AppEnvironment = utils.Env("APP_ENV")
	AppName = utils.Env("APP_NAME")
	AppUrl = utils.Env("APP_URL")

	// Enable if you use envenc
	// if err := intializeEnvEncVariables(AppEnvironment); err != nil {
	// 	return err
	// }

	CmsUserTemplateID = utils.Env("CMS_TEMPLATE_ID")
	DbDriver = utils.Env("DB_DRIVER")
	DbHost = utils.Env("DB_HOST")
	DbPort = utils.Env("DB_PORT")
	DbName = utils.EnvMust("DB_DATABASE")
	DbUser = utils.Env("DB_USERNAME")
	DbPass = utils.Env("DB_PASSWORD")
	Debug = utils.Env("DEBUG") == "yes"
	MailDriver = utils.Env("MAIL_DRIVER")
	MailFromEmailAddress = utils.Env("EMAIL_FROM_ADDRESS")
	MailFromName = utils.Env("EMAIL_FROM_NAME")
	MailHost = utils.Env("MAIL_HOST")
	MailPassword = utils.Env("MAIL_PASSWORD")
	MailPort = utils.Env("MAIL_PORT")
	MailUsername = utils.Env("MAIL_USERNAME")
	MediaBucket = utils.Env("MEDIA_BUCKET")
	MediaDriver = utils.Env("MEDIA_DRIVER")
	MediaEndpoint = utils.Env("MEDIA_ENDPOINT")
	MediaKey = utils.Env("MEDIA_KEY")
	MediaRoot = utils.Env("MEDIA_ROOT")
	MediaSecret = utils.Env("MEDIA_SECRET")
	MediaRegion = utils.Env("MEDIA_REGION")
	MediaUrl = utils.Env("MEDIA_URL")
	OpenAiApiKey = utils.Env("OPENAI_API_KEY")
	StripeKeyPrivate = utils.Env("STRIPE_KEY_PRIVATE")
	StripeKeyPublic = utils.Env("STRIPE_KEY_PUBLIC")
	VaultKey = utils.Env("VAULT_KEY")
	VertexModelID = utils.Env("VERTEX_MODEL_ID")
	VertexProjectID = utils.Env("VERTEX_PROJECT_ID")
	VertexRegionID = utils.Env("VERTEX_REGION_ID")
	WebServerHost = utils.Env("SERVER_HOST")
	WebServerPort = utils.Env("SERVER_PORT")

	// Check required variables

	if AppEnvironment == "" {
		return errors.New("APP_ENV is required")
	}

	// Enable if you use CMS template
	// if CmsUserTemplateID == "" {
	// 	return errors.New("CMS_TEMPLATE_ID is required")
	// }

	if DbDriver == "" {
		return errors.New("DB_DRIVER is required")
	}

	if DbDriver != "sqlite" && DbHost == "" {
		return errors.New("DB_HOST is required")
	}

	if DbDriver != "sqlite" && DbPort == "" {
		return errors.New("DB_PORT is required")
	}

	if DbName == "" {
		return errors.New("DB_DATABASE is required")
	}

	if DbDriver != "sqlite" && DbUser == "" {
		return errors.New("DB_USERNAME is required")
	}

	if DbDriver != "sqlite" && DbPass == "" {
		return errors.New("DB_PASSWORD is required")
	}

	if OpenAiUsed && OpenAiApiKey == "" {
		return errors.New("OPENAI_API_KEY is required")
	}

	if StripeUsed && StripeKeyPrivate == "" {
		return errors.New("STRIPE_KEY_PRIVATE is required")
	}

	if StripeUsed && StripeKeyPublic == "" {
		return errors.New("STRIPE_KEY_PUBLIC is required")
	}

	if VaultStoreUsed && VaultKey == "" {
		return errors.New("VAULT_KEY is required")
	}

	if VertexUsed && VertexModelID == "" {
		return errors.New("VERTEX_MODEL_ID is required")
	}
	if VertexUsed && VertexProjectID == "" {
		return errors.New("VERTEX_PROJECT_ID is required")
	}
	if VertexUsed && VertexRegionID == "" {
		return errors.New("VERTEX_REGION_ID is required")
	}

	if WebServerHost == "" {
		return errors.New("SERVER_HOST is required")
	}

	if WebServerPort == "" {
		return errors.New("SERVER_PORT is required")
	}

	return nil
}

// initializeEnvEncVariables initializes the envenc variables
// based on the app environment
//
// Business logic:
//   - checkd if the app environment is testing, skipped as not needed
//   - requires the ENV_ENCRYPTION_KEY env variable
//   - looks for file the file name is .env.<app_environment>.vault
//     both in the local file system and in the resources folder
//   - if none found, it will panic
//   - if it fails for other reasons, it will panic
//
// Parameters:
// - appEnvironment: the app environment
//
// Returns:
// - none
func intializeEnvEncVariables(appEnvironment string) error {
	if appEnvironment == APP_ENVIRONMENT_TESTING {
		return nil
	}

	appEnvironment = strings.ToLower(appEnvironment)
	envEncryptionKey := utils.EnvMust("ENV_ENCRYPTION_KEY")

	vaultFilePath := ".env." + appEnvironment + ".vault"

	vaultContent := resources.Resource(".env." + appEnvironment + ".vault")

	derivedEnvEncKey, err := deriveEnvEncKey(envEncryptionKey)

	if err != nil {
		return err
	}

	err = utils.EnvEncInitialize(struct {
		Password      string
		VaultFilePath string
		VaultContent  string
	}{
		Password:      derivedEnvEncKey,
		VaultFilePath: lo.Ternary(vaultContent == "", vaultFilePath, ""),
		VaultContent:  lo.Ternary(vaultContent != "", vaultContent, ""),
	})

	if err != nil {
		return err
	}

	return nil
}

// initializeCache initializes the cache
func initializeCache() {
	CacheMemory = ttlcache.New[string, any]()
	// create a new directory
	_ = os.MkdirAll(".cache", os.ModePerm)
	CacheFile = file.New(".cache")
}

// initializeDatabase initializes the database
//
// Business logic:
//   - opens the database
//   - initializes the required stores
//
// Parameters:
// - none
//
// Returns:
// - error: the error if any
func initializeDatabase() error {
	db, err := database.Open(database.Options().
		SetDatabaseType(DbDriver).
		SetDatabaseHost(DbHost).
		SetDatabasePort(DbPort).
		SetDatabaseName(DbName).
		SetCharset(`utf8mb4`).
		SetUserName(DbUser).
		SetPassword(DbPass))

	if err != nil {
		return err
	}

	if db == nil {
		return errors.New("db is nil")
	}

	dbInstance := sb.NewDatabase(db, DbDriver)

	if dbInstance == nil {
		return errors.New("dbInstance is nil")
	}

	Database = dbInstance

	for _, init := range databaseInits {
		err = init(db)

		if err != nil {
			return err
		}
	}

	return nil
}

// migrateDatabase migrates the database
//
// Business logic:
//   - migrates the database for each store
//   - a store is only assigned if it is not nil
//
// Parameters:
// - none
//
// Returns:
// - error: the error if any
func migrateDatabase() (err error) {
	for _, migrate := range databaseMigrations {
		err = migrate(context.Background())

		if err != nil {
			return err
		}
	}

	return nil
}
