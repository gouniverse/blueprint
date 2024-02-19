package widgets

import (
	"net/http"
	"project/internal/helpers"
)

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

// unauthenticatedWidget used to render the unauthenticatedWidget shortcode.
//
// It displays the content of the shortcode if the user is not authenticated.
type unauthenticatedWidget struct{}

// Render implements the shortcode interface.
func (w *unauthenticatedWidget) Render(req *http.Request, content string, data map[string]string) string {
	authUser := helpers.GetAuthUser(req)

	if authUser == nil {
		return content
	}

	return ""
}
