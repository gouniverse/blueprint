package config

import (
	"os"
)

// TestsConfigureAndInitialize configures the test environment
// variables and initializes the test environment
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

	os.Setenv("VAULT_KEY", "abcdefghijklmnopqrstuvwxyz1234567890")

	os.Setenv("STRIPE_KEY_PRIVATE", "123")
	os.Setenv("STRIPE_KEY_PUBLIC", "345")
	os.Setenv("VERTEX_PROJECT_ID", "123")
	os.Setenv("VERTEX_REGION_ID", "345")
	os.Setenv("VERTEX_MODEL_ID", "678")

	Initialize()
}
