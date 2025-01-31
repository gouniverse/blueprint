package config

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"project/internal/resources"
	"strings"

	"github.com/faabiosr/cachego/file"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/jellydator/ttlcache/v3"
	"github.com/samber/lo"

	// _ "github.com/lib/pq" // Enable PostgreSQL driver if needed
	_ "modernc.org/sqlite"
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
func Initialize() {
	initializeEnvVariables()

	os.Setenv("TZ", "UTC")

	err := initializeDatabase()

	if err != nil {
		panic(err.Error())
	}

	err = migrateDatabase()

	if err != nil {
		panic(err.Error())
	}

	initializeCache()

	Logger = *slog.New(logstore.NewSlogHandler(&LogStore))
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
func initializeEnvVariables() {
	utils.EnvInitialize(".env")

	AppEnvironment = utils.EnvMust("APP_ENV")

	// Enable if you use envenc
	// intializeEnvEncVariables(AppEnvironment)

	AppName = utils.Env("APP_NAME")
	AppUrl = utils.Env("APP_URL")

	// Enable if you use CMS template
	//CmsUserTemplateID = utils.EnvMust("CMS_TEMPLATE_ID")

	DbDriver = utils.EnvMust("DB_DRIVER")
	DbHost = lo.TernaryF(DbDriver == "sqlite", func() string {
		return utils.Env("DB_HOST")
	}, func() string {
		return utils.EnvMust("DB_HOST")
	})
	DbPort = lo.TernaryF(DbDriver == "sqlite", func() string {
		return utils.Env("DB_PORT")
	}, func() string {
		return utils.EnvMust("DB_PORT")
	})
	DbName = utils.EnvMust("DB_DATABASE")
	DbUser = lo.TernaryF(DbDriver == "sqlite", func() string {
		return utils.Env("DB_USERNAME")
	}, func() string {
		return utils.EnvMust("DB_USERNAME")
	})
	DbPass = lo.TernaryF(DbDriver == "sqlite", func() string {
		return utils.Env("DB_PASSWORD")
	}, func() string {
		return utils.EnvMust("DB_PASSWORD")
	})
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

	// Enable if you use OpenAI
	// OpenAiApiKey = utils.EnvMust("OPENAI_API_KEY")

	// Enable if you use Stripe
	// StripeKeyPrivate = utils.EnvMust("STRIPE_KEY_PRIVATE")
	// StripeKeyPublic = utils.EnvMust("STRIPE_KEY_PUBLIC")

	VaultKey = lo.TernaryF(VaultStoreUsed, func() string {
		return utils.EnvMust("VAULT_KEY")
	}, func() string {
		return utils.Env("VAULT_KEY")
	})

	// Enable if you use Vertex
	// VertexModelID = utils.EnvMust("VERTEX_MODEL_ID")
	// VertexProjectID = utils.EnvMust("VERTEX_PROJECT_ID")
	// VertexRegionID = utils.EnvMust("VERTEX_REGION_ID")

	WebServerHost = utils.EnvMust("SERVER_HOST")
	WebServerPort = utils.EnvMust("SERVER_PORT")
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

	err := utils.EnvEncInitialize(struct {
		Password      string
		VaultFilePath string
		VaultContent  string
	}{
		Password:      buildEnvEncKey(envEncryptionKey),
		VaultFilePath: lo.Ternary(vaultContent == "", vaultFilePath, ""),
		VaultContent:  lo.Ternary(vaultContent != "", vaultContent, ""),
	})

	if err != nil {
		panic(err.Error())
	}
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
