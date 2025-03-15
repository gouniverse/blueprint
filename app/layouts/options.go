package layouts

import (
	"net/http"

	"github.com/gouniverse/hb"
)

// Options defines the options for the layout
type Options struct {
	Request        *http.Request
	WebsiteSection string
	Title          string
	Content        hb.TagInterface
	ScriptURLs     []string
	Scripts        []string
	StyleURLs      []string
	Styles         []string
}
