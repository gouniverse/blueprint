package config

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"project/internal/resources"
	"strings"

	"github.com/faabiosr/cachego/file"
	"github.com/gouniverse/blindindexstore"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/envenc"
	"github.com/gouniverse/filesystem"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/metastore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/shopstore"
	"github.com/gouniverse/statsstore"
	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
	"github.com/gouniverse/vaultstore"
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

	initializeInMemoryCache()

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

	VaultKey = utils.EnvMust("VAULT_KEY")

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
	db, err := openDb(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass)

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

	blindIndexStoreEmailInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          Database.DB(),
		TableName:   "snv_bindx_email",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreEmailInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreEmailInstance is nil")
	}

	BlindIndexStoreEmail = *blindIndexStoreEmailInstance

	blindIndexStoreFirstNameInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          Database.DB(),
		TableName:   "snv_bindx_first_name",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreFirstNameInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreFirstNameInstance is nil")
	}

	BlindIndexStoreFirstName = *blindIndexStoreFirstNameInstance

	blindIndexStoreLastNameInstance, err := blindindexstore.NewStore(blindindexstore.NewStoreOptions{
		DB:          Database.DB(),
		TableName:   "snv_bindx_last_name",
		Transformer: &blindindexstore.Sha256Transformer{},
	})

	if err != nil {
		return errors.Join(errors.New("blindindexstore.NewStore"), err)
	}

	if blindIndexStoreLastNameInstance == nil {
		return errors.New("blindindexstore.NewStore: blindIndexStoreLastNameInstance is nil")
	}

	BlindIndexStoreLastName = *blindIndexStoreLastNameInstance

	blogStoreInstance, err := blogstore.NewStore(blogstore.NewStoreOptions{
		DB:            Database.DB(),
		PostTableName: "snv_blogs_post",
	})

	if err != nil {
		return errors.Join(errors.New("blogstore.NewStore"), err)
	}

	if blogStoreInstance == nil {
		return errors.New("blogstore.NewStore: blogStoreInstance is nil")
	}

	BlogStore = *blogStoreInstance

	cacheStoreInstance, err := cachestore.NewStore(cachestore.NewStoreOptions{
		DB:             db,
		CacheTableName: "snv_caches_cache",
	})

	if err != nil {
		return errors.Join(errors.New("cachestore.NewStore"), err)
	}

	if cacheStoreInstance == nil {
		return errors.New("cachestore.NewStore: cacheStoreInstance is nil")
	}

	CacheStore = *cacheStoreInstance

	cmsInstance, err := cms.NewCms(cms.Config{
		Database:        Database,
		Prefix:          "cms_",
		TemplatesEnable: true,
		PagesEnable:     true,
		MenusEnable:     true,
		BlocksEnable:    true,
		//CacheAutomigrate:    true,
		//CacheEnable:         true,
		EntitiesAutomigrate: true,
		//LogsEnable:          true,
		//LogsAutomigrate:     true,
		// SettingsEnable: true,
		//SettingsAutomigrate: true,
		//SessionAutomigrate:  true,
		//SessionEnable:       true,
		Shortcodes: []cms.ShortcodeInterface{},
		//TasksEnable:         true,
		//TasksAutomigrate:    true,
		// TranslationsEnable:  true,
		// TranslationLanguageDefault: TRANSLATION_LANGUAGE_DEFAULT,
		// TranslationLanguages:       TRANSLATION_LANGUAGE_LIST,
		// CustomEntityList:    entityList(),
	})

	if err != nil {
		return errors.Join(errors.New("cms.NewCms"), err)
	}

	if cmsInstance == nil {
		panic("cmsInstance is nil")
	}

	Cms = *cmsInstance

	customStoreInstance, err := customstore.NewStore(customstore.NewStoreOptions{
		DB:        db,
		TableName: "snv_custom_record",
	})

	if err != nil {
		return errors.Join(errors.New("customstore.NewStore"), err)
	}

	if customStoreInstance == nil {
		panic("customStoreInstance is nil")
	}

	CustomStore = *customStoreInstance

	geoStoreInstance, err := geostore.NewStore(geostore.NewStoreOptions{
		DB:                db,
		CountryTableName:  "snv_geo_country",
		StateTableName:    "snv_geo_state",
		TimezoneTableName: "snv_geo_timezone",
	})

	if err != nil {
		return errors.Join(errors.New("geostore.NewStore"), err)
	}

	if geoStoreInstance == nil {
		panic("GeoStore is nil")
	}

	GeoStore = *geoStoreInstance

	logStoreInstance, err := logstore.NewStore(logstore.NewStoreOptions{
		DB:           db,
		LogTableName: "snv_logs_log",
	})

	if err != nil {
		return errors.Join(errors.New("logstore.NewStore"), err)
	}

	if logStoreInstance == nil {
		panic("logStoreInstance is nil")
	}

	LogStore = *logStoreInstance

	metaStoreInstance, err := metastore.NewStore(metastore.NewStoreOptions{
		DB:            db,
		MetaTableName: "snv_metas_meta",
	})

	if err != nil {
		return errors.Join(errors.New("metastore.NewStore"), err)
	}

	if metaStoreInstance == nil {
		panic("MetaStore is nil")
	}

	MetaStore = *metaStoreInstance

	sessionStoreInstance, err := sessionstore.NewStore(sessionstore.NewStoreOptions{
		DB:               db,
		SessionTableName: "snv_sessions_session",
		TimeoutSeconds:   7200,
	})

	if err != nil {
		return errors.Join(errors.New("sessionstore.NewStore"), err)
	}

	if sessionStoreInstance == nil {
		panic("sessionStoreInstance is nil")
	}

	SessionStore = *sessionStoreInstance

	shopStoreInstance, err := shopstore.NewStore(shopstore.NewStoreOptions{
		DB:                     Database.DB(),
		DiscountTableName:      "snv_shop_discount",
		OrderTableName:         "snv_shop_order",
		OrderLineItemTableName: "snv_shop_order_line_item",
		ProductTableName:       "snv_shop_product",
	})

	if err != nil {
		return errors.Join(errors.New("shopstore.NewStore"), err)
	}

	if shopStoreInstance == nil {
		panic("ShopStore is nil")
	}

	ShopStore = *shopStoreInstance

	sqlFileStorageInstance, err := filesystem.NewStorage(filesystem.Disk{
		DiskName:  filesystem.DRIVER_SQL,
		Driver:    filesystem.DRIVER_SQL,
		Url:       "/files",
		DB:        db,
		TableName: "snv_media_file",
	})

	if err != nil {
		return errors.Join(errors.New("filesystem.NewStorage"), err)
	}

	if sqlFileStorageInstance == nil {
		panic("sqlFileStorageInstance is nil")
	}

	SqlFileStorage = sqlFileStorageInstance

	statsStoreInstance, err := statsstore.NewStore(statsstore.NewStoreOptions{
		VisitorTableName: "snv_stats_visitor",
		DB:               db,
	})

	if err != nil {
		return errors.Join(errors.New("statsstore.NewStore"), err)
	}

	if statsStoreInstance == nil {
		panic("StatsStore is nil")
	}

	StatsStore = *statsStoreInstance

	taskStoreInstance, err := taskstore.NewStore(taskstore.NewStoreOptions{
		DB:             db,
		TaskTableName:  "snv_tasks_task",
		QueueTableName: "snv_tasks_queue",
	})

	if err != nil {
		return errors.Join(errors.New("taskstore.NewStore"), err)
	}

	if taskStoreInstance == nil {
		panic("TaskStore is nil")
	}

	TaskStore = *taskStoreInstance

	userStoreInstance, err := userstore.NewStore(userstore.NewStoreOptions{
		DB:            db,
		UserTableName: "snv_users_user",
	})

	if err != nil {
		return errors.Join(errors.New("userstore.NewStore"), err)
	}

	if userStoreInstance == nil {
		panic("UserStore is nil")
	}

	UserStore = *userStoreInstance

	vaultStoreInstance, err := vaultstore.NewStore(vaultstore.NewStoreOptions{
		DB:             db,
		VaultTableName: "snv_vault_vault",
	})

	if err != nil {
		return errors.Join(errors.New("vaultstore.NewStore"), err)
	}

	if vaultStoreInstance == nil {
		panic("VaultStore is nil")
	}

	VaultStore = *vaultStoreInstance

	return nil
}

