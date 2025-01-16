package config

import (
	"os"
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
func TestsConfigureAndInitialize() {
	os.Setenv("APP_NAME", "Blueprint")
	os.Setenv("APP_URL", "http://localhost:8080")
	os.Setenv("APP_ENV", APP_ENVIRONMENT_TESTING)
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_DATABASE", "file::memory:?cache=shared")
	os.Setenv("DB_USERNAME", "")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("ENV_ENCRYPTION_KEY", "123456")
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	// os.Setenv("DEBUG", "yes")
	// os.Setenv("MAIL_DRIVER", "smtp")
	// os.Setenv("MAIL_HOST", "127.0.0.1")
	// os.Setenv("MAIL_PORT", "32435")
	// os.Setenv("MAIL_USERNAME", "")
	// os.Setenv("MAIL_PASSWORD", "")
	// os.Setenv("MAIL_ENCRYPTION", "")
	// os.Setenv("MAIL_HOST", "smtp.mailtrap.io")
	// os.Setenv("MAIL_PORT", "2525")
	// os.Setenv("MAIL_USERNAME", "")
	// os.Setenv("MAIL_PASSWORD", "")
	// os.Setenv("MAIL_ENCRYPTION", "")

	// os.Setenv("EMAIL_FROM_ADDRESS", "")
	// os.Setenv("EMAIL_FROM_NAME", "")

	os.Setenv("CMS_TEMPLATE_ID", "default")

	os.Setenv("VAULT_KEY", "abcdefghijklmnopqrstuvwxyz1234567890")

	os.Setenv("OPENAI_API_KEY", "openai_api_key")
	os.Setenv("STRIPE_KEY_PRIVATE", "stripe_key_private")
	os.Setenv("STRIPE_KEY_PUBLIC", "stripe_key_public")
	os.Setenv("VERTEX_PROJECT_ID", "vertex_project_id")
	os.Setenv("VERTEX_REGION_ID", "vertex_region_id")
	os.Setenv("VERTEX_MODEL_ID", "vertex_model_id")

	Initialize()
}
