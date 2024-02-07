package config

import (
	"log"
	"net/http"
	"os"
	"project/pkg/userstore"

	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/metastore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/utils"
)

func Initialize() {
	// log.Println("1. Initializing environment variables...")
	utils.EnvInitialize()

	ServerHost = utils.Env("SERVER_HOST")
	ServerPort = utils.Env("SERVER_PORT")
	AppName = utils.Env("APP_NAME")
	Debug = utils.Env("DEBUG") == "yes"
	AppUrl = utils.Env("APP_URL")
	AppEnvironment = utils.Env("APP_ENV")
	DbDriver = utils.Env("DB_DRIVER")
	DbHost = utils.Env("DB_HOST")
	DbPort = utils.Env("DB_PORT")
	DbName = utils.Env("DB_DATABASE")
	DbUser = utils.Env("DB_USERNAME")
	DbPass = utils.Env("DB_PASSWORD")
	MailDriver = utils.Env("MAIL_DRIVER")
	MailHost = utils.Env("MAIL_HOST")
	MailPort = utils.Env("MAIL_PORT")
	MailUsername = utils.Env("MAIL_USERNAME")
	MailPassword = utils.Env("MAIL_PASSWORD")
	MailFromEmailAddress = utils.Env("EMAIL_FROM_ADDRESS")
	MailFromName = utils.Env("EMAIL_FROM_NAME")

	MediaDriver = utils.Env("MEDIA_DRIVER")
	MediaKey = utils.Env("MEDIA_KEY")
	MediaSecret = utils.Env("MEDIA_SECRET")
	MediaEndpoint = utils.Env("MEDIA_ENDPOINT")
	MediaRegion = utils.Env("MEDIA_REGION")
	MediaBucket = utils.Env("MEDIA_BUCKET")
	MediaUrl = utils.Env("MEDIA_URL")

	StripeKeyPrivate = utils.Env("STRIPE_KEY_PRIVATE")
	StripeKeyPublic = utils.Env("STRIPE_KEY_PUBLIC")

	debug := utils.Env("DEBUG")

	if debug == "yes" {
		Debug = true
	}

	if ServerHost == "" {
		panic("Environment variable SERVER_HOST is required")
	}

	if ServerPort == "" {
		panic("Environment variable SERVER_PORT is required")
	}

	if AppEnvironment == "" {
		panic("Environment variable APP_ENV is required")
	}

	if DbDriver == "" {
		panic("Environment variable DB_DRIVER is required")
	}

	if DbDriver != "sqlite" && DbHost == "" {
		panic("Environment variable DB_HOST is required")
	}

	if DbDriver != "sqlite" && DbPort == "" {
		panic("Environment variable DB_PORT is required")
	}

	if DbName == "" {
		panic("Environment variable DB_DATABASE is required")
	}

	if DbDriver != "sqlite" && DbUser == "" {
		panic("Environment variable DB_USERNAME is required")
	}

	if DbDriver != "sqlite" && DbPass == "" {
		panic("Environment variable DB_PASSWORD is required")
	}

	// Enable if you use Stripe
	// if StripeKeyPrivate == "" {
	// 	panic("Environment variable STRIPE_KEY_PRIVATE is required")
	// }

	// Enable if you use Stripe
	// if StripeKeyPublic == "" {
	// 	panic("Environment variable STRIPE_KEY_PUBLIC is required")
	// }

	os.Setenv("TZ", "UTC")

	db, err := openDb(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass)

	if err != nil {
		log.Fatal(err.Error())
	}

	Database = sb.NewDatabase(db, DbDriver)

	err = initializeDatabase()

	if err != nil {
		panic(err.Error())
	}

	err = migrateDatabase()

	if err != nil {
		panic(err.Error())
	}

	var errCms error
	Cms, errCms = cms.NewCms(cms.Config{
		Database:        sb.NewDatabase(db, DbDriver),
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
		SettingsEnable: true,
		//SettingsAutomigrate: true,
		//SessionAutomigrate:  true,
		//SessionEnable:       true,
		Shortcodes: map[string]func(*http.Request, string, map[string]string) string{},
		//TasksEnable:         true,
		//TasksAutomigrate:    true,
		// TranslationsEnable:  true,
		// TranslationLanguageDefault: TRANSLATION_LANGUAGE_DEFAULT,
		// TranslationLanguages:       TRANSLATION_LANGUAGE_LIST,
		// CustomEntityList:    entityList(),
	})

	if errCms != nil {
		panic(errCms.Error())
	}
}

func initializeDatabase() error {
	db, err := openDb(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass)

	if err != nil {
		return err
	}

	Database = sb.NewDatabase(db, DbDriver)

	CacheStore, err = cachestore.NewStore(cachestore.NewStoreOptions{
		DB:             db,
		CacheTableName: "snv_caches_cache",
	})

	CustomStore, err = customstore.NewStore(customstore.NewStoreOptions{
		DB:        db,
		TableName: "snv_custom_record",
	})

	if err != nil {
		return err
	}

	GeoStore, err = geostore.NewStore(geostore.NewStoreOptions{
		DB:                db,
		CountryTableName:  "snv_geo_country",
		TimezoneTableName: "snv_geo_timezone",
	})

	if err != nil {
		return err
	}

	LogStore, err = logstore.NewStore(logstore.NewStoreOptions{
		DB:           db,
		LogTableName: "snv_logs_log",
	})

	if err != nil {
		return err
	}

	MetaStore, err = metastore.NewStore(metastore.NewStoreOptions{
		DB:            db,
		MetaTableName: "snv_metas_meta",
	})

	if err != nil {
		return err
	}

	SessionStore, err = sessionstore.NewStore(sessionstore.NewStoreOptions{
		DB:               db,
		SessionTableName: "snv_sessions_session",
	})

	if err != nil {
		return err
	}

	TaskStore, err = taskstore.NewStore(taskstore.NewStoreOptions{
		DB:             db,
		TaskTableName:  "snv_tasks_task",
		QueueTableName: "snv_tasks_queue",
	})

	// UserStore, err = entitystore.NewStore(entitystore.NewStoreOptions{
	//	DB:                      db,
	//	EntityTableName:         "user_entity",
	//	EntityTrashTableName:    "user_entity_trash",
	//	AttributeTableName:      "user_attribute",
	//	AttributeTrashTableName: "user_attribute_trash",
	// })
	UserStore, err = userstore.NewStore(userstore.NewStoreOptions{
		DB:            db,
		UserTableName: "snv_users_user",
	})

	if err != nil {
		return err
	}

	return nil
}

func migrateDatabase() (err error) {
	err = CacheStore.AutoMigrate()

	if err != nil {
		return err
	}

	err = CustomStore.AutoMigrate()

	if err != nil {
		return err
	}

	err = GeoStore.AutoMigrate()

	if err != nil {
		return err
	}

	err = LogStore.AutoMigrate()

	if err != nil {
		return err
	}

	MetaStore.AutoMigrate()

	err = SessionStore.AutoMigrate()

	if err != nil {
		return err
	}

	err = TaskStore.AutoMigrate()

	if err != nil {
		return err
	}

	err = UserStore.AutoMigrate()

	if err != nil {
		return err
	}

	return nil
}
