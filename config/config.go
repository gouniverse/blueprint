package config

import (
	"project/internal/server"
	"project/pkg/userstore"

	// "github.com/gouniverse/cms"
	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/taskstore"

	// "github.com/gouniverse/entitystore"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/metastore"
	"github.com/gouniverse/sb"
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
var Database *sb.Database
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
var CacheStore *cachestore.Store
var CustomStore *customstore.Store
var GeoStore *geostore.Store
var LogStore *logstore.Store
var MetaStore *metastore.Store
var SessionStore *sessionstore.Store
var TaskStore *taskstore.Store
var UserStore *userstore.Store

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
