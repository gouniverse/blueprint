package widgets

import (
	"net/http"

	"github.com/mingrammer/cfmt"
	"github.com/robertkrimen/otto"
)

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

// print is the struct that will be used to render the print shortcode.
//
// This shortcode is used to evaluate the result of the provided content
// and return it.
//
// It uses Otto as the engine.
type printWidget struct{}

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
