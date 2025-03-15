package config_v2

import (
	"os"

	_ "modernc.org/sqlite"
)

// TestsConfigureAndInitialize configures the test environment
// variables and initializes the test environment
//
// Business logic:
//   - configures the test environment variables
//   - initializes the test environment
//
// Parameters:
// - none
//
// Returns:
// - none
func TestsConfigureAndInitialize() (*Config, error) {
	TestsSetEnvironmentVariables()

	// cfg := &config{
	// 	// AppEnvironment:   APP_ENVIRONMENT_TESTING,
	// 	// AppName:          "TEST APP NAME",
	// 	// AppUrl:           "http://localhost:8080",
	// 	// DbDriver:         "sqlite",
	// 	// DbHost:           "",
	// 	// DbPort:           "",
	// 	// DbName:           "file::memory:?cache=shared",
	// 	// DbUser:           "",
	// 	// DbPass:           "",
	// 	// Debug:            false,
	// 	// MailDriver:       "smtp",
	// 	// MailHost:         "127.0.0.1",
	// 	// MailPort:         "32435",
	// 	// MailUsername:     "",
	// 	// MailPassword:     "",
	// 	// EmailFromAddress: "admintest@test.com",
	// 	// EmailFromName:    "Admin Test",
	// 	// CmsTemplateId:    "default",
	// 	// VaultKey:         "abcdefghijklmnopqrstuvwxyz1234567890",
	// 	// OpenAiApiKey:     "openai_api_key",
	// 	// StripeKeyPrivate: "sk_test_yoursecretkey",
	// 	// StripeKeyPublic:  "pk_test_yourpublickey",
	// 	// VertexProjectId:  "vertex_project_id",
	// 	// VertexRegionId:   "vertex_region_id",
	// 	// VertexModelId:    "vertex_model_id",
	// }

	// cfg.SetAppName("TEST APP NAME").
	// 	SetAppUrl("http://localhost:8080").
	// 	SetDbDriver("sqlite").
	// 	SetDbHost("").
	// 	SetDbPort("").
	// 	SetDbName("file::memory:?cache=shared").
	// 	SetDbUser("").
	// 	SetDbPass("").
	// 	SetDebug(false).
	// 	SetMailDriver("smtp").
	// 	SetMailHost("127.0.0.1").
	// 	SetMailPort("32435").
	// 	SetMailUsername("").
	// 	SetMailPassword("").
	// 	SetMailFromEmailAddress("admintest@test.com").
	// 	SetMailFromName("Admin Test").
	// 	SetCmsUserTemplateId("default").
	// 	SetVaultKey("abcdefghijklmnopqrstuvwxyz1234567890").
	// 	SetOpenAiApiKey("openai_api_key").
	// 	SetStripeKeyPrivate("sk_test_yoursecretkey").
	// 	SetStripeKeyPublic("pk_test_yourpublickey").
	// 	SetVertexProjectId("vertex_project_id").
	// 	SetVertexRegionId("vertex_region_id").
	// 	SetVertexModelId("vertex_model_id")

	cfg, err := New()
	return cfg, err
}

// TestsSetEnvironmentVariables sets the test environment variables
// Parameters:
// - none
//
// Returns:
// - none
func TestsSetEnvironmentVariables() {
	os.Setenv("APP_NAME", "TEST APP NAME")
	os.Setenv("APP_URL", "http://localhost:8080")
	os.Setenv("APP_ENV", APP_ENVIRONMENT_TESTING)

	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_DATABASE", "file::memory:?cache=shared")
	os.Setenv("DB_USERNAME", "")
	os.Setenv("DB_PASSWORD", "")

	// os.Setenv("DEBUG", "yes")

	os.Setenv("ENV_ENCRYPTION_KEY", "123456")

	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("MAIL_DRIVER", "smtp")
	os.Setenv("MAIL_HOST", "127.0.0.1")
	os.Setenv("MAIL_PORT", "32435")
	os.Setenv("MAIL_USERNAME", "")
	os.Setenv("MAIL_PASSWORD", "")

	os.Setenv("EMAIL_FROM_ADDRESS", "admintest@test.com")
	os.Setenv("EMAIL_FROM_NAME", "Admin Test")

	os.Setenv("CMS_TEMPLATE_ID", "default")

	os.Setenv("VAULT_KEY", "abcdefghijklmnopqrstuvwxyz1234567890")

	os.Setenv("OPENAI_API_KEY", "openai_api_key")

	os.Setenv("STRIPE_KEY_PRIVATE", "sk_test_yoursecretkey")
	os.Setenv("STRIPE_KEY_PUBLIC", "pk_test_yourpublickey")

	os.Setenv("VERTEX_PROJECT_ID", "vertex_project_id")
	os.Setenv("VERTEX_REGION_ID", "vertex_region_id")
	os.Setenv("VERTEX_MODEL_ID", "vertex_model_id")
}
