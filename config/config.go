package config

import (
	"project/pkg/userstore"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/logstore"
	"github.com/gouniverse/metastore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/webserver"
	"github.com/jellydator/ttlcache/v3"
)

type AuthenticatedUserKey struct{}
type AuthenticatedSessionKey struct{}

const APP_ENVIRONMENT_DEVELOPMENT = "development"
const APP_ENVIRONMENT_LOCAL = "local"
const APP_ENVIRONMENT_PRODUCTION = "production"
const APP_ENVIRONMENT_STAGING = "staging"
const APP_ENVIRONMENT_TESTING = "testing"
const ENV1 = "c54f"         // CHANGE
const ENV2 = "c54f"         // CHANGE
const ENV3 = "4663a0642dc6" // CHANGE

var AppEnvironment string
var AppName = "Blueprint"
var AppUrl string
var AppVersion string
var AuthEndpoint = "/auth"
var Database *sb.Database
var DbDriver string
var DbHost string
var DbName string
var DbPass string
var DbPort string
var DbUser string
var Debug bool
var MailDriver string
var MailFromEmailAddress string
var MailFromName string
var MailHost string
var MailPort string
var MailPassword string
var MailUsername string
var MediaBucket string
var MediaDriver string
var MediaKey string
var MediaEndpoint string
var MediaRegion string
var MediaRoot string = "/"
var MediaSecret string
var MediaUrl string = "/files"
var OpenAiApiKey string
var StripeKeyPrivate string
var StripeKeyPublic string
var WebServer *webserver.Server
var WebServerHost string
var WebServerPort string

// InMem is an in-memory cache.
var InMem *ttlcache.Cache[string, any]

// Cms is the CMS instance.
var Cms *cms.Cms

// CmsUserTemplateID is the CMS user template ID.
var CmsUserTemplateID string = ""

// ===================================== //
var BlogStore *blogstore.Store
var CacheStore *cachestore.Store

// var CommentStore *commentstore.Store
var CustomStore *customstore.Store
var GeoStore *geostore.Store
var LogStore *logstore.Store
var MetaStore *metastore.Store
var SessionStore *sessionstore.Store

// var SubscriptionStore *subscriptionstore.Store
var TaskStore *taskstore.Store
var UserStore *userstore.Store

func init() {
	AppVersion = "0.0.1" // default
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
