package config

import (
	"project/internal/server"

	"github.com/gouniverse/cms"
	"github.com/gouniverse/entitystore"
	"github.com/gouniverse/sql"
)

type AuthenticatedUserKey struct{}

const APP_ENVIRONMENT_DEVELOPMENT = "development"
const APP_ENVIRONMENT_LOCAL = "local"
const APP_ENVIRONMENT_PRODUCTION = "production"
const APP_ENVIRONMENT_STAGING = "staging"
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
var StripeKeyPublic string
var StripeKeyPrivate string
var WebServer *server.Server

var Cms *cms.Cms
var UserStore *entitystore.Store

func init() {
	AppVersion = "0.0.1" // default
	Debug = false        // default
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
