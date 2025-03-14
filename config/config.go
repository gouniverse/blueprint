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
	"github.com/gouniverse/entitystore"
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

// == TYPES ================================================================= //

// AuthenticatedUserContextKey is a context key for the authenticated user.
type AuthenticatedUserContextKey struct{}

// AuthenticatedSessionContextKey is a context key for the authenticated session.
type AuthenticatedSessionContextKey struct{}

// == CONSTANTS ============================================================= //

const APP_ENVIRONMENT_DEVELOPMENT = "development"
const APP_ENVIRONMENT_LOCAL = "local"
const APP_ENVIRONMENT_PRODUCTION = "production"
const APP_ENVIRONMENT_STAGING = "staging"
const APP_ENVIRONMENT_TESTING = "testing"

// == VARIABLES ============================================================= //

// AppEnvironment is the environment the application is running in
// (e.g., development, production, testing).
var AppEnvironment string

// AppName is the name of the application.
var AppName string

// AppUrl is the URL of the application.
var AppUrl string

// AuthEndpoint is the authentication endpoint.
var AuthEndpoint = "/auth"

// Database is the database interface.
var Database sb.DatabaseInterface

// DbDriver is the database driver to use (e.g., mysql, postgres, sqlite).
var DbDriver string

// DbHost is the database host.
var DbHost string

// DbName is the database name.
var DbName string

// DbPass is the database password.
var DbPass string

// DbPort is the database port.
var DbPort string

// DbUser is the database user.
var DbUser string

// Debug is a boolean indicating whether the application is in debug mode.
var Debug bool

// MailDriver is the mail driver to use (e.g., smtp, sendgrid).
var MailDriver string

// MailFromEmailAddress is the email address to send emails from.
var MailFromEmailAddress string

// MailFromName is the name to send emails from.
var MailFromName string

// MailHost is the mail host.
var MailHost string

// MailPort is the mail port.
var MailPort string

// MailPassword is the mail password.
var MailPassword string

// MailUsername is the mail username.
var MailUsername string

// MediaBucket is the media bucket to use.
var MediaBucket string

// MediaDriver is the media driver to use (e.g., s3, gcs, filesystem).
var MediaDriver string

// MediaKey is the media key.
var MediaKey string

// MediaEndpoint is the media endpoint.
var MediaEndpoint string

// MediaRegion is the media region.
var MediaRegion string

// MediaRoot is the media root.
var MediaRoot string = "/"

// MediaSecret is the media secret.
var MediaSecret string

// MediaUrl is the media URL.
var MediaUrl string = "/files"

// OpenAiApiKey is the OpenAI API key.
var OpenAiApiKey string

// StripeKeyPrivate is the Stripe private key.
var StripeKeyPrivate string

// StripeKeyPublic is the Stripe public key.
var StripeKeyPublic string

// TranslationLanguageDefault is the default translation language.
var TranslationLanguageDefault string = "en"

// TranslationLanguageList is the list of supported translation languages.
var TranslationLanguageList map[string]string = map[string]string{"en": "English", "bg": "Bulgarian", "de": "German"}

// VaultKey is the Vault key.
var VaultKey string

// VertexModelID is the Vertex model ID.
var VertexModelID string

// VertexProjectID is the Vertex project ID.
var VertexProjectID string

// VertexRegionID is the Vertex region ID.
var VertexRegionID string

// WebServer is the web server.
var WebServer *webserver.Server

// WebServerHost is the web server host.
var WebServerHost string

// WebServerPort is the web server port.
var WebServerPort string

// == CACHE ================================================================ //

var CacheMemory *ttlcache.Cache[string, any]
var CacheFile cachego.Cache

// == CMS OLD ============================================================== //

// Cms is the old CMS package (replaced by CmsStore).
var CmsUsed = true
var Cms cms.Cms

// == STORES =============================================================== //

var BlindIndexStoreUsed = true
var BlindIndexStoreEmail blindindexstore.Store
var BlindIndexStoreFirstName blindindexstore.Store
var BlindIndexStoreLastName blindindexstore.Store

var BlogStoreUsed = true
var BlogStore blogstore.StoreInterface

var CmsStoreUsed = false
var CmsStore cmsstore.StoreInterface

// CmsUserTemplateID is the CMS user template ID.
var CmsUserTemplateID string

var CacheStoreUsed = true
var CacheStore cachestore.Store

// var CommentStore *commentstore.Store

var CustomStoreUsed = false
var CustomStore customstore.Store

// used by the testimonials package
var EntityStoreUsed = true
var EntityStore entitystore.Store

var GeoStoreUsed = true
var GeoStore geostore.Store

var LogStoreUsed = true
var LogStore logstore.Store

var MetaStoreUsed = false
var MetaStore metastore.Store

var SessionStoreUsed = true
var SessionStore sessionstore.StoreInterface

var ShopStoreUsed = false
var ShopStore shopstore.StoreInterface

var SqlFileStoreUsed = false
var SqlFileStorage filesystem.StorageInterface

var StatsStoreUsed = true
var StatsStore statsstore.StoreInterface

// var SubscriptionStore *subscriptionstore.Store

var TaskStoreUsed = true
var TaskStore taskstore.StoreInterface

var UserStoreUsed = true
var UserStore userstore.StoreInterface

var VaultStoreUsed = false
var VaultStore vaultstore.StoreInterface

var Logger slog.Logger
