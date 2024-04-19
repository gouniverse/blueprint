package widgets

import (
	"net/http"

	"github.com/mingrammer/cfmt"
	"github.com/robertkrimen/otto"
)

var _ Widget = (*printWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewPrintWidget creates a new instance of the print struct.
//
// Parameters:
//   - None
//
// Returns:
//   - *print - A pointer to the print struct
func NewPrintWidget() *printWidget {
	return &printWidget{}
}

// == WIDGET ================================================================

// print is the struct that will be used to render the print shortcode.
//
// This shortcode is used to evaluate the result of the provided content
// and return it.
//
// It uses Otto as the engine.
type printWidget struct{}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (t *printWidget) Alias() string {
	return "x-print"
}

// Description a user-friendly description of the shortcode.
func (t *printWidget) Description() string {
	return "Renders the result of the provided content"
}

// Render implements the shortcode interface.
func (t *printWidget) Render(r *http.Request, content string, params map[string]string) string {
	path := r.URL.Path

	vm := otto.New()

	vm.Set("path", path)

	result, err := vm.Run("result = " + content)

	if err != nil {
		cfmt.Errorln(err)
	}

	return result.String()
}
