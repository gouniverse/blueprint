package config

import (
	"log/slog"

	"github.com/faabiosr/cachego"
	"github.com/gouniverse/blindindexstore"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/cachestore"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/cmsstore"
	"github.com/gouniverse/customstore"
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
	"github.com/gouniverse/vaultstore"
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
const ENV_ENCRYPTION_SALT = "YOUR_OBFUSCATED_SALT"

var AppEnvironment string
var AppName string
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
var TranslationLanguageDefault string = "en"
var TranslationLanguageList map[string]string = map[string]string{"en": "English", "bg": "Bulgarian", "de": "German"}
var VaultKey string
var VertexModelID string
var VertexProjectID string
var VertexRegionID string
var WebServer *webserver.Server
var WebServerHost string
var WebServerPort string

var CacheMemory *ttlcache.Cache[string, any]
var CacheFile cachego.Cache

// Cms is the old CMS package (replaced by CmsStore).
var CmsUsed = true
var Cms cms.Cms

// CmsUserTemplateID is the CMS user template ID.
var CmsUserTemplateID string

// ===================================== //
var BlindIndexStoreUsed = true
var BlindIndexStoreEmail blindindexstore.Store
var BlindIndexStoreFirstName blindindexstore.Store
var BlindIndexStoreLastName blindindexstore.Store

var BlogStoreUsed = true
var BlogStore blogstore.Store

var CmsStoreUsed = true
var CmsStore cmsstore.StoreInterface

var CacheStoreUsed = true
var CacheStore cachestore.Store

// var CommentStore *commentstore.Store

var CustomStoreUsed = false
var CustomStore customstore.Store

var GeoStoreUsed = true
var GeoStore geostore.Store

var LogStoreUsed = true
var LogStore logstore.Store

var MetaStoreUsed = false
var MetaStore metastore.Store

var SessionStoreUsed = true
var SessionStore sessionstore.Store

var ShopStoreUsed = false
var ShopStore shopstore.Store

var SqlFileStoreUsed = false
var SqlFileStorage filesystem.StorageInterface

var StatsStoreUsed = true
var StatsStore statsstore.Store

// var SubscriptionStore *subscriptionstore.Store

var TaskStoreUsed = true
var TaskStore taskstore.StoreInterface

var UserStoreUsed = false
var UserStore userstore.StoreInterface

var VaultStoreUsed = false
var VaultStore vaultstore.Store

var Logger slog.Logger

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
