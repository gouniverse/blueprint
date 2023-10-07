package config

import (
	"log"
	"net/http"

	"github.com/gouniverse/cms"
	"github.com/gouniverse/sql"
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

	db, err := openDb(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass)
	Database = sql.NewDatabase(db, DbDriver)

	if err != nil {
		log.Fatal(err.Error())
	}

	var errCms error
	Cms, errCms = cms.NewCms(cms.Config{
		Database:            Database,
		Prefix:              "cms_",
		TemplatesEnable:     true,
		PagesEnable:         true,
		MenusEnable:         true,
		BlocksEnable:        true,
		CacheAutomigrate:    true,
		CacheEnable:         true,
		EntitiesAutomigrate: true,
		LogsEnable:          true,
		LogsAutomigrate:     true,
		SettingsEnable:      true,
		SessionAutomigrate:  true,
		SessionEnable:       true,
		Shortcodes:          map[string]func(*http.Request, string, map[string]string) string{},
		TranslationsEnable:  true,
		// TranslationLanguageDefault: TRANSLATION_LANGUAGE_DEFAULT,
		// TranslationLanguages:       TRANSLATION_LANGUAGE_LIST,
		// CustomEntityList:    entityList(),
	})

	if errCms != nil {
		panic(errCms.Error())
	}
}
