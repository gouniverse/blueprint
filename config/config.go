package config

import (
	// "database/sql"

	"project/internal/server"

	"github.com/gouniverse/sql"
)

const APP_ENVIRONMENT_DEVELOPMENT = "development"
const APP_ENVIRONMENT_PRODUCTION = "production"
const APP_ENVIRONMENT_TESTING = "testing"

var AppName = "Blueprint"
var AppUrl string
var AppEnvironment string
var AuthEndpoint = "/auth"
var ServerHost string
var ServerPort string
var AppVersion string
var DbDriver string
var DbHost string
var DbPort string
var DbName string
var DbUser string
var DbPass string
var Database *sql.Database
var Debug = false
var MailDriver string
var MailHost string
var MailPort string
var MailUsername string
var MailPassword string
var MailFromEmailAddress string
var MailFromName string
var MediaDriver string
var MediaKey string
var MediaSecret string
var MediaEndpoint string
var MediaRegion string
var MediaBucket string
var MediaUrl string
var WebServer *server.Server

func init() {
	AppVersion = "0.0.1" // default
	Debug = false        // default
}

func IsEnvTesting() bool {
	return AppEnvironment == APP_ENVIRONMENT_TESTING
}

func IsEnvDevelopment() bool {
	return AppEnvironment == APP_ENVIRONMENT_DEVELOPMENT
}

func IsEnvProduction() bool {
	return AppEnvironment == APP_ENVIRONMENT_PRODUCTION
}

func IsDebugEnabled() bool {
	return Debug
}
