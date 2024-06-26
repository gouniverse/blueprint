package config

import (
	"errors"
	"os"
	"project/pkg/userstore"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/filesystem"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/metastore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/shopstore"
	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/utils"
	"github.com/jellydator/ttlcache/v3"
)

func Initialize() {
	utils.EnvInitialize()
	utils.EnvEncInitialize(ENV1 + ENV2 + ENV3)

	AppName = utils.Env("APP_NAME")
	AppUrl = utils.Env("APP_URL")
	AppEnvironment = utils.Env("APP_ENV")
	DbDriver = utils.Env("DB_DRIVER")
	DbHost = utils.Env("DB_HOST")
	DbPort = utils.Env("DB_PORT")
	DbName = utils.Env("DB_DATABASE")
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
	WebServerHost = utils.Env("SERVER_HOST")
	WebServerPort = utils.Env("SERVER_PORT")

	CmsUserTemplateID = utils.Env("CMS_TEMPLATE_ID")

	debug := utils.Env("DEBUG")

	if debug == "yes" {
		Debug = true
	}

	if WebServerHost == "" {
		panic("Environment variable SERVER_HOST is required")
	}

	if WebServerPort == "" {
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

	err := initializeDatabase()

	if err != nil {
		panic(err.Error())
	}

	err = migrateDatabase()

	if err != nil {
		panic(err.Error())
	}

	initializeInMemoryCache()
}

func initializeDatabase() error {
	db, err := openDb(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass)

	if err != nil {
		return err
	}

	Database = sb.NewDatabase(db, DbDriver)

	BlogStore, err = blogstore.NewStore(blogstore.NewStoreOptions{
		DB:            Database.DB(),
		PostTableName: "snv_blogs_post",
	})

	if err != nil {
		return errors.Join(errors.New("blogstore.NewStore"), err)
	}

	CacheStore, err = cachestore.NewStore(cachestore.NewStoreOptions{
		DB:             db,
		CacheTableName: "snv_caches_cache",
	})

	if err != nil {
		return errors.Join(errors.New("cachestore.NewStore"), err)
	}

	Cms, err = cms.NewCms(cms.Config{
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
		SettingsEnable: true,
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

	CustomStore, err = customstore.NewStore(customstore.NewStoreOptions{
		DB:        db,
		TableName: "snv_custom_record",
	})

	if err != nil {
		return errors.Join(errors.New("customstore.NewStore"), err)
	}

	GeoStore, err = geostore.NewStore(geostore.NewStoreOptions{
		DB:                db,
		CountryTableName:  "snv_geo_country",
		StateTableName:    "snv_geo_state",
		TimezoneTableName: "snv_geo_timezone",
	})

	if err != nil {
		return errors.Join(errors.New("geostore.NewStore"), err)
	}

	if GeoStore == nil {
		panic("GeoStore is nil")
	}

	LogStore, err = logstore.NewStore(logstore.NewStoreOptions{
		DB:           db,
		LogTableName: "snv_logs_log",
	})

	if err != nil {
		return errors.Join(errors.New("logstore.NewStore"), err)
	}

	MetaStore, err = metastore.NewStore(metastore.NewStoreOptions{
		DB:            db,
		MetaTableName: "snv_metas_meta",
	})

	if err != nil {
		return errors.Join(errors.New("metastore.NewStore"), err)
	}

	if MetaStore == nil {
		panic("MetaStore is nil")
	}

	SessionStore, err = sessionstore.NewStore(sessionstore.NewStoreOptions{
		DB:               db,
		SessionTableName: "snv_sessions_session",
		TimeoutSeconds:   7200,
	})

	if err != nil {
		return errors.Join(errors.New("sessionstore.NewStore"), err)
	}

	ShopStore, err = shopstore.NewStore(shopstore.NewStoreOptions{
		DB:                Database.DB(),
		DiscountTableName: "snv_shop_discount",
		OrderTableName:    "snv_shop_order",
	})

	if err != nil {
		return errors.Join(errors.New("shopstore.NewStore"), err)
	}

	if ShopStore == nil {
		panic("ShopStore is nil")
	}

	SqlFileStorage, err = filesystem.NewStorage(filesystem.Disk{
		DiskName:  filesystem.DRIVER_SQL,
		Driver:    filesystem.DRIVER_SQL,
		Url:       "/file",
		DB:        db,
		TableName: "snv_media_file",
	})

	if err != nil {
		return errors.Join(errors.New("filesystem.NewStorage"), err)
	}

	TaskStore, err = taskstore.NewStore(taskstore.NewStoreOptions{
		DB:             db,
		TaskTableName:  "snv_tasks_task",
		QueueTableName: "snv_tasks_queue",
	})

	if err != nil {
		return errors.Join(errors.New("taskstore.NewStore"), err)
	}

	if TaskStore == nil {
		panic("TaskStore is nil")
	}

	UserStore, err = userstore.NewStore(userstore.NewStoreOptions{
		DB:            db,
		UserTableName: "snv_users_user",
	})

	if err != nil {
		return errors.Join(errors.New("userstore.NewStore"), err)
	}

	if UserStore == nil {
		panic("UserStore is nil")
	}

	return nil
}

func initializeInMemoryCache() {
	InMem = ttlcache.New[string, any]()
}

func migrateDatabase() (err error) {
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

	return nil
}
