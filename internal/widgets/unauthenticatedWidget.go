package widgets

import (
	"net/http"
	"project/internal/helpers"

	"github.com/gouniverse/cms"
)

var _ cms.ShortcodeInterface = (*unauthenticatedWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewUnauthenticatedWidget creates a new instance of unauthenticatedWidget.
//
// Parameters:
// - none
//
// Returns:
// - *unauthenticatedWidget: a pointer to an unauthenticatedWidget.
func NewUnauthenticatedWidget() *unauthenticatedWidget {
	return &unauthenticatedWidget{}
}

// == WIDGET =================================================================

// unauthenticatedWidget used to render the unauthenticatedWidget shortcode.
//
// It displays the content of the shortcode if the user is not authenticated.
type unauthenticatedWidget struct{}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (w *unauthenticatedWidget) Alias() string {
	return "x-unauthenticated"
}

// Description a user-friendly description of the shortcode.
func (w *unauthenticatedWidget) Description() string {
	return "Renders the content if the user is not authenticated"
}

// Render implements the shortcode interface.
func (w *unauthenticatedWidget) Render(req *http.Request, content string, data map[string]string) string {
	authUser := helpers.GetAuthUser(req)

	if authUser == nil {
		return content
	}

	return ""
}
