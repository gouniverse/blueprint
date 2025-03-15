package config_v2

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"project/resources"
	"strings"

	"github.com/dracory/base/env"
	"github.com/faabiosr/cachego/file"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/jellydator/ttlcache/v3"
	"github.com/samber/lo"
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
func (c *Config) initialize() error {
	err := c.initializeEnvVariables()
	if err != nil {
		return err
	}

	os.Setenv("TZ", "UTC")

	err = c.initializeDatabase()

	if err != nil {
		panic(err.Error())
	}

	err = c.migrateDatabase()

	if err != nil {
		panic(err.Error())
	}

	c.initializeCache()

	c.Logger = *slog.New(logstore.NewSlogHandler(c.LogStore))

	return nil
}

func (c *Config) initializeEnvVariables() error {
	utils.EnvInitialize(".env")

	appEnvironment, err := env.ValueOrError("APP_ENV")

	if err != nil {
		return errors.New("APP_ENV is required")
	}

	// Enable if you use envenc
	intializeEnvEncVariables(appEnvironment)

	c.AppEnvironment = env.Value("APP_ENV")
	c.AppName = env.Value("APP_NAME")
	c.AppUrl = env.Value("APP_URL")
	c.CmsUserTemplateID = env.Value("CMS_TEMPLATE_ID")
	c.DbDriver = env.Value("DB_DRIVER")
	c.DbHost = env.Value("DB_HOST")
	c.DbPort = env.Value("DB_PORT")
	c.DbName = env.Value("DB_DATABASE")
	c.DbUser = env.Value("DB_USERNAME")
	c.DbPass = env.Value("DB_PASSWORD")
	c.Debug = env.Value("DEBUG") == "yes"
	c.MailDriver = env.Value("MAIL_DRIVER")
	c.MailFromEmailAddress = env.Value("EMAIL_FROM_ADDRESS")
	c.MailFromName = env.Value("EMAIL_FROM_NAME")
	c.MailHost = env.Value("MAIL_HOST")
	c.MailPassword = env.Value("MAIL_PASSWORD")
	c.MailPort = env.Value("MAIL_PORT")
	c.MailUsername = env.Value("MAIL_USERNAME")
	c.MediaBucket = env.Value("MEDIA_BUCKET")
	c.MediaDriver = env.Value("MEDIA_DRIVER")
	c.MediaEndpoint = env.Value("MEDIA_ENDPOINT")
	c.MediaKey = env.Value("MEDIA_KEY")
	c.MediaRoot = env.Value("MEDIA_ROOT")
	c.MediaSecret = env.Value("MEDIA_SECRET")
	c.MediaRegion = env.Value("MEDIA_REGION")
	c.MediaUrl = env.Value("MEDIA_URL")
	c.OpenAiApiKey = env.Value("OPENAI_API_KEY")
	c.VaultKey = env.Value("VAULT_KEY")
	c.VertexModelID = env.Value("VERTEX_MODEL_ID")
	c.VertexProjectID = env.Value("VERTEX_PROJECT_ID")
	c.VertexRegionID = env.Value("VERTEX_REGION_ID")
	c.WebServerHost = env.Value("SERVER_HOST")
	c.WebServerPort = env.Value("SERVER_PORT")

	if c.DbDriver != "sqlite" && c.DbHost == "" {
		return errors.New("DB_HOST is required")
	}

	if c.DbDriver != "sqlite" && c.DbPort == "" {
		return errors.New("DB_PORT is required")
	}

	if c.DbDriver != "sqlite" && c.DbName == "" {
		return errors.New("DB_DATABASE is required")
	}

	if c.DbDriver != "sqlite" && c.DbUser == "" {
		return errors.New("DB_USERNAME is required")
	}

	if c.DbDriver != "sqlite" && c.DbPass == "" {
		return errors.New("DB_PASSWORD is required")
	}

	if c.VaultStoreUsed && c.VaultKey == "" {
		return errors.New("VAULT_KEY is required")
	}

	if c.WebServerHost == "" {
		return errors.New("SERVER_HOST is required")
	}

	if c.WebServerPort == "" {
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
func intializeEnvEncVariables(appEnvironment string) {
	if appEnvironment == APP_ENVIRONMENT_TESTING {
		return
	}

	appEnvironment = strings.ToLower(appEnvironment)
	envEncryptionKey := utils.EnvMust("ENV_ENCRYPTION_KEY")

	vaultFilePath := ".env." + appEnvironment + ".vault"

	vaultContent := resources.Resource(".env." + appEnvironment + ".vault")

	derivedEnvEncKey, err := deriveEnvEncKey(envEncryptionKey)

	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}
}

// if c.Database == nil {
// 	return errors.New("database is not initialized")
// }

// if c.BlindIndexStoreUsed {
// 	c.databaseInits = append(c.databaseInits, c.BlindIndexStoreInitialize)
// 	c.databaseMigrations = append(c.databaseMigrations, c.BlindIndexStoreAutoMigrate)
// }

// if len(c.databaseInits) == 0 {
// 	return nil
// }

// for _, init := range c.databaseInits {
// 	if err := init(c.Database.DB()); err != nil {
// 		return err
// 	}
// }

// initializeCache initializes the cache
func (c *Config) initializeCache() {
	c.CacheMemory = ttlcache.New[string, any]()
	// create a new directory
	_ = os.MkdirAll(".cache", os.ModePerm)
	c.CacheFile = file.New(".cache")
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
func (c *Config) initializeDatabase() error {
	db, err := database.Open(database.Options().
		SetDatabaseType(c.DbDriver).
		SetDatabaseHost(c.DbHost).
		SetDatabasePort(c.DbPort).
		SetDatabaseName(c.DbName).
		SetCharset(`utf8mb4`).
		SetUserName(c.DbUser).
		SetPassword(c.DbPass))

	if err != nil {
		return err
	}

	if db == nil {
		return errors.New("db is nil")
	}

	dbInstance := sb.NewDatabase(db, c.DbDriver)

	if dbInstance == nil {
		return errors.New("dbInstance is nil")
	}

	c.Database = dbInstance

	for _, init := range c.databaseInits {
		err = init(db)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) migrateDatabase() error {
	for _, migration := range c.databaseMigrations {
		if err := migration(context.Background()); err != nil {
			return err
		}
	}
	return nil
}
