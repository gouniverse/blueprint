package layouts

import (
	"net/http"

	"github.com/gouniverse/hb"
)

type Options struct {
	Request        *http.Request
	WebsiteSection string
	Title          string
	Content        *hb.Tag
	ScriptURLs     []string
	Scripts        []string
	StyleURLs      []string
	Styles         []string
}
