package website

import (
	"net/http"
	"project/app/links"

	"github.com/gouniverse/router"
)

type robotsTxtController struct{}

// NewRobotsTxtController creates a new instance of the robotsTxtController struct.
//
// Returns:
// - *robotsTxtController: a pointer to the newly created robotsTxtController.
func NewRobotsTxtController() *robotsTxtController {
	return &robotsTxtController{}
}

var _ router.HTMLControllerInterface = (*robotsTxtController)(nil)

func (c robotsTxtController) Handler(w http.ResponseWriter, r *http.Request) string {

	// Allow: /contact
	// Allow: /faq
	// Allow: /marketplace

	webpage := `
User-agent: *
Allow: /
Allow: /about
Allow: /blog
Allow: /blog/post/*
Allow: /blog/post/*/*
Allow: /privacy-policy
Allow: /terms-of-use

Disallow: /admin/
Disallow: /api/
Disallow: /auth/
Disallow: /f/
Disallow: /c/
Disallow: /certificate/
Disallow: /files/
Disallow: /flash
Disallow: /media/
Disallow: /message
Disallow: /theme
Disallow: /user/
Disallow: /*-ajax$

Sitemap: ` + links.RootURL() + `/sitemap.xml
	`

	w.Header().Set("Content-Type", "text/plain")
	return webpage
}
