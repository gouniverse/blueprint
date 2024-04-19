package widgets

import (
	"net/http"
	"project/config"

	"github.com/samber/lo"
)

var _ Widget = (*visibleWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewVisibleWidget creates a new instance of the show widget
//
// Parameters:
//   - None
//
// Returns:
//   - *visibleWidget - A pointer to the show widget
func NewVisibleWidget() *visibleWidget {
	return &visibleWidget{}
}

// == WIDGET ================================================================

// print is the struct that will be used to render the print shortcode.
//
// This shortcode is used to show the result of the provided content
// if a condition is met.
//
// Example:
// <x-visible environment="production">content</x-visible>
type visibleWidget struct{}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (w *visibleWidget) Alias() string {
	return "x-visible"
}

// Description a user-friendly description of the shortcode.
func (w *visibleWidget) Description() string {
	return "Renders the content if the condition is met"
}

// Render implements the shortcode interface.
func (w *visibleWidget) Render(r *http.Request, content string, params map[string]string) string {
	environment := lo.ValueOr(params, "environment", "")

	if environment != "" {
		if w.isEnvironmentMatch(environment) {
			return content
		}
	}

	return "" // No content is shown by default
}

// == PRIVATE METHODS ========================================================

func (t *visibleWidget) isEnvironmentMatch(environment string) bool {
	if environment == "" {
		return false
	}

	if environment == config.APP_ENVIRONMENT_DEVELOPMENT && config.IsEnvDevelopment() {
		return true
	}

	if environment == config.APP_ENVIRONMENT_LOCAL && config.IsEnvLocal() {
		return true
	}

	if environment == config.APP_ENVIRONMENT_PRODUCTION && config.IsEnvProduction() {
		return true
	}

	if environment == config.APP_ENVIRONMENT_STAGING && config.IsEnvStaging() {
		return true
	}

	if environment == config.APP_ENVIRONMENT_TESTING && config.IsEnvTesting() {
		return true
	}

	return false

}
