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

	Initialize()
}
