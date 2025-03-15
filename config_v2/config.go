package config_v2

import (
	"context"
	"database/sql"
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

// Config holds all application configuration
type Config struct {
	databaseInits      []func(db *sql.DB) error
	databaseMigrations []func(ctx context.Context) error

	// AppEnvironment is the environment the application is running in
	// (e.g., development, production, testing).
	AppEnvironment string

	// AppName is the name of the application.
	AppName string

	// AppUrl is the URL of the application.
	AppUrl string

	// AppVersion is the version of the application.
	AppVersion string

	// AuthEndpoint is the authentication endpoint.
	AuthEndpoint string

	// Database is the database interface.
	Database sb.DatabaseInterface

	// DbDriver is the database driver to use (e.g., mysql, postgres, sqlite).
	DbDriver string

	// DbHost is the database host.
	DbHost string

	// DbName is the database name.
	DbName string

	// DbPass is the database password.
	DbPass string

	// DbPort is the database port.
	DbPort string

	// DbUser is the database user.
	DbUser string

	// Debug is a boolean indicating whether the application is in debug mode.
	Debug bool

	// Logger is the logger.
	Logger slog.Logger

	// MailDriver is the mail driver to use (e.g., smtp, sendgrid).
	MailDriver string

	// MailFromEmailAddress is the email address to send emails from.
	MailFromEmailAddress string

	// MailFromName is the name to send emails from.
	MailFromName string

	// MailHost is the mail host.
	MailHost string

	// MailPort is the mail port.
	MailPort string

	// MailPassword is the mail password.
	MailPassword string

	// MailUsername is the mail username.
	MailUsername string

	// MediaBucket is the media bucket to use.
	MediaBucket string

	// MediaDriver is the media driver to use (e.g., s3, gcs, filesystem).
	MediaDriver string

	// MediaKey is the media key.
	MediaKey string

	// MediaEndpoint is the media endpoint.
	MediaEndpoint string

	// MediaRegion is the media region.
	MediaRegion string

	// MediaRoot is the media root.
	MediaRoot string

	// MediaSecret is the media secret.
	MediaSecret string

	// MediaUrl is the media URL.
	MediaUrl string

	// OpenAiApiKey is the OpenAI API key.
	OpenAiApiKey string

	// StripeKeyPrivate is the Stripe private key.
	StripeKeyPrivate string

	// StripeKeyPublic is the Stripe public key.
	StripeKeyPublic string

	// TranslationLanguageDefault is the default translation language.
	TranslationLanguageDefault string

	// TranslationLanguageList is the list of supported translation languages.
	TranslationLanguageList map[string]string

	// VaultKey is the Vault key.
	VaultKey string

	// VertexModelID is the Vertex model ID.
	VertexModelID string

	// VertexProjectID is the Vertex project ID.
	VertexProjectID string

	// VertexRegionID is the Vertex region ID.
	VertexRegionID string

	// WebServer is the web server.
	WebServer *webserver.Server

	// WebServerHost is the web server host.
	WebServerHost string

	// WebServerPort is the web server port.
	WebServerPort string

	// == CACHE ================================================================ //

	// CacheMemory is the memory cache.
	CacheMemory *ttlcache.Cache[string, any]

	// CacheFile is the file cache.
	CacheFile cachego.Cache

	// == CMS OLD ============================================================== //

	// Cms is the old CMS package (replaced by CmsStore).
	CmsUsed bool
	Cms     cms.Cms

	// == STORES =============================================================== //

	BlindIndexStoreUsed      bool
	BlindIndexStoreEmail     blindindexstore.StoreInterface
	BlindIndexStoreFirstName blindindexstore.StoreInterface
	BlindIndexStoreLastName  blindindexstore.StoreInterface

	BlogStoreUsed bool
	BlogStore     blogstore.StoreInterface

	CmsStoreUsed bool
	CmsStore     cmsstore.StoreInterface

	// CmsUserTemplateID is the CMS user template ID.
	CmsUserTemplateID string

	CacheStoreUsed bool
	CacheStore     cachestore.StoreInterface

	// var CommentStore *commentstore.Store

	CustomStoreUsed bool
	CustomStore     customstore.StoreInterface

	// used by the testimonials package
	EntityStoreUsed bool
	EntityStore     entitystore.StoreInterface

	GeoStoreUsed bool
	GeoStore     geostore.StoreInterface

	LogStoreUsed bool
	LogStore     logstore.StoreInterface

	MetaStoreUsed bool
	MetaStore     metastore.StoreInterface

	SessionStoreUsed bool
	SessionStore     sessionstore.StoreInterface

	ShopStoreUsed bool
	ShopStore     shopstore.StoreInterface

	SqlFileStoreUsed bool
	SqlFileStorage   filesystem.StorageInterface

	StatsStoreUsed bool
	StatsStore     statsstore.StoreInterface

	// var SubscriptionStore *subscriptionstore.Store

	TaskStoreUsed bool
	TaskStore     taskstore.StoreInterface

	UserStoreUsed bool
	UserStore     userstore.StoreInterface

	VaultStoreUsed bool
	VaultStore     vaultstore.StoreInterface
}

// New creates a new configuration instance
func New() (*Config, error) {
	c := &Config{}
	err := c.initialize()
	return c, err
}

// ========================================================================
// == METHODS
// ========================================================================

// Close closes any resources associated with the configuration
func (c *Config) Close() error {
	if c.Database != nil {
		if closer, ok := c.Database.(interface{ Close() error }); ok {
			return closer.Close()
		}
	}

	return nil
}

// IsDebugEnabled returns whether debug mode is enabled
func (c *Config) IsDebugEnabled() bool {
	return c.Debug
}

// IsEnvDevelopment returns whether the environment is development
func (c *Config) IsEnvDevelopment() bool {
	return c.AppEnvironment == APP_ENVIRONMENT_DEVELOPMENT
}

// IsEnvProduction returns whether the environment is production
func (c *Config) IsEnvProduction() bool {
	return c.AppEnvironment == APP_ENVIRONMENT_PRODUCTION
}

// IsEnvLocal returns whether the environment is local
func (c *Config) IsEnvLocal() bool {
	return c.AppEnvironment == APP_ENVIRONMENT_LOCAL
}

// IsEnvTesting returns whether the environment is testing
func (c *Config) IsEnvTesting() bool {
	return c.AppEnvironment == APP_ENVIRONMENT_TESTING
}

// ========================================================================
// == SETTERS AND GETTERS
// ========================================================================

func (c *Config) GetAppEnvironment() string {
	return c.AppEnvironment
}

func (c *Config) SetAppEnvironment(env string) *Config {
	c.AppEnvironment = env
	return c
}

// GetAppName returns the application name from the configuration
func (c *Config) GetAppName() string {
	return c.AppName
}

// SetAppName sets the application name for the configuration
func (c *Config) SetAppName(name string) *Config {
	c.AppName = name
	return c
}

// GetAppUrl returns the application URL from the configuration
func (c *Config) GetAppUrl() string {
	return c.AppUrl
}

func (c *Config) SetAppUrl(s string) *Config {
	c.AppUrl = s
	return c
}

// GetAppVersion returns the application version from the configuration
func (c *Config) GetAppVersion() string {
	return c.AppVersion
}

// SetAppVersion sets the application version for the configuration
func (c *Config) SetAppVersion(version string) *Config {
	c.AppVersion = version
	return c
}

// GetDatabase returns the database from the configuration
func (c *Config) GetDatabase() sb.DatabaseInterface {
	return c.Database
}

// SetDatabase sets the database for the configuration
func (c *Config) SetDatabase(db sb.DatabaseInterface) *Config {
	c.Database = db
	return c
}

// GetDbDriver returns the database driver from the configuration
func (c *Config) GetDbDriver() string {
	return c.DbDriver
}

// SetDbDriver sets the database driver for the configuration
func (c *Config) SetDbDriver(driver string) *Config {
	c.DbDriver = driver
	return c
}

// GetDbHost returns the database host from the configuration
func (c *Config) GetDbHost() string {
	return c.DbHost
}

// SetDbHost sets the database host for the configuration
func (c *Config) SetDbHost(host string) *Config {
	c.DbHost = host
	return c
}

// GetDbPort returns the database port from the configuration
func (c *Config) GetDbPort() string {
	return c.DbPort
}

// SetDbPort sets the database port for the configuration
func (c *Config) SetDbPort(port string) *Config {
	c.DbPort = port
	return c
}

// GetDbUser returns the database user from the configuration
func (c *Config) GetDbUser() string {
	return c.DbUser
}

// SetDbUser sets the database user for the configuration
func (c *Config) SetDbUser(user string) *Config {
	c.DbUser = user
	return c
}

// GetDbPass returns the database password from the configuration
func (c *Config) GetDbPass() string {
	return c.DbPass
}

// SetDbPass sets the database password for the configuration
func (c *Config) SetDbPass(pass string) *Config {
	c.DbPass = pass
	return c
}

// GetDbName returns the database name from the configuration
func (c *Config) GetDbName() string {
	return c.DbName
}

// SetDbName sets the database name for the configuration
func (c *Config) SetDbName(name string) *Config {
	c.DbName = name
	return c
}

// SetUserStore sets the user store for the configuration
func (c *Config) SetUserStore(store userstore.StoreInterface) *Config {
	c.UserStore = store
	return c
}

// GetUserStore returns the user store from the configuration
func (c *Config) GetUserStore() userstore.StoreInterface {
	return c.UserStore
}

// SetBlogStore sets the blog store for the configuration
func (c *Config) SetBlogStore(store blogstore.StoreInterface) *Config {
	c.BlogStore = store
	return c
}

// GetBlogStore returns the blog store from the configuration
func (c *Config) GetBlogStore() blogstore.StoreInterface {
	return c.BlogStore
}

// SetSessionStore sets the session store for the configuration
func (c *Config) SetSessionStore(store sessionstore.StoreInterface) *Config {
	c.SessionStore = store
	return c
}

// GetSessionStore returns the session store from the configuration
func (c *Config) GetSessionStore() sessionstore.StoreInterface {
	return c.SessionStore
}

// GetDebug returns the debug mode from the configuration
func (c *Config) GetDebug() bool {
	return c.Debug
}

// SetDebug sets the debug mode for the configuration
func (c *Config) SetDebug(debug bool) *Config {
	c.Debug = debug
	return c
}

// GetLogger returns the logger from the configuration
func (c *Config) GetLogger() slog.Logger {
	return c.Logger
}

// SetLogger sets the logger for the configuration
func (c *Config) SetLogger(logger slog.Logger) *Config {
	c.Logger = logger
	return c
}

// GetMailDriver returns the mail driver from the configuration
func (c *Config) GetMailDriver() string {
	return c.MailDriver
}

// SetMailDriver sets the mail driver for the configuration
func (c *Config) SetMailDriver(driver string) *Config {
	c.MailDriver = driver
	return c
}

func (c *Config) GetMailHost() string {
	return c.MailHost
}

// SetMailHost sets the mail host for the configuration
func (c *Config) SetMailHost(host string) *Config {
	c.MailHost = host
	return c
}

// GetMailPort returns the mail port from the configuration
func (c *Config) GetMailPort() string {
	return c.MailPort
}

// SetMailPort sets the mail port for the configuration
func (c *Config) SetMailPort(port string) *Config {
	c.MailPort = port
	return c
}

// GetMailUsername returns the mail username from the configuration
func (c *Config) GetMailUsername() string {
	return c.MailUsername
}

// SetMailUsername sets the mail username for the configuration
func (c *Config) SetMailUsername(username string) *Config {
	c.MailUsername = username
	return c
}

// GetMailPassword returns the mail password from the configuration
func (c *Config) GetMailPassword() string {
	return c.MailPassword
}

// SetMailPassword sets the mail password for the configuration
func (c *Config) SetMailPassword(password string) *Config {
	c.MailPassword = password
	return c
}

// GetMailFromEmailAddress returns the mail from email address from the configuration
func (c *Config) GetMailFromEmailAddress() string {
	return c.MailFromEmailAddress
}

// SetMailFromEmailAddress sets the mail from email address for the configuration
func (c *Config) SetMailFromEmailAddress(address string) *Config {
	c.MailFromEmailAddress = address
	return c
}

// GetMailFromName returns the mail from name from the configuration
func (c *Config) GetMailFromName() string {
	return c.MailFromName
}

// SetMailFromName sets the mail from name for the configuration
func (c *Config) SetMailFromName(name string) *Config {
	c.MailFromName = name
	return c
}

// GetCmsUserTemplateId returns the CMS template ID from the configuration
func (c *Config) GetCmsUserTemplateId() string {
	return c.CmsUserTemplateID
}

// SetCmsTemplateId sets the CMS template ID for the configuration
func (c *Config) SetCmsUserTemplateId(id string) *Config {
	c.CmsUserTemplateID = id
	return c
}

// GetVaultKey returns the vault key from the configuration
func (c *Config) GetVaultKey() string {
	return c.VaultKey
}

// SetVaultKey sets the vault key for the configuration
func (c *Config) SetVaultKey(key string) *Config {
	c.VaultKey = key
	return c
}

// GetOpenAiApiKey returns the OpenAI API key from the configuration
func (c *Config) GetOpenAiApiKey() string {
	return c.OpenAiApiKey
}

// SetOpenAiApiKey sets the OpenAI API key for the configuration
func (c *Config) SetOpenAiApiKey(key string) *Config {
	c.OpenAiApiKey = key
	return c
}

// GetStripeKeyPrivate returns the Stripe private key from the configuration
func (c *Config) GetStripeKeyPrivate() string {
	return c.StripeKeyPrivate
}

// SetStripeKeyPrivate sets the Stripe private key for the configuration
func (c *Config) SetStripeKeyPrivate(key string) *Config {
	c.StripeKeyPrivate = key
	return c
}

// GetStripeKeyPublic returns the Stripe public key from the configuration
func (c *Config) GetStripeKeyPublic() string {
	return c.StripeKeyPublic
}

// SetStripeKeyPublic sets the Stripe public key for the configuration
func (c *Config) SetStripeKeyPublic(key string) *Config {
	c.StripeKeyPublic = key
	return c
}

// GetVertexProjectId returns the Vertex project ID from the configuration
func (c *Config) GetVertexProjectId() string {
	return c.VertexProjectID
}

// SetVertexProjectId sets the Vertex project ID for the configuration
func (c *Config) SetVertexProjectId(id string) *Config {
	c.VertexProjectID = id
	return c
}

// GetVertexRegionId returns the Vertex region ID from the configuration
func (c *Config) GetVertexRegionId() string {
	return c.VertexRegionID
}

// SetVertexRegionId sets the Vertex region ID for the configuration
func (c *Config) SetVertexRegionId(id string) *Config {
	c.VertexRegionID = id
	return c
}

// GetVertexModelId returns the Vertex model ID from the configuration
func (c *Config) GetVertexModelId() string {
	return c.VertexModelID
}

// SetVertexModelId sets the Vertex model ID for the configuration
func (c *Config) SetVertexModelId(id string) *Config {
	c.VertexModelID = id
	return c
}
