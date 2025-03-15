package config_v2

// == TYPES ================================================================= //

// AuthenticatedUserContextKey is a context key for the authenticated user.
type AuthenticatedUserContextKey struct{}

// AuthenticatedSessionContextKey is a context key for the authenticated session.
type AuthenticatedSessionContextKey struct{}

// configContextKey is a context key for the configuration.
// it is deliberatly not exported
type configContextKey struct{}

// == CONSTANTS ============================================================= //

const APP_ENVIRONMENT_DEVELOPMENT = "development"
const APP_ENVIRONMENT_LOCAL = "local"
const APP_ENVIRONMENT_PRODUCTION = "production"
const APP_ENVIRONMENT_STAGING = "staging"
const APP_ENVIRONMENT_TESTING = "testing"
