package widgets

import (
	"net/http"
	"project/internal/helpers"
)

// NewAuthenticatedWidget returns a new instance of authenticatedWidget.
//
// Parameters:
// - none
//
// Returns:
// - *authenticatedWidget: a pointer to an authenticatedWidget.
func NewAuthenticatedWidget() *authenticatedWidget {
	return &authenticatedWidget{}
}

// authenticatedWidget used to render the authenticatedWidget shortcode.
//
// It displays the content of the shortcode if the user is authenticated.
type authenticatedWidget struct{}

// Render implements the shortcode interface.
func (w *authenticatedWidget) Render(req *http.Request, content string, data map[string]string) string {
	authUser := helpers.GetAuthUser(req)

	if authUser != nil {
		return content
	}

	return ""
}