// initializeInMemoryCache initializes the in memory cache
func initializeInMemoryCache() {
	CacheMemory = ttlcache.New[string, any]()
	// create a new directory
	_ = os.MkdirAll(".cache", os.ModePerm)
	CacheFile = file.New(".cache")
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
	err = BlindIndexStoreEmail.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstoreemail.AutoMigrate"), err)
	}

	err = BlindIndexStoreFirstName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstorefirstname.AutoMigrate"), err)
	}

	err = BlindIndexStoreLastName.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blindindexstorelastname.AutoMigrate"), err)
	}

	err = BlogStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("blogstore.AutoMigrate"), err)
	}

	err = CacheStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("cachestore.AutoMigrate"), err)
	}

	err = CustomStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("customstore.AutoMigrate"), err)
	}

	err = GeoStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("geostore.AutoMigrate"), err)
	}

	err = LogStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("logstore.AutoMigrate"), err)
	}

	err = MetaStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("metastore.AutoMigrate"), err)
	}

	err = SessionStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("sessionstore.AutoMigrate"), err)
	}

	err = ShopStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("shopstore.AutoMigrate"), err)
	}

	err = TaskStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("taskstore.AutoMigrate"), err)
	}

	err = UserStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("userstore.AutoMigrate"), err)
	}

	err = VaultStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("vaultstore.AutoMigrate"), err)
	}

	return nil
}
