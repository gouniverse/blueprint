package website

import (
	"net/http"
	"strings"

	"github.com/gouniverse/router"
)

type securityTxtController struct{}

// NewSecurityTxtController creates a new instance of the securityTxtController struct.
//
// Returns:
// - *securityTxtController: a pointer to the newly created securityTxtController.
func NewSecurityTxtController() *securityTxtController {
	return &securityTxtController{}
}

var _ router.HTMLControllerInterface = &securityTxtController{}

func (c securityTxtController) Handler(w http.ResponseWriter, r *http.Request) string {
	text := `
# Our security contact form
Contact: https://tiny.vip/BlCe

# Date this document expires
Expires: 2029-12-31T23:59:59z
	`

	w.Header().Set("Content-Type", "text/plain")

	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", "\r\n")

	return text
}
